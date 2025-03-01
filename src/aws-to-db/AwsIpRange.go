package main

import "fmt"

type AwsIpRange struct {
	CidrIp      string `json:"CidrIp,omitempty"`
	Description string `json:"Description,omitempty"`
}

func (d *AwsIpRange) String() string {
	return fmt.Sprintf("AwsIpRange cidrIp(%s) description(%s)", d.CidrIp, d.Description)
}
