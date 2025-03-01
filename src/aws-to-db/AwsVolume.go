package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type AwsVolume struct {
	AvailabilityZone string          `json:"AvailabilityZone,omitempty"`
	Attachments      []Awsattachment `json:"Attachments,omitempty"`
	Tags             []AwsTag        `json:"Tags,omitempty"`
	Encrypted        bool            `json:"Encrypted,omitempty"`
	VolumeType       string          `json:"VolumeType,omitempty"`
	VolumeId         string          `json:"VolumeId,omitempty"`
	State            string          `json:"State,omitempty"`
	Iops             int16           `json:"Iops,omitempty"`
	SnapshotId       string          `json:"SnapshotId,omitempty"`
	CreateTime       string          `json:"CreateTime,omitempty"`
	Size             int16           `json:"Size,omitempty"`
}

type AwsVolumes struct {
	Volumes []AwsVolume `json:"Volumes,omitempty"`
}

func (d AwsVolume) String() string {
	return fmt.Sprintf("AwsVolume volumeId(%s) snapShotId(%s)", d.VolumeId, d.SnapshotId)
}

func awsVolumeParseAndPersist(db *sql.DB, jsonString string) error {
	fmt.Printf("awsVolumeParseAndPersist with json length %d\n", len(jsonString))
	awsVolumes := new(AwsVolumes)
	errUnmarshal := json.Unmarshal([]byte(jsonString), awsVolumes)
	if errUnmarshal == nil {
		for _, awsVolume := range awsVolumes.Volumes {
			fmt.Printf("AwsVolume volumeId(%s)\n", awsVolume.VolumeId)
			volume := createVolume(db,
				awsVolume.AvailabilityZone,
				awsVolume.Encrypted,
				awsVolume.VolumeType,
				awsVolume.VolumeId,
				awsVolume.State,
				awsVolume.Iops,
				awsVolume.SnapshotId,
				awsVolume.CreateTime,
				awsVolume.Size)

			if volume != nil {
				fmt.Printf("    Volume %s loaded\n", awsVolume.VolumeId)
				for _, awsVolumeTag := range awsVolume.Tags {
					tag := createVolumetag(db,
						volume.Id,
						awsVolumeTag.Key,
						awsVolumeTag.Value)

					if tag != nil {
						fmt.Printf("    		Tag %s has loaded\n", tag.Key)
					} else {
						fmt.Printf("     		ERROR Tag %s has not loaded\n", awsVolumeTag.Key)
					}
				}

				for _, awsVolumeAtt := range awsVolume.Attachments {
					attachment := createVolumeAttachment(db,
						volume.Id,
						awsVolumeAtt.AttachTime,
						awsVolumeAtt.InstanceId,
						awsVolumeAtt.State,
						awsVolumeAtt.DeleteOnTermination,
						awsVolumeAtt.Device)
					if attachment != nil {
						fmt.Printf("    	Volume's %s attachment %s has loaded\n", awsVolume.VolumeId, attachment.InstanceId)
					} else {
						fmt.Printf("     	ERROR Volume's %s attachment %s has not loaded\n", awsVolume.VolumeId, awsVolumeAtt.InstanceId)
					}
				}

			} else {
				fmt.Printf("  ERROR  Volume %s not loaded\n", awsVolume.VolumeId)
			}
		}
	} else {
		return errUnmarshal
	}
	return nil
}
