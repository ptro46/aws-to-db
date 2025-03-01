package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type AwsSnapshot struct {
	SnapshotId  string `json:"SnapshotId,omitempty"`
	OwnerId     string `json:"OwnerId,omitempty"`
	VolumeId    string `json:"VolumeId,omitempty"`
	Encrypted   bool   `json:"Encrypted,omitempty"`
	Description string `json:"Description,omitempty"`
	State       string `json:"State,omitempty"`
	VolumeSize  int16  `json:"VolumeSize,omitempty"`
	StartTime   string `json:"StartTime,omitempty"`
	Progress    string `json:"Progress,omitempty"`
}

type AwsSnapshots struct {
	Snapshots []AwsSnapshot `json:"Snapshots,omitempty"`
}

func (d AwsSnapshot) String() string {
	return fmt.Sprintf("AwsSnapshot volumeId(%s) snapShotId(%s)", d.VolumeId, d.SnapshotId)
}

func awsSnapshotParseAndPersist(db *sql.DB, jsonString string) error {
	fmt.Printf("awsSnapshotParseAndPersist with json length %d\n", len(jsonString))
	awsSnapshots := new(AwsSnapshots)
	errUnmarshal := json.Unmarshal([]byte(jsonString), awsSnapshots)
	if errUnmarshal == nil {
		for _, awsSnapshot := range awsSnapshots.Snapshots {
			fmt.Printf("AwsSnapshot snapshotId(%s)\n", awsSnapshot.SnapshotId)
			snapshot := createSnapshot(db,
				awsSnapshot.SnapshotId,
				awsSnapshot.OwnerId,
				awsSnapshot.VolumeId,
				awsSnapshot.Encrypted,
				awsSnapshot.Description,
				awsSnapshot.State,
				awsSnapshot.VolumeSize,
				awsSnapshot.StartTime,
				awsSnapshot.Progress)

			if snapshot != nil {
				fmt.Printf("    Snapshot %s loaded\n", awsSnapshot.SnapshotId)
			} else {
				fmt.Printf("  ERROR  Snapshot %s not loaded\n", awsSnapshot.SnapshotId)
			}
		}
	} else {
		return errUnmarshal
	}
	return nil
}
