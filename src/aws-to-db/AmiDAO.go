package main

import (
	"database/sql" // package SQL
	"fmt"
	"log"

	_ "github.com/lib/pq" // driver Postgres
)

func rowResultSetToAmi(row *sql.Row) (*Ami, error) {
	var err error
	var Id int64
	var VirtualizationType string
	var Description string
	var Hypervisor string
	var EnaSupport bool
	var ImageId string
	var State string
	var DeviceName string
	var SnapShotId string
	var DeleteOnTermination bool
	var VolumeType string
	var VolumeSize int16
	var Encrypted bool
	var Architecture string
	var ImageLocation string
	var RootDeviceType string
	var OwnerId string
	var RootDeviceName string
	var CreationDate string
	var Public bool
	var ImageType string
	var Name string

	err = row.Scan(&Id, &VirtualizationType, &Description, &Hypervisor, &EnaSupport,
		&ImageId, &State, &DeviceName, &SnapShotId, &DeleteOnTermination,
		&VolumeType, &VolumeSize, &Encrypted, &Architecture, &ImageLocation,
		&RootDeviceType, &OwnerId, &RootDeviceName, &CreationDate, &Public,
		&ImageType, &Name)
	if err != nil {
		return nil, err
	}
	return NewAmi(Id, VirtualizationType, Description, Hypervisor, EnaSupport,
		ImageId, State, DeviceName, SnapShotId, DeleteOnTermination,
		VolumeType, VolumeSize, Encrypted, Architecture, ImageLocation,
		RootDeviceType, OwnerId, RootDeviceName, CreationDate, Public,
		ImageType, Name), nil
}

func rowsNoFetchResultSetToAmi(rows *sql.Rows) (*Ami, error) {
	var err error
	var Id int64
	var VirtualizationType string
	var Description string
	var Hypervisor string
	var EnaSupport bool
	var ImageId string
	var State string
	var DeviceName string
	var SnapShotId string
	var DeleteOnTermination bool
	var VolumeType string
	var VolumeSize int16
	var Encrypted bool
	var Architecture string
	var ImageLocation string
	var RootDeviceType string
	var OwnerId string
	var RootDeviceName string
	var CreationDate string
	var Public bool
	var ImageType string
	var Name string

	err = rows.Scan(&Id, &VirtualizationType, &Description, &Hypervisor, &EnaSupport,
		&ImageId, &State, &DeviceName, &SnapShotId, &DeleteOnTermination,
		&VolumeType, &VolumeSize, &Encrypted, &Architecture, &ImageLocation,
		&RootDeviceType, &OwnerId, &RootDeviceName, &CreationDate, &Public,
		&ImageType, &Name)
	if err != nil {
		return nil, err
	}
	return NewAmi(Id, VirtualizationType, Description, Hypervisor, EnaSupport,
		ImageId, State, DeviceName, SnapShotId, DeleteOnTermination,
		VolumeType, VolumeSize, Encrypted, Architecture, ImageLocation,
		RootDeviceType, OwnerId, RootDeviceName, CreationDate, Public,
		ImageType, Name), nil
}

func rowsResultSetToAmi(rows *sql.Rows) (*Ami, error) {
	var err error
	if rows.Next() {
		var err error
		var Id int64
		var VirtualizationType string
		var Description string
		var Hypervisor string
		var EnaSupport bool
		var ImageId string
		var State string
		var DeviceName string
		var SnapShotId string
		var DeleteOnTermination bool
		var VolumeType string
		var VolumeSize int16
		var Encrypted bool
		var Architecture string
		var ImageLocation string
		var RootDeviceType string
		var OwnerId string
		var RootDeviceName string
		var CreationDate string
		var Public bool
		var ImageType string
		var Name string

		err = rows.Scan(&Id, &VirtualizationType, &Description, &Hypervisor, &EnaSupport,
			&ImageId, &State, &DeviceName, &SnapShotId, &DeleteOnTermination,
			&VolumeType, &VolumeSize, &Encrypted, &Architecture, &ImageLocation,
			&RootDeviceType, &OwnerId, &RootDeviceName, &CreationDate, &Public,
			&ImageType, &Name)
		if err != nil {
			return nil, err
		}
		return NewAmi(Id, VirtualizationType, Description, Hypervisor, EnaSupport,
			ImageId, State, DeviceName, SnapShotId, DeleteOnTermination,
			VolumeType, VolumeSize, Encrypted, Architecture, ImageLocation,
			RootDeviceType, OwnerId, RootDeviceName, CreationDate, Public,
			ImageType, Name), nil
	}
	return nil, err
}

