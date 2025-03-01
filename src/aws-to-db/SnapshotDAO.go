package main

import (
	"database/sql" // package SQL
	"fmt"
	"log"

	_ "github.com/lib/pq" // driver Postgres
)

func rowResultSetToSnapshot(row *sql.Row) (*Snapshot, error) {
	var err error
	var Id int64
	var SnapshotId string
	var OwnerId string
	var VolumeId string
	var Encrypted bool
	var Description string
	var State string
	var VolumeSize int16
	var StartTime string
	var Progress string

	err = row.Scan(&Id, &SnapshotId, &OwnerId, &VolumeId, &Encrypted, &Description, &State, &VolumeSize, &StartTime, &Progress)
	if err != nil {
		return nil, err
	}
	return NewSnapshot(Id, SnapshotId, OwnerId, VolumeId, Encrypted, Description, State, VolumeSize, StartTime, Progress), nil
}

func rowsNoFetchResultSetToSnapshot(rows *sql.Rows) (*Snapshot, error) {
	var err error
	var Id int64
	var SnapshotId string
	var OwnerId string
	var VolumeId string
	var Encrypted bool
	var Description string
	var State string
	var VolumeSize int16
	var StartTime string
	var Progress string

	err = rows.Scan(&Id, &SnapshotId, &OwnerId, &VolumeId, &Encrypted, &Description, &State, &VolumeSize, &StartTime, &Progress)
	if err != nil {
		return nil, err
	}
	return NewSnapshot(Id, SnapshotId, OwnerId, VolumeId, Encrypted, Description, State, VolumeSize, StartTime, Progress), nil
}

func rowsResultSetToSnapshot(rows *sql.Rows) (*Snapshot, error) {
	var err error
	if rows.Next() {
		var Id int64
		var SnapshotId string
		var OwnerId string
		var VolumeId string
		var Encrypted bool
		var Description string
		var State string
		var VolumeSize int16
		var StartTime string
		var Progress string

		err = rows.Scan(&Id, &SnapshotId, &OwnerId, &VolumeId, &Encrypted, &Description, &State, &VolumeSize, &StartTime, &Progress)
		if err != nil {
			return nil, err
		}
		return NewSnapshot(Id, SnapshotId, OwnerId, VolumeId, Encrypted, Description, State, VolumeSize, StartTime, Progress), nil
	}
	return nil, err
}

func createSnapshot(db *sql.DB, snapshotId string, ownerId string, volumeId string, encrypted bool, description string, state string, volumeSize int16, startTime string, progress string) *Snapshot {
	rows := db.QueryRow("insert into snapshot(snapshot_id,owner_id,volume_id,encrypted,description,state,volume_size,start_time,progress) values($1,$2,$3,$4,$5,$6,$7,$8,$9) returning id,snapshot_id,owner_id,volume_id,encrypted,description,state,volume_size,start_time,progress",
		snapshotId, ownerId, volumeId, encrypted, description, state, volumeSize, startTime, progress)

	snapshot, err := rowResultSetToSnapshot(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return snapshot
}

func loadAllSnapshot(db *sql.DB) ([]*Snapshot, error) {
	rows, err := db.Query("select id,snapshot_id,owner_id,volume_id,encrypted,description,state,volume_size,start_time,progress from snapshot order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfSnapshot := make([]*Snapshot, 0, 0)
	for rows.Next() {
		snapshot, err := rowsNoFetchResultSetToSnapshot(rows)
		if err == nil {
			arrayOfSnapshot = append(arrayOfSnapshot, snapshot)
		} else {
			log.Println(err)
		}
	}

	return arrayOfSnapshot, nil
}

func loadAllSnapshotByOrderUnique(db *sql.DB) ([]*Snapshot, error) {
	rows, err := db.Query("select id,snapshot_id,owner_id,volume_id,encrypted,description,state,volume_size,start_time,progress from snapshot order by snapshot_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfSnapshot := make([]*Snapshot, 0, 0)
	for rows.Next() {
		snapshot, err := rowsNoFetchResultSetToSnapshot(rows)
		if err == nil {
			arrayOfSnapshot = append(arrayOfSnapshot, snapshot)
		} else {
			log.Println(err)
		}
	}

	return arrayOfSnapshot, nil
}

func loadSnapshotbyAwsReferenceId(db *sql.DB, snapshotId string) (*Snapshot, error) {
	rows, err := db.Query("select id,snapshot_id,owner_id,volume_id,encrypted,description,state,volume_size,start_time,progress from snapshot where snapshot_id=$1", snapshotId)
	if err != nil {
		fmt.Println(err, "SDSFEFEF")

		return nil, err
	}
	snapshot, err := rowsResultSetToSnapshot(rows)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return snapshot, nil

}

func updateSnapshotVolumeId(db *sql.DB, snapshot *Snapshot, volumeId int64) *Snapshot {
	rows := db.QueryRow("update snapshot set ref_volume_id=$1 where id=$2 returning id,snapshot_id,owner_id,volume_id,encrypted,description,state,volume_size,start_time,progress", volumeId, snapshot.Id)

	snapshot, err := rowResultSetToSnapshot(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return snapshot
}

func deleteAllSnapshot(db *sql.DB) error {
	_, err := db.Exec("delete from snapshot")

	return err
}
