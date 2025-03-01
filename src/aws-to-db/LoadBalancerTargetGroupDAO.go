package main

import (
	"database/sql" // package SQL
	"fmt"

	_ "github.com/lib/pq" // driver Postgres
)

//////////////////////////////////////////////////////////
//RowResult
//////////////////////////////////////////////////////////

func rowResultSetToLoadBalancerTargetGroup(row *sql.Row) (*LoadBalancerTargetGroup, error) {
	var err error
	var Id int64
	var TargetGroupArn string
	var TargetGroupName string
	var Protocol string
	var Port int
	var VpcId string
	var RefVpcId sql.NullInt64
	var HealthCheckProtocol string
	var HealthCheckPort string
	var HealthCheckEnabled bool
	var HealthCheckIntervalSeconds int
	var HealthCheckTimeoutSeconds int
	var HealthyThresholdCount int
	var UnhealthyThresholdCount int
	var HealthCheckPath string
	var HttpCode string
	var TargetType string

	err = row.Scan(&Id, &TargetGroupArn, &TargetGroupName, &Protocol, &Port, &VpcId, &RefVpcId, &HealthCheckProtocol, &HealthCheckPort, &HealthCheckEnabled, &HealthCheckIntervalSeconds,
		&HealthCheckTimeoutSeconds, &HealthyThresholdCount, &UnhealthyThresholdCount, &HealthCheckPath, &HttpCode, &TargetType)
	if err != nil {
		return nil, err
	}
	return NewLoadBalancerTargetGroup(Id, TargetGroupArn, TargetGroupName, Protocol, Port, VpcId, RefVpcId, HealthCheckProtocol, HealthCheckPort, HealthCheckEnabled, HealthCheckIntervalSeconds,
		HealthCheckTimeoutSeconds, HealthyThresholdCount, UnhealthyThresholdCount, HealthCheckPath, HttpCode, TargetType), nil
}

func rowResultSetToLoadBalancerTargetGroupLoadBalancerARN(row *sql.Row) (*LoadBalancerTargetGroupLoadBalancerARN, error) {
	var err error
	var Id int64
	var LoadBalancerTargetGroupId int64
	var LoadBalancerArn string
	var RefLoadBalancerId sql.NullInt64

	err = row.Scan(&Id, &LoadBalancerTargetGroupId, &LoadBalancerArn, &RefLoadBalancerId)
	if err != nil {
		return nil, err
	}
	return NewLoadBalancerTargetGroupLoadBalancerARN(Id, LoadBalancerTargetGroupId, LoadBalancerArn, RefLoadBalancerId), nil
}

func rowResultSetToLoadBalancerTargetGroupHealth(row *sql.Row) (*LoadBalancerTargetGroupHealth, error) {
	var err error
	var Id int64
	var LoadBalancerTargetGroupId int64
	var HealthCheckPort string
	var TargetId string
	var RefInstanceId sql.NullInt64
	var Port int
	var TargetHealthState string
	var TargetHealthReason string
	var TargetHealthDescription string

	err = row.Scan(&Id, &LoadBalancerTargetGroupId, &HealthCheckPort, &TargetId, &RefInstanceId, &Port, &TargetHealthState, &TargetHealthReason, &TargetHealthDescription)
	if err != nil {
		return nil, err
	}
	return NewLoadBalancerTargetGroupHealth(Id, LoadBalancerTargetGroupId, HealthCheckPort, TargetId, RefInstanceId, Port, TargetHealthState, TargetHealthReason, TargetHealthDescription), nil
}

//////////////////////////////////////////////////////////
//RowsNoFetch
//////////////////////////////////////////////////////////

