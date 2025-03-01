package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type AwsAmi struct {
	VirtualizationType  string                   `json:"VirtualizationType,omitempty"`
	Description         string                   `json:"Description,omitempty"`
	Tags                []AwsTag                 `json:"Tags,omitempty"`
	Hypervisor          string                   `json:"Hypervisor,omitempty"`
	EnaSupport          bool                     `json:"EnaSupport,omitempty"`
	SriovNetSupport     string                   `json:"SriovNetSupport,omitempty"`
	ImageId             string                   `json:"ImageId,omitempty"`
	State               string                   `json:"State,omitempty"`
	BlockDeviceMappings []AwsBlockDeviceMappings `json:"BlockDeviceMappings,omitempty"`
	Architecture        string                   `json:"Architecture,omitempty"`
	ImageLocation       string                   `json:"ImageLocation,omitempty"`
	RootDeviceType      string                   `json:"RootDeviceType,omitempty"`
	OwnerId             string                   `json:"OwnerId,omitempty"`
	RootDeviceName      string                   `json:"RootDeviceName,omitempty"`
	CreationDate        string                   `json:"CreationDate,omitempty"`
	Public              bool                     `json:"Public,omitempty"`
	ImageType           string                   `json:"ImageType,omitempty"`
	Name                string                   `json:"Name,omitempty"`
}

type AwsAmis struct {
	Amis []AwsAmi `json:"Images,omitempty"`
}

func (d AwsAmi) String() string {
	return fmt.Sprintf("AwsImage Description(%s) ImageId(%s) OwnerId(%s) Name(%s)", d.Description, d.ImageId, d.OwnerId, d.Name)
}

func awsAmiParseAndPersist(db *sql.DB, jsonString string) error {
	fmt.Printf("awsAmiParseAndPersist with json length %d\n", len(jsonString))
	AwsAmis := new(AwsAmis)
	errUnmarshal := json.Unmarshal([]byte(jsonString), AwsAmis)
	if errUnmarshal == nil {
		for _, awsAmi := range AwsAmis.Amis {
			fmt.Printf("Description(%s)\n", awsAmi.Description)
			ami := createAmi(db,
				awsAmi.VirtualizationType,
				awsAmi.Description,
				awsAmi.Hypervisor,
				awsAmi.EnaSupport,
				awsAmi.ImageId,
				awsAmi.State,
				awsAmi.BlockDeviceMappings[0].DeviceName,
				awsAmi.BlockDeviceMappings[0].Ebs.SnapShotId,
				awsAmi.BlockDeviceMappings[0].Ebs.DeleteOnTermination,
				awsAmi.BlockDeviceMappings[0].Ebs.VolumeType,
				awsAmi.BlockDeviceMappings[0].Ebs.VolumeSize,
				awsAmi.BlockDeviceMappings[0].Ebs.Encrypted,
				awsAmi.Architecture,
				awsAmi.ImageLocation,
				awsAmi.RootDeviceType,
				awsAmi.OwnerId,
				awsAmi.RootDeviceName,
				awsAmi.CreationDate,
				awsAmi.Public,
				awsAmi.ImageType,
				awsAmi.Name)

			if ami != nil {
				fmt.Printf("    Image %s loaded\n", awsAmi.Description)
			} else {
				fmt.Printf("  ERROR  Image %s not loaded (%+v)\n", awsAmi.Description)
			}
		}
	} else {
		return errUnmarshal
	}
	return nil
}
