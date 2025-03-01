package main

import (
	"database/sql" // package SQL
	"fmt"
	"log"

	_ "github.com/lib/pq" // driver Postgres
)

//////////////////////////////////////////////////////////
//RowResult
//////////////////////////////////////////////////////////

func rowResultSetToAutoScalingGroupLauchTemplateOverride(row *sql.Row) (*AutoScalingGroupLauchTemplateOverride, error) {
	var err error
	var Id int64
	var AutoScalingGroupId int64
	var InstanceType string

	err = row.Scan(&Id, &AutoScalingGroupId, &InstanceType)
	if err != nil {
		return nil, err
	}
	return NewAutoScalingGroupLauchTemplateOverride(Id, AutoScalingGroupId, InstanceType), nil
}

func rowResultSetToAutoScalingGroupAvailabilityZone(row *sql.Row) (*AutoScalingGroupAvailabilityZone, error) {
	var err error
	var Id int64
	var AutoScalingGroupId int64
	var AvailabilityZone string

	err = row.Scan(&Id, &AutoScalingGroupId, &AvailabilityZone)
	if err != nil {
		return nil, err
	}
	return NewAutoScalingGroupAvailabilityZone(Id, AutoScalingGroupId, AvailabilityZone), nil
}

func rowResultSetToAutoScalingGroupLoadBalancerName(row *sql.Row) (*AutoScalingGroupLoadBalancerName, error) {
	var err error
	var Id int64
	var AutoScalingGroupId int64
	var LoadBalancerName string

	err = row.Scan(&Id, &AutoScalingGroupId, &LoadBalancerName)
	if err != nil {
		return nil, err
	}
	return NewAutoScalingGroupLoadBalancerName(Id, AutoScalingGroupId, LoadBalancerName), nil
}

func rowResultSetToAutoScalingGroupTargetGroupARN(row *sql.Row) (*AutoScalingGroupTargetGroupARN, error) {
	var err error
	var Id int64
	var AutoScalingGroupId int64
	var ArnName string

	err = row.Scan(&Id, &AutoScalingGroupId, &ArnName)
	if err != nil {
		return nil, err
	}
	return NewAutoScalingGroupTargetGroupARN(Id, AutoScalingGroupId, ArnName), nil
}

func rowResultSetToAutoScalingInstance(row *sql.Row) (*AutoScalingInstance, error) {
	var err error
	var Id int64
	var AutoScalingGroupId int64
	var InstanceId string
	var RefInstanceId sql.NullFloat64
	var AvailabilityZone string
	var LifecycleState string
	var HealthStatus string
	var LaunchConfigurationName string
	var LaunchTemplateId string
	var LaunchTemplateName string
	var LaunchTemplateVersion string
	var ProtectedFromScaleIn bool

	err = row.Scan(&Id, &AutoScalingGroupId, &InstanceId, &RefInstanceId, &AvailabilityZone, &LifecycleState, &HealthStatus,
		&LaunchConfigurationName, &LaunchTemplateId, &LaunchTemplateName, &LaunchTemplateVersion, &ProtectedFromScaleIn)
	if err != nil {
		return nil, err
	}
	return NewAutoScalingInstance(Id, AutoScalingGroupId, InstanceId, RefInstanceId, AvailabilityZone, LifecycleState, HealthStatus,
		LaunchConfigurationName, LaunchTemplateId, LaunchTemplateName, LaunchTemplateVersion, ProtectedFromScaleIn), nil
}

func rowResultSetToAutoScalingGroupsSuspendedProcesses(row *sql.Row) (*AutoScalingGroupsSuspendedProcesses, error) {
	var err error
	var Id int64
	var AutoScalingGroupId int64
	var ProcessName string
	var SuspensionReason string

	err = row.Scan(&Id, &AutoScalingGroupId, &ProcessName, &SuspensionReason)
	if err != nil {
		return nil, err
	}
	return NewAutoScalingGroupsSuspendedProcesses(Id, AutoScalingGroupId, ProcessName, SuspensionReason), nil
}

