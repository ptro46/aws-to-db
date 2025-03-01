package main

import (
	"database/sql" // package SQL
	"fmt"
	"log"

	_ "github.com/lib/pq" // driver Postgres
)

func rowResultSetToVolumeAttachment(row *sql.Row) (*VolumeAttachment, error) {
	var err error
	var Id int64
	var VolumeId int64
	var AttachTime string
	var InstanceId string
	var RefInstanceId sql.NullInt64
	var State string
	var DeleteOnTermination bool
	var Device string

	err = row.Scan(&Id, &VolumeId, &AttachTime, &InstanceId, &RefInstanceId, &State, &DeleteOnTermination, &Device)
	if err != nil {
		return nil, err
	}
	return NewVolumeAttachment(Id, VolumeId, AttachTime, InstanceId, RefInstanceId, State, DeleteOnTermination, Device), nil
}

func rowsNoFetchResultSetToVolumeAttachment(rows *sql.Rows) (*VolumeAttachment, error) {
	var err error
	var Id int64
	var VolumeId int64
	var AttachTime string
	var InstanceId string
	var RefInstanceId sql.NullInt64
	var State string
	var DeleteOnTermination bool
	var Device string

	err = rows.Scan(&Id, &VolumeId, &AttachTime, &InstanceId, &RefInstanceId, &State, &DeleteOnTermination, &Device)
	if err != nil {
		return nil, err
	}
	return NewVolumeAttachment(Id, VolumeId, AttachTime, InstanceId, RefInstanceId, State, DeleteOnTermination, Device), nil
}

func rowsResultSetToVolumeAttachment(rows *sql.Rows) (*VolumeAttachment, error) {
	var err error
	if rows.Next() {
		var Id int64
		var VolumeId int64
		var AttachTime string
		var InstanceId string
		var RefInstanceId sql.NullInt64
		var State string
		var DeleteOnTermination bool
		var Device string

		err = rows.Scan(&Id, &VolumeId, &AttachTime, &InstanceId, &RefInstanceId, &State, &DeleteOnTermination, &Device)
		if err != nil {
			return nil, err
		}
		return NewVolumeAttachment(Id, VolumeId, AttachTime, InstanceId, RefInstanceId, State, DeleteOnTermination, Device), nil
	}
	return nil, err
}

func createVolumeAttachment(db *sql.DB, VolumeId int64, AttachTime string, InstanceId string, State string,
	DeleteOnTermination bool, Device string) *VolumeAttachment {
	rows := db.QueryRow("insert into volume_attachment(volume_id,attach_time,instance_id,state,delete_on_termination,device) values($1,$2,$3,$4,$5,$6) returning id,volume_id,attach_time,instance_id,ref_instance_id,state,delete_on_termination,device",
		VolumeId, AttachTime, InstanceId, State, DeleteOnTermination, Device)

	volumeAtt, err := rowResultSetToVolumeAttachment(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return volumeAtt
}

func updateVolumeAttachmentId(db *sql.DB, volumeAttachment *VolumeAttachment, instanceID int64) *VolumeAttachment {
	rows := db.QueryRow("update volume_attachment set ref_instance_id=$1 where id=$2 returning id,volume_id,attach_time,instance_id,ref_instance_id,state,delete_on_termination,device", instanceID, volumeAttachment.Id)

	volumeAttachment, err := rowResultSetToVolumeAttachment(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return volumeAttachment
}

func loadAllVolumeAttachment(db *sql.DB) ([]*VolumeAttachment, error) {
	rows, err := db.Query("select id,volume_id,attach_time,instance_id,ref_instance_id,state,delete_on_termination,device from volume_attachment order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfVolumeAtt := make([]*VolumeAttachment, 0, 0)
	for rows.Next() {
		volumeAtt, err := rowsNoFetchResultSetToVolumeAttachment(rows)
		if err == nil {
			arrayOfVolumeAtt = append(arrayOfVolumeAtt, volumeAtt)
		} else {
			log.Println(err)
		}
	}

	return arrayOfVolumeAtt, nil
}

func loadAllVolumeAttachmentByOrderUnique(db *sql.DB) ([]*VolumeAttachment, error) {
	rows, err := db.Query("select id,volume_id,attach_time,instance_id,ref_instance_id,state,delete_on_termination,device from volume_attachment order by instance_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfVolumeAtt := make([]*VolumeAttachment, 0, 0)
	for rows.Next() {
		volumeAtt, err := rowsNoFetchResultSetToVolumeAttachment(rows)
		if err == nil {
			arrayOfVolumeAtt = append(arrayOfVolumeAtt, volumeAtt)
		} else {
			log.Println(err)
		}
	}

	return arrayOfVolumeAtt, nil
}

func deleteAllVolumeAttachment(db *sql.DB) error {
	_, err := db.Exec("delete from volume_attachment")

	return err
}
