package main

import (
	"database/sql" // package SQL
	"log"

	_ "github.com/lib/pq" // driver Postgres
)

func rowResultSetToStackParameter(row *sql.Row) (*StackParameter, error) {
	var err error
	var Id int64
	var StackId int64
	var ParameterKey string
	var ParameterValue string
	var UsePreviousValue bool
	var ResolvedValue string

	err = row.Scan(&Id, &StackId, &ParameterKey, &ParameterValue, &UsePreviousValue, &ResolvedValue)
	if err != nil {
		return nil, err
	}
	return NewStackParameter(Id, StackId, ParameterKey, ParameterValue, UsePreviousValue, ResolvedValue), nil
}

func rowsNoFetchResultSetToStackStackParameter(rows *sql.Rows) (*StackParameter, error) {
	var err error
	var Id int64
	var StackId int64
	var ParameterKey string
	var ParameterValue string
	var UsePreviousValue bool
	var ResolvedValue string

	err = rows.Scan(&Id, &StackId, &ParameterKey, &ParameterValue, &UsePreviousValue, &ResolvedValue)
	if err != nil {
		return nil, err
	}
	return NewStackParameter(Id, StackId, ParameterKey, ParameterValue, UsePreviousValue, ResolvedValue), nil
}

func rowsResultSetToStackParameter(rows *sql.Rows) (*StackParameter, error) {
	var err error
	if rows.Next() {
		var Id int64
		var StackId int64
		var ParameterKey string
		var ParameterValue string
		var UsePreviousValue bool
		var ResolvedValue string

		err = rows.Scan(&Id, &StackId, &ParameterKey, &ParameterValue, &UsePreviousValue, &ResolvedValue)
		if err != nil {
			return nil, err
		}
		return NewStackParameter(Id, StackId, ParameterKey, ParameterValue, UsePreviousValue, ResolvedValue), nil
	}
	return nil, err
}

func createStackParameter(db *sql.DB, StackId int64, ParameterKey string, ParameterValue string, UsePreviousValue bool, ResolvedValue string) *StackParameter {
	rows := db.QueryRow("insert into stack_parameter(stack_id,parameter_key,parameter_value,use_previous_value,resolved_value) values($1,$2,$3,$4,$5) returning id,stack_id,parameter_key,parameter_value,use_previous_value,resolved_value	", StackId, ParameterKey, ParameterValue, UsePreviousValue, ResolvedValue)

	StackParameter, err := rowResultSetToStackParameter(rows)
	if err != nil {
		return nil
	}
	return StackParameter
}

func loadAllStackParameter(db *sql.DB) ([]*StackParameter, error) {
	rows, err := db.Query("select id,stack_id,parameter_key,parameter_value,use_previous_value,resolved_value from stack_parameter order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfStackParameter := make([]*StackParameter, 0, 0)
	for rows.Next() {
		stackParameter, err := rowsNoFetchResultSetToStackStackParameter(rows)
		if err == nil {
			arrayOfStackParameter = append(arrayOfStackParameter, stackParameter)
		} else {
			log.Println(err)
		}
	}

	return arrayOfStackParameter, nil
}

func deleteAllStackParameter(db *sql.DB) error {
	_, err := db.Exec("delete from stack_parameter")

	return err
}
