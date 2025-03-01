package main

import (
	"database/sql" // package SQL
	"fmt"
	"log"

	_ "github.com/lib/pq" // driver Postgres
)

func rowResultSetToViewInstanceByCidr(row *sql.Row) (*ViewInstanceByCidr, error) {
	var err error
	var InstanceId int64
	var InstanceKey string
	var InstanceName string
	var CidrBlock string

	err = row.Scan(&InstanceId, &InstanceKey, &InstanceName, &CidrBlock)
	if err != nil {
		return nil, err
	}
	return NewViewInstanceByCidr(InstanceId, InstanceKey, InstanceName, CidrBlock), nil
}

func rowsNoFetchResultSetToViewInstanceByCidr(rows *sql.Rows) (*ViewInstanceByCidr, error) {
	var err error
	var InstanceId int64
	var InstanceKey string
	var InstanceName string
	var CidrBlock string

	err = rows.Scan(&InstanceId, &InstanceKey, &InstanceName, &CidrBlock)
	if err != nil {
		return nil, err
	}
	return NewViewInstanceByCidr(InstanceId, InstanceKey, InstanceName, CidrBlock), nil
}

func rowsResultSetToViewInstanceByCidr(rows *sql.Rows) (*ViewInstanceByCidr, error) {
	var err error
	if rows.Next() {
		var InstanceId int64
		var InstanceKey string
		var InstanceName string
		var CidrBlock string

		err = rows.Scan(&InstanceId, &InstanceKey, &InstanceName, &CidrBlock)
		if err != nil {
			return nil, err
		}
		return NewViewInstanceByCidr(InstanceId, InstanceKey, InstanceName, CidrBlock), nil
	}
	return nil, err
}

func createViewInstanceByCidr(db *sql.DB, InstanceId int64, InstanceKey string, InstanceName string, CidrBlock string) *ViewInstanceByCidr {
	rows := db.QueryRow("insert into view_instance_by_cidr(instance_id,instance_key,instance_name,cidr_block) values($1,$2,$3,$4) returning instance_id,instance_key,instance_name,cidr_block",
		InstanceId, InstanceKey, InstanceName, CidrBlock)

	viewInstanceByCidr, err := rowResultSetToViewInstanceByCidr(rows)
	if err != nil {
		return nil
	}
	return viewInstanceByCidr
}

func loadAllViewInstanceByCidr(db *sql.DB) ([]*ViewInstanceByCidr, error) {
	rows, err := db.Query("select instance_id,instance_key,instance_name,cidr_block from view_instance_by_cidr order by instance_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfViewInstanceByCidr := make([]*ViewInstanceByCidr, 0, 0)
	for rows.Next() {
		viewInstanceByCidr, err := rowsNoFetchResultSetToViewInstanceByCidr(rows)
		if err == nil {
			arrayOfViewInstanceByCidr = append(arrayOfViewInstanceByCidr, viewInstanceByCidr)
		} else {
			log.Println(err)
		}
	}

	return arrayOfViewInstanceByCidr, nil
}

func loadAllViewInstanceByCidrInstanceName(db *sql.DB, cidrIp string) (string, error) {
	query := fmt.Sprintf("select distinct(instance_id),instance_name from view_instance_by_cidr where cidr_block like '%s", cidrIp)
	rows, err := db.Query(query + "%' limit 1")
	if err != nil {
		return "", err
	}
	defer rows.Close()
	instanceName := ""
	if rows.Next() {
		var InstanceId int64

		err = rows.Scan(&InstanceId, &instanceName)
		if err != nil {
			return instanceName, err
		}
	}

	return instanceName, nil
}

func deleteAllViewInstanceByCidr(db *sql.DB) error {
	_, err := db.Exec("delete from ViewInstanceByCidr")

	return err
}
