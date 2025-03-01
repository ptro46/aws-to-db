package main

import "fmt"

type AwsEbs struct {
	SnapShotId          string `json:"SnapShotId,omitempty"`
	DeleteOnTermination bool   `json:"DeleteOnTermination,omitempty"`
	VolumeType          string `json:"VolumeType,omitempty"`
	VolumeSize          int16  `json:"VolumeSize,omitempty"`
	Encrypted           bool   `json:"Encrypted,omitempty"`
}

func (d AwsEbs) String() string {
	return fmt.Sprintf("AwsEbs snapShotId(%s) deleteOnTermination(%b) volumeType(%s) volumeSize(%d) encrypted(%b)", d.SnapShotId, d.DeleteOnTermination, d.VolumeType, d.VolumeSize, d.Encrypted)
}
