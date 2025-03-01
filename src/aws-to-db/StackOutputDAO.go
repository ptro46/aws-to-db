package main

import (
	"database/sql" // package SQL
	"fmt"
	"log"

	_ "github.com/lib/pq" // driver Postgres
)

func rowResultSetToStackOutput(row *sql.Row) (*StackOutput, error) {
	var err error
	var Id int64
	var StackId int64
	var OutputKey string
	var OutputValue string
	var Description string
	var ExportName string

	err = row.Scan(&Id, &StackId, &OutputKey, &OutputValue, &Description, &ExportName)
	if err != nil {
		return nil, err
	}
	return NewStackOutput(Id, StackId, OutputKey, OutputValue, Description, ExportName), nil
}

func rowsNoFetchResultSetToStackStackOutput(rows *sql.Rows) (*StackOutput, error) {
	var err error
	var Id int64
	var StackId int64
	var OutputKey string
	var OutputValue string
	var Description string
	var ExportName string

	err = rows.Scan(&Id, &StackId, &OutputKey, &OutputValue, &Description, &ExportName)
	if err != nil {
		return nil, err
	}
	return NewStackOutput(Id, StackId, OutputKey, OutputValue, Description, ExportName), nil
}

func rowsResultSetToStackOutput(rows *sql.Rows) (*StackOutput, error) {
	var err error
	if rows.Next() {
		var Id int64
		var StackId int64
		var OutputKey string
		var OutputValue string
		var Description string
		var ExportName string

		err = rows.Scan(&Id, &StackId, &OutputKey, &OutputValue, &Description, &ExportName)
		if err != nil {
			return nil, err
		}
		return NewStackOutput(Id, StackId, OutputKey, OutputValue, Description, ExportName), nil
	}
	return nil, err
}

func createStackOutput(db *sql.DB, StackId int64, OutputKey string, OutputValue string, Description string, ExportName string) *StackOutput {
	rows := db.QueryRow("insert into stack_output(stack_id,output_key,output_value,description,export_name) values($1,$2,$3,$4,$5) returning id,stack_id,output_key,output_value,description,export_name", StackId, OutputKey, OutputValue, Description, ExportName)

	stackOutput, err := rowResultSetToStackOutput(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return stackOutput
}

func loadAllStackOutput(db *sql.DB) ([]*StackOutput, error) {
	rows, err := db.Query("select id,stack_id,output_key,output_value,description,export_name from stack_output order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfStackOutput := make([]*StackOutput, 0, 0)
	for rows.Next() {
		stackOutput, err := rowsNoFetchResultSetToStackStackOutput(rows)
		if err == nil {
			arrayOfStackOutput = append(arrayOfStackOutput, stackOutput)
		} else {
			log.Println(err)
		}
	}

	return arrayOfStackOutput, nil
}

func deleteAllStackOutput(db *sql.DB) error {
	_, err := db.Exec("delete from stack_output")

	return err
}
