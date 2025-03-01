package main

import (
	"database/sql" // package SQL
	"fmt"

	_ "github.com/lib/pq" // driver Postgres
)

//////////////////////////////////////////////////////////
//RowResult
//////////////////////////////////////////////////////////

func rowResultSetToReservation(row *sql.Row) (*Reservation, error) {
	var err error
	var Id int64
	var OwnerId string
	var RequesterId string
	var ReservationId string

	err = row.Scan(&Id, &OwnerId, &RequesterId, &ReservationId)
	if err != nil {
		return nil, err
	}
	return NewReservation(Id, OwnerId, RequesterId, ReservationId), nil
}

func rowResultSetToGroup(row *sql.Row) (*Group, error) {
	var err error
	var Id int64
	var ReservationId int64
	var GroupName string
	var GroupId string

	err = row.Scan(&Id, &ReservationId, &GroupName, &GroupId)
	if err != nil {
		return nil, err
	}
	return NewGroup(Id, ReservationId, GroupName, GroupId), nil
}

func rowResultSetToInstance(row *sql.Row) (*Instance, error) {
	var err error
	var Id int64
	var ReservationId int64
	var AmiLaunchIndex int16
	var ImageId string
	var RefImageId sql.NullInt64
	var InstanceId string
	var InstanceType string
	var KernelId string
	var KeyName string
	var LaunchTime string
	var MonitoringState string
	var PlacementAvailabilityZone string
	var PlacementAffinity string
	var PlacementGroupName string
	var PlacementPartitionNumber int16
	var PlacementHostId string
	var PlacementTenancy string
	var PlacementSpreadDomain string
	var Platform string
	var PrivateDnsName string
	var PrivateIpAddress string
	var PublicDnsName string
	var PublicIpAddress string
	var RamdiskId string
	var StateCode int16
	var StateName string
	var StateTransitionReason string
	var SubnetId string
	var RefSubnetId sql.NullInt64
	var VpcId string
	var RefVpcId sql.NullInt64
	var Architecture string
	var ClientToken string
	var EbsOptimized bool
	var EnaSupport bool
	var Hypervisor string
	var IamInstanceProfileId string
	var IamInstanceProfileArn string
	var InstanceLifecycle string
	var RootDeviceName string
	var RootDeviceType string
	var SourceDestCheck bool
	var SpotInstanceRequestId string
	var SriovNetSupport string
	var StateReasonCode string
	var StateReasonMessage string
	var VirtualizationType string
	var CpuOptionCoreCount int16
	var CpuOptionThreadsPerCore int16
	var CapacityReservationId string
	var CapacityReservationPreference string
	var CapacityReservationTargetId string
	var HibernationOptionsConfigured bool

	err = row.Scan(&Id, &ReservationId, &AmiLaunchIndex, &ImageId, &RefImageId, &InstanceId, &InstanceType, &KernelId, &KeyName, &LaunchTime, &MonitoringState, &PlacementAvailabilityZone, &PlacementAffinity, &PlacementGroupName, &PlacementPartitionNumber, &PlacementHostId, &PlacementTenancy, &PlacementSpreadDomain, &Platform, &PrivateDnsName, &PrivateIpAddress, &PublicDnsName, &PublicIpAddress, &RamdiskId, &StateCode, &StateName, &StateTransitionReason, &SubnetId, &RefSubnetId, &VpcId, &RefVpcId, &Architecture, &ClientToken, &EbsOptimized, &EnaSupport, &Hypervisor, &IamInstanceProfileId, &IamInstanceProfileArn, &InstanceLifecycle, &RootDeviceName, &RootDeviceType, &SourceDestCheck, &SpotInstanceRequestId, &SriovNetSupport, &StateReasonCode, &StateReasonMessage, &VirtualizationType, &CpuOptionCoreCount, &CpuOptionThreadsPerCore, &CapacityReservationId, &CapacityReservationPreference, &CapacityReservationTargetId, &HibernationOptionsConfigured)
	if err != nil {
		return nil, err
	}
	return NewInstance(Id, ReservationId, AmiLaunchIndex, ImageId, RefImageId, InstanceId, InstanceType, KernelId, KeyName, LaunchTime, MonitoringState, PlacementAvailabilityZone, PlacementAffinity, PlacementGroupName, PlacementPartitionNumber, PlacementHostId, PlacementTenancy, PlacementSpreadDomain, Platform, PrivateDnsName, PrivateIpAddress, PublicDnsName, PublicIpAddress, RamdiskId, StateCode, StateName, StateTransitionReason, SubnetId, RefSubnetId, VpcId, RefVpcId, Architecture, ClientToken, EbsOptimized, EnaSupport, Hypervisor, IamInstanceProfileId, IamInstanceProfileArn, InstanceLifecycle, RootDeviceName, RootDeviceType, SourceDestCheck, SpotInstanceRequestId, SriovNetSupport, StateReasonCode, StateReasonMessage, VirtualizationType, CpuOptionCoreCount, CpuOptionThreadsPerCore, CapacityReservationId, CapacityReservationPreference, CapacityReservationTargetId, HibernationOptionsConfigured), nil
}

