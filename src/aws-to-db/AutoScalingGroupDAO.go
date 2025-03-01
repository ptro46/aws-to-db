package main

import (
	"database/sql" // package SQL
	"fmt"
	"log"

	_ "github.com/lib/pq" // driver Postgres
)

func rowResultSetToAutoScalingGroup(row *sql.Row) (*AutoScalingGroup, error) {
	var err error
	var Id int64
	var AutoScalingGroupName string
	var AutoScalingGroupARN string
	var LaunchConfigurationName string
	var LaunchTemplateId string
	var LaunchTemplateName string
	var LaunchTemplateVersion string
	var MixedInstancesPolicyLaunchTemplateSpecificationId string
	var MixedInstancesPolicyLaunchTemplateSpecificationName string
	var MixedInstancesPolicyLaunchTemplateSpecificationVersion string
	var MixedInstancesPolicyinstancesDistributionOnDemandAllocationStrategy string
	var MixedInstancesPolicyinstancesDistributionOnDemandBaseCapacity int16
	var MixedInstancesPolicyinstancesDistributionOnDemandPercentageAboveBaseCapacity int16
	var MixedInstancesPolicyinstancesDistributionOnDemandSpotAllocationStrategy string
	var MixedInstancesPolicyinstancesDistributionOnDemandSpotInstancePools int16
	var MixedInstancesPolicyinstancesDistributionOnDemandSpotMaxPrice string
	var MinSize int16
	var MaxSize int16
	var DesiredCapacity int16
	var DefaultCooldown int16
	var HealthCheckType string
	var HealthCheckGracePeriod int16
	var CreatedTime string
	var PlacementGroup string
	var VPCZoneIdentifier string
	var RefSubnetId sql.NullInt64
	var Status string
	var NewInstancesProtectedFromScaleIn bool
	var ServiceLinkedRoleARN string

	err = row.Scan(&Id, &AutoScalingGroupName, &AutoScalingGroupARN, &LaunchConfigurationName, &LaunchTemplateId, &LaunchTemplateName, &LaunchTemplateVersion, &MixedInstancesPolicyLaunchTemplateSpecificationId, &MixedInstancesPolicyLaunchTemplateSpecificationName, &MixedInstancesPolicyLaunchTemplateSpecificationVersion, &MixedInstancesPolicyinstancesDistributionOnDemandAllocationStrategy, &MixedInstancesPolicyinstancesDistributionOnDemandBaseCapacity, &MixedInstancesPolicyinstancesDistributionOnDemandPercentageAboveBaseCapacity, &MixedInstancesPolicyinstancesDistributionOnDemandSpotAllocationStrategy, &MixedInstancesPolicyinstancesDistributionOnDemandSpotInstancePools, &MixedInstancesPolicyinstancesDistributionOnDemandSpotMaxPrice, &MinSize, &MaxSize, &DesiredCapacity, &DefaultCooldown, &HealthCheckType, &HealthCheckGracePeriod, &CreatedTime, &PlacementGroup, &VPCZoneIdentifier, &RefSubnetId, &Status, &NewInstancesProtectedFromScaleIn, &ServiceLinkedRoleARN)
	if err != nil {
		return nil, err
	}
	return NewAutoScalingGroup(Id, AutoScalingGroupName, AutoScalingGroupARN, LaunchConfigurationName, LaunchTemplateId, LaunchTemplateName, LaunchTemplateVersion, MixedInstancesPolicyLaunchTemplateSpecificationId, MixedInstancesPolicyLaunchTemplateSpecificationName, MixedInstancesPolicyLaunchTemplateSpecificationVersion, MixedInstancesPolicyinstancesDistributionOnDemandAllocationStrategy, MixedInstancesPolicyinstancesDistributionOnDemandBaseCapacity, MixedInstancesPolicyinstancesDistributionOnDemandPercentageAboveBaseCapacity, MixedInstancesPolicyinstancesDistributionOnDemandSpotAllocationStrategy, MixedInstancesPolicyinstancesDistributionOnDemandSpotInstancePools, MixedInstancesPolicyinstancesDistributionOnDemandSpotMaxPrice, MinSize, MaxSize, DesiredCapacity, DefaultCooldown, HealthCheckType, HealthCheckGracePeriod, CreatedTime, PlacementGroup, VPCZoneIdentifier, RefSubnetId, Status, NewInstancesProtectedFromScaleIn, ServiceLinkedRoleARN), nil
}