func rowsNoFetchResultSetToLoadBalancerTargetGroup(rows *sql.Rows) (*LoadBalancerTargetGroup, error) {
	var err error
	var Id int64
	var TargetGroupArn string
	var TargetGroupName string
	var Protocol string
	var Port int
	var VpcId string
	var RefVpcId sql.NullInt64
	var HealthCheckProtocol string
	var HealthCheckPort string
	var HealthCheckEnabled bool
	var HealthCheckIntervalSeconds int
	var HealthCheckTimeoutSeconds int
	var HealthyThresholdCount int
	var UnhealthyThresholdCount int
	var HealthCheckPath string
	var HttpCode string
	var TargetType string

	err = rows.Scan(&Id, &TargetGroupArn, &TargetGroupName, &Protocol, &Port, &VpcId, &RefVpcId, &HealthCheckProtocol, &HealthCheckPort, &HealthCheckEnabled, &HealthCheckIntervalSeconds,
		&HealthCheckTimeoutSeconds, &HealthyThresholdCount, &UnhealthyThresholdCount, &HealthCheckPath, &HttpCode, &TargetType)
	if err != nil {
		return nil, err
	}
	return NewLoadBalancerTargetGroup(Id, TargetGroupArn, TargetGroupName, Protocol, Port, VpcId, RefVpcId, HealthCheckProtocol, HealthCheckPort, HealthCheckEnabled, HealthCheckIntervalSeconds,
		HealthCheckTimeoutSeconds, HealthyThresholdCount, UnhealthyThresholdCount, HealthCheckPath, HttpCode, TargetType), nil

}

func rowsNoFetchResultSetToLoadBalancerTargetGroupLoadBalancerARN(rows *sql.Rows) (*LoadBalancerTargetGroupLoadBalancerARN, error) {
	var err error
	var Id int64
	var LoadBalancerTargetGroupId int64
	var LoadBalancerArn string
	var RefLoadBalancerId sql.NullInt64

	err = rows.Scan(&Id, &LoadBalancerTargetGroupId, &LoadBalancerArn, &RefLoadBalancerId)
	if err != nil {
		return nil, err
	}
	return NewLoadBalancerTargetGroupLoadBalancerARN(Id, LoadBalancerTargetGroupId, LoadBalancerArn, RefLoadBalancerId), nil
}

func rowsNoFetchResultSetToLoadBalancerTargetGroupHealth(rows *sql.Rows) (*LoadBalancerTargetGroupHealth, error) {
	var err error
	var Id int64
	var LoadBalancerTargetGroupId int64
	var HealthCheckPort string
	var TargetId string
	var RefInstanceId sql.NullInt64
	var Port int
	var TargetHealthState string
	var TargetHealthReason string
	var TargetHealthDescription string

	err = rows.Scan(&Id, &LoadBalancerTargetGroupId, &HealthCheckPort, &TargetId, &RefInstanceId, &Port, &TargetHealthState, &TargetHealthReason, &TargetHealthDescription)
	if err != nil {
		return nil, err
	}
	return NewLoadBalancerTargetGroupHealth(Id, LoadBalancerTargetGroupId, HealthCheckPort, TargetId, RefInstanceId, Port, TargetHealthState, TargetHealthReason, TargetHealthDescription), nil
}

//////////////////////////////////////////////////////////
//RowsResult
//////////////////////////////////////////////////////////

func rowsResultSetToLoadBalancerTargetGroup(rows *sql.Rows) (*LoadBalancerTargetGroup, error) {
	var err error
	if rows.Next() {
		var err error
		var Id int64
		var TargetGroupArn string
		var TargetGroupName string
		var Protocol string
		var Port int
		var VpcId string
		var RefVpcId sql.NullInt64
		var HealthCheckProtocol string
		var HealthCheckPort string
		var HealthCheckEnabled bool
		var HealthCheckIntervalSeconds int
		var HealthCheckTimeoutSeconds int
		var HealthyThresholdCount int
		var UnhealthyThresholdCount int
		var HealthCheckPath string
		var HttpCode string
		var TargetType string

		err = rows.Scan(&Id, &TargetGroupArn, &TargetGroupName, &Protocol, &Port, &VpcId, &RefVpcId, &HealthCheckProtocol, &HealthCheckPort, &HealthCheckEnabled, &HealthCheckIntervalSeconds,
			&HealthCheckTimeoutSeconds, &HealthyThresholdCount, &UnhealthyThresholdCount, &HealthCheckPath, &HttpCode, &TargetType)
		if err != nil {
			return nil, err
		}
		return NewLoadBalancerTargetGroup(Id, TargetGroupArn, TargetGroupName, Protocol, Port, VpcId, RefVpcId, HealthCheckProtocol, HealthCheckPort, HealthCheckEnabled, HealthCheckIntervalSeconds,
			HealthCheckTimeoutSeconds, HealthyThresholdCount, UnhealthyThresholdCount, HealthCheckPath, HttpCode, TargetType), nil

	}
	return nil, err
}

