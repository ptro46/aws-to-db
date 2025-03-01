package main

import (
	"database/sql" // package SQL
	"fmt"
	"log"

	_ "github.com/lib/pq" // driver Postgres
)

func rowResultSetToAutoScalingGroupActivities(row *sql.Row) (*AutoScalingGroupActivities, error) {
	var err error
	var Id int64
	var ActivityId string
	var Description string
	var AutoScalingGroupName string
	var Details string
	var StartTime string
	var Progress int
	var EndTime string
	var Cause string
	var StatusCode string

	err = row.Scan(&Id, &ActivityId, &Description, &AutoScalingGroupName, &Details, &StartTime, &Progress, &EndTime, &Cause, &StatusCode)
	if err != nil {
		return nil, err
	}
	return NewAutoScalingGroupActivities(Id, ActivityId, Description, AutoScalingGroupName, Details, StartTime, Progress, EndTime, Cause, StatusCode), nil
}

func rowsNoFetchResultSetToAutoScalingGroupActivities(rows *sql.Rows) (*AutoScalingGroupActivities, error) {
	var err error
	var Id int64
	var ActivityId string
	var Description string
	var AutoScalingGroupName string
	var Details string
	var StartTime string
	var Progress int
	var EndTime string
	var Cause string
	var StatusCode string

	err = rows.Scan(&Id, &ActivityId, &Description, &AutoScalingGroupName, &Details, &StartTime, &Progress, &EndTime, &Cause, &StatusCode)
	if err != nil {
		return nil, err
	}
	return NewAutoScalingGroupActivities(Id, ActivityId, Description, AutoScalingGroupName, Details, StartTime, Progress, EndTime, Cause, StatusCode), nil

}

func rowsResultSetToAutoScalingGroupActivities(rows *sql.Rows) (*AutoScalingGroupActivities, error) {
	var err error
	if rows.Next() {
		var Id int64
		var ActivityId string
		var Description string
		var AutoScalingGroupName string
		var Details string
		var StartTime string
		var Progress int
		var EndTime string
		var Cause string
		var StatusCode string

		err = rows.Scan(&Id, &ActivityId, &Description, &AutoScalingGroupName, &Details, &StartTime, &Progress, &EndTime, &Cause, &StatusCode)
		if err != nil {
			return nil, err
		}
		return NewAutoScalingGroupActivities(Id, ActivityId, Description, AutoScalingGroupName, Details, StartTime, Progress, EndTime, Cause, StatusCode), nil

	}
	return nil, err
}

func createAutoScalingGroupActivities(db *sql.DB, ActivityId string, Description string, AutoScalingGroupName string, Details string, StartTime string, Progress int, EndTime string, Cause string, StatusCode string) *AutoScalingGroupActivities {
	rows := db.QueryRow("insert into auto_scaling_activities(activity_id,description,auto_scaling_group_name,details,start_time,progress,end_time,cause,status_code) values($1,$2,$3,$4,$5,$6,$7,$8,$9) returning id,activity_id,description,auto_scaling_group_name,details,start_time,progress,end_time,cause,status_code",
		ActivityId,
		Description,
		AutoScalingGroupName,
		Details,
		StartTime,
		Progress,
		EndTime,
		Cause,
		StatusCode)

	autoScaling, err := rowResultSetToAutoScalingGroupActivities(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return autoScaling
}

func loadAllAutoScalingGroupActivitiesByOrderUnique(db *sql.DB) ([]*AutoScalingGroupActivities, error) {
	rows, err := db.Query("select id,activity_id,description,auto_scaling_group_name,details,start_time,progress,end_time,cause,status_code from auto_scaling_activities order by activity_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfAuto := make([]*AutoScalingGroupActivities, 0, 0)
	for rows.Next() {
		auto, err := rowsNoFetchResultSetToAutoScalingGroupActivities(rows)
		if err == nil {
			arrayOfAuto = append(arrayOfAuto, auto)
		} else {
			log.Println(err)
		}
	}

	return arrayOfAuto, nil
}

func loadAllAutoScalingGroupActivities(db *sql.DB) ([]*AutoScalingGroupActivities, error) {
	rows, err := db.Query("select id,activity_id,description,auto_scaling_group_name,details,start_time,progress,end_time,cause,status_code from auto_scaling_activities order by start_time")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfAuto := make([]*AutoScalingGroupActivities, 0, 0)
	for rows.Next() {
		auto, err := rowsNoFetchResultSetToAutoScalingGroupActivities(rows)
		if err == nil {
			arrayOfAuto = append(arrayOfAuto, auto)
		} else {
			log.Println(err)
		}
	}

	return arrayOfAuto, nil
}

func loadAllAutoScalingGroupActivitiesOrderByASGName(db *sql.DB) ([]*AutoScalingGroupActivities, error) {
	rows, err := db.Query("select id,activity_id,description,auto_scaling_group_name,details,start_time,progress,end_time,cause,status_code from auto_scaling_activities order by auto_scaling_group_name,start_time")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfAuto := make([]*AutoScalingGroupActivities, 0, 0)
	for rows.Next() {
		auto, err := rowsNoFetchResultSetToAutoScalingGroupActivities(rows)
		if err == nil {
			arrayOfAuto = append(arrayOfAuto, auto)
		} else {
			log.Printf("loadAllAutoScalingGroupActivitiesOrderByASGName %+v\n", err)
		}
	}

	return arrayOfAuto, nil
}

func loadAllAutoScalingGroupActivitiesWithASGName(db *sql.DB, autoScalingGroupName string) ([]*AutoScalingGroupActivities, error) {
	rows, err := db.Query("select id,activity_id,description,auto_scaling_group_name,details,start_time,progress,end_time,cause,status_code from auto_scaling_activities where auto_scaling_group_name = $1 order by start_time", autoScalingGroupName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfAuto := make([]*AutoScalingGroupActivities, 0, 0)
	for rows.Next() {
		auto, err := rowsNoFetchResultSetToAutoScalingGroupActivities(rows)
		if err == nil {
			arrayOfAuto = append(arrayOfAuto, auto)
		} else {
			log.Printf("loadAllAutoScalingGroupActivitiesWithASGName %+v\n", err)
		}
	}

	return arrayOfAuto, nil
}

func updateAutoScalingGroupActivities(db *sql.DB, autoScalingGroup *AutoScalingGroupActivities, autoScalingGroupId int64) *AutoScalingGroupActivities {
	rows := db.QueryRow("update auto_scaling_activities set ref_auto_scaling_group_id=$1 where id=$2 returning id,activity_id,description,auto_scaling_group_name,details,start_time,progress,end_time,cause,status_code", autoScalingGroupId, autoScalingGroup.Id)

	autoScalingGroup, err := rowResultSetToAutoScalingGroupActivities(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return autoScalingGroup
}

func deleteAllAutoScalingGroupActivities(db *sql.DB) error {
	_, err := db.Exec("delete from auto_scaling_activities")

	return err
}
