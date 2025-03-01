package main

import "fmt"

type AwsBlockDeviceMappings struct {
	DeviceName string `json:"DeviceName,omitempty"`
	Ebs        AwsEbs `json:"Ebs,omitempty"`
}

func (d AwsBlockDeviceMappings) String() string {
	return fmt.Sprintf("AwsBlockDeviceMappings deviceName(%s)", d.DeviceName)
}
