package main

import (
	"database/sql" // package SQL
	"fmt"
	"log"

	_ "github.com/lib/pq" // driver Postgres
)

func rowResultSetToStackTag(row *sql.Row) (*StackTag, error) {
	var err error
	var Id int64
	var StackId int64
	var Key string
	var Value string
	err = row.Scan(&Id, &StackId, &Key, &Value)
	if err != nil {
		return nil, err
	}
	return NewStackTag(Id, StackId, Key, Value), nil
}

func rowsNoFetchResultSetToStackStackTag(rows *sql.Rows) (*StackTag, error) {
	var err error
	var Id int64
	var StackId int64
	var Key string
	var Value string

	err = rows.Scan(&Id, &StackId, &Key, &Value)
	if err != nil {
		return nil, err
	}
	return NewStackTag(Id, StackId, Key, Value), nil
}

func rowsResultSetToStackTag(rows *sql.Rows) (*StackTag, error) {
	var err error
	if rows.Next() {
		var Id int64
		var StackId int64
		var Key string
		var Value string

		err = rows.Scan(&Id, &StackId, &Key, &Value)
		if err != nil {
			return nil, err
		}
		return NewStackTag(Id, StackId, Key, Value), nil
	}
	return nil, err
}

func createStackTag(db *sql.DB, StackId int64, Key string, Value string) *StackTag {
	rows := db.QueryRow("insert into stack_tag(stack_id,key,value) values($1,$2,$3) returning id,stack_id,key,value", StackId, Key, Value)

	stackTag, err := rowResultSetToStackTag(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return stackTag
}

func loadAllStackTag(db *sql.DB) ([]*StackTag, error) {
	rows, err := db.Query("select id,stack_id,key,value from stack_tag order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfStackTag := make([]*StackTag, 0, 0)
	for rows.Next() {
		stackTag, err := rowsNoFetchResultSetToStackStackTag(rows)
		if err == nil {
			arrayOfStackTag = append(arrayOfStackTag, stackTag)
		} else {
			log.Println(err)
		}
	}

	return arrayOfStackTag, nil
}

func loadAllStackTagByOrderUnique(db *sql.DB) ([]*StackTag, error) {
	rows, err := db.Query("select id,stack_id,key,value from stack_tag order by value")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfStackTag := make([]*StackTag, 0, 0)
	for rows.Next() {
		stackTag, err := rowsNoFetchResultSetToStackStackTag(rows)
		if err == nil {
			arrayOfStackTag = append(arrayOfStackTag, stackTag)
		} else {
			log.Println(err)
		}
	}

	return arrayOfStackTag, nil
}

func deleteAllStackTag(db *sql.DB) error {
	_, err := db.Exec("delete from stack_tag")

	return err
}