func rowsResultSetToLoadBalancerTargetGroupLoadBalancerARN(rows *sql.Rows) (*LoadBalancerTargetGroupLoadBalancerARN, error) {
	var err error
	if rows.Next() {
		var Id int64
		var LoadBalancerTargetGroupId int64
		var LoadBalancerArn string
		var RefLoadBalancerId sql.NullInt64

		err = rows.Scan(&Id, &LoadBalancerTargetGroupId, &LoadBalancerArn, &RefLoadBalancerId)
		if err != nil {
			return nil, err
		}
		return NewLoadBalancerTargetGroupLoadBalancerARN(Id, LoadBalancerTargetGroupId, LoadBalancerArn, RefLoadBalancerId), nil
	}
	return nil, err
}

func rowsResultSetToLoadBalancerTargetGroupHealth(rows *sql.Rows) (*LoadBalancerTargetGroupHealth, error) {
	var err error
	if rows.Next() {
		var Id int64
		var LoadBalancerTargetGroupId int64
		var HealthCheckPort string
		var TargetId string
		var RefInstanceId sql.NullInt64
		var Port int
		var TargetHealthState string
		var TargetHealthReason string
		var TargetHealthDescription string

		err = rows.Scan(&Id, &LoadBalancerTargetGroupId, &HealthCheckPort, &TargetId, &RefInstanceId, &Port, &TargetHealthState, &TargetHealthReason, &TargetHealthDescription)
		if err != nil {
			return nil, err
		}
		return NewLoadBalancerTargetGroupHealth(Id, LoadBalancerTargetGroupId, HealthCheckPort, TargetId, RefInstanceId, Port, TargetHealthState, TargetHealthReason, TargetHealthDescription), nil
	}
	return nil, err
}

//////////////////////////////////////////////////////////
//Create
//////////////////////////////////////////////////////////

