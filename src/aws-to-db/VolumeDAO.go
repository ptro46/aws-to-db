package main

import (
	"database/sql" // package SQL
	"fmt"
	"log"

	_ "github.com/lib/pq" // driver Postgres
)

func rowResultSetToVolume(row *sql.Row) (*Volume, error) {
	var err error
	var Id int64
	var AvailabilityZone string
	var Encrypted bool
	var VolumeType string
	var VolumeId string
	var State string
	var Iops int16
	var SnapShotId string
	var CreateTime string
	var Size int16

	err = row.Scan(&Id, &AvailabilityZone, &Encrypted, &VolumeType, &VolumeId, &State, &Iops, &SnapShotId, &CreateTime, &Size)
	if err != nil {
		return nil, err
	}
	return NewVolume(Id, AvailabilityZone, Encrypted, VolumeType, VolumeId, State, Iops, SnapShotId, CreateTime, Size), nil
}

func rowsNoFetchResultSetToVolume(rows *sql.Rows) (*Volume, error) {
	var err error
	var Id int64
	var AvailabilityZone string
	var Encrypted bool
	var VolumeType string
	var VolumeId string
	var State string
	var Iops int16
	var SnapShotId string
	var CreateTime string
	var Size int16

	err = rows.Scan(&Id, &AvailabilityZone, &Encrypted, &VolumeType, &VolumeId, &State, &Iops, &SnapShotId, &CreateTime, &Size)
	if err != nil {
		return nil, err
	}
	return NewVolume(Id, AvailabilityZone, Encrypted, VolumeType, VolumeId, State, Iops, SnapShotId, CreateTime, Size), nil
}

func rowsResultSetToVolume(rows *sql.Rows) (*Volume, error) {
	var err error
	if rows.Next() {
		var Id int64
		var AvailabilityZone string
		var Encrypted bool
		var VolumeType string
		var VolumeId string
		var State string
		var Iops int16
		var SnapShotId string
		var CreateTime string
		var Size int16

		err = rows.Scan(&Id, &AvailabilityZone, &Encrypted, &VolumeType, &VolumeId, &State, &Iops, &SnapShotId, &CreateTime, &Size)
		if err != nil {
			return nil, err
		}
		return NewVolume(Id, AvailabilityZone, Encrypted, VolumeType, VolumeId, State, Iops, SnapShotId, CreateTime, Size), nil
	}
	return nil, err
}

func createVolume(db *sql.DB, availabilityZone string, encrypted bool, volumeType string,
	volumeId string, state string, iops int16, snapShotId string, createTime string, size int16) *Volume {
	rows := db.QueryRow("insert into volume(availability_zone,encrypted,volume_type,volume_id,state,iops,snap_shot_id,create_time,size) values($1,$2,$3,$4,$5,$6,$7,$8,$9) returning id,availability_zone,encrypted,volume_type,volume_id,state,iops,snap_shot_id,create_time,size",
		availabilityZone, encrypted, volumeType,
		volumeId, state, iops, snapShotId, createTime, size)

	volume, err := rowResultSetToVolume(rows)
	if err != nil {
		return nil
	}
	return volume
}

func loadAllVolume(db *sql.DB) ([]*Volume, error) {
	rows, err := db.Query("select id,availability_zone,encrypted,volume_type,volume_id,state,iops,snap_shot_id,create_time,size from volume order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfVolume := make([]*Volume, 0, 0)
	for rows.Next() {
		volume, err := rowsNoFetchResultSetToVolume(rows)
		if err == nil {
			arrayOfVolume = append(arrayOfVolume, volume)
		} else {
			log.Println(err)
		}
	}

	return arrayOfVolume, nil
}

func loadAllVolumeByOrderUnique(db *sql.DB) ([]*Volume, error) {
	rows, err := db.Query("select id,availability_zone,encrypted,volume_type,volume_id,state,iops,snap_shot_id,create_time,size from volume order by volume_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfVolume := make([]*Volume, 0, 0)
	for rows.Next() {
		volume, err := rowsNoFetchResultSetToVolume(rows)
		if err == nil {
			arrayOfVolume = append(arrayOfVolume, volume)
		} else {
			log.Println(err)
		}
	}

	return arrayOfVolume, nil
}

func loadVolumebyAwsReferenceId(db *sql.DB, volumeId string) (*Volume, error) {
	rows, err := db.Query("select id,availability_zone,encrypted,volume_type,volume_id,state,iops,snap_shot_id,create_time,size from volume where volume_id=$1", volumeId)
	if err != nil {
		fmt.Println(err, "SDSFEFEF")

		return nil, err
	}
	volume, err := rowsResultSetToVolume(rows)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return volume, nil

}

func deleteAllVolume(db *sql.DB) error {
	_, err := db.Exec("delete from volume")

	return err
}