func rowResultSetToAutoScalingGroupsEnabledMetric(row *sql.Row) (*AutoScalingGroupsEnabledMetric, error) {
	var err error
	var Id int64
	var AutoScalingGroupId int64
	var Metric string
	var Granularity string

	err = row.Scan(&Id, &AutoScalingGroupId, &Metric, &Granularity)
	if err != nil {
		return nil, err
	}
	return NewAutoScalingGroupsEnabledMetric(Id, AutoScalingGroupId, Metric, Granularity), nil
}

func rowResultSetToAutoScalingGroupsTag(row *sql.Row) (*AutoScalingGroupsTag, error) {
	var err error
	var Id int64
	var AutoScalingGroupId int64
	var ResourceId string
	var ResourceType string
	var Key string
	var Value string
	var PropagateAtLaunch bool

	err = row.Scan(&Id, &AutoScalingGroupId, &ResourceId, &ResourceType, &Key, &Value, &PropagateAtLaunch)
	if err != nil {
		return nil, err
	}
	return NewAutoScalingGroupsTag(Id, AutoScalingGroupId, ResourceId, ResourceType, Key, Value, PropagateAtLaunch), nil
}

func rowResultSetToAutoScalingGroupsTerminationPolicy(row *sql.Row) (*AutoScalingGroupsTerminationPolicy, error) {
	var err error
	var Id int64
	var AutoScalingGroupId int64
	var TerminationPolicy string

	err = row.Scan(&Id, &AutoScalingGroupId, &TerminationPolicy)
	if err != nil {
		return nil, err
	}
	return NewAutoScalingGroupsTerminationPolicy(Id, AutoScalingGroupId, TerminationPolicy), nil
}

//////////////////////////////////////////////////////////
//RowsNoFetch
//////////////////////////////////////////////////////////

func rowsNoFetchResultSetToAutoScalingGroupLauchTemplateOverride(rows *sql.Rows) (*AutoScalingGroupLauchTemplateOverride, error) {
	var err error
	var Id int64
	var AutoScalingGroupId int64
	var InstanceType string

	err = rows.Scan(&Id, &AutoScalingGroupId, &InstanceType)
	if err != nil {
		return nil, err
	}
	return NewAutoScalingGroupLauchTemplateOverride(Id, AutoScalingGroupId, InstanceType), nil
}

func rowsNoFetchResultSetToAutoScalingGroupAvailabilityZone(rows *sql.Rows) (*AutoScalingGroupAvailabilityZone, error) {
	var err error
	var Id int64
	var AutoScalingGroupId int64
	var AvailabilityZone string

	err = rows.Scan(&Id, &AutoScalingGroupId, &AvailabilityZone)
	if err != nil {
		return nil, err
	}
	return NewAutoScalingGroupAvailabilityZone(Id, AutoScalingGroupId, AvailabilityZone), nil
}

func rowsNoFetchResultSetToAutoScalingGroupLoadBalancerName(rows *sql.Rows) (*AutoScalingGroupLoadBalancerName, error) {
	var err error
	var Id int64
	var AutoScalingGroupId int64
	var LoadBalancerName string

	err = rows.Scan(&Id, &AutoScalingGroupId, &LoadBalancerName)
	if err != nil {
		return nil, err
	}
	return NewAutoScalingGroupLoadBalancerName(Id, AutoScalingGroupId, LoadBalancerName), nil
}

func rowsNoFetchResultSetToAutoScalingGroupTargetGroupARN(rows *sql.Rows) (*AutoScalingGroupTargetGroupARN, error) {
	var err error
	var Id int64
	var AutoScalingGroupId int64
	var ArnName string

	err = rows.Scan(&Id, &AutoScalingGroupId, &ArnName)
	if err != nil {
		return nil, err
	}
	return NewAutoScalingGroupTargetGroupARN(Id, AutoScalingGroupId, ArnName), nil
}