//////////////////////////////////////////////////////////
//RowsNoFetch
//////////////////////////////////////////////////////////

func rowsNoFetchResultSetToReservation(rows *sql.Rows) (*Reservation, error) {
	var err error
	var Id int64
	var OwnerId string
	var RequesterId string
	var ReservationId string

	err = rows.Scan(&Id, &OwnerId, &RequesterId, &ReservationId)
	if err != nil {
		return nil, err
	}
	return NewReservation(Id, OwnerId, RequesterId, ReservationId), nil

}

func rowsNoFetchResultSetToGroup(rows *sql.Rows) (*Group, error) {
	var err error
	var Id int64
	var ReservationId int64
	var GroupName string
	var GroupId string

	err = rows.Scan(&Id, &ReservationId, &GroupName, &GroupId)
	if err != nil {
		return nil, err
	}
	return NewGroup(Id, ReservationId, GroupName, GroupId), nil
}

func rowsNoFetchResultSetToInstance(rows *sql.Rows) (*Instance, error) {
	var err error
	var Id int64
	var ReservationId int64
	var AmiLaunchIndex int16
	var ImageId string
	var RefImageId sql.NullInt64
	var InstanceId string
	var InstanceType string
	var KernelId string
	var KeyName string
	var LaunchTime string
	var MonitoringState string
	var PlacementAvailabilityZone string
	var PlacementAffinity string
	var PlacementGroupName string
	var PlacementPartitionNumber int16
	var PlacementHostId string
	var PlacementTenancy string
	var PlacementSpreadDomain string
	var Platform string
	var PrivateDnsName string
	var PrivateIpAddress string
	var PublicDnsName string
	var PublicIpAddress string
	var RamdiskId string
	var StateCode int16
	var StateName string
	var StateTransitionReason string
	var SubnetId string
	var RefSubnetId sql.NullInt64
	var VpcId string
	var RefVpcId sql.NullInt64
	var Architecture string
	var ClientToken string
	var EbsOptimized bool
	var EnaSupport bool
	var Hypervisor string
	var IamInstanceProfileId string
	var IamInstanceProfileArn string
	var InstanceLifecycle string
	var RootDeviceName string
	var RootDeviceType string
	var SourceDestCheck bool
	var SpotInstanceRequestId string
	var SriovNetSupport string
	var StateReasonCode string
	var StateReasonMessage string
	var VirtualizationType string
	var CpuOptionCoreCount int16
	var CpuOptionThreadsPerCore int16
	var CapacityReservationId string
	var CapacityReservationPreference string
	var CapacityReservationTargetId string
	var HibernationOptionsConfigured bool

	err = rows.Scan(&Id, &ReservationId, &AmiLaunchIndex, &ImageId, &RefImageId, &InstanceId, &InstanceType, &KernelId, &KeyName, &LaunchTime, &MonitoringState, &PlacementAvailabilityZone, &PlacementAffinity, &PlacementGroupName, &PlacementPartitionNumber, &PlacementHostId, &PlacementTenancy, &PlacementSpreadDomain, &Platform, &PrivateDnsName, &PrivateIpAddress, &PublicDnsName, &PublicIpAddress, &RamdiskId, &StateCode, &StateName, &StateTransitionReason, &SubnetId, &RefSubnetId, &VpcId, &RefVpcId, &Architecture, &ClientToken, &EbsOptimized, &EnaSupport, &Hypervisor, &IamInstanceProfileId, &IamInstanceProfileArn, &InstanceLifecycle, &RootDeviceName, &RootDeviceType, &SourceDestCheck, &SpotInstanceRequestId, &SriovNetSupport, &StateReasonCode, &StateReasonMessage, &VirtualizationType, &CpuOptionCoreCount, &CpuOptionThreadsPerCore, &CapacityReservationId, &CapacityReservationPreference, &CapacityReservationTargetId, &HibernationOptionsConfigured)
	if err != nil {
		return nil, err
	}
	return NewInstance(Id, ReservationId, AmiLaunchIndex, ImageId, RefImageId, InstanceId, InstanceType, KernelId, KeyName, LaunchTime, MonitoringState, PlacementAvailabilityZone, PlacementAffinity, PlacementGroupName, PlacementPartitionNumber, PlacementHostId, PlacementTenancy, PlacementSpreadDomain, Platform, PrivateDnsName, PrivateIpAddress, PublicDnsName, PublicIpAddress, RamdiskId, StateCode, StateName, StateTransitionReason, SubnetId, RefSubnetId, VpcId, RefVpcId, Architecture, ClientToken, EbsOptimized, EnaSupport, Hypervisor, IamInstanceProfileId, IamInstanceProfileArn, InstanceLifecycle, RootDeviceName, RootDeviceType, SourceDestCheck, SpotInstanceRequestId, SriovNetSupport, StateReasonCode, StateReasonMessage, VirtualizationType, CpuOptionCoreCount, CpuOptionThreadsPerCore, CapacityReservationId, CapacityReservationPreference, CapacityReservationTargetId, HibernationOptionsConfigured), nil

}

