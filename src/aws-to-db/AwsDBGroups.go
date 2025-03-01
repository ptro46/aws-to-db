package main

type AwsDBSecurityGroup struct {
	DBSecurityGroupName string `json:"DBSecurityGroupName,omitempty"`
	Status              string `json:"Status,omitempty"`
}

type AwsVpcSecurityGroup struct {
	VpcSecurityGroupId string `json:"VpcSecurityGroupId,omitempty"`
	Status             string `json:"Status,omitempty"`
}

type AwsDBParameterGroup struct {
	DBParameterGroupName string `json:"DBParameterGroupName,omitempty"`
	ParameterApplyStatus string `json:"ParameterApplyStatus,omitempty"`
}

type AwsSubnetAvailabilityZone struct {
	Name string `json:"Name,omitempty"`
}

type AwsDBSubnet struct {
	SubnetIdentifier       string `json:"SubnetIdentifier,omitempty"`
	SubnetAvailabilityZone AwsSubnetAvailabilityZone `json:"SubnetAvailabilityZone,omitempty"`
	SubnetStatus           string `json:"SubnetStatus,omitempty"`
}

type AwsDBSubnetGroup struct {
	DBSubnetGroupName        string `json:"DBSubnetGroupName,omitempty"`
	DBSubnetGroupDescription string `json:"DBSubnetGroupDescription,omitempty"`
	VpcId                    string `json:"VpcId,omitempty"`
	SubnetGroupStatus        string `json:"SubnetGroupStatus,omitempty"`
	Subnets                  []AwsDBSubnet `json:"Subnets,omitempty"`
	DBSubnetGroupArn         string `json:"DBSubnetGroupArn,omitempty"`
}