func rowsNoFetchResultSetToAutoScalingInstance(rows *sql.Rows) (*AutoScalingInstance, error) {
	var err error
	var Id int64
	var AutoScalingGroupId int64
	var InstanceId string
	var RefInstanceId sql.NullFloat64
	var AvailabilityZone string
	var LifecycleState string
	var HealthStatus string
	var LaunchConfigurationName string
	var LaunchTemplateId string
	var LaunchTemplateName string
	var LaunchTemplateVersion string
	var ProtectedFromScaleIn bool

	err = rows.Scan(&Id, &AutoScalingGroupId, &InstanceId, &RefInstanceId, &AvailabilityZone, &LifecycleState, &HealthStatus,
		&LaunchConfigurationName, &LaunchTemplateId, &LaunchTemplateName, &LaunchTemplateVersion, &ProtectedFromScaleIn)
	if err != nil {
		return nil, err
	}
	return NewAutoScalingInstance(Id, AutoScalingGroupId, InstanceId, RefInstanceId, AvailabilityZone, LifecycleState, HealthStatus,
		LaunchConfigurationName, LaunchTemplateId, LaunchTemplateName, LaunchTemplateVersion, ProtectedFromScaleIn), nil
}

func rowsNoFetchResultSetToAutoScalingGroupsSuspendedProcesses(rows *sql.Rows) (*AutoScalingGroupsSuspendedProcesses, error) {
	var err error
	var Id int64
	var AutoScalingGroupId int64
	var ProcessName string
	var SuspensionReason string

	err = rows.Scan(&Id, &AutoScalingGroupId, &ProcessName, &SuspensionReason)
	if err != nil {
		return nil, err
	}
	return NewAutoScalingGroupsSuspendedProcesses(Id, AutoScalingGroupId, ProcessName, SuspensionReason), nil
}

func rowsNoFetchResultSetToAutoScalingGroupsEnabledMetric(rows *sql.Rows) (*AutoScalingGroupsEnabledMetric, error) {
	var err error
	var Id int64
	var AutoScalingGroupId int64
	var Metric string
	var Granularity string

	err = rows.Scan(&Id, &AutoScalingGroupId, &Metric, &Granularity)
	if err != nil {
		return nil, err
	}
	return NewAutoScalingGroupsEnabledMetric(Id, AutoScalingGroupId, Metric, Granularity), nil
}

func rowsNoFetchResultSetToAutoScalingGroupsTag(rows *sql.Rows) (*AutoScalingGroupsTag, error) {
	var err error
	var Id int64
	var AutoScalingGroupId int64
	var ResourceId string
	var ResourceType string
	var Key string
	var Value string
	var PropagateAtLaunch bool

	err = rows.Scan(&Id, &AutoScalingGroupId, &ResourceId, &ResourceType, &Key, &Value, &PropagateAtLaunch)
	if err != nil {
		return nil, err
	}
	return NewAutoScalingGroupsTag(Id, AutoScalingGroupId, ResourceId, ResourceType, Key, Value, PropagateAtLaunch), nil
}

func rowsNoFetchResultSetToAutoScalingGroupsTerminationPolicy(rows *sql.Rows) (*AutoScalingGroupsTerminationPolicy, error) {
	var err error
	var Id int64
	var AutoScalingGroupId int64
	var TerminationPolicy string

	err = rows.Scan(&Id, &AutoScalingGroupId, &TerminationPolicy)
	if err != nil {
		return nil, err
	}
	return NewAutoScalingGroupsTerminationPolicy(Id, AutoScalingGroupId, TerminationPolicy), nil
}

//////////////////////////////////////////////////////////
//RowsResult
//////////////////////////////////////////////////////////

func rowsResultSetToAutoScalingGroupLauchTemplateOverride(rows *sql.Rows) (*AutoScalingGroupLauchTemplateOverride, error) {
	var err error

	if rows.Next() {
		var Id int64
		var AutoScalingGroupId int64
		var InstanceType string

		err = rows.Scan(&Id, &AutoScalingGroupId, &InstanceType)
		if err != nil {
			return nil, err
		}
		return NewAutoScalingGroupLauchTemplateOverride(Id, AutoScalingGroupId, InstanceType), nil
	}

	return nil, err

}

func rowsResultSetToAutoScalingGroupAvailabilityZone(rows *sql.Rows) (*AutoScalingGroupAvailabilityZone, error) {
	var err error
	if rows.Next() {
		var Id int64
		var AutoScalingGroupId int64
		var AvailabilityZone string

		err = rows.Scan(&Id, &AutoScalingGroupId, &AvailabilityZone)
		if err != nil {
			return nil, err
		}
		return NewAutoScalingGroupAvailabilityZone(Id, AutoScalingGroupId, AvailabilityZone), nil
	}

	return nil, err

}

