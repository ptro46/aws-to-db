package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
)

type AwsLoadBalancer struct {
	LoadBalancerArn       string                `json:"LoadBalancerArn,omitempty"`
	DNSName               string                `json:"DNSName,omitempty"`
	CanonicalHostedZoneId string                `json:"CanonicalHostedZoneId,omitempty"`
	CreatedTime           string                `json:"CreatedTime,omitempty"`
	LoadBalancerName      string                `json:"LoadBalancerName,omitempty"`
	Scheme                string                `json:"Scheme,omitempty"`
	VpcId                 string                `json:"VpcId,omitempty"`
	State                 AwsLoadBalancerState  `json:"State,omitempty"`
	Type                  string                `json:"Type,omitempty"`
	AvailabilityZones     []AwsAvailabilityZone `json:"AvailabilityZones,omitempty"`
	SecurityGroups        []string              `json:"SecurityGroups,omitempty"`
	IpAddressType         string                `json:"IpAddressType,omitempty"`
}

type AwsLoadBalancers struct {
	LoadBalancers []AwsLoadBalancer `json:"LoadBalancers,omitempty"`
}

type AwsLoadBalancerAddress struct {
	IpAddress    string `json:"IpAddress,omitempty"`
	AllocationId string `json:"AllocationId,omitempty"`
}

type AwsLoadBalancerState struct {
	Code   string `json:"Code,omitempty"`
	Reason string `json:"Reason,omitempty"`
}

func awsLoadBalancerParseAndPersist(db *sql.DB, jsonString string) error {
	fmt.Printf("awsLoadBalancerParseAndPersist with json lenght %d\n", len(jsonString))
	awsLoadBalancers := new(AwsLoadBalancers)
	errUnmarshal := json.Unmarshal([]byte(jsonString), awsLoadBalancers)
	if errUnmarshal == nil {
		for _, awsLoadBalancer := range awsLoadBalancers.LoadBalancers {
			fmt.Printf("LoadBalancer DNS(%s)\n", awsLoadBalancer.DNSName)
			load := createLoadBalancer(db, awsLoadBalancer.LoadBalancerArn, awsLoadBalancer.DNSName, awsLoadBalancer.CanonicalHostedZoneId, awsLoadBalancer.CreatedTime, awsLoadBalancer.LoadBalancerName, awsLoadBalancer.Scheme, awsLoadBalancer.VpcId, awsLoadBalancer.State.Code, awsLoadBalancer.State.Reason, awsLoadBalancer.Type, awsLoadBalancer.IpAddressType)
			if load != nil {
				fmt.Printf("    LoadBalancer %s loaded\n", load.DNSName)

				for _, awsLoadSecGroup := range awsLoadBalancer.SecurityGroups {
					loadSecGroup := createLoadBalancerSecurityGroup(db, load.Id, awsLoadSecGroup)
					if loadSecGroup != nil {
						fmt.Printf("    	LoadBalancerSecurityGroup %d loaded\n", loadSecGroup.LoadBalancerId)
					} else {
						fmt.Printf("    	ERROR LoadBalancerSecurityGroup %d not loaded\n", awsLoadSecGroup)
					}
				}

				for _, awsLoadZone := range awsLoadBalancer.AvailabilityZones {
					loadZone := createLoadBalancerAvailabilityZone(db, load.Id, awsLoadZone.ZoneName, awsLoadZone.SubnetId)
					if loadZone != nil {
						fmt.Printf("    	LoadBalancerAvailabilityZone %s loaded\n", loadZone.ZoneName)
						for _, awsLoadAddress := range awsLoadZone.LoadBalancerAddresses {
							loadAddress := createLoadBalancerAddress(db, loadZone.Id, awsLoadAddress.IpAddress, awsLoadAddress.AllocationId)
							if loadAddress != nil {
								fmt.Printf("    		LoadBalancerAvailabilityZoneAddress %s loaded\n", loadAddress.IpAddress)
							} else {
								fmt.Printf("    		ERROR LoadBalancerAvailabilityZone %s not loaded\n", awsLoadAddress.IpAddress)
							}
						}

					} else {
						fmt.Printf("    	ERROR LoadBalancerAvailabilityZone %s not loaded\n", awsLoadZone.ZoneName)
					}
				}

				err := loadLoadBalancerListeners(db, load)
				if err != nil {
					fmt.Println(err)
				}

			} else {
				fmt.Printf("  ERROR  LoadBalancer %s::%s not loaded\n", awsLoadBalancer.DNSName)
			}
		}
	} else {
		return errUnmarshal
	}
	return nil
}

func loadLoadBalancerListeners(db *sql.DB, loadBalancer *LoadBalancer) error {
	jsonString, err := execLoadBalancerListenerToJSON(loadBalancer.LoadBalancerArn)
	if err == nil {
		err = awsLoadBalancerListenersParseAndPersist(db, jsonString, loadBalancer.Id)
		if err != nil {
			return err
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
		return err
	}
	return nil
}

func execLoadBalancerListenerToJSON(arn string) (string, error) {
	cmd := exec.Command("/bin/bash", "command-load-balancers-listeners.sh", arn)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	_, err = cmd.StderrPipe()
	if err != nil {
		return "", err
	}
	if err := cmd.Start(); err != nil {
		return "", err
	}

	bufOut, _ := ioutil.ReadAll(stdout)
	//bufErr, _ := ioutil.ReadAll(stderr)

	if err := cmd.Wait(); err != nil {
		return "", err
	}
	jsonReturn := string(bufOut)

	return jsonReturn, nil
}
