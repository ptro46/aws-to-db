package main

import (
	"database/sql"
	"fmt"
)

type VolumeAttachment struct {
	Id                  int64
	VolumeId            int64
	AttachTime          string
	InstanceId          string
	RefInstanceId       sql.NullInt64
	Volume_Id           string
	State               string
	DeleteOnTermination bool
	Device              string
}

func NewVolumeAttachment(id int64, volumeId int64, attachTime string, instanceId string, refInstanceId sql.NullInt64, state string, deleteOnTermination bool, device string) *VolumeAttachment {
	return &VolumeAttachment{
		Id:                  id,
		VolumeId:            volumeId,
		AttachTime:          attachTime,
		InstanceId:          instanceId,
		RefInstanceId:       refInstanceId,
		State:               state,
		DeleteOnTermination: deleteOnTermination,
		Device:              device}
}

func (d *VolumeAttachment) String() string {
	return fmt.Sprintf("VolumeAttachment volumeId(%d) state(%s)", d.VolumeId, d.State)
}

func (d *VolumeAttachment) Dump() string {
	return fmt.Sprintf("AttachTime(%s) InstanceId(%s) RefInstanceId(%d) Volume_Id(%s) State(%s) DeleteOnTermination(%b) Device(%s)",
		d.AttachTime,
		d.InstanceId,
		d.RefInstanceId,
		d.Volume_Id,
		d.State,
		d.DeleteOnTermination,
		d.Device)
}
