package main

import (
	"database/sql" // package SQL
	"log"

	_ "github.com/lib/pq" // driver Postgres
)

func rowResultSetToViewElbTargetsHealth(row *sql.Row) (*ViewElbTargetsHealth, error) {
	var err error
	var Id int64
	var LoadBalancerArn string
	var LoadBalancerName string
	var LoadBalancerPort int64
	var LoadBalancerProtocol string
	var Type string
	var HealthCheckIntervalSeconds int
	var HealthCheckTimeoutSeconds int
	var HealthyThresholdCount int
	var UnhealthyThresholdCount int
	var HealthCheckPath string
	var Httpcode string
	var TargetType string
	var TargetHealthState string
	var InstanceId string
	var InstanceName string

	err = row.Scan(&Id, &LoadBalancerArn, &LoadBalancerName, &LoadBalancerPort, &LoadBalancerProtocol, &Type, &HealthCheckIntervalSeconds, &HealthCheckTimeoutSeconds, &HealthyThresholdCount, &UnhealthyThresholdCount, &HealthCheckPath, &Httpcode, &TargetType, &TargetHealthState, &InstanceId, &InstanceName)
	if err != nil {
		return nil, err
	}
	return NewViewElbTargetsHealth(Id, LoadBalancerArn, LoadBalancerName, LoadBalancerPort, LoadBalancerProtocol, Type, HealthCheckIntervalSeconds, HealthCheckTimeoutSeconds, HealthyThresholdCount, UnhealthyThresholdCount, HealthCheckPath, Httpcode, TargetType, TargetHealthState, InstanceId, InstanceName), nil
}

func rowsNoFetchResultSetToViewElbTargetsHealth(rows *sql.Rows) (*ViewElbTargetsHealth, error) {
	var err error
	var Id int64
	var LoadBalancerArn string
	var LoadBalancerName string
	var LoadBalancerPort int64
	var LoadBalancerProtocol string
	var Type string
	var HealthCheckIntervalSeconds int
	var HealthCheckTimeoutSeconds int
	var HealthyThresholdCount int
	var UnhealthyThresholdCount int
	var HealthCheckPath string
	var Httpcode string
	var TargetType string
	var TargetHealthState string
	var InstanceId string
	var InstanceName string

	err = rows.Scan(&Id, &LoadBalancerArn, &LoadBalancerName, &LoadBalancerPort, &LoadBalancerProtocol, &Type, &HealthCheckIntervalSeconds, &HealthCheckTimeoutSeconds, &HealthyThresholdCount, &UnhealthyThresholdCount, &HealthCheckPath, &Httpcode, &TargetType, &TargetHealthState, &InstanceId, &InstanceName)
	if err != nil {
		return nil, err
	}
	return NewViewElbTargetsHealth(Id, LoadBalancerArn, LoadBalancerName, LoadBalancerPort, LoadBalancerProtocol, Type, HealthCheckIntervalSeconds, HealthCheckTimeoutSeconds, HealthyThresholdCount, UnhealthyThresholdCount, HealthCheckPath, Httpcode, TargetType, TargetHealthState, InstanceId, InstanceName), nil
}

func rowsResultSetToViewElbTargetsHealth(rows *sql.Rows) (*ViewElbTargetsHealth, error) {
	var err error
	if rows.Next() {
		var Id int64
		var LoadBalancerArn string
		var LoadBalancerName string
		var LoadBalancerPort int64
		var LoadBalancerProtocol string
		var Type string
		var HealthCheckIntervalSeconds int
		var HealthCheckTimeoutSeconds int
		var HealthyThresholdCount int
		var UnhealthyThresholdCount int
		var HealthCheckPath string
		var Httpcode string
		var TargetType string
		var TargetHealthState string
		var InstanceId string
		var InstanceName string

		err = rows.Scan(&Id, &LoadBalancerArn, &LoadBalancerName, &LoadBalancerPort, &LoadBalancerProtocol, &Type, &HealthCheckIntervalSeconds, &HealthCheckTimeoutSeconds, &HealthyThresholdCount, &UnhealthyThresholdCount, &HealthCheckPath, &Httpcode, &TargetType, &TargetHealthState, &InstanceId, &InstanceName)
		if err != nil {
			return nil, err
		}
		return NewViewElbTargetsHealth(Id, LoadBalancerArn, LoadBalancerName, LoadBalancerPort, LoadBalancerProtocol, Type, HealthCheckIntervalSeconds, HealthCheckTimeoutSeconds, HealthyThresholdCount, UnhealthyThresholdCount, HealthCheckPath, Httpcode, TargetType, TargetHealthState, InstanceId, InstanceName), nil
	}
	return nil, err
}

func loadAllViewElbTargetsHealth(db *sql.DB) ([]*ViewElbTargetsHealth, error) {
	rows, err := db.Query("select id,load_balancer_arn,load_balancer_name,load_balancer_port,load_balancer_protocol,type,health_check_interval_seconds,health_check_timeout_seconds,healthy_threshold_count,unhealthy_threshold_count,health_check_path,http_code,target_type,target_health_state,instance_id,instance_name from public.view_elb_targets_health order by load_balancer_name,instance_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfViewElbTargetsHealth := make([]*ViewElbTargetsHealth, 0, 0)
	for rows.Next() {
		viewElbTargetsHealth, err := rowsNoFetchResultSetToViewElbTargetsHealth(rows)
		if err == nil {
			arrayOfViewElbTargetsHealth = append(arrayOfViewElbTargetsHealth, viewElbTargetsHealth)
		} else {
			log.Println(err)
		}
	}
	return arrayOfViewElbTargetsHealth, nil
}

func loadViewElbTargetsHealthByElbName(db *sql.DB, elbName string) ([]*ViewElbTargetsHealth, error) {
	rows, err := db.Query("select id,load_balancer_arn,load_balancer_name,load_balancer_port,load_balancer_protocol,type,health_check_interval_seconds,health_check_timeout_seconds,healthy_threshold_count,unhealthy_threshold_count,health_check_path,http_code,target_type,target_health_state,instance_id,instance_name from public.view_elb_targets_health where upper(load_balancer_name) like upper($1) order by id", elbName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfViewElbTargetsHealth := make([]*ViewElbTargetsHealth, 0, 0)
	for rows.Next() {
		viewElbTargetsHealth, err := rowsNoFetchResultSetToViewElbTargetsHealth(rows)
		if err == nil {
			arrayOfViewElbTargetsHealth = append(arrayOfViewElbTargetsHealth, viewElbTargetsHealth)
		} else {
			log.Println(err)
		}
	}
	return arrayOfViewElbTargetsHealth, nil
}