func rowsResultSetToAutoScalingGroupLoadBalancerName(rows *sql.Rows) (*AutoScalingGroupLoadBalancerName, error) {
	var err error
	if rows.Next() {
		var Id int64
		var AutoScalingGroupId int64
		var LoadBalancerName string

		err = rows.Scan(&Id, &AutoScalingGroupId, &LoadBalancerName)
		if err != nil {
			return nil, err
		}
		return NewAutoScalingGroupLoadBalancerName(Id, AutoScalingGroupId, LoadBalancerName), nil
	}

	return nil, err

}

func rowsResultSetToAutoScalingGroupTargetGroupARN(rows *sql.Rows) (*AutoScalingGroupTargetGroupARN, error) {
	var err error
	if rows.Next() {
		var Id int64
		var AutoScalingGroupId int64
		var ArnName string

		err = rows.Scan(&Id, &AutoScalingGroupId, &ArnName)
		if err != nil {
			return nil, err
		}
		return NewAutoScalingGroupTargetGroupARN(Id, AutoScalingGroupId, ArnName), nil
	}

	return nil, err

}

func rowsResultSetToAutoScalingInstance(rows *sql.Rows) (*AutoScalingInstance, error) {
	var err error
	if rows.Next() {
		var Id int64
		var AutoScalingGroupId int64
		var InstanceId string
		var RefInstanceId sql.NullFloat64
		var AvailabilityZone string
		var LifecycleState string
		var HealthStatus string
		var LaunchConfigurationName string
		var LaunchTemplateId string
		var LaunchTemplateName string
		var LaunchTemplateVersion string
		var ProtectedFromScaleIn bool

		err = rows.Scan(&Id, &AutoScalingGroupId, &InstanceId, &RefInstanceId, &AvailabilityZone, &LifecycleState, &HealthStatus,
			&LaunchConfigurationName, &LaunchTemplateId, &LaunchTemplateName, &LaunchTemplateVersion, &ProtectedFromScaleIn)
		if err != nil {
			return nil, err
		}
		return NewAutoScalingInstance(Id, AutoScalingGroupId, InstanceId, RefInstanceId, AvailabilityZone, LifecycleState, HealthStatus,
			LaunchConfigurationName, LaunchTemplateId, LaunchTemplateName, LaunchTemplateVersion, ProtectedFromScaleIn), nil
	}

	return nil, err

}

func rowsResultSetToAutoScalingGroupsSuspendedProcesses(rows *sql.Rows) (*AutoScalingGroupsSuspendedProcesses, error) {
	var err error
	if rows.Next() {
		var Id int64
		var AutoScalingGroupId int64
		var ProcessName string
		var SuspensionReason string

		err = rows.Scan(&Id, &AutoScalingGroupId, &ProcessName, &SuspensionReason)
		if err != nil {
			return nil, err
		}
		return NewAutoScalingGroupsSuspendedProcesses(Id, AutoScalingGroupId, ProcessName, SuspensionReason), nil
	}

	return nil, err

}

func rowsResultSetToAutoScalingGroupsEnabledMetric(rows *sql.Rows) (*AutoScalingGroupsEnabledMetric, error) {
	var err error
	if rows.Next() {
		var Id int64
		var AutoScalingGroupId int64
		var Metric string
		var Granularity string

		err = rows.Scan(&Id, &AutoScalingGroupId, &Metric, &Granularity)
		if err != nil {
			return nil, err
		}
		return NewAutoScalingGroupsEnabledMetric(Id, AutoScalingGroupId, Metric, Granularity), nil
	}

	return nil, err

}

func rowsResultSetToAutoScalingGroupsTag(rows *sql.Rows) (*AutoScalingGroupsTag, error) {
	var err error
	if rows.Next() {
		var Id int64
		var AutoScalingGroupId int64
		var ResourceId string
		var ResourceType string
		var Key string
		var Value string
		var PropagateAtLaunch bool

		err = rows.Scan(&Id, &AutoScalingGroupId, &ResourceId, &ResourceType, &Key, &Value, &PropagateAtLaunch)
		if err != nil {
			return nil, err
		}
		return NewAutoScalingGroupsTag(Id, AutoScalingGroupId, ResourceId, ResourceType, Key, Value, PropagateAtLaunch), nil
	}

	return nil, err

}

