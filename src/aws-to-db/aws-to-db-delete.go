package main

import (
	"database/sql"
	"fmt"
)

func awsPurge(db *sql.DB) error {
	var err error
	names := []string{"auto_scaling_activities", "auto_scaling_groups", "load_balancers", "buckets", "stacks", "db_instances", "instances", "amis", "snapshots", "volumes", "security_groups", "subnets", "vpc"}
	//{"subnets", "vpc", "buckets", "amis", "stacks", "volumes", "security_groups", "db_instances", "instances", "load_balancers", "auto_scaling_groups"}
	fmt.Println("Start Purge")
	for _, name := range names {
		err = deleteContent(db, name)
		if err != nil {
			fmt.Printf("	ERROR Purge %s failed\n", name)
			break
		}
		fmt.Printf("	Purge %s success\n", name)
	}
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Purge succesful")
	}
	return nil
}

func awsPrepareUpdate(db *sql.DB) error {
	var err error
	namesForUpdate := []string{"auto_scaling_activities", "auto_scaling_groups", "load_balancers", "instances"}
	//{"subnets", "vpc", "buckets", "amis", "stacks", "volumes", "security_groups", "db_instances", "instances", "load_balancers", "auto_scaling_groups"}
	fmt.Println("Start Purge")
	for _, name := range namesForUpdate {
		err = deleteContent(db, name)
		if err != nil {
			fmt.Printf("	ERROR Purge %s failed\n", name)
			break
		}
		fmt.Printf("	Purge %s success\n", name)
	}
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Purge succesful")
	}
	return nil
}