func rowsNoFetchResultSetToAutoScalingGroup(rows *sql.Rows) (*AutoScalingGroup, error) {
	var err error
	var Id int64
	var AutoScalingGroupName string
	var AutoScalingGroupARN string
	var LaunchConfigurationName string
	var LaunchTemplateId string
	var LaunchTemplateName string
	var LaunchTemplateVersion string
	var MixedInstancesPolicyLaunchTemplateSpecificationId string
	var MixedInstancesPolicyLaunchTemplateSpecificationName string
	var MixedInstancesPolicyLaunchTemplateSpecificationVersion string
	var MixedInstancesPolicyinstancesDistributionOnDemandAllocationStrategy string
	var MixedInstancesPolicyinstancesDistributionOnDemandBaseCapacity int16
	var MixedInstancesPolicyinstancesDistributionOnDemandPercentageAboveBaseCapacity int16
	var MixedInstancesPolicyinstancesDistributionOnDemandSpotAllocationStrategy string
	var MixedInstancesPolicyinstancesDistributionOnDemandSpotInstancePools int16
	var MixedInstancesPolicyinstancesDistributionOnDemandSpotMaxPrice string
	var MinSize int16
	var MaxSize int16
	var DesiredCapacity int16
	var DefaultCooldown int16
	var HealthCheckType string
	var HealthCheckGracePeriod int16
	var CreatedTime string
	var PlacementGroup string
	var VPCZoneIdentifier string
	var RefSubnetId sql.NullInt64
	var Status string
	var NewInstancesProtectedFromScaleIn bool
	var ServiceLinkedRoleARN string

	err = rows.Scan(&Id, &AutoScalingGroupName, &AutoScalingGroupARN, &LaunchConfigurationName, &LaunchTemplateId, &LaunchTemplateName, &LaunchTemplateVersion, &MixedInstancesPolicyLaunchTemplateSpecificationId, &MixedInstancesPolicyLaunchTemplateSpecificationName, &MixedInstancesPolicyLaunchTemplateSpecificationVersion, &MixedInstancesPolicyinstancesDistributionOnDemandAllocationStrategy, &MixedInstancesPolicyinstancesDistributionOnDemandBaseCapacity, &MixedInstancesPolicyinstancesDistributionOnDemandPercentageAboveBaseCapacity, &MixedInstancesPolicyinstancesDistributionOnDemandSpotAllocationStrategy, &MixedInstancesPolicyinstancesDistributionOnDemandSpotInstancePools, &MixedInstancesPolicyinstancesDistributionOnDemandSpotMaxPrice, &MinSize, &MaxSize, &DesiredCapacity, &DefaultCooldown, &HealthCheckType, &HealthCheckGracePeriod, &CreatedTime, &PlacementGroup, &VPCZoneIdentifier, &RefSubnetId, &Status, &NewInstancesProtectedFromScaleIn, &ServiceLinkedRoleARN)
	if err != nil {
		return nil, err
	}
	return NewAutoScalingGroup(Id, AutoScalingGroupName, AutoScalingGroupARN, LaunchConfigurationName, LaunchTemplateId, LaunchTemplateName, LaunchTemplateVersion, MixedInstancesPolicyLaunchTemplateSpecificationId, MixedInstancesPolicyLaunchTemplateSpecificationName, MixedInstancesPolicyLaunchTemplateSpecificationVersion, MixedInstancesPolicyinstancesDistributionOnDemandAllocationStrategy, MixedInstancesPolicyinstancesDistributionOnDemandBaseCapacity, MixedInstancesPolicyinstancesDistributionOnDemandPercentageAboveBaseCapacity, MixedInstancesPolicyinstancesDistributionOnDemandSpotAllocationStrategy, MixedInstancesPolicyinstancesDistributionOnDemandSpotInstancePools, MixedInstancesPolicyinstancesDistributionOnDemandSpotMaxPrice, MinSize, MaxSize, DesiredCapacity, DefaultCooldown, HealthCheckType, HealthCheckGracePeriod, CreatedTime, PlacementGroup, VPCZoneIdentifier, RefSubnetId, Status, NewInstancesProtectedFromScaleIn, ServiceLinkedRoleARN), nil

}