func rowsResultSetToAutoScalingGroupsTerminationPolicy(rows *sql.Rows) (*AutoScalingGroupsTerminationPolicy, error) {
	var err error
	if rows.Next() {
		var Id int64
		var AutoScalingGroupId int64
		var TerminationPolicy string

		err = rows.Scan(&Id, &AutoScalingGroupId, &TerminationPolicy)
		if err != nil {
			return nil, err
		}
		return NewAutoScalingGroupsTerminationPolicy(Id, AutoScalingGroupId, TerminationPolicy), nil
	}

	return nil, err

}

//////////////////////////////////////////////////////////
//Create
//////////////////////////////////////////////////////////

func createAutoScalingGroupLauchTemplateOverride(db *sql.DB, AutoScalingGroupId int64, InstanceType string) *AutoScalingGroupLauchTemplateOverride {
	rows := db.QueryRow("insert into auto_scaling_group_launch_template_override(auto_scaling_group_id,instance_type) values($1,$2) returning id,auto_scaling_group_id,instance_type",
		AutoScalingGroupId, InstanceType)

	autoScalingGroupLauchTemplateOverride, err := rowResultSetToAutoScalingGroupLauchTemplateOverride(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return autoScalingGroupLauchTemplateOverride
}

func createAutoScalingGroupAvailabilityZone(db *sql.DB, AutoScalingGroupId int64, AvailabilityZone string) *AutoScalingGroupAvailabilityZone {
	rows := db.QueryRow("insert into auto_scaling_group_availability_zone(auto_scaling_group_id,availability_zone) values($1,$2) returning id,auto_scaling_group_id,availability_zone",
		AutoScalingGroupId, AvailabilityZone)

	autoScalingGroupAvailabilityZone, err := rowResultSetToAutoScalingGroupAvailabilityZone(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return autoScalingGroupAvailabilityZone
}

func createAutoScalingGroupLoadBalancerName(db *sql.DB, AutoScalingGroupId int64, LoadBalancerName string) *AutoScalingGroupLoadBalancerName {
	rows := db.QueryRow("insert into auto_scaling_group_load_balancer_name(auto_scaling_group_id,load_balancer_name) values($1,$2) returning id,auto_scaling_group_id,load_balancer_name",
		AutoScalingGroupId, LoadBalancerName)

	autoScalingGroupLoadBalancerName, err := rowResultSetToAutoScalingGroupLoadBalancerName(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return autoScalingGroupLoadBalancerName
}

func createAutoScalingGroupTargetGroupARN(db *sql.DB, AutoScalingGroupId int64, ArnName string) *AutoScalingGroupTargetGroupARN {
	rows := db.QueryRow("insert into auto_scaling_group_target_group_arn(auto_scaling_group_id,arn_name) values($1,$2) returning id,auto_scaling_group_id,arn_name",
		AutoScalingGroupId, ArnName)

	sutoScalingGroupTargetGroupARN, err := rowResultSetToAutoScalingGroupTargetGroupARN(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return sutoScalingGroupTargetGroupARN
}

func createAutoScalingInstance(db *sql.DB, AutoScalingGroupId int64, InstanceId string, AvailabilityZone string, LifecycleState string, HealthStatus string, LaunchConfigurationName string, LaunchTemplateId string, LaunchTemplateName string, LaunchTemplateVersion string, ProtectedFromScaleIn bool) *AutoScalingInstance {
	rows := db.QueryRow("insert into auto_scaling_group_instance(auto_scaling_group_id,instance_id,availability_zone,lifecycle_state,health_status,launch_configuration_name,launch_template_id,launch_template_name,launch_template_version,protected_from_scale_in) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) returning id,auto_scaling_group_id,instance_id,ref_instance_id,availability_zone,lifecycle_state,health_status,launch_configuration_name,launch_template_id,launch_template_name,launch_template_version,protected_from_scale_in",
		AutoScalingGroupId, InstanceId, AvailabilityZone, LifecycleState, HealthStatus, LaunchConfigurationName, LaunchTemplateId, LaunchTemplateName, LaunchTemplateVersion, ProtectedFromScaleIn)

	autoScalingInstance, err := rowResultSetToAutoScalingInstance(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return autoScalingInstance
}

func createAutoScalingGroupsSuspendedProcesses(db *sql.DB, AutoScalingGroupId int64, ProcessName string, SuspensionReason string) *AutoScalingGroupsSuspendedProcesses {
	rows := db.QueryRow("insert into auto_scaling_group_suspended_process(auto_scaling_group_id,process_name,suspension_reason) values($1,$2,$3) returning id,auto_scaling_group_id,process_name,suspension_reason",
		AutoScalingGroupId, ProcessName, SuspensionReason)

	autoScalingGroupsSuspendedProcesses, err := rowResultSetToAutoScalingGroupsSuspendedProcesses(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return autoScalingGroupsSuspendedProcesses
}

func createAutoScalingGroupsEnabledMetric(db *sql.DB, AutoScalingGroupId int64, Metric string, Granularity string) *AutoScalingGroupsEnabledMetric {
	rows := db.QueryRow("insert into auto_scaling_group_enabled_metric(auto_scaling_group_id,metric,granularity) values($1,$2,$3) returning id,auto_scaling_group_id,metric,granularity",
		AutoScalingGroupId, Metric, Granularity)

	autoScalingGroupsEnabledMetric, err := rowResultSetToAutoScalingGroupsEnabledMetric(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return autoScalingGroupsEnabledMetric
}

func createAutoScalingGroupsTag(db *sql.DB, AutoScalingGroupId int64, ResourceId string, ResourceType string, Key string, Value string, PropagateAtLaunch bool) *AutoScalingGroupsTag {
	rows := db.QueryRow("insert into auto_scaling_group_tag(auto_scaling_group_id,resource_id,resource_type,key,value,propagate_at_launch) values($1,$2,$3,$4,$5,$6) returning id,auto_scaling_group_id,resource_id,resource_type,key,value,propagate_at_launch",
		AutoScalingGroupId, ResourceId, ResourceType, Key, Value, PropagateAtLaunch)

	autoScalingGroupsTag, err := rowResultSetToAutoScalingGroupsTag(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return autoScalingGroupsTag
}

func createAutoScalingGroupsTerminationPolicy(db *sql.DB, AutoScalingGroupId int64, TerminationPolicy string) *AutoScalingGroupsTerminationPolicy {
	rows := db.QueryRow("insert into auto_scaling_group_termination_policy(auto_scaling_group_id,termination_policy) values($1,$2) returning id,auto_scaling_group_id,termination_policy",
		AutoScalingGroupId, TerminationPolicy)

	autoScalingGroupsTerminationPolicy, err := rowResultSetToAutoScalingGroupsTerminationPolicy(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return autoScalingGroupsTerminationPolicy
}

//////////////////////////////////////////////////////////
//Load Unique
//////////////////////////////////////////////////////////

func loadAllAutoScalingGroupLauchTemplateOverrideByOrderUnique(db *sql.DB) ([]*AutoScalingGroupLauchTemplateOverride, error) {
	rows, err := db.Query("select id,auto_scaling_group_id,instance_type from auto_scaling_group_launch_template_override order by instance_type")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfAuto := make([]*AutoScalingGroupLauchTemplateOverride, 0, 0)
	for rows.Next() {
		auto, err := rowsNoFetchResultSetToAutoScalingGroupLauchTemplateOverride(rows)
		if err == nil {
			arrayOfAuto = append(arrayOfAuto, auto)
		} else {
			log.Println(err)
		}
	}

	return arrayOfAuto, nil
}
func loadAllAutoScalingGroupAvailabilityZoneByOrderUnique(db *sql.DB) ([]*AutoScalingGroupAvailabilityZone, error) {
	rows, err := db.Query("select id,auto_scaling_group_id,availability_zone from auto_scaling_group_availability_zone order by availability_zone")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfAuto := make([]*AutoScalingGroupAvailabilityZone, 0, 0)
	for rows.Next() {
		auto, err := rowsNoFetchResultSetToAutoScalingGroupAvailabilityZone(rows)
		if err == nil {
			arrayOfAuto = append(arrayOfAuto, auto)
		} else {
			log.Println(err)
		}
	}

	return arrayOfAuto, nil
}
func loadAllAutoScalingGroupLoadBalancerNameByOrderUnique(db *sql.DB) ([]*AutoScalingGroupLoadBalancerName, error) {
	rows, err := db.Query("select id,auto_scaling_group_id,load_balancer_name from auto_scaling_group_load_balancer_name order by load_balancer_name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfAuto := make([]*AutoScalingGroupLoadBalancerName, 0, 0)
	for rows.Next() {
		auto, err := rowsNoFetchResultSetToAutoScalingGroupLoadBalancerName(rows)
		if err == nil {
			arrayOfAuto = append(arrayOfAuto, auto)
		} else {
			log.Println(err)
		}
	}

	return arrayOfAuto, nil
}
func loadAllAutoScalingGroupTargetGroupARNByOrderUnique(db *sql.DB) ([]*AutoScalingGroupTargetGroupARN, error) {
	rows, err := db.Query("select id,auto_scaling_group_id,arn_name from auto_scaling_group_target_group_arn order by arn_name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfAuto := make([]*AutoScalingGroupTargetGroupARN, 0, 0)
	for rows.Next() {
		auto, err := rowsNoFetchResultSetToAutoScalingGroupTargetGroupARN(rows)
		if err == nil {
			arrayOfAuto = append(arrayOfAuto, auto)
		} else {
			log.Println(err)
		}
	}

	return arrayOfAuto, nil
}
func loadAllAutoScalingInstanceByOrderUnique(db *sql.DB) ([]*AutoScalingInstance, error) {
	rows, err := db.Query("select id,auto_scaling_group_id,instance_id,ref_instance_id,availability_zone,lifecycle_state,health_status,launch_configuration_name,launch_template_id,launch_template_name,launch_template_version,protected_from_scale_in from auto_scaling_group_instance order by instance_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfAuto := make([]*AutoScalingInstance, 0, 0)
	for rows.Next() {
		auto, err := rowsNoFetchResultSetToAutoScalingInstance(rows)
		if err == nil {
			arrayOfAuto = append(arrayOfAuto, auto)
		} else {
			log.Println(err)
		}
	}

	return arrayOfAuto, nil
}
func loadAllAutoScalingGroupsSuspendedProcessesByOrderUnique(db *sql.DB) ([]*AutoScalingGroupsSuspendedProcesses, error) {
	rows, err := db.Query("select id,auto_scaling_group_id,process_name,suspension_reason from auto_scaling_group_suspended_process order by suspension_reason")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfAuto := make([]*AutoScalingGroupsSuspendedProcesses, 0, 0)
	for rows.Next() {
		auto, err := rowsNoFetchResultSetToAutoScalingGroupsSuspendedProcesses(rows)
		if err == nil {
			arrayOfAuto = append(arrayOfAuto, auto)
		} else {
			log.Println(err)
		}
	}

	return arrayOfAuto, nil
}
func loadAllAutoScalingGroupsEnabledMetricByOrderUnique(db *sql.DB) ([]*AutoScalingGroupsEnabledMetric, error) {
	rows, err := db.Query("select id,auto_scaling_group_id,metric,granularity from auto_scaling_group_enabled_metric order by metric")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfAuto := make([]*AutoScalingGroupsEnabledMetric, 0, 0)
	for rows.Next() {
		auto, err := rowsNoFetchResultSetToAutoScalingGroupsEnabledMetric(rows)
		if err == nil {
			arrayOfAuto = append(arrayOfAuto, auto)
		} else {
			log.Println(err)
		}
	}

	return arrayOfAuto, nil
}
func loadAllAutoScalingGroupsTagByOrderUnique(db *sql.DB) ([]*AutoScalingGroupsTag, error) {
	rows, err := db.Query("select id,auto_scaling_group_id,resource_id,resource_type,key,value,propagate_at_launch from auto_scaling_group_tag order by value")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfAuto := make([]*AutoScalingGroupsTag, 0, 0)
	for rows.Next() {
		auto, err := rowsNoFetchResultSetToAutoScalingGroupsTag(rows)
		if err == nil {
			arrayOfAuto = append(arrayOfAuto, auto)
		} else {
			log.Println(err)
		}
	}

	return arrayOfAuto, nil
}
func loadAllAutoScalingGroupsTerminationPolicyByOrderUnique(db *sql.DB) ([]*AutoScalingGroupsTerminationPolicy, error) {
	rows, err := db.Query("select id,auto_scaling_group_id,termination_policy from auto_scaling_group_termination_policy order by termination_policy")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfAuto := make([]*AutoScalingGroupsTerminationPolicy, 0, 0)
	for rows.Next() {
		auto, err := rowsNoFetchResultSetToAutoScalingGroupsTerminationPolicy(rows)
		if err == nil {
			arrayOfAuto = append(arrayOfAuto, auto)
		} else {
			log.Println(err)
		}
	}

	return arrayOfAuto, nil
}

//////////////////////////////////////////////////////////
//Load
//////////////////////////////////////////////////////////

func loadAllAutoScalingGroupInstance(db *sql.DB) []*AutoScalingInstance {
	results := make([]*AutoScalingInstance, 0, 0)
	sqlSelect := "SELECT id,auto_scaling_group_id,instance_id,ref_instance_id,availability_zone,lifecycle_state,health_status,launch_configuration_name,launch_template_id,launch_template_name,launch_template_version,protected_from_scale_in from auto_scaling_group_instance"

	rows, err := db.Query(sqlSelect)
	if err != nil {
		return results
	}
	defer rows.Close()

	for rows.Next() {
		autoScalingInstance, err := rowsNoFetchResultSetToAutoScalingInstance(rows)
		if err != nil {
			return make([]*AutoScalingInstance, 0, 0)
		}
		results = append(results, autoScalingInstance)
	}
	return results
}

func updateAutoScalingGroupInstance(db *sql.DB, autoScalingInstance *AutoScalingInstance, instanceId int64) *AutoScalingInstance {
	rows := db.QueryRow("update auto_scaling_group_instance set ref_instance_id=$1 where id=$2 returning id,auto_scaling_group_id,instance_id,ref_instance_id,availability_zone,lifecycle_state,health_status,launch_configuration_name,launch_template_id,launch_template_name,launch_template_version,protected_from_scale_in", instanceId, autoScalingInstance.Id)

	autoScalingInstance, err := rowResultSetToAutoScalingInstance(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return autoScalingInstance
}

//////////////////////////////////////////////////////////
//DeleteAll
//////////////////////////////////////////////////////////

func deleteAllAutoScalingGroupLauchTemplateOverride(db *sql.DB) error {
	_, err := db.Exec("delete from auto_scaling_group_launch_template_override")

	return err
}

func deleteAllAutoScalingGroupAvailabilityZone(db *sql.DB) error {
	_, err := db.Exec("delete from auto_scaling_group_availability_zone")

	return err
}

func deleteAllAutoScalingGroupLoadBalancerName(db *sql.DB) error {
	_, err := db.Exec("delete from auto_scaling_group_load_balancer_name")

	return err
}

func deleteAllAutoScalingGroupTargetGroupARN(db *sql.DB) error {
	_, err := db.Exec("delete from auto_scaling_group_target_group_arn")

	return err
}

func deleteAllAutoScalingInstance(db *sql.DB) error {
	_, err := db.Exec("delete from auto_scaling_group_instance")

	return err
}

func deleteAllAutoScalingGroupsSuspendedProcesses(db *sql.DB) error {
	_, err := db.Exec("delete from auto_scaling_group_suspended_process")

	return err
}

func deleteAllAutoScalingGroupsEnabledMetric(db *sql.DB) error {
	_, err := db.Exec("delete from auto_scaling_group_enabled_metric")

	return err
}

func deleteAllAutoScalingGroupsTag(db *sql.DB) error {
	_, err := db.Exec("delete from auto_scaling_group_tag")

	return err
}

func deleteAllAutoScalingGroupsTerminationPolicy(db *sql.DB) error {
	_, err := db.Exec("delete from auto_scaling_group_termination_policy")

	return err
}