//////////////////////////////////////////////////////////
//RowsResult
//////////////////////////////////////////////////////////

func rowsResultSetToReservation(rows *sql.Rows) (*Reservation, error) {
	var err error
	if rows.Next() {
		var err error
		var Id int64
		var OwnerId string
		var RequesterId string
		var ReservationId string

		err = rows.Scan(&Id, &OwnerId, &RequesterId, &ReservationId)
		if err != nil {
			return nil, err
		}
		return NewReservation(Id, OwnerId, RequesterId, ReservationId), nil

	}
	return nil, err
}

func rowsResultSetToGroup(rows *sql.Rows) (*Group, error) {
	var err error
	if rows.Next() {
		var err error
		var Id int64
		var ReservationId int64
		var GroupName string
		var GroupId string

		err = rows.Scan(&Id, &ReservationId, &GroupName, &GroupId)
		if err != nil {
			return nil, err
		}
		return NewGroup(Id, ReservationId, GroupName, GroupId), nil

	}
	return nil, err
}

func rowsResultSetToInstance(rows *sql.Rows) (*Instance, error) {
	var err error
	if rows.Next() {
		var err error
		var Id int64
		var ReservationId int64
		var AmiLaunchIndex int16
		var ImageId string
		var RefImageId sql.NullInt64
		var InstanceId string
		var InstanceType string
		var KernelId string
		var KeyName string
		var LaunchTime string
		var MonitoringState string
		var PlacementAvailabilityZone string
		var PlacementAffinity string
		var PlacementGroupName string
		var PlacementPartitionNumber int16
		var PlacementHostId string
		var PlacementTenancy string
		var PlacementSpreadDomain string
		var Platform string
		var PrivateDnsName string
		var PrivateIpAddress string
		var PublicDnsName string
		var PublicIpAddress string
		var RamdiskId string
		var StateCode int16
		var StateName string
		var StateTransitionReason string
		var SubnetId string
		var RefSubnetId sql.NullInt64
		var VpcId string
		var RefVpcId sql.NullInt64
		var Architecture string
		var ClientToken string
		var EbsOptimized bool
		var EnaSupport bool
		var Hypervisor string
		var IamInstanceProfileId string
		var IamInstanceProfileArn string
		var InstanceLifecycle string
		var RootDeviceName string
		var RootDeviceType string
		var SourceDestCheck bool
		var SpotInstanceRequestId string
		var SriovNetSupport string
		var StateReasonCode string
		var StateReasonMessage string
		var VirtualizationType string
		var CpuOptionCoreCount int16
		var CpuOptionThreadsPerCore int16
		var CapacityReservationId string
		var CapacityReservationPreference string
		var CapacityReservationTargetId string
		var HibernationOptionsConfigured bool

		err = rows.Scan(&Id, &ReservationId, &AmiLaunchIndex, &ImageId, &RefImageId, &InstanceId, &InstanceType, &KernelId, &KeyName, &LaunchTime, &MonitoringState, &PlacementAvailabilityZone, &PlacementAffinity, &PlacementGroupName, &PlacementPartitionNumber, &PlacementHostId, &PlacementTenancy, &PlacementSpreadDomain, &Platform, &PrivateDnsName, &PrivateIpAddress, &PublicDnsName, &PublicIpAddress, &RamdiskId, &StateCode, &StateName, &StateTransitionReason, &SubnetId, &RefSubnetId, &VpcId, &RefVpcId, &Architecture, &ClientToken, &EbsOptimized, &EnaSupport, &Hypervisor, &IamInstanceProfileId, &IamInstanceProfileArn, &InstanceLifecycle, &RootDeviceName, &RootDeviceType, &SourceDestCheck, &SpotInstanceRequestId, &SriovNetSupport, &StateReasonCode, &StateReasonMessage, &VirtualizationType, &CpuOptionCoreCount, &CpuOptionThreadsPerCore, &CapacityReservationId, &CapacityReservationPreference, &CapacityReservationTargetId, &HibernationOptionsConfigured)
		if err != nil {
			return nil, err
		}
		return NewInstance(Id, ReservationId, AmiLaunchIndex, ImageId, RefImageId, InstanceId, InstanceType, KernelId, KeyName, LaunchTime, MonitoringState, PlacementAvailabilityZone, PlacementAffinity, PlacementGroupName, PlacementPartitionNumber, PlacementHostId, PlacementTenancy, PlacementSpreadDomain, Platform, PrivateDnsName, PrivateIpAddress, PublicDnsName, PublicIpAddress, RamdiskId, StateCode, StateName, StateTransitionReason, SubnetId, RefSubnetId, VpcId, RefVpcId, Architecture, ClientToken, EbsOptimized, EnaSupport, Hypervisor, IamInstanceProfileId, IamInstanceProfileArn, InstanceLifecycle, RootDeviceName, RootDeviceType, SourceDestCheck, SpotInstanceRequestId, SriovNetSupport, StateReasonCode, StateReasonMessage, VirtualizationType, CpuOptionCoreCount, CpuOptionThreadsPerCore, CapacityReservationId, CapacityReservationPreference, CapacityReservationTargetId, HibernationOptionsConfigured), nil

	}
	fmt.Println("No Next")
	return nil, err
}

