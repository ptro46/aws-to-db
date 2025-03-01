package main

import (
	"database/sql" // package SQL
	"log"

	_ "github.com/lib/pq" // driver Postgres
)

func rowResultSetToSecurityGroupTag(row *sql.Row) (*SecurityGroupTag, error) {
	var err error
	var Id int64
	var SecurityGroupId int64
	var Key string
	var Value string

	err = row.Scan(&Id, &SecurityGroupId, &Key, &Value)
	if err != nil {
		return nil, err
	}
	return NewSecurityGroupTag(Id, SecurityGroupId, Key, Value), nil
}

func rowsNoFetchResultSetToSecurityGroupTag(rows *sql.Rows) (*SecurityGroupTag, error) {
	var err error
	var Id int64
	var SecurityGroupId int64
	var Key string
	var Value string

	err = rows.Scan(&Id, &SecurityGroupId, &Key, &Value)
	if err != nil {
		return nil, err
	}
	return NewSecurityGroupTag(Id, SecurityGroupId, Key, Value), nil
}

func rowsResultSetToSecurityGroupTag(rows *sql.Rows) (*SecurityGroupTag, error) {
	var err error
	if rows.Next() {
		var Id int64
		var SecurityGroupId int64
		var Key string
		var Value string

		err = rows.Scan(&Id, &SecurityGroupId, &Key, &Value)
		if err != nil {
			return nil, err
		}
		return NewSecurityGroupTag(Id, SecurityGroupId, Key, Value), nil
	}
	return nil, err
}

func createSecurityGroupTag(db *sql.DB, SecurityGroupId int64, Key string, Value string) *SecurityGroupTag {
	rows := db.QueryRow("insert into security_group_tag(security_group_id,key,value) values($1,$2,$3) returning id,security_group_id,key,value", SecurityGroupId, Key, Value)

	secGroupTag, err := rowResultSetToSecurityGroupTag(rows)
	if err != nil {
		return nil
	}
	return secGroupTag
}

func loadAllSecurityGroupTag(db *sql.DB) ([]*SecurityGroupTag, error) {
	rows, err := db.Query("select id,security_group_id,key,value from security_group_tag order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfSGTag := make([]*SecurityGroupTag, 0, 0)
	for rows.Next() {
		secGroupTag, err := rowsNoFetchResultSetToSecurityGroupTag(rows)
		if err == nil {
			arrayOfSGTag = append(arrayOfSGTag, secGroupTag)
		} else {
			log.Println(err)
		}
	}

	return arrayOfSGTag, nil
}

func loadAllSecurityGroupTagByOrderUnique(db *sql.DB) ([]*SecurityGroupTag, error) {
	rows, err := db.Query("select id,security_group_id,key,value from security_group_tag order by value")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfSGTag := make([]*SecurityGroupTag, 0, 0)
	for rows.Next() {
		secGroupTag, err := rowsNoFetchResultSetToSecurityGroupTag(rows)
		if err == nil {
			arrayOfSGTag = append(arrayOfSGTag, secGroupTag)
		} else {
			log.Println(err)
		}
	}

	return arrayOfSGTag, nil
}

func deleteAllSecurityGroupTag(db *sql.DB) error {
	_, err := db.Exec("delete from security_group_tag")

	return err
}
