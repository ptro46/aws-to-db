package main

import (
	"fmt"
)

type Snapshot struct {
	Id          int64
	SnapshotId  string
	OwnerId     string
	VolumeId    string
	Encrypted   bool
	Description string
	State       string
	VolumeSize  int16
	StartTime   string
	Progress    string
}

func NewSnapshot(id int64, snapshotId string, ownerId string, volumeId string, encrypted bool, description string, state string, volumeSize int16, startTime string, progress string) *Snapshot {
	return &Snapshot{Id: id,
		SnapshotId:  snapshotId,
		OwnerId:     ownerId,
		VolumeId:    volumeId,
		Encrypted:   encrypted,
		Description: description,
		State:       state,
		VolumeSize:  volumeSize,
		StartTime:   startTime,
		Progress:    progress}
}

func (d *Snapshot) String() string {
	return fmt.Sprintf("Snapshot snapshotId(%s) volumeId(%s)", d.SnapshotId, d.VolumeId)
}

func (v *Snapshot) Dump() string {
	// Effectue le dump du Snapshot

	snapshotDump := fmt.Sprintf("Snapshot SnapshotId(%s) Encrypted(%b) VolumeId(%s) State(%s) StartTime(%s) VolumeSize(%d)",
		v.SnapshotId,
		v.Encrypted,
		v.VolumeId,
		v.State,
		v.StartTime,
		v.VolumeSize)

	return snapshotDump
}
