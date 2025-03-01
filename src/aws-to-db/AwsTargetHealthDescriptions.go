package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type AwsTargeHealthDescription struct {
	HealthCheckPort string          `json:"HealthCheckPort,omitempty"`
	Target          AwsTarget       `json:"Target,omitempty"`
	TargetHealth    AwsTargetHealth `json:"TargetHealth,omitempty"`
}

type AwsTarget struct {
	Id   string `json:"Id,omitempty"`
	Port int    `json:"Port,omitempty"`
}

type AwsTargetHealth struct {
	State       string `json:"State,omitempty"`
	Reason      string `json:"Reason,omitempty"`
	Description string `json:"Description,omitempty"`
}

type AwsTargeHealthDescriptions struct {
	TargetHealthDescriptions []AwsTargeHealthDescription `json:"TargetHealthDescriptions,omitempty"`
}

func awsTargeHealthDescription(db *sql.DB, jsonString string, targetGroupId int64) error {
	fmt.Printf("	awsTargeHealthDescription with json lenght %d\n", len(jsonString))
	awsTargeHealthDescriptions := new(AwsTargeHealthDescriptions)
	errUnmarshal := json.Unmarshal([]byte(jsonString), awsTargeHealthDescriptions)
	if errUnmarshal == nil {
		for _, awsHealth := range awsTargeHealthDescriptions.TargetHealthDescriptions {
			fmt.Printf("		Port (%s)\n", awsHealth.HealthCheckPort)
			health := createLoadBalancerTargetGroupHealth(db, targetGroupId, awsHealth.HealthCheckPort, awsHealth.Target.Id, awsHealth.Target.Port, awsHealth.TargetHealth.State, awsHealth.TargetHealth.Reason, awsHealth.TargetHealth.Description)
			if health != nil {
				fmt.Printf("    		TargeHealthDescription %s loaded\n", health.HealthCheckPort)
			} else {
				fmt.Printf("  ERROR  		TargeHealthDescription %s not loaded\n", awsHealth.HealthCheckPort)
			}
		}
	} else {
		return errUnmarshal
	}
	return nil
}
