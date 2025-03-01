package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type AwsVpc struct {
	VpcId                   string                       `json:"VpcId,omitempty"`
	InstanceTenancy         string                       `json:"InstanceTenancy,omitempty"`
	Tags                    []AwsTag                     `json:"Tags,omitempty"`
	CidrBlockAssociationSet []AwsCidrBlockAssociationSet `json:"CidrBlockAssociationSet,omitempty"`
	State                   string                       `json:"State,omitempty"`
	DhcpOptionsId           string                       `json:"DhcpOptionsId,omitempty"`
	OwnerId                 string                       `json:"OwnerId,omitempty"`
	CidrBlock               string                       `json:"CidrBlock,omitempty"`
	IsDefault               bool                         `json:"IsDefault,omitempty"`
}

type AwsVpcs struct {
	Vpcs []AwsVpc `json:"Vpcs,omitempty"`
}

func (d AwsVpc) String() string {
	return fmt.Sprintf("AwsVpc VpcId(%s) CidrBlock(%s)", d.VpcId, d.CidrBlock)
}

// func createVPC(db *sql.DB, VpcId string, VpcName string, CidrAssociationId string, CidrBlock string, CidrBlockState string, VpcState string, OwnerId string) *VPC {

func awsVpcParseAndPersist(db *sql.DB, jsonString string) error {
	fmt.Printf("awsVpcParseAndPersist with json lenght %d\n", len(jsonString))
	awsVpcs := new(AwsVpcs)
	errUnmarshal := json.Unmarshal([]byte(jsonString), awsVpcs)
	if errUnmarshal == nil {
		for _, awsVpc := range awsVpcs.Vpcs {
			fmt.Printf("VpcId(%s) CidrBlock(%s)\n", awsVpc.VpcId, awsVpc.CidrBlock)
			vpcName := findVpcName(&awsVpc)
			vpc := createVPC(db,
				awsVpc.VpcId,
				vpcName,
				awsVpc.CidrBlockAssociationSet[0].AssociationId,
				awsVpc.CidrBlockAssociationSet[0].CidrBlock,
				awsVpc.CidrBlockAssociationSet[0].CidrBlockState.State,
				awsVpc.State,
				awsVpc.OwnerId)
			if vpc != nil {
				fmt.Printf("    VPC %s::%s loaded\n", awsVpc.VpcId, vpcName)
			} else {
				fmt.Printf("  ERROR  VPC %s::%s not loaded\n", awsVpc.VpcId, vpcName)
			}
		}
	} else {
		return errUnmarshal
	}
	return nil
}

func findVpcName(vpc *AwsVpc) string {
	for _, tag := range vpc.Tags {
		if tag.Key == "Name" {
			return tag.Value
		}
	}
	return ""
}
