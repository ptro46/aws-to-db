package main

import (
	"fmt"
)

type AutoScalingGroupActivities struct {
	Id                   int64
	ActivityId           string
	Description          string
	AutoScalingGroupName string
	Details              string
	StartTime            string
	Progress             int
	EndTime              string
	Cause                string
	StatusCode           string
}

func (d *AutoScalingGroupActivities) Dump() string {
	dumpAuto := fmt.Sprintf("ActivityId(%s) Description(%s) AutoScalingGroupName(%s) StartTime(%s) StatusCode(%s)",
		d.ActivityId,
		d.Description,
		d.AutoScalingGroupName,
		d.StartTime,
		d.StatusCode)

	return dumpAuto
}

func NewAutoScalingGroupActivities(Id int64, ActivityId string, Description string, AutoScalingGroupName string, Details string, StartTime string, Progress int, EndTime string, Cause string, StatusCode string) *AutoScalingGroupActivities {
	return &AutoScalingGroupActivities{Id: Id,
		ActivityId:           ActivityId,
		Description:          Description,
		AutoScalingGroupName: AutoScalingGroupName,
		Details:              Details,
		StartTime:            StartTime,
		Progress:             Progress,
		EndTime:              EndTime,
		Cause:                Cause,
		StatusCode:           StatusCode}
}