func rowsResultSetToAutoScalingGroup(rows *sql.Rows) (*AutoScalingGroup, error) {
	var err error
	if rows.Next() {
		var Id int64
		var AutoScalingGroupName string
		var AutoScalingGroupARN string
		var LaunchConfigurationName string
		var LaunchTemplateId string
		var LaunchTemplateName string
		var LaunchTemplateVersion string
		var MixedInstancesPolicyLaunchTemplateSpecificationId string
		var MixedInstancesPolicyLaunchTemplateSpecificationName string
		var MixedInstancesPolicyLaunchTemplateSpecificationVersion string
		var MixedInstancesPolicyinstancesDistributionOnDemandAllocationStrategy string
		var MixedInstancesPolicyinstancesDistributionOnDemandBaseCapacity int16
		var MixedInstancesPolicyinstancesDistributionOnDemandPercentageAboveBaseCapacity int16
		var MixedInstancesPolicyinstancesDistributionOnDemandSpotAllocationStrategy string
		var MixedInstancesPolicyinstancesDistributionOnDemandSpotInstancePools int16
		var MixedInstancesPolicyinstancesDistributionOnDemandSpotMaxPrice string
		var MinSize int16
		var MaxSize int16
		var DesiredCapacity int16
		var DefaultCooldown int16
		var HealthCheckType string
		var HealthCheckGracePeriod int16
		var CreatedTime string
		var PlacementGroup string
		var VPCZoneIdentifier string
		var RefSubnetId sql.NullInt64
		var Status string
		var NewInstancesProtectedFromScaleIn bool
		var ServiceLinkedRoleARN string

		err = rows.Scan(&Id, &AutoScalingGroupName, &AutoScalingGroupARN, &LaunchConfigurationName, &LaunchTemplateId, &LaunchTemplateName, &LaunchTemplateVersion, &MixedInstancesPolicyLaunchTemplateSpecificationId, &MixedInstancesPolicyLaunchTemplateSpecificationName, &MixedInstancesPolicyLaunchTemplateSpecificationVersion, &MixedInstancesPolicyinstancesDistributionOnDemandAllocationStrategy, &MixedInstancesPolicyinstancesDistributionOnDemandBaseCapacity, &MixedInstancesPolicyinstancesDistributionOnDemandPercentageAboveBaseCapacity, &MixedInstancesPolicyinstancesDistributionOnDemandSpotAllocationStrategy, &MixedInstancesPolicyinstancesDistributionOnDemandSpotInstancePools, &MixedInstancesPolicyinstancesDistributionOnDemandSpotMaxPrice, &MinSize, &MaxSize, &DesiredCapacity, &DefaultCooldown, &HealthCheckType, &HealthCheckGracePeriod, &CreatedTime, &PlacementGroup, &VPCZoneIdentifier, &RefSubnetId, &Status, &NewInstancesProtectedFromScaleIn, &ServiceLinkedRoleARN)
		if err != nil {
			return nil, err
		}
		return NewAutoScalingGroup(Id, AutoScalingGroupName, AutoScalingGroupARN, LaunchConfigurationName, LaunchTemplateId, LaunchTemplateName, LaunchTemplateVersion, MixedInstancesPolicyLaunchTemplateSpecificationId, MixedInstancesPolicyLaunchTemplateSpecificationName, MixedInstancesPolicyLaunchTemplateSpecificationVersion, MixedInstancesPolicyinstancesDistributionOnDemandAllocationStrategy, MixedInstancesPolicyinstancesDistributionOnDemandBaseCapacity, MixedInstancesPolicyinstancesDistributionOnDemandPercentageAboveBaseCapacity, MixedInstancesPolicyinstancesDistributionOnDemandSpotAllocationStrategy, MixedInstancesPolicyinstancesDistributionOnDemandSpotInstancePools, MixedInstancesPolicyinstancesDistributionOnDemandSpotMaxPrice, MinSize, MaxSize, DesiredCapacity, DefaultCooldown, HealthCheckType, HealthCheckGracePeriod, CreatedTime, PlacementGroup, VPCZoneIdentifier, RefSubnetId, Status, NewInstancesProtectedFromScaleIn, ServiceLinkedRoleARN), nil

	}
	return nil, err
}