func createAmi(db *sql.DB, VirtualizationType string, Description string, Hypervisor string, EnaSupport bool,
	ImageId string, State string, DeviceName string, SnapShotId string, DeleteOnTermination bool,
	VolumeType string, VolumeSize int16, Encrypted bool, Architecture string, ImageLocation string,
	RootDeviceType string, OwnerId string, RootDeviceName string, CreationDate string, Public bool,
	ImageType string, Name string) *Ami {
	rows := db.QueryRow("insert into ami(virtualization_type,description,hypervisor,ena_support,image_id,state,device_name,snap_shot_id,delete_on_termination,volume_type,volume_size,encrypted,architecture,image_location,root_device_type,owner_id,root_device_name,creation_date,public,image_type,name) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21) returning id,virtualization_type,description,hypervisor,ena_support,image_id,state,device_name,snap_shot_id,delete_on_termination,volume_type,volume_size,encrypted,architecture,image_location,root_device_type,owner_id,root_device_name,creation_date,public,image_type,name", VirtualizationType, Description, Hypervisor, EnaSupport,
		ImageId, State, DeviceName, SnapShotId, DeleteOnTermination,
		VolumeType, VolumeSize, Encrypted, Architecture, ImageLocation,
		RootDeviceType, OwnerId, RootDeviceName, CreationDate, Public,
		ImageType, Name)

	ami, err := rowResultSetToAmi(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return ami
}

func loadAmiById(db *sql.DB, id int64) (*Ami, error) {
	rows, err := db.Query("select id,virtualization_type,description,hypervisor,ena_support,image_id,state,device_name,snap_shot_id,delete_on_termination,volume_type,volume_size,encrypted,architecture,image_location,root_device_type,owner_id,root_device_name,creation_date,public,image_type,name from ami where id=$1", id)
	if err != nil {
		return nil, err
	}

	ami, err := rowsResultSetToAmi(rows)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	return ami, nil
}

func loadAmiByName(db *sql.DB, name string) (*Ami, error) {
	rows, err := db.Query("select id,virtualization_type,description,hypervisor,ena_support,image_id,state,device_name,snap_shot_id,delete_on_termination,volume_type,volume_size,encrypted,architecture,image_location,root_device_type,owner_id,root_device_name,creation_date,public,image_type,name from ami where name=$1", name)
	if err != nil {
		return nil, err
	}

	ami, err := rowsResultSetToAmi(rows)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	return ami, nil
}

func loadAllAmi(db *sql.DB) ([]*Ami, error) {
	rows, err := db.Query("select id,virtualization_type,description,hypervisor,ena_support,image_id,state,device_name,snap_shot_id,delete_on_termination,volume_type,volume_size,encrypted,architecture,image_location,root_device_type,owner_id,root_device_name,creation_date, public,image_type,name from ami order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfAmi := make([]*Ami, 0, 0)
	for rows.Next() {
		ami, err := rowsNoFetchResultSetToAmi(rows)
		if err == nil {
			arrayOfAmi = append(arrayOfAmi, ami)
		} else {
			log.Println(err)
		}
	}

	return arrayOfAmi, nil
}

func loadAllAmiByOrderUnique(db *sql.DB) ([]*Ami, error) {
	rows, err := db.Query("select id,virtualization_type,description,hypervisor,ena_support,image_id,state,device_name,snap_shot_id,delete_on_termination,volume_type,volume_size,encrypted,architecture,image_location,root_device_type,owner_id,root_device_name,creation_date, public,image_type,name from ami order by image_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfAmi := make([]*Ami, 0, 0)
	for rows.Next() {
		ami, err := rowsNoFetchResultSetToAmi(rows)
		if err == nil {
			arrayOfAmi = append(arrayOfAmi, ami)
		} else {
			log.Println(err)
		}
	}

	return arrayOfAmi, nil
}

func loadAmiByAWSReferenceId(db *sql.DB, imageId string) (*Ami, error) {
	rows, err := db.Query("select id,virtualization_type,description,hypervisor,ena_support,image_id,state,device_name,snap_shot_id,delete_on_termination,volume_type,volume_size,encrypted,architecture,image_location,root_device_type,owner_id,root_device_name,creation_date, public,image_type,name from ami where image_id=$1", imageId)
	if err != nil {
		return nil, err
	}

	image, err := rowsResultSetToAmi(rows)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	return image, nil
}

func updateAmiSnapshotId(db *sql.DB, ami *Ami, snapshotId int64) *Ami {
	rows := db.QueryRow("update ami set ref_snapshot_id=$1 where id=$2 returning id,virtualization_type,description,hypervisor,ena_support,image_id,state,device_name,snap_shot_id,delete_on_termination,volume_type,volume_size,encrypted,architecture,image_location,root_device_type,owner_id,root_device_name,creation_date, public,image_type,name", snapshotId, ami.Id)

	ami, err := rowResultSetToAmi(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return ami
}

func deleteAllAmi(db *sql.DB) error {
	_, err := db.Exec("delete from ami")

	return err
}
