package main

import (
	"database/sql" // package SQL
	"log"

	_ "github.com/lib/pq" // driver Postgres
)

func rowResultSetToVolumeTag(row *sql.Row) (*VolumeTag, error) {
	var err error
	var Id int64
	var VolumeId int64
	var Key string
	var Value string
	err = row.Scan(&Id, &VolumeId, &Key, &Value)
	if err != nil {
		return nil, err
	}
	return NewVolumeTag(Id, VolumeId, Key, Value), nil
}

func rowsNoFetchResultSetToVolumeTag(rows *sql.Rows) (*VolumeTag, error) {
	var err error
	var Id int64
	var VolumeId int64
	var Key string
	var Value string

	err = rows.Scan(&Id, &VolumeId, &Key, &Value)
	if err != nil {
		return nil, err
	}
	return NewVolumeTag(Id, VolumeId, Key, Value), nil
}

func rowsResultSetToVolumeTag(rows *sql.Rows) (*VolumeTag, error) {
	var err error
	if rows.Next() {
		var Id int64
		var VolumeId int64
		var Key string
		var Value string

		err = rows.Scan(&Id, &VolumeId, &Key, &Value)
		if err != nil {
			return nil, err
		}
		return NewVolumeTag(Id, VolumeId, Key, Value), nil
	}
	return nil, err
}

func createVolumetag(db *sql.DB, VolumeId int64, Key string, Value string) *VolumeTag {
	rows := db.QueryRow("insert into volume_tag(volume_id,key,value) values($1,$2,$3) returning id,volume_id,key,value", VolumeId, Key, Value)

	volumeTag, err := rowResultSetToVolumeTag(rows)
	if err != nil {
		return nil
	}
	return volumeTag
}

func loadAllVolumeTag(db *sql.DB) ([]*VolumeTag, error) {
	rows, err := db.Query("select id,volume_id,key,value from volume_tag order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfVolumeTag := make([]*VolumeTag, 0, 0)
	for rows.Next() {
		tag, err := rowsNoFetchResultSetToVolumeTag(rows)
		if err == nil {
			arrayOfVolumeTag = append(arrayOfVolumeTag, tag)
		} else {
			log.Println(err)
		}
	}

	return arrayOfVolumeTag, nil
}

func loadAllVolumeTagByOrderUnique(db *sql.DB) ([]*VolumeTag, error) {
	rows, err := db.Query("select id,volume_id,key,value from volume_tag order by value")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfVolumeTag := make([]*VolumeTag, 0, 0)
	for rows.Next() {
		tag, err := rowsNoFetchResultSetToVolumeTag(rows)
		if err == nil {
			arrayOfVolumeTag = append(arrayOfVolumeTag, tag)
		} else {
			log.Println(err)
		}
	}

	return arrayOfVolumeTag, nil
}

func deleteAllVolumeTag(db *sql.DB) error {
	_, err := db.Exec("delete from volume_tag")

	return err
}