func deleteContent(db *sql.DB, option string) error {

	switch option {
	case "vpc":
		return deleteAllVPC(db)

	case "subnets":
		return deleteAllSubnet(db)

	case "buckets":
		return deleteAllBucket(db)

	case "amis":
		return deleteAllAmi(db)

	case "stacks":
		err := deleteAllStackOutput(db)
		if err != nil {
			return err
		}
		err = deleteAllStackParameter(db)
		if err != nil {
			return err
		}
		err = deleteAllStackTag(db)
		if err != nil {
			return err
		}
		return deleteAllStack(db)

	case "snapshots":
		return deleteAllSnapshot(db)

	case "volumes":
		err := deleteAllVolumeTag(db)
		if err != nil {
			return err
		}
		err = deleteAllVolumeAttachment(db)
		if err != nil {
			return err
		}
		return deleteAllVolume(db)

	case "security_groups":
		err := deleteAllSecurityGroupIpPermissionEgressUserIdGroupPair(db)
		if err != nil {
			return err
		}
		err = deleteAllSecurityGroupIpPermissionUserIdGroupPair(db)
		if err != nil {
			return err
		}
		err = deleteAllSecurityGroupIpPermissionEgressPrefixListId(db)
		if err != nil {
			return err
		}
		err = deleteAllSecurityGroupIpPermissionPrefixListId(db)
		if err != nil {
			return err
		}
		err = deleteAllSecurityGroupIpPermissionEgressIpRange(db)
		if err != nil {
			return err
		}
		err = deleteAllSecurityGroupIpPermissionIpRange(db)
		if err != nil {
			return err
		}
		err = deleteAllSecurityGroupIpPermission(db)
		if err != nil {
			return err
		}
		err = deleteAllSecurityGroupIpPermissionEgress(db)
		if err != nil {
			return err
		}
		err = deleteAllSecurityGroupTag(db)
		if err != nil {
			return err
		}
		return deleteAllSecurityGroup(db)

	case "db_instances":
		err := deleteAllDBInstanceAssociatedRoles(db)
		if err != nil {
			return err
		}
		err = deleteAllDBInstanceDBParameterGroups(db)
		if err != nil {
			return err
		}
		err = deleteAllDBInstanceDBSecurityGroups(db)
		if err != nil {
			return err
		}
		err = deleteAllDBInstanceDBSubnetGroupSubnet(db)
		if err != nil {
			return err
		}
		err = deleteAllDBInstanceDBSubnetGroup(db)
		if err != nil {
			return err
		}
		err = deleteAllDBInstanceDomainMemberships(db)
		if err != nil {
			return err
		}
		err = deleteAllDBInstanceOptionGroupMemberships(db)
		if err != nil {
			return err
		}
		err = deleteAllDBInstanceProcessorFeatures(db)
		if err != nil {
			return err
		}
		err = deleteAllDBInstanceVpcSecurityGroups(db)
		if err != nil {
			return err
		}
		err = deleteAllDBInstanceStatusInfos(db)
		if err != nil {
			return err
		}
		return deleteAllDBInstance(db)

	case "instances":
		err := deleteAllInstanceProductCode(db)
		if err != nil {
			return err
		}
		err = deleteAllInstanceBlockDeviceMapping(db)
		if err != nil {
			return err
		}
		err = deleteAllInstanceElasticGpuAssociation(db)
		if err != nil {
			return err
		}
		err = deleteAllInstanceElasticInferenceAcceleratorAssociation(db)
		if err != nil {
			return err
		}
		err = deleteAllInstanceNetworkInterfaceGroup(db)
		if err != nil {
			return err
		}
		err = deleteAllInstanceNetworkInterfacePrivateIpAddress(db)
		if err != nil {
			return err
		}
		err = deleteAllInstanceNetworkInterface(db)
		if err != nil {
			return err
		}
		err = deleteAllInstanceSecurityGroup(db)
		if err != nil {
			return err
		}
		err = deleteAllInstanceTag(db)
		if err != nil {
			return err
		}
		err = deleteAllInstanceLicense(db)
		if err != nil {
			return err
		}
		err = deleteAllVolumeAttachment(db)
		if err != nil {
			return err
		}
		err = deleteAllInstance(db)
		if err != nil {
			return err
		}
		err = deleteAllGroup(db)
		if err != nil {
			return err
		}
		return deleteAllReservation(db)

	case "load_balancers":
		err := deleteAllLoadBalancerTargetGroupHealth(db)
		if err != nil {
			return err
		}
		err = deleteAllLoadBalancerTargetGroupLoadBalancerARN(db)
		if err != nil {
			return err
		}
		err = deleteAllLoadBalancerListenerCertificat(db)
		if err != nil {
			return err
		}
		err = deleteAllLoadBalancerListenerDefaultAction(db)
		if err != nil {
			return err
		}
		err = deleteAllLoadBalancerTargetGroup(db)
		if err != nil {
			return err
		}

		err = deleteAllLoadBalancerListener(db)
		if err != nil {
			return err
		}

		err = deleteAllLoadBalancerAddress(db)
		if err != nil {
			return err
		}
		err = deleteAllLoadBalancerAvailabilityZone(db)
		if err != nil {
			return err
		}
		err = deleteAllLoadBalancerSecurityGroup(db)
		if err != nil {
			return err
		}

		return deleteAllLoadBalancer(db)

	case "auto_scaling_activities":
		return deleteAllAutoScalingGroupActivities(db)

	case "auto_scaling_groups":
		err := deleteAllAutoScalingGroupAvailabilityZone(db)
		if err != nil {
			return err
		}
		err = deleteAllAutoScalingGroupLauchTemplateOverride(db)
		if err != nil {
			return err
		}
		err = deleteAllAutoScalingGroupLoadBalancerName(db)
		if err != nil {
			return err
		}
		err = deleteAllAutoScalingGroupTargetGroupARN(db)
		if err != nil {
			return err
		}
		err = deleteAllAutoScalingGroupsEnabledMetric(db)
		if err != nil {
			return err
		}
		err = deleteAllAutoScalingGroupsSuspendedProcesses(db)
		if err != nil {
			return err
		}
		err = deleteAllAutoScalingGroupsTag(db)
		if err != nil {
			return err
		}
		err = deleteAllAutoScalingGroupsTerminationPolicy(db)
		if err != nil {
			return err
		}
		err = deleteAllAutoScalingInstance(db)
		if err != nil {
			return err
		}
		return deleteAllAutoScalingGroup(db)

	default:
		fmt.Println("option did not match with any delete function")
		return nil
	}
}
