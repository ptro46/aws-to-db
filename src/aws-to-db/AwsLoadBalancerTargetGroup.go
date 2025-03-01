package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
)

type AwsMatcher struct {
	HttpCode string `json:"HttpCode,omitempty"`
}

type AwsLoadBalancerTargetGroup struct {
	TargetGroupArn             string     `json:"TargetGroupArn,omitempty"`
	TargetGroupName            string     `json:"TargetGroupName,omitempty"`
	Protocol                   string     `json:"Protocol,omitempty"`
	Port                       int        `json:"Port,omitempty"`
	VpcId                      string     `json:"VpcId,omitempty"`
	HealthCheckProtocol        string     `json:"HealthCheckProtocol,omitempty"`
	HealthCheckPort            string     `json:"HealthCheckPort,omitempty"`
	HealthCheckEnabled         bool       `json:"HealthCheckEnabled,omitempty"`
	HealthCheckIntervalSeconds int        `json:"HealthCheckIntervalSeconds,omitempty"`
	HealthCheckTimeoutSeconds  int        `json:"HealthCheckTimeoutSeconds,omitempty"`
	HealthyThresholdCount      int        `json:"HealthyThresholdCount,omitempty"`
	UnhealthyThresholdCount    int        `json:"UnhealthyThresholdCount,omitempty"`
	HealthCheckPath            string     `json:"HealthCheckPath,omitempty"`
	Matcher                    AwsMatcher `json:"Matcher,omitempty"`
	LoadBalancerArns           []string   `json:"LoadBalancerArns,omitempty"`
	TargetType                 string     `json:"TargetType,omitempty"`
}

type AwsLoadBalancerTargetGroups struct {
	TargetGroups []AwsLoadBalancerTargetGroup `json:"TargetGroups,omitempty"`
}

func awsLoadBalancerTargetGroupParseAndPersist(db *sql.DB, jsonString string) error {
	fmt.Printf("awsLoadBalancerTargetGroupParseAndPersist with json length %d\n", len(jsonString))
	awsLoadBalancerTargetGroups := new(AwsLoadBalancerTargetGroups)
	errUnmarshal := json.Unmarshal([]byte(jsonString), awsLoadBalancerTargetGroups)
	if errUnmarshal == nil {
		for _, awsTargetGroup := range awsLoadBalancerTargetGroups.TargetGroups {
			fmt.Printf("Name(%s) \n", awsTargetGroup.TargetGroupName)
			targetGroup := createLoadBalancerTargetGroup(db, awsTargetGroup.TargetGroupArn, awsTargetGroup.TargetGroupName, awsTargetGroup.Protocol, awsTargetGroup.Port, awsTargetGroup.VpcId, awsTargetGroup.HealthCheckProtocol, awsTargetGroup.HealthCheckPort, awsTargetGroup.HealthCheckEnabled, awsTargetGroup.HealthCheckIntervalSeconds, awsTargetGroup.HealthCheckTimeoutSeconds, awsTargetGroup.HealthyThresholdCount, awsTargetGroup.UnhealthyThresholdCount, awsTargetGroup.HealthCheckPath, awsTargetGroup.Matcher.HttpCode, awsTargetGroup.TargetType)
			if targetGroup != nil {
				fmt.Printf("    LoadBalancerTargetGroup (%s) loaded\n", targetGroup.TargetGroupName)
				for _, awsARN := range awsTargetGroup.LoadBalancerArns {
					arn := createLoadBalancerTargetGroupLoadBalancerARN(db, targetGroup.Id, awsARN)
					if arn != nil {
						fmt.Printf("    	LoadBalancerARN(%s) loaded\n", arn.LoadBalancerArn)
					} else {
						fmt.Printf("  ERROR  	LoadBalancerARN (%s) not loaded\n", awsARN)
					}
				}
				//Time to make the command health
				err := loadTargetHealth(db, targetGroup)
				if err != nil {
					fmt.Println(err)
				}

			} else {
				fmt.Printf("  ERROR  LoadBalancerTargetGroup (%s) not loaded\n", awsTargetGroup.TargetGroupName)
			}
		}
	} else {
		fmt.Println(errUnmarshal)
		return errUnmarshal
	}
	return nil
}

func loadTargetHealth(db *sql.DB, targetGroup *LoadBalancerTargetGroup) error {
	jsonString, err := execTargetHealthToJSON(targetGroup.TargetGroupArn)
	if err == nil {
		err = awsTargeHealthDescription(db, jsonString, targetGroup.Id)
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

func execTargetHealthToJSON(arn string) (string, error) {
	cmd := exec.Command("/bin/bash", "command-load-balancers-target-health.sh", arn)
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
