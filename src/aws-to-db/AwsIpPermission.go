package main

import "fmt"

type AwsIpPermission struct {
	FromPort         int16                `json:"FromPort,omitempty"`
	IpProtocol       string               `json:"IpProtocol,omitempty"`
	IpRanges         []AwsIpRange         `json:"IpRanges,omitempty"`
	PrefixListIds    []AwsPrefixListId    `json:"PrefixListIds,omitempty"`
	ToPort           int16                `json:"ToPort,omitempty"`
	UserIdGroupPairs []AwsUserIdGroupPair `json:"UserIdGroupPairs,omitempty"`
}

func (d *AwsIpPermission) String() string {
	return fmt.Sprintf("AwsIpPermission fromPort(%d) IpProtocol(%s)", d.FromPort, d.IpProtocol)
}

