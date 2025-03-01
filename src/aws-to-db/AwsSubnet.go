package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type AwsSubnet struct {
	MapPublicIpOnLaunch         bool                         `json:"MapPublicIpOnLaunch,omitempty"`
	AvailabilityZoneId          string                       `json:"AvailabilityZoneId,omitempty"`
	Tags                        []AwsTag                     `json:"Tags,omitempty"`
	AvailableIpAddressCount     int16                        `json:"AvailableIpAddressCount,omitempty"`
	DefaultForAz                bool                         `json:"DefaultForAz,omitempty"`
	SubnetArn                   string                       `json:"SubnetArn,omitempty"`
	Ipv6CidrBlockAssociationSet []AwsCidrBlockAssociationSet `json:"Ipv6CidrBlockAssociationSet,omitempty"`
	VpcId                       string                       `json:"VpcId,omitempty"`
	State                       string                       `json:"State,omitempty"`
	AvailabilityZone            string                       `json:"AvailabilityZone,omitempty"`
	SubnetId                    string                       `json:"SubnetId,omitempty"`
	OwnerId                     string                       `json:"OwnerId,omitempty"`
	CidrBlock                   string                       `json:"CidrBlock,omitempty"`
	AssignIpv6AddressOnCreation bool                         `json:"AssignIpv6AddressOnCreation,omitempty"`
}

type AwsSubnets struct {
	Subnets []AwsSubnet `json:"Subnets,omitempty"`
}

func (d AwsSubnet) String() string {
	return fmt.Sprintf("AwsSubnet VpcId(%s) CidrBlock(%s)", d.VpcId, d.CidrBlock)
}

func awsSubnetParseAndPersist(db *sql.DB, jsonString string) error {
	fmt.Printf("awsSubnetParseAndPersist with json length %d\n", len(jsonString))
	awsSubnets := new(AwsSubnets)
	errUnmarshal := json.Unmarshal([]byte(jsonString), awsSubnets)
	if errUnmarshal == nil {
		for _, awsSubnet := range awsSubnets.Subnets {
			fmt.Printf("VpcId(%s) CidrBlock(%s)\n", awsSubnet.VpcId, awsSubnet.CidrBlock)
			SubentName := findSubnetName(&awsSubnet)
			CloudFormationStackName := findSubnetCloudStackName(&awsSubnet)
			Vpc, err := loadVPCByVPCId(db, awsSubnet.VpcId)
			if Vpc == nil || err != nil {
				fmt.Printf("    Skip VPC %s::%s not loaded %+v\n", awsSubnet.VpcId, SubentName, err)
				continue
			}

			subnet := createSubnet(db,
				Vpc.Id,
				awsSubnet.SubnetId,
				awsSubnet.MapPublicIpOnLaunch,
				awsSubnet.AvailabilityZoneId,
				SubentName,
				CloudFormationStackName,
				awsSubnet.AvailableIpAddressCount,
				awsSubnet.State,
				awsSubnet.AvailabilityZone,
				awsSubnet.OwnerId,
				awsSubnet.CidrBlock)

			if subnet != nil {
				fmt.Printf("    VPC %s::%s loaded\n", awsSubnet.VpcId, SubentName)
			} else {
				fmt.Printf("    Skip VPC %s::%s not loaded\n", awsSubnet.VpcId, SubentName)
			}
		}
	} else {
		return errUnmarshal
	}
	return nil
}

/*func FindAttachedVPCId(subnet *AwsSubnet, db *sql.DB) int64 {
	vpcs, err := loadAllVPC(db)
	if err == nil {
		for _, vpc := range vpcs {
			if vpc.VpcId == subnet.VpcId {
				return vpc.Id
			}
		}
	}
	return 0
}*/

func findSubnetCloudStackName(subnet *AwsSubnet) string {
	for _, tag := range subnet.Tags {
		if tag.Key == "aws:cloudformation:stack-name" {
			return tag.Value
		}
	}
	return ""
}

func findSubnetName(subnet *AwsSubnet) string {
	for _, tag := range subnet.Tags {
		if tag.Key == "Name" {
			return tag.Value
		}
	}
	return ""
}
