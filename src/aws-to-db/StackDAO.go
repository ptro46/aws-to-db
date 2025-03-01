package main

import (
	"database/sql" // package SQL
	"log"

	_ "github.com/lib/pq" // driver Postgres
)

func rowResultSetToStack(row *sql.Row) (*Stack, error) {
	var err error
	var Id int64
	var StackId string
	var StackName string
	var ChangeSetId string
	var Description string
	var CreationTime string
	var DeletionTime string
	var LastUpdateTime string
	var StackStatus string
	var StackStatusRaison string
	var DisableRollback bool
	var TimeoutInMinutes int64
	var RoleArn string
	var EnableTerminationProtection bool
	var ParentId string
	var RootId string
	var NextToken string

	err = row.Scan(&Id, &StackId, &StackName, &ChangeSetId, &Description, &CreationTime, &DeletionTime, &LastUpdateTime,
		&StackStatus, &StackStatusRaison, &DisableRollback, &TimeoutInMinutes, &RoleArn, &EnableTerminationProtection,
		&ParentId, &RootId, &NextToken)
	if err != nil {
		return nil, err
	}
	return NewStack(Id, StackId, StackName, ChangeSetId, Description, CreationTime, DeletionTime, LastUpdateTime,
		StackStatus, StackStatusRaison, DisableRollback, TimeoutInMinutes, RoleArn, EnableTerminationProtection,
		ParentId, RootId, NextToken), nil
}

func rowsNoFetchResultSetToStack(rows *sql.Rows) (*Stack, error) {
	var err error
	var Id int64
	var StackId string
	var StackName string
	var ChangeSetId string
	var Description string
	var CreationTime string
	var DeletionTime string
	var LastUpdateTime string
	var StackStatus string
	var StackStatusRaison string
	var DisableRollback bool
	var TimeoutInMinutes int64
	var RoleArn string
	var EnableTerminationProtection bool
	var ParentId string
	var RootId string
	var NextToken string

	err = rows.Scan(&Id, &StackId, &StackName, &ChangeSetId, &Description, &CreationTime, &DeletionTime, &LastUpdateTime,
		&StackStatus, &StackStatusRaison, &DisableRollback, &TimeoutInMinutes, &RoleArn, &EnableTerminationProtection,
		&ParentId, &RootId, &NextToken)
	if err != nil {
		return nil, err
	}
	return NewStack(Id, StackId, StackName, ChangeSetId, Description, CreationTime, DeletionTime, LastUpdateTime,
		StackStatus, StackStatusRaison, DisableRollback, TimeoutInMinutes, RoleArn, EnableTerminationProtection,
		ParentId, RootId, NextToken), nil
}

func rowsResultSetToStack(rows *sql.Rows) (*Stack, error) {
	var err error
	if rows.Next() {
		var err error
		var Id int64
		var StackId string
		var StackName string
		var ChangeSetId string
		var Description string
		var CreationTime string
		var DeletionTime string
		var LastUpdateTime string
		var StackStatus string
		var StackStatusRaison string
		var DisableRollback bool
		var TimeoutInMinutes int64
		var RoleArn string
		var EnableTerminationProtection bool
		var ParentId string
		var RootId string
		var NextToken string

		err = rows.Scan(&Id, &StackId, &StackName, &ChangeSetId, &Description, &CreationTime, &DeletionTime, &LastUpdateTime,
			&StackStatus, &StackStatusRaison, &DisableRollback, &TimeoutInMinutes, &RoleArn, &EnableTerminationProtection,
			&ParentId, &RootId, &NextToken)
		if err != nil {
			return nil, err
		}
		return NewStack(Id, StackId, StackName, ChangeSetId, Description, CreationTime, DeletionTime, LastUpdateTime,
			StackStatus, StackStatusRaison, DisableRollback, TimeoutInMinutes, RoleArn, EnableTerminationProtection,
			ParentId, RootId, NextToken), nil
	}
	return nil, err
}

func createStack(db *sql.DB, StackId string, StackName string, ChangeSetId string, Description string, CreationTime string,
	DeletionTime string, LastUpdateTime string, StackStatus string, StackStatusRaison string, DisableRollback bool,
	TimeoutInMinutes int64, RoleArn string, EnableTerminationProtection bool, ParentId string, RootId string,
	NextToken string) *Stack {
	rows := db.QueryRow("insert into stack(stack_id,stack_name,change_set_id,description,creation_time,deletion_time,last_update_time,stack_status,stack_status_raison,disable_rollback,timeout_in_minutes,role_arn,enable_termination_protection,parent_id,root_id,next_token) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16) returning id,stack_id,stack_name,change_set_id,description,creation_time,deletion_time,last_update_time,stack_status,stack_status_raison,disable_rollback,timeout_in_minutes,role_arn,enable_termination_protection,parent_id,root_id,next_token", StackId, StackName, ChangeSetId, Description, CreationTime, DeletionTime, LastUpdateTime,
		StackStatus, StackStatusRaison, DisableRollback, TimeoutInMinutes, RoleArn, EnableTerminationProtection,
		ParentId, RootId, NextToken)

	stack, err := rowResultSetToStack(rows)
	if err != nil {
		return nil
	}
	return stack
}

func loadStackById(db *sql.DB, id int64) (*Stack, error) {
	rows, err := db.Query("select id,stack_id,stack_name,change_set_id,description,creation_time,deletion_time,last_update_time,stack_status,stack_status_raison,disable_rollback,timeout_in_minutes,role_arn,enable_termination_protection,parent_id,root_id,next_token from stack where id=$1", id)
	if err != nil {
		return nil, err
	}

	stack, err := rowsResultSetToStack(rows)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	return stack, nil
}

func loadAllStack(db *sql.DB) ([]*Stack, error) {
	rows, err := db.Query("select id,stack_id,stack_name,change_set_id,description,creation_time,deletion_time,last_update_time,stack_status,stack_status_raison,disable_rollback,timeout_in_minutes,role_arn,enable_termination_protection,parent_id,root_id,next_tokenfrom stack order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfStack := make([]*Stack, 0, 0)
	for rows.Next() {
		stack, err := rowsNoFetchResultSetToStack(rows)
		if err == nil {
			arrayOfStack = append(arrayOfStack, stack)
		} else {
			log.Println(err)
		}
	}

	return arrayOfStack, nil
}

func loadAllStackByOrderUnique(db *sql.DB) ([]*Stack, error) {
	rows, err := db.Query("select id,stack_id,stack_name,change_set_id,description,creation_time,deletion_time,last_update_time,stack_status,stack_status_raison,disable_rollback,timeout_in_minutes,role_arn,enable_termination_protection,parent_id,root_id,next_token from stack order by stack_name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfStack := make([]*Stack, 0, 0)
	for rows.Next() {
		stack, err := rowsNoFetchResultSetToStack(rows)
		if err == nil {
			arrayOfStack = append(arrayOfStack, stack)
		} else {
			log.Println(err)
		}
	}

	return arrayOfStack, nil
}

func deleteAllStack(db *sql.DB) error {
	_, err := db.Exec("delete from stack")

	return err
}
