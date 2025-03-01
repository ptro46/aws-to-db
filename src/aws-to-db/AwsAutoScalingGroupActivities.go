package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type AwsAutoScalingGroupActivity struct {
	ActivityId           string `json:"ActivityId,omitempty"`
	Description          string `json:"Description,omitempty"`
	AutoScalingGroupName string `json:"AutoScalingGroupName,omitempty"`
	Details              string `json:"Details,omitempty"`
	StartTime            string `json:"StartTime,omitempty"`
	Progress             int    `json:"Progress,omitempty"`
	EndTime              string `json:"EndTime,omitempty"`
	Cause                string `json:"Cause,omitempty"`
	StatusCode           string `json:"StatusCode,omitempty"`
}

type AwsAutoScalingGroupActivities struct {
	Activities []AwsAutoScalingGroupActivity `json:"Activities,omitempty"`
}

func awsAutoScalingGroupActivitiesParseAndPersist(db *sql.DB, jsonString string) error {
	fmt.Printf("awsAutoScalingGroupActivitiesParseAndPersist with json lenght %d\n", len(jsonString))
	awsAutoScalingGroupActivities := new(AwsAutoScalingGroupActivities)
	errUnmarshal := json.Unmarshal([]byte(jsonString), awsAutoScalingGroupActivities)
	if errUnmarshal == nil {
		for _, awsAutoScalingGroupActivity := range awsAutoScalingGroupActivities.Activities {
			fmt.Printf("AutoScalingGroupActivity %s\n", awsAutoScalingGroupActivity.ActivityId)
			autoGroup := createAutoScalingGroupActivities(db,
				awsAutoScalingGroupActivity.ActivityId,
				awsAutoScalingGroupActivity.Description,
				awsAutoScalingGroupActivity.AutoScalingGroupName,
				awsAutoScalingGroupActivity.Details,
				awsAutoScalingGroupActivity.StartTime,
				awsAutoScalingGroupActivity.Progress,
				awsAutoScalingGroupActivity.EndTime,
				awsAutoScalingGroupActivity.Cause,
				awsAutoScalingGroupActivity.StatusCode)
			if autoGroup != nil {
				fmt.Printf("    AutoScalingGroupActivity %s loaded\n", autoGroup.ActivityId)
			} else {
				fmt.Printf("  ERROR  AutoScalingGroupActivity %s not loaded\n", awsAutoScalingGroupActivity.ActivityId)
			}
		}
	} else {
		return errUnmarshal
	}

	return nil
}
