package main

import (
	"database/sql"
	"fmt"
)

type Volume struct {
	Id               int64
	AvailabilityZone string
	Attachments      []*VolumeAttachment
	Tags             []*VolumeTag
	Encrypted        bool
	VolumeType       string
	VolumeId         string
	State            string
	Iops             int16
	SnapShotId       string
	CreateTime       string
	Size             int16
}

func NewVolume(id int64, availabilityZone string, encrypted bool, volumeType string,
	volumeId string, state string, iops int16, snapShotId string, createTime string, size int16) *Volume {
	return &Volume{Id: id,
		AvailabilityZone: availabilityZone,
		Encrypted:        encrypted,
		VolumeType:       volumeType,
		VolumeId:         volumeId,
		State:            state,
		Iops:             iops,
		SnapShotId:       snapShotId,
		CreateTime:       createTime,
		Size:             size}
}

func (d *Volume) String() string {
	return fmt.Sprintf("Volume availabilityZone(%s) volumeId(%s)", d.AvailabilityZone, d.VolumeId)
}

func (v *Volume) loadDependencies(db *sql.DB) {
	// Ici on va charger les Attachments et les Tags pour alimenter les tableaux
	arrayOfVolumeAttachment, err := loadAllVolumeAttachmentByOrderUnique(db)
	if err == nil {
		for _, attachment := range arrayOfVolumeAttachment {
			if attachment.VolumeId == v.Id {
				v.Attachments = append(v.Attachments, attachment)
			}
		}
	} else {
		fmt.Println(err)
	}

	arrayOfVolumeTag, err := loadAllVolumeTagByOrderUnique(db)
	if err == nil {
		for _, tag := range arrayOfVolumeTag {
			if tag.VolumeId == v.Id {
				v.Tags = append(v.Tags, tag)
			}
		}
	} else {
		fmt.Println(err)
	}
}

func (v *Volume) Dump() string {
	// Effectue le dump du Volume
	// Boucle sur le dump des attachments
	// Boucle sur le dump des Tags

	volumeDump := fmt.Sprintf("Volume AvailabilityZone(%s) Encrypted(%b) VolumeType(%s) VolumeId(%s) State(%s) Iops(%d) SnapShotId(%s) CreateTime(%s) Size(%d)",
		v.AvailabilityZone,
		v.Encrypted,
		v.VolumeType,
		v.VolumeId,
		v.State,
		v.Iops,
		v.SnapShotId,
		v.CreateTime,
		v.Size)

	volumeDump += "\n	 Volume Attachments ["
	for _, attachment := range v.Attachments {
		volumeDump += "{" + attachment.Dump() + "}"
	}

	volumeDump += "]\n	 Volume Tags ["
	for _, tag := range v.Tags {
		volumeDump += fmt.Sprintf("{Key(%s) Value(%s)}", tag.Key, tag.Value)
	}
	volumeDump += "]"
	return volumeDump
}