func createLoadBalancerTargetGroup(db *sql.DB, TargetGroupArn string, TargetGroupName string, Protocol string, Port int, VpcId string, HealthCheckProtocol string, HealthCheckPort string, HealthCheckEnabled bool, HealthCheckIntervalSeconds int, HealthCheckTimeoutSeconds int, HealthyThresholdCount int, UnhealthyThresholdCount int, HealthCheckPath string, HttpCode string, TargetType string) *LoadBalancerTargetGroup {
	rows := db.QueryRow("insert into load_balancer_target_group(target_group_arn,target_group_name,protocol,port,vpc_id,health_check_protocol,health_check_port,health_check_enabled,health_check_interval_seconds,health_check_timeout_seconds,healthy_threshold_count,unhealthy_threshold_count,health_check_path,http_code,target_type) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15) returning id,target_group_arn,target_group_name,protocol,port,vpc_id,ref_vpc_id,health_check_protocol,health_check_port,health_check_enabled,health_check_interval_seconds,health_check_timeout_seconds,healthy_threshold_count,unhealthy_threshold_count,health_check_path,http_code,target_type",
		TargetGroupArn, TargetGroupName, Protocol, Port, VpcId, HealthCheckProtocol, HealthCheckPort, HealthCheckEnabled, HealthCheckIntervalSeconds, HealthCheckTimeoutSeconds, HealthyThresholdCount, UnhealthyThresholdCount, HealthCheckPath, HttpCode, TargetType)

	loadBalancerTargetGroup, err := rowResultSetToLoadBalancerTargetGroup(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return loadBalancerTargetGroup
}

func createLoadBalancerTargetGroupLoadBalancerARN(db *sql.DB, LoadBalancerTargetGroupId int64, LoadBalancerArn string) *LoadBalancerTargetGroupLoadBalancerARN {
	rows := db.QueryRow("insert into load_balancer_target_group_load_balancer_arn(load_balancer_target_group_id,load_balancer_arn) values($1,$2) returning id,load_balancer_target_group_id,load_balancer_arn,ref_load_balancer_id", LoadBalancerTargetGroupId, LoadBalancerArn)

	loadBalancerTargetGroupLoadBalancerARN, err := rowResultSetToLoadBalancerTargetGroupLoadBalancerARN(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return loadBalancerTargetGroupLoadBalancerARN
}

func createLoadBalancerTargetGroupHealth(db *sql.DB, LoadBalancerTargetGroupId int64, HealthCheckPort string, TargetId string, Port int, TargetHealthState string, TargetHealthReason string, TargetHealthDescription string) *LoadBalancerTargetGroupHealth {
	rows := db.QueryRow("insert into load_balancer_target_group_health(load_balancer_target_group_id,health_check_port,target_id,port,target_health_state,target_health_reason,target_health_description) values($1,$2,$3,$4,$5,$6,$7) returning id,load_balancer_target_group_id,health_check_port,target_id,ref_instance_id,port,target_health_state,target_health_reason,target_health_description",
		LoadBalancerTargetGroupId, HealthCheckPort, TargetId, Port, TargetHealthState, TargetHealthReason, TargetHealthDescription)

	loadBalancerTargetGroupHealth, err := rowResultSetToLoadBalancerTargetGroupHealth(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return loadBalancerTargetGroupHealth
}

//////////////////////////////////////////////////////////
//Load
//////////////////////////////////////////////////////////

func loadLoadBalancerTargetGroupByTargetGroupARN(db *sql.DB, target_group_arn string) (*LoadBalancerTargetGroup, error) {
	sqlSelect := "SELECT id,target_group_arn,target_group_name,protocol,port,vpc_id,ref_vpc_id,health_check_protocol,health_check_port,health_check_enabled,health_check_interval_seconds,health_check_timeout_seconds,healthy_threshold_count,unhealthy_threshold_count,health_check_path,http_code,target_type from load_balancer_target_group where target_group_arn=$1"

	rows, err := db.Query(sqlSelect, target_group_arn)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	loadBalancerTargetGroup, err := rowsResultSetToLoadBalancerTargetGroup(rows)
	if err != nil {
		return nil, err
	}
	return loadBalancerTargetGroup, nil
}

func loadAllLoadBalancerTargetGroupByOrderUnique(db *sql.DB) ([]*LoadBalancerTargetGroup, error) {
	results := make([]*LoadBalancerTargetGroup, 0, 0)
	sqlSelect := "SELECT id,target_group_arn,target_group_name,protocol,port,vpc_id,ref_vpc_id,health_check_protocol,health_check_port,health_check_enabled,health_check_interval_seconds,health_check_timeout_seconds,healthy_threshold_count,unhealthy_threshold_count,health_check_path,http_code,target_type from load_balancer_target_group order by port"

	rows, err := db.Query(sqlSelect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		loadBalancerTargetGroup, err := rowsNoFetchResultSetToLoadBalancerTargetGroup(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, loadBalancerTargetGroup)
	}
	return results, nil
}

func loadAllLoadBalancerTargetGroupLoadBalancerARNByOrderUnique(db *sql.DB) ([]*LoadBalancerTargetGroupLoadBalancerARN, error) {
	results := make([]*LoadBalancerTargetGroupLoadBalancerARN, 0, 0)
	sqlSelect := "SELECT id,load_balancer_target_group_id,load_balancer_arn,ref_load_balancer_id from load_balancer_target_group_load_balancer_arn order by load_balancer_arn"

	rows, err := db.Query(sqlSelect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		loadBalancerTargetGroupLoadBalancerARN, err := rowsNoFetchResultSetToLoadBalancerTargetGroupLoadBalancerARN(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, loadBalancerTargetGroupLoadBalancerARN)
	}
	return results, nil
}

func loadAllLoadBalancerTargetGroupHealthByOrderUnique(db *sql.DB) ([]*LoadBalancerTargetGroupHealth, error) {
	results := make([]*LoadBalancerTargetGroupHealth, 0, 0)
	sqlSelect := "SELECT id,load_balancer_target_group_id,health_check_port,target_id,ref_instance_id,port,target_health_state,target_health_reason,target_health_description from load_balancer_target_group_health order by port"

	rows, err := db.Query(sqlSelect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		loadBalancerTargetGroupHealth, err := rowsNoFetchResultSetToLoadBalancerTargetGroupHealth(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, loadBalancerTargetGroupHealth)
	}
	return results, nil
}

func updateLoadBalancerTargetGroupId(db *sql.DB, loadBalancerTargetGroup *LoadBalancerTargetGroup, vpcId int64) *LoadBalancerTargetGroup {
	rows := db.QueryRow("update load_balancer_target_group set ref_vpc_id=$1 where id=$2 returning id,target_group_arn,target_group_name,protocol,port,vpc_id,ref_vpc_id,health_check_protocol,health_check_port,health_check_enabled,health_check_interval_seconds,health_check_timeout_seconds,healthy_threshold_count,unhealthy_threshold_count,health_check_path,http_code,target_type",
		vpcId, loadBalancerTargetGroup.Id)

	loadBalancerTargetGroup, err := rowResultSetToLoadBalancerTargetGroup(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return loadBalancerTargetGroup
}

func updateLoadBalancerTargetGroupLoadBalancerARNId(db *sql.DB, loadBalancerTargetGroupLoadBalancerARN *LoadBalancerTargetGroupLoadBalancerARN, loadBalancerId int64) *LoadBalancerTargetGroupLoadBalancerARN {
	rows := db.QueryRow("update load_balancer_target_group_load_balancer_arn set ref_load_balancer_id=$1 where id=$2 returning id,load_balancer_target_group_id,load_balancer_arn,ref_load_balancer_id",
		loadBalancerId, loadBalancerTargetGroupLoadBalancerARN.Id)
	loadBalancerTargetGroupLoadBalancerARN, err := rowResultSetToLoadBalancerTargetGroupLoadBalancerARN(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return loadBalancerTargetGroupLoadBalancerARN
}

func updateLoadBalancerTargetGroupHealthRefInstanceId(db *sql.DB, loadBalancerTargetGroupHealth *LoadBalancerTargetGroupHealth, refInstanceId int64) (*LoadBalancerTargetGroupHealth, error) {
	row := db.QueryRow("update load_balancer_target_group_health set ref_instance_id=$1 where id=$2 returning id,load_balancer_target_group_id,health_check_port,target_id,ref_instance_id,port,target_health_state,target_health_reason,target_health_description", refInstanceId, loadBalancerTargetGroupHealth.Id)
	loadBalancerTargetGroupHealth, err := rowResultSetToLoadBalancerTargetGroupHealth(row)
	if err != nil {
		return nil, err
	}
	return loadBalancerTargetGroupHealth, nil
}

//////////////////////////////////////////////////////////
//DeleteAll
//////////////////////////////////////////////////////////

func deleteAllLoadBalancerTargetGroup(db *sql.DB) error {
	_, err := db.Exec("delete from load_balancer_target_group")

	return err
}
func deleteAllLoadBalancerTargetGroupHealth(db *sql.DB) error {
	_, err := db.Exec("delete from load_balancer_target_group_health")

	return err
}

func deleteAllLoadBalancerTargetGroupLoadBalancerARN(db *sql.DB) error {
	_, err := db.Exec("delete from load_balancer_target_group_load_balancer_arn")

	return err
}
