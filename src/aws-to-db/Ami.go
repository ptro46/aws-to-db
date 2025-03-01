package main

import "fmt"

type Ami struct {
	Id                  int64
	VirtualizationType  string
	Description         string
	Hypervisor          string
	EnaSupport          bool
	ImageId             string
	State               string
	DeviceName          string
	SnapShotId          string
	DeleteOnTermination bool
	VolumeType          string
	VolumeSize          int16
	Encrypted           bool
	Architecture        string
	ImageLocation       string
	RootDeviceType      string
	OwnerId             string
	RootDeviceName      string
	CreationDate        string
	Public              bool
	ImageType           string
	Name                string
}

func NewAmi(id int64, virtualizationType string, description string, hypervisor string, enaSupport bool,
	imageId string, state string, deviceName string, snapShotId string, deleteOnTermination bool,
	volumeType string, volumeSize int16, encrypted bool, architecture string, imageLocation string,
	rootDeviceType string, ownerId string, rootDeviceName string, creationDate string, public bool,
	imageType string, name string) *Ami {
	return &Ami{
		Id:                  id,
		VirtualizationType:  virtualizationType,
		Description:         description,
		Hypervisor:          hypervisor,
		EnaSupport:          enaSupport,
		ImageId:             imageId,
		State:               state,
		DeviceName:          deviceName,
		SnapShotId:          snapShotId,
		DeleteOnTermination: deleteOnTermination,
		VolumeType:          volumeType,
		VolumeSize:          volumeSize,
		Encrypted:           encrypted,
		Architecture:        architecture,
		ImageLocation:       imageLocation,
		RootDeviceType:      rootDeviceType,
		OwnerId:             ownerId,
		RootDeviceName:      rootDeviceName,
		CreationDate:        creationDate,
		Public:              public,
		ImageType:           imageType,
		Name:                name,
	}
}

func (d *Ami) String() string {
	return fmt.Sprintf("Image id(%d) imageId(%d) state(%s) deviceName(%s) name(%s)", d.Id, d.ImageId, d.State, d.DeviceName, d.Name)
}

func (d *Ami) Dump() string {
	return fmt.Sprintf("Image VirtualizationType(%s)  Description(%s) Hypervisor(%s) EnaSupport(%b) ImageId(%s) State(%s) DeviceName(%s) SnapShotId(%s) DeleteOnTermination(%b) VolumeType(%s) VolumeSize(%d) Encrypted(%b) Architecture(%s) ImageLocation(%s) RootDeviceType(%s) OwnerId(%s) RootDeviceName(%s) CreationDate(%s) Public(%b) ImageType(%s) Name(%s)",
		d.VirtualizationType,
		d.Description,
		d.Hypervisor,
		d.EnaSupport,
		d.ImageId,
		d.State,
		d.DeviceName,
		d.SnapShotId,
		d.DeleteOnTermination,
		d.VolumeType,
		d.VolumeSize,
		d.Encrypted,
		d.Architecture,
		d.ImageLocation,
		d.RootDeviceType,
		d.OwnerId,
		d.RootDeviceName,
		d.CreationDate,
		d.Public,
		d.ImageType,
		d.Name)
}
