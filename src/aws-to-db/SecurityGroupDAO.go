package main

import (
	"database/sql" // package SQL
	"fmt"
	"log"

	_ "github.com/lib/pq" // driver Postgres
)

func rowResultSetToSecurityGroup(row *sql.Row) (*SecurityGroup, error) {
	var err error
	var Id int64
	var Description string
	var GroupName string
	var VpcId string
	var RefVpcId sql.NullInt64
	var OwnerId string
	var GroupId string

	err = row.Scan(&Id, &Description, &GroupName, &VpcId, &RefVpcId, &OwnerId, &GroupId)
	if err != nil {
		return nil, err
	}
	return NewSecurityGroup(Id, Description, GroupName, VpcId, RefVpcId, OwnerId, GroupId), nil
}

func rowsNoFetchResultSetToSecurityGroup(rows *sql.Rows) (*SecurityGroup, error) {
	var err error
	var Id int64
	var Description string
	var GroupName string
	var VpcId string
	var RefVpcId sql.NullInt64
	var OwnerId string
	var GroupId string

	err = rows.Scan(&Id, &Description, &GroupName, &VpcId, &RefVpcId, &OwnerId, &GroupId)
	if err != nil {
		return nil, err
	}
	return NewSecurityGroup(Id, Description, GroupName, VpcId, RefVpcId, OwnerId, GroupId), nil
}

func rowsResultSetToSecurityGroup(rows *sql.Rows) (*SecurityGroup, error) {
	var err error
	if rows.Next() {
		var Id int64
		var Description string
		var GroupName string
		var VpcId string
		var RefVpcId sql.NullInt64
		var OwnerId string
		var GroupId string

		err = rows.Scan(&Id, &Description, &GroupName, &VpcId, &RefVpcId, &OwnerId, &GroupId)
		if err != nil {
			return nil, err
		}
		return NewSecurityGroup(Id, Description, GroupName, VpcId, RefVpcId, OwnerId, GroupId), nil
	}
	return nil, err
}

func createSecurityGroup(db *sql.DB, Description string, GroupName string, VpcId string, OwnerId string, GroupId string) *SecurityGroup {
	rows := db.QueryRow("insert into security_group(description,group_name,vpc_id,owner_id,group_id) values($1,$2,$3,$4,$5) returning id,description,group_name,vpc_id,ref_vpc_id,owner_id,group_id", Description, GroupName, VpcId, OwnerId, GroupId)

	secGroup, err := rowResultSetToSecurityGroup(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return secGroup
}

func updateSecurityVPCId(db *sql.DB, securityGroup *SecurityGroup, VpcId int64) *SecurityGroup {
	rows := db.QueryRow("update security_group set ref_vpc_id=$1 where id=$2 returning id,description,group_name,vpc_id,ref_vpc_id,owner_id,group_id", VpcId, securityGroup.Id)

	securityGroup, err := rowResultSetToSecurityGroup(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return securityGroup
}

func loadAllSecurityGroup(db *sql.DB) ([]*SecurityGroup, error) {
	rows, err := db.Query("select id,description,group_name,vpc_id,ref_vpc_id,owner_id,group_id from security_group order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfSG := make([]*SecurityGroup, 0, 0)
	for rows.Next() {
		secGroup, err := rowsNoFetchResultSetToSecurityGroup(rows)
		if err == nil {
			arrayOfSG = append(arrayOfSG, secGroup)
		} else {
			log.Println(err)
		}
	}

	return arrayOfSG, nil
}

func loadAllSecurityGroupByOrderUnique(db *sql.DB) ([]*SecurityGroup, error) {
	rows, err := db.Query("select id,description,group_name,vpc_id,ref_vpc_id,owner_id,group_id from security_group order by group_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfSG := make([]*SecurityGroup, 0, 0)
	for rows.Next() {
		secGroup, err := rowsNoFetchResultSetToSecurityGroup(rows)
		if err == nil {
			arrayOfSG = append(arrayOfSG, secGroup)
		} else {
			log.Println(err)
		}
	}

	return arrayOfSG, nil
}

func loadSecurityGroupByAWSReferenceId(db *sql.DB, groupId string) (*SecurityGroup, error) {
	rows, err := db.Query("select id,description,group_name,vpc_id,ref_vpc_id,owner_id,group_id from security_group where group_id=$1", groupId)
	if err != nil {
		return nil, err
	}

	secGroup, err := rowsResultSetToSecurityGroup(rows)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	return secGroup, nil
}

func deleteAllSecurityGroup(db *sql.DB) error {
	_, err := db.Exec("delete from security_group")

	return err
}
