package main

import "fmt"

type Awsattachment struct {
	AttachTime          string `json:"AttachTime,omitempty"`
	InstanceId          string `json:"InstanceId,omitempty"`
	VolumeId            string `json:"VolumeId,omitempty"`
	State               string `json:"State,omitempty"`
	DeleteOnTermination bool   `json:"DeleteOnTermination,omitempty"`
	Device              string `json:"Device,omitempty"`
}

func (d *Awsattachment) String() string {
	return fmt.Sprintf("AwsAttachment volumeId(%s) instanceId(%s)", d.VolumeId, d.InstanceId)
}
