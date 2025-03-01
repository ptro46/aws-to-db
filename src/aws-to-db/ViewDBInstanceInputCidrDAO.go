package main

import (
	"database/sql" // package SQL
	"fmt"
	"log"

	_ "github.com/lib/pq" // driver Postgres
)

func rowResultSetToViewDBInstanceInputCidr(row *sql.Row) (*ViewDBInstanceInputCidr, error) {
	var err error
	var DBInstanceId int64
	var DBInstanceName string
	var SubnetCidrBlock string

	err = row.Scan(&DBInstanceId, &DBInstanceName, &SubnetCidrBlock)
	if err != nil {
		return nil, err
	}
	return NewViewDBInstanceInputCidr(DBInstanceId, DBInstanceName, SubnetCidrBlock), nil
}

func rowsNoFetchResultSetToViewDBInstanceInputCidr(rows *sql.Rows) (*ViewDBInstanceInputCidr, error) {
	var err error
	var DBInstanceId int64
	var DBInstanceName string
	var SubnetCidrBlock string

	err = rows.Scan(&DBInstanceId, &DBInstanceName, &SubnetCidrBlock)
	if err != nil {
		return nil, err
	}
	return NewViewDBInstanceInputCidr(DBInstanceId, DBInstanceName, SubnetCidrBlock), nil
}

func rowsResultSetToViewDBInstanceInputCidr(rows *sql.Rows) (*ViewDBInstanceInputCidr, error) {
	var err error
	if rows.Next() {
		var DBInstanceId int64
		var DBInstanceName string
		var SubnetCidrBlock string

		err = rows.Scan(&DBInstanceId, &DBInstanceName, &SubnetCidrBlock)
		if err != nil {
			return nil, err
		}
		return NewViewDBInstanceInputCidr(DBInstanceId, DBInstanceName, SubnetCidrBlock), nil
	}
	return nil, err
}

func createViewDBInstanceInputCidr(db *sql.DB, SubnetId int64, SubnetCidrBlock string, DBInstanceId int64, DBInstanceName string) *ViewDBInstanceInputCidr {
	rows := db.QueryRow("insert into view_db_instance_input_cidr(db_instance_id,db_instance_name,subnet_cidr_block) values($1,$2,3$) returning db_instance_id,db_instance_name,subnet_cidr_block",
		SubnetId, SubnetCidrBlock, DBInstanceId, DBInstanceName)

	viewDBInstanceInputCidr, err := rowResultSetToViewDBInstanceInputCidr(rows)
	if err != nil {
		return nil
	}
	return viewDBInstanceInputCidr
}

func loadAllViewDBInstanceInputCidr(db *sql.DB) ([]*ViewDBInstanceInputCidr, error) {
	rows, err := db.Query("select db_instance_id,db_instance_name,subnet_cidr_block from view_db_instance_input_cidr")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfViewDBInstanceInputCidr := make([]*ViewDBInstanceInputCidr, 0, 0)
	for rows.Next() {
		viewDBInstanceInputCidr, err := rowsNoFetchResultSetToViewDBInstanceInputCidr(rows)
		if err == nil {
			arrayOfViewDBInstanceInputCidr = append(arrayOfViewDBInstanceInputCidr, viewDBInstanceInputCidr)
		} else {
			log.Println(err)
		}
	}

	return arrayOfViewDBInstanceInputCidr, nil
}

func loadAllViewDBInstanceByCidrDBInstanceName(db *sql.DB, cidrIp string) (string, error) {
	query := fmt.Sprintf("select distinct(db_instance_name) from view_db_instance_input_cidr where subnet_cidr_block like '%s", cidrIp)
	rows, err := db.Query(query + "%'")
	if err != nil {
		return "", err
	}
	defer rows.Close()

	viewDBInstanceByCidrDBInstanceName := ""
	if rows.Next() {
		err = rows.Scan(&viewDBInstanceByCidrDBInstanceName)
		if err != nil {
			return "", err
		}
	}

	return viewDBInstanceByCidrDBInstanceName, nil
}

func deleteAllViewDBInstanceInputCidr(db *sql.DB) error {
	_, err := db.Exec("delete from view_db_instance_input_cidr")

	return err
}