func createAutoScalingGroup(db *sql.DB, AutoScalingGroupName string, AutoScalingGroupARN string, LaunchConfigurationName string, LaunchTemplateId string, LaunchTemplateName string, LaunchTemplateVersion string, MixedInstancesPolicyLaunchTemplateSpecificationId string, MixedInstancesPolicyLaunchTemplateSpecificationName string, MixedInstancesPolicyLaunchTemplateSpecificationVersion string, MixedInstancesPolicyinstancesDistributionOnDemandAllocationStrategy string, MixedInstancesPolicyinstancesDistributionOnDemandBaseCapacity int16, MixedInstancesPolicyinstancesDistributionOnDemandPercentageAboveBaseCapacity int16, MixedInstancesPolicyinstancesDistributionOnDemandSpotAllocationStrategy string, MixedInstancesPolicyinstancesDistributionOnDemandSpotInstancePools int16, MixedInstancesPolicyinstancesDistributionOnDemandSpotMaxPrice string, MinSize int16, MaxSize int16, DesiredCapacity int16, DefaultCooldown int16, HealthCheckType string, HealthCheckGracePeriod int16, CreatedTime string, PlacementGroup string, VPCZoneIdentifier string, Status string, NewInstancesProtectedFromScaleIn bool, ServiceLinkedRoleARN string) *AutoScalingGroup {
	rows := db.QueryRow("insert into auto_scaling_group(auto_scaling_group_name,auto_scaling_group_arn,launch_configuration_name,launch_template_id,launch_template_name,launch_template_version,mx_inst_pol_launch_template_specification_id,mx_inst_pol_launch_template_specification_name,mx_inst_pol_launch_template_specification_version,mx_inst_pol_inst_dist_on_demand_allocation_strategy,mx_inst_pol_inst_dist_on_demand_base_capacity,mx_inst_pol_inst_dist_on_demand_percentage_above_base_capacity,mx_inst_pol_inst_dist_on_demand_spot_allocation_strategy,mx_inst_pol_inst_dist_on_demand_spot_instance_pools,mx_inst_pol_inst_dist_on_demand_spot_max_price,minSize,maxSize,desired_capacity,default_cooldown,health_check_type,health_check_grace_period,created_time,placement_group,vpc_zone_identifier,status,new_instances_protected_from_scale_in,service_linked_role_arn) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26,$27) returning id,auto_scaling_group_name,auto_scaling_group_arn,launch_configuration_name,launch_template_id,launch_template_name,launch_template_version,mx_inst_pol_launch_template_specification_id,mx_inst_pol_launch_template_specification_name,mx_inst_pol_launch_template_specification_version,mx_inst_pol_inst_dist_on_demand_allocation_strategy,mx_inst_pol_inst_dist_on_demand_base_capacity,mx_inst_pol_inst_dist_on_demand_percentage_above_base_capacity,mx_inst_pol_inst_dist_on_demand_spot_allocation_strategy,mx_inst_pol_inst_dist_on_demand_spot_instance_pools,mx_inst_pol_inst_dist_on_demand_spot_max_price,minSize,maxSize,desired_capacity,default_cooldown,health_check_type,health_check_grace_period,created_time,placement_group,vpc_zone_identifier,ref_subnet_id,status,new_instances_protected_from_scale_in,service_linked_role_arn",
		AutoScalingGroupName, AutoScalingGroupARN, LaunchConfigurationName, LaunchTemplateId, LaunchTemplateName, LaunchTemplateVersion, MixedInstancesPolicyLaunchTemplateSpecificationId, MixedInstancesPolicyLaunchTemplateSpecificationName, MixedInstancesPolicyLaunchTemplateSpecificationVersion, MixedInstancesPolicyinstancesDistributionOnDemandAllocationStrategy, MixedInstancesPolicyinstancesDistributionOnDemandBaseCapacity, MixedInstancesPolicyinstancesDistributionOnDemandPercentageAboveBaseCapacity, MixedInstancesPolicyinstancesDistributionOnDemandSpotAllocationStrategy, MixedInstancesPolicyinstancesDistributionOnDemandSpotInstancePools, MixedInstancesPolicyinstancesDistributionOnDemandSpotMaxPrice, MinSize, MaxSize, DesiredCapacity, DefaultCooldown, HealthCheckType, HealthCheckGracePeriod, CreatedTime, PlacementGroup, VPCZoneIdentifier, Status, NewInstancesProtectedFromScaleIn, ServiceLinkedRoleARN)

	autoScaling, err := rowResultSetToAutoScalingGroup(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return autoScaling
}

func loadAllAutoScalingGroupByOrderUnique(db *sql.DB) ([]*AutoScalingGroup, error) {
	rows, err := db.Query("select id,auto_scaling_group_name,auto_scaling_group_arn,launch_configuration_name,launch_template_id,launch_template_name,launch_template_version,mx_inst_pol_launch_template_specification_id,mx_inst_pol_launch_template_specification_name,mx_inst_pol_launch_template_specification_version,mx_inst_pol_inst_dist_on_demand_allocation_strategy,mx_inst_pol_inst_dist_on_demand_base_capacity,mx_inst_pol_inst_dist_on_demand_percentage_above_base_capacity,mx_inst_pol_inst_dist_on_demand_spot_allocation_strategy,mx_inst_pol_inst_dist_on_demand_spot_instance_pools,mx_inst_pol_inst_dist_on_demand_spot_max_price,minSize,maxSize,desired_capacity,default_cooldown,health_check_type,health_check_grace_period,created_time,placement_group,vpc_zone_identifier,ref_subnet_id,status,new_instances_protected_from_scale_in,service_linked_role_arn from auto_scaling_group order by auto_scaling_group_name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfAuto := make([]*AutoScalingGroup, 0, 0)
	for rows.Next() {
		auto, err := rowsNoFetchResultSetToAutoScalingGroup(rows)
		if err == nil {
			arrayOfAuto = append(arrayOfAuto, auto)
		} else {
			log.Println(err)
		}
	}

	return arrayOfAuto, nil
}

func loadAllAutoScalingGroup(db *sql.DB) ([]*AutoScalingGroup, error) {
	rows, err := db.Query("select id,auto_scaling_group_name,auto_scaling_group_arn,launch_configuration_name,launch_template_id,launch_template_name,launch_template_version,mx_inst_pol_launch_template_specification_id,mx_inst_pol_launch_template_specification_name,mx_inst_pol_launch_template_specification_version,mx_inst_pol_inst_dist_on_demand_allocation_strategy,mx_inst_pol_inst_dist_on_demand_base_capacity,mx_inst_pol_inst_dist_on_demand_percentage_above_base_capacity,mx_inst_pol_inst_dist_on_demand_spot_allocation_strategy,mx_inst_pol_inst_dist_on_demand_spot_instance_pools,mx_inst_pol_inst_dist_on_demand_spot_max_price,minSize,maxSize,desired_capacity,default_cooldown,health_check_type,health_check_grace_period,created_time,placement_group,vpc_zone_identifier,ref_subnet_id,status,new_instances_protected_from_scale_in,service_linked_role_arn from auto_scaling_group order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfAuto := make([]*AutoScalingGroup, 0, 0)
	for rows.Next() {
		auto, err := rowsNoFetchResultSetToAutoScalingGroup(rows)
		if err == nil {
			arrayOfAuto = append(arrayOfAuto, auto)
		} else {
			log.Println(err)
		}
	}

	return arrayOfAuto, nil
}

func loadAllAutoScalingGroupOrderByASGName(db *sql.DB) ([]*AutoScalingGroup, error) {
	rows, err := db.Query("select id,auto_scaling_group_name,auto_scaling_group_arn,launch_configuration_name,launch_template_id,launch_template_name,launch_template_version,mx_inst_pol_launch_template_specification_id,mx_inst_pol_launch_template_specification_name,mx_inst_pol_launch_template_specification_version,mx_inst_pol_inst_dist_on_demand_allocation_strategy,mx_inst_pol_inst_dist_on_demand_base_capacity,mx_inst_pol_inst_dist_on_demand_percentage_above_base_capacity,mx_inst_pol_inst_dist_on_demand_spot_allocation_strategy,mx_inst_pol_inst_dist_on_demand_spot_instance_pools,mx_inst_pol_inst_dist_on_demand_spot_max_price,minSize,maxSize,desired_capacity,default_cooldown,health_check_type,health_check_grace_period,created_time,placement_group,vpc_zone_identifier,ref_subnet_id,status,new_instances_protected_from_scale_in,service_linked_role_arn from auto_scaling_group order by auto_scaling_group_name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfAuto := make([]*AutoScalingGroup, 0, 0)
	for rows.Next() {
		auto, err := rowsNoFetchResultSetToAutoScalingGroup(rows)
		if err == nil {
			arrayOfAuto = append(arrayOfAuto, auto)
		} else {
			log.Println(err)
		}
	}

	return arrayOfAuto, nil
}

func loadAllAutoScalingGroupWithASGName(db *sql.DB, autoScalingGroupName string) ([]*AutoScalingGroup, error) {
	rows, err := db.Query("select id,auto_scaling_group_name,auto_scaling_group_arn,launch_configuration_name,launch_template_id,launch_template_name,launch_template_version,mx_inst_pol_launch_template_specification_id,mx_inst_pol_launch_template_specification_name,mx_inst_pol_launch_template_specification_version,mx_inst_pol_inst_dist_on_demand_allocation_strategy,mx_inst_pol_inst_dist_on_demand_base_capacity,mx_inst_pol_inst_dist_on_demand_percentage_above_base_capacity,mx_inst_pol_inst_dist_on_demand_spot_allocation_strategy,mx_inst_pol_inst_dist_on_demand_spot_instance_pools,mx_inst_pol_inst_dist_on_demand_spot_max_price,minSize,maxSize,desired_capacity,default_cooldown,health_check_type,health_check_grace_period,created_time,placement_group,vpc_zone_identifier,ref_subnet_id,status,new_instances_protected_from_scale_in,service_linked_role_arn from auto_scaling_group where auto_scaling_group_name=$1", autoScalingGroupName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfAuto := make([]*AutoScalingGroup, 0, 0)
	for rows.Next() {
		auto, err := rowsNoFetchResultSetToAutoScalingGroup(rows)
		if err == nil {
			arrayOfAuto = append(arrayOfAuto, auto)
		} else {
			log.Println(err)
		}
	}

	return arrayOfAuto, nil
}

func updateAutoScalingGroup(db *sql.DB, autoScalingGroup *AutoScalingGroup, subnetId int64) *AutoScalingGroup {
	rows := db.QueryRow("update auto_scaling_group set ref_subnet_id=$1 where id=$2 returning id,auto_scaling_group_name,auto_scaling_group_arn,launch_configuration_name,launch_template_id,launch_template_name,launch_template_version,mx_inst_pol_launch_template_specification_id,mx_inst_pol_launch_template_specification_name,mx_inst_pol_launch_template_specification_version,mx_inst_pol_inst_dist_on_demand_allocation_strategy,mx_inst_pol_inst_dist_on_demand_base_capacity,mx_inst_pol_inst_dist_on_demand_percentage_above_base_capacity,mx_inst_pol_inst_dist_on_demand_spot_allocation_strategy,mx_inst_pol_inst_dist_on_demand_spot_instance_pools,mx_inst_pol_inst_dist_on_demand_spot_max_price,minSize,maxSize,desired_capacity,default_cooldown,health_check_type,health_check_grace_period,created_time,placement_group,vpc_zone_identifier,ref_subnet_id,status,new_instances_protected_from_scale_in,service_linked_role_arn", subnetId, autoScalingGroup.Id)

	autoScalingGroup, err := rowResultSetToAutoScalingGroup(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return autoScalingGroup
}

func deleteAllAutoScalingGroup(db *sql.DB) error {
	_, err := db.Exec("delete from auto_scaling_group")

	return err
}
