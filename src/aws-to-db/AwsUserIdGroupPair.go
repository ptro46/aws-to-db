package main

import "fmt"

type AwsUserIdGroupPair struct {
	Description            string `json:"Description,omitempty"`
	GroupId                string `json:"GroupId,omitempty"`
	GroupName              string `json:"GroupName,omitempty"`
	PeeringStatus          string `json:"PeeringStatus,omitempty"`
	UserId                 string `json:"UserId,omitempty"`
	VpcId                  string `json:"VpcId,omitempty"`
	VpcPeeringConnectionId string `json:"VpcPeeringConnectionId,omitempty"`
}

func (d *AwsUserIdGroupPair) String() string {
	return fmt.Sprintf("AwsUserIdGroupPair description(%s) groupName(%s)", d.Description, d.GroupName)
}