//////////////////////////////////////////////////////////
//Create
//////////////////////////////////////////////////////////

func createReservation(db *sql.DB, OwnerId string, RequesterId string, ReservationId string) *Reservation {
	rows := db.QueryRow("insert into reservation(owner_id,requester_id,reservation_id) values($1,$2,$3) returning id,owner_id,requester_id,reservation_id", OwnerId, RequesterId, ReservationId)

	reservation, err := rowResultSetToReservation(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return reservation
}

func createGroup(db *sql.DB, ReservationId int64, GroupName string, GroupId string) *Group {
	rows := db.QueryRow("insert into reservation_group(reservation_id,group_name,group_id) values($1,$2,$3) returning id,reservation_id,group_name,group_id", ReservationId, GroupName, GroupId)

	group, err := rowResultSetToGroup(rows)
	if err != nil {
		return nil
	}
	return group
}

func createInstance(db *sql.DB, ReservationId int64, AmiLaunchIndex int16, ImageId string, InstanceId string, InstanceType string, KernelId string, KeyName string, LaunchTime string, MonitoringState string, PlacementAvailabilityZone string, PlacementAffinity string, PlacementGroupName string, PlacementPartitionNumber int16, PlacementHostId string, PlacementTenancy string, PlacementSpreadDomain string, Platform string, PrivateDnsName string, PrivateIpAddress string, PublicDnsName string, PublicIpAddress string, RamdiskId string, StateCode int16, StateName string, StateTransitionReason string, SubnetId string, VpcId string, Architecture string, ClientToken string, EbsOptimized bool, EnaSupport bool, Hypervisor string, IamInstanceProfileId string, IamInstanceProfileArn string, InstanceLifecycle string, RootDeviceName string, RootDeviceType string, SourceDestCheck bool, SpotInstanceRequestId string, SriovNetSupport string, StateReasonCode string, StateReasonMessage string, VirtualizationType string, CpuOptionCoreCount int16, CpuOptionThreadsPerCore int16, CapacityReservationId string, CapacityReservationPreference string, CapacityReservationTargetId string, HibernationOptionsConfigured bool) *Instance {
	rows := db.QueryRow("insert into instance(reservation_id,ami_launch_index,image_id,instance_id,instance_type,kernel_id,key_name,launch_time,monitoring_state,placement_availability_zone,placement_affinity,placement_group_name,placement_partition_number,placement_host_id,placement_tenancy,placement_spread_domain,platform,private_dns_name,private_ip_address,public_dns_name,public_ip_address,ramdisk_id,state_code,state_name,state_transition_reason,subnet_id,vpc_id,architecture,client_token,ebs_optimized,ena_support,hypervisor,iam_instance_profile_id,iam_instance_profile_arn,instance_lifecycle,root_device_name,root_device_type,source_dest_check,spot_instance_request_id,sriov_net_support,state_reason_code,state_reason_message,virtualization_type,cpu_option_core_count,cpu_option_threads_per_core,capacity_reservation_id,capacity_reservation_preference,capacity_reservation_target_id,hibernation_options_configured) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26,$27,$28,$29,$30,$31,$32,$33,$34,$35,$36,$37,$38,$39,$40,$41,$42,$43,$44,$45,$46,$47,$48,$49) returning id,reservation_id,ami_launch_index,image_id,ref_image_id,instance_id,instance_type,kernel_id,key_name,launch_time,monitoring_state,placement_availability_zone,placement_affinity,placement_group_name,placement_partition_number,placement_host_id,placement_tenancy,placement_spread_domain,platform,private_dns_name,private_ip_address,public_dns_name,public_ip_address,ramdisk_id,state_code,state_name,state_transition_reason,subnet_id,ref_subnet_id,vpc_id,ref_vpc_id,architecture,client_token,ebs_optimized,ena_support,hypervisor,iam_instance_profile_id,iam_instance_profile_arn,instance_lifecycle,root_device_name,root_device_type,source_dest_check,spot_instance_request_id,sriov_net_support,state_reason_code,state_reason_message,virtualization_type,cpu_option_core_count,cpu_option_threads_per_core,capacity_reservation_id,capacity_reservation_preference,capacity_reservation_target_id,hibernation_options_configured",
		ReservationId, AmiLaunchIndex, ImageId, InstanceId, InstanceType, KernelId, KeyName, LaunchTime, MonitoringState, PlacementAvailabilityZone, PlacementAffinity, PlacementGroupName, PlacementPartitionNumber, PlacementHostId, PlacementTenancy, PlacementSpreadDomain, Platform, PrivateDnsName, PrivateIpAddress, PublicDnsName,
		PublicIpAddress, RamdiskId, StateCode, StateName, StateTransitionReason, SubnetId, VpcId, Architecture, ClientToken, EbsOptimized, EnaSupport, Hypervisor, IamInstanceProfileId, IamInstanceProfileArn, InstanceLifecycle, RootDeviceName, RootDeviceType, SourceDestCheck, SpotInstanceRequestId, SriovNetSupport, StateReasonCode, StateReasonMessage,
		VirtualizationType, CpuOptionCoreCount, CpuOptionThreadsPerCore, CapacityReservationId, CapacityReservationPreference, CapacityReservationTargetId, HibernationOptionsConfigured)

	instance, err := rowResultSetToInstance(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return instance
}

//////////////////////////////////////////////////////////
//Load
//////////////////////////////////////////////////////////

func loadAllInstance(db *sql.DB) []*Instance {
	results := make([]*Instance, 0, 0)
	sqlSelect := "SELECT id,reservation_id,ami_launch_index,image_id,ref_image_id,instance_id,instance_type,kernel_id,key_name,launch_time,monitoring_state,placement_availability_zone,placement_affinity,placement_group_name,placement_partition_number,placement_host_id,placement_tenancy,placement_spread_domain,platform,private_dns_name,private_ip_address,public_dns_name,public_ip_address,ramdisk_id,state_code,state_name,state_transition_reason,subnet_id,ref_subnet_id,vpc_id,ref_vpc_id,architecture,client_token,ebs_optimized,ena_support,hypervisor,iam_instance_profile_id,iam_instance_profile_arn,instance_lifecycle,root_device_name,root_device_type,source_dest_check,spot_instance_request_id,sriov_net_support,state_reason_code,state_reason_message,virtualization_type,cpu_option_core_count,cpu_option_threads_per_core,capacity_reservation_id,capacity_reservation_preference,capacity_reservation_target_id,hibernation_options_configured from instance"

	rows, err := db.Query(sqlSelect)
	if err != nil {
		return results
	}
	defer rows.Close()

	for rows.Next() {
		instance, err := rowsNoFetchResultSetToInstance(rows)
		if err != nil {
			return make([]*Instance, 0, 0)
		}
		results = append(results, instance)
	}
	return results
}

func loadAllReservationsByOrderUnique(db *sql.DB) ([]*Reservation, error) {
	results := make([]*Reservation, 0, 0)
	sqlSelect := "SELECT id,owner_id,requester_id,reservation_id from reservation order by reservation_id"

	rows, err := db.Query(sqlSelect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		reservation, err := rowsNoFetchResultSetToReservation(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, reservation)
	}
	return results, nil
}

func loadAllGroupByOrderUnique(db *sql.DB) ([]*Group, error) {
	results := make([]*Group, 0, 0)
	sqlSelect := "SELECT id,reservation_id,group_name,group_id from reservation_group order by group_name"

	rows, err := db.Query(sqlSelect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		group, err := rowsNoFetchResultSetToGroup(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, group)
	}
	return results, nil
}

func loadAllInstanceByOrderUnique(db *sql.DB) ([]*Instance, error) {
	results := make([]*Instance, 0, 0)
	sqlSelect := "SELECT id,reservation_id,ami_launch_index,image_id,ref_image_id,instance_id,instance_type,kernel_id,key_name,launch_time,monitoring_state,placement_availability_zone,placement_affinity,placement_group_name,placement_partition_number,placement_host_id,placement_tenancy,placement_spread_domain,platform,private_dns_name,private_ip_address,public_dns_name,public_ip_address,ramdisk_id,state_code,state_name,state_transition_reason,subnet_id,ref_subnet_id,vpc_id,ref_vpc_id,architecture,client_token,ebs_optimized,ena_support,hypervisor,iam_instance_profile_id,iam_instance_profile_arn,instance_lifecycle,root_device_name,root_device_type,source_dest_check,spot_instance_request_id,sriov_net_support,state_reason_code,state_reason_message,virtualization_type,cpu_option_core_count,cpu_option_threads_per_core,capacity_reservation_id,capacity_reservation_preference,capacity_reservation_target_id,hibernation_options_configured from instance order by instance_id"

	rows, err := db.Query(sqlSelect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		instance, err := rowsNoFetchResultSetToInstance(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, instance)
	}
	return results, nil
}

func loadInstanceByAWSReferenceId(db *sql.DB, instanceRefID string) (*Instance, error) {
	rows, err := db.Query("select id,reservation_id,ami_launch_index,image_id,ref_image_id,instance_id,instance_type,kernel_id,key_name,launch_time,monitoring_state,placement_availability_zone,placement_affinity,placement_group_name,placement_partition_number,placement_host_id,placement_tenancy,placement_spread_domain,platform,private_dns_name,private_ip_address,public_dns_name,public_ip_address,ramdisk_id,state_code,state_name,state_transition_reason,subnet_id,ref_subnet_id,vpc_id,ref_vpc_id,architecture,client_token,ebs_optimized,ena_support,hypervisor,iam_instance_profile_id,iam_instance_profile_arn,instance_lifecycle,root_device_name,root_device_type,source_dest_check,spot_instance_request_id,sriov_net_support,state_reason_code,state_reason_message,virtualization_type,cpu_option_core_count,cpu_option_threads_per_core,capacity_reservation_id,capacity_reservation_preference,capacity_reservation_target_id,hibernation_options_configured from instance where instance_id=$1", instanceRefID)
	if err != nil {
		return nil, err
	}

	instance, err := rowsResultSetToInstance(rows)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	return instance, nil
}

/*
select I.*
from instance I, instance_tag IT
where I.id=IT.instance_id and IT.key='Name' and IT.value='logs-rsyslogd'
group by I.id;
*/
func loadInstanceByName(db *sql.DB, name string) (*Instance, error) {
	rows, err := db.Query("select I.id,reservation_id,ami_launch_index,image_id,ref_image_id,I.instance_id,instance_type,kernel_id,key_name,launch_time,monitoring_state,placement_availability_zone,placement_affinity,placement_group_name,placement_partition_number,placement_host_id,placement_tenancy,placement_spread_domain,platform,private_dns_name,private_ip_address,public_dns_name,public_ip_address,ramdisk_id,state_code,state_name,state_transition_reason,subnet_id,ref_subnet_id,vpc_id,ref_vpc_id,architecture,client_token,ebs_optimized,ena_support,hypervisor,iam_instance_profile_id,iam_instance_profile_arn,instance_lifecycle,root_device_name,root_device_type,source_dest_check,spot_instance_request_id,sriov_net_support,state_reason_code,state_reason_message,virtualization_type,cpu_option_core_count,cpu_option_threads_per_core,capacity_reservation_id,capacity_reservation_preference,capacity_reservation_target_id,hibernation_options_configured from instance I, instance_tag IT where I.id=IT.instance_id and IT.key='Name' and IT.value=$1", name)
	if err != nil {
		return nil, err
	}

	instance, err := rowsResultSetToInstance(rows)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	return instance, nil
}

func updateInstanceId(db *sql.DB, instance *Instance, imageId int64, vpcId int64, subnetId int64) *Instance {
	rows := db.QueryRow("update instance set ref_image_id=$1,ref_subnet_id=$2,ref_vpc_id=$3 where id=$4 returning id,reservation_id,ami_launch_index,image_id,ref_image_id,instance_id,instance_type,kernel_id,key_name,launch_time,monitoring_state,placement_availability_zone,placement_affinity,placement_group_name,placement_partition_number,placement_host_id,placement_tenancy,placement_spread_domain,platform,private_dns_name,private_ip_address,public_dns_name,public_ip_address,ramdisk_id,state_code,state_name,state_transition_reason,subnet_id,ref_subnet_id,vpc_id,ref_vpc_id,architecture,client_token,ebs_optimized,ena_support,hypervisor,iam_instance_profile_id,iam_instance_profile_arn,instance_lifecycle,root_device_name,root_device_type,source_dest_check,spot_instance_request_id,sriov_net_support,state_reason_code,state_reason_message,virtualization_type,cpu_option_core_count,cpu_option_threads_per_core,capacity_reservation_id,capacity_reservation_preference,capacity_reservation_target_id,hibernation_options_configured", imageId, subnetId, vpcId, instance.Id)

	instance, err := rowResultSetToInstance(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return instance
}

//////////////////////////////////////////////////////////
//DeleteAll
//////////////////////////////////////////////////////////

func deleteAllReservation(db *sql.DB) error {
	_, err := db.Exec("delete from reservation")

	return err
}

func deleteAllGroup(db *sql.DB) error {
	_, err := db.Exec("delete from reservation_group")

	return err
}

func deleteAllInstance(db *sql.DB) error {
	_, err := db.Exec("delete from instance")

	return err
}
