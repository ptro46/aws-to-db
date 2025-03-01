package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os/exec"
)

func executeAWSCommandAndReturnJSON(awsCommand string) (string, error) {
	fmt.Printf("/bin/bash %s\n", awsCommand)
	cmd := exec.Command("/bin/bash", awsCommand)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "1", err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return "2", err
	}
	if err := cmd.Start(); err != nil {
		return "3", err
	}

	bufOut, _ := ioutil.ReadAll(stdout)
	bufErr, _ := ioutil.ReadAll(stderr)

	if err := cmd.Wait(); err != nil {
		return string(bufOut) + " " + string(bufErr), err
	}
	jsonReturn := string(bufOut)
	return jsonReturn, nil
}

func awsLoadAll(db *sql.DB, cmdByName map[string]pLoadFunc, commands []CommandConfig) {
	for _, command := range commands {
		pFunc := cmdByName[command.Name]
		if pFunc != nil {
			jsonResponse, err := executeAWSCommandAndReturnJSON(command.Cmd)
			if err == nil {
				err = pFunc(db, jsonResponse)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Printf("\nBypass %s\n", command.Name)
		}
		fmt.Println("\n------------------------------------\n")

	}
	awsCorrelation(db)
}

func awsLoadOption(db *sql.DB, commands []CommandConfig, cmdByName map[string]pLoadFunc, option string) {
	var cmd = ""
	for _, command := range commands {
		if command.Name == option {
			cmd = command.Cmd
			break
		}
	}

	if cmd != "" {
		deleteContent(db, option)
		jsonResponse, err := executeAWSCommandAndReturnJSON(cmd)
		if err == nil {
			pFunc := cmdByName[option]
			if pFunc != nil {
				err = pFunc(db, jsonResponse)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	} else {
		fmt.Println("Invalid load option entered\n Acceptable Inputs:\n		subnets, vpc, buckets, amis, stacks, volumes, security_groups, db_instances, instances, load_balancers, auto_scaling_groups")
	}
}

func awsCorrelation(db *sql.DB) {

	fmt.Println("\nBuild relation for load_balancer_target_group to vpc")
	arrayOfTargetGroup, _ := loadAllLoadBalancerTargetGroupByOrderUnique(db)
	fmt.Println("  #Found:", len(arrayOfTargetGroup))
	for _, oneTargetGroup := range arrayOfTargetGroup {
		vpc, err := loadVPCByVPCId(db, oneTargetGroup.VpcId)
		if err == nil && vpc != nil {
			updateLoadBalancerTargetGroupId(db, oneTargetGroup, vpc.Id)
			fmt.Println("	", oneTargetGroup.VpcId, "Loaded")
		} else {
			fmt.Println("	ERROR", oneTargetGroup.VpcId, "Not Loaded")
		}
	}

	fmt.Println("\nBuild relation for load_balancer_target_group_arn to load_balancer")
	arrayOfTargetGroupArn, _ := loadAllLoadBalancerTargetGroupLoadBalancerARNByOrderUnique(db)
	fmt.Println("  #Found:", len(arrayOfTargetGroupArn))
	for _, oneArn := range arrayOfTargetGroupArn {
		loadBalancer, err := loadAllLoadBalancerByARN(db, oneArn.LoadBalancerArn)
		if err == nil && loadBalancer != nil {
			updateLoadBalancerTargetGroupLoadBalancerARNId(db, oneArn, loadBalancer.Id)
			fmt.Println("	", oneArn.LoadBalancerArn, "Loaded")
		} else {
			fmt.Println("	ERROR", oneArn.LoadBalancerArn, "Not Loaded")
		}
	}

	fmt.Println("\nBuild relation for db_instance_vpc_security_groups to security_group")
	arrayOfVpcSecGroup, _ := loadAllDBInstanceVpcSecurityGroup(db)
	fmt.Println("  #Found:", len(arrayOfVpcSecGroup))
	for _, oneSecGroup := range arrayOfVpcSecGroup {
		secGroup, err := loadSecurityGroupByAWSReferenceId(db, oneSecGroup.VpcSecurityGroupId)
		if err == nil && secGroup != nil {
			updateDBInstanceVpcSecurityGroup(db, oneSecGroup, secGroup.Id)
			fmt.Println("	", oneSecGroup.VpcSecurityGroupId, "Loaded")
		} else {
			fmt.Println("	ERROR", oneSecGroup.VpcSecurityGroupId, "Not Loaded")
		}
	}

	fmt.Println("\nBuild relation for load_balancer_security_group to security_group")
	arrayOfLoadBalancerSecurityGroup, _ := loadAllLoadBalancerSecurityGroup(db)
	fmt.Println("  #Found:", len(arrayOfLoadBalancerSecurityGroup))
	for _, oneLoadBalancerSecurityGroup := range arrayOfLoadBalancerSecurityGroup {
		if oneLoadBalancerSecurityGroup == nil {
			continue
		}
		secGroup, err := loadSecurityGroupByAWSReferenceId(db, oneLoadBalancerSecurityGroup.SecurityGroupId)
		if err == nil && secGroup != nil {
			updateLoadBalancerSecurityGroupId(db, oneLoadBalancerSecurityGroup, secGroup.Id)
			fmt.Println("	", oneLoadBalancerSecurityGroup.SecurityGroupId, "Loaded")
		} else {
			fmt.Println("	ERROR", oneLoadBalancerSecurityGroup.SecurityGroupId, "Not Loaded")
		}
	}

	fmt.Println("\nBuild relation for security_group to vpc")
	arrayOfSecGroup, _ := loadAllSecurityGroup(db)
	fmt.Println("  #Found:", len(arrayOfSecGroup))
	for _, oneSecGroup := range arrayOfSecGroup {
		vpc, err := loadVPCByVPCId(db, oneSecGroup.VpcId)
		if err == nil && vpc != nil {
			updateSecurityVPCId(db, oneSecGroup, vpc.Id)
			fmt.Println("	", oneSecGroup.VpcId, "Loaded")
		} else {
			fmt.Println("	ERROR", oneSecGroup.VpcId, "Not Loaded")
		}
	}

	fmt.Println("\nBuild relation for security_group_ip_permissions_egress_user_id_group_pairs to vpc")
	arrayOfEgress, _ := loadAllSecurityGroupIpPermissionEgressUserIdGroupPair(db)
	fmt.Println("  #Found:", len(arrayOfEgress))
	for _, oneEgress := range arrayOfEgress {
		vpc, err := loadVPCByVPCId(db, oneEgress.VpcId)
		if err == nil && vpc != nil {
			updateSecurityGroupIpPermissionEgressUserIdGroupPair(db, oneEgress, vpc.Id)
			fmt.Println("	", oneEgress.VpcId, "Loaded")
		} else {
			fmt.Println("	ERROR", oneEgress.VpcId, "Not Loaded")
		}
	}

	fmt.Println("\nBuild relation for security_group_ip_permissions_user_id_group_pairs to vpc")
	arrayOfPair, _ := loadAllSecurityGroupIpPermissionUserIdGroupPair(db)
	fmt.Println("  #Found:", len(arrayOfPair))
	for _, onePair := range arrayOfPair {
		vpc, err := loadVPCByVPCId(db, onePair.VpcId)
		if err == nil && vpc != nil {
			updateSecurityGroupIpPermissionUserIdGroupPair(db, onePair, vpc.Id)
			fmt.Println("	", onePair.VpcId, "Loaded")
		} else {
			fmt.Println("	ERROR", onePair.VpcId, "Not Loaded")
		}
	}

	fmt.Println("\nBuild relation for db_instance_db_subnet_group to vpc")
	arrayOfDBInstanceDBSubnetGroup := loadAllDBInstanceDBSubnetGroup(db)
	fmt.Println("  #Found:", len(arrayOfDBInstanceDBSubnetGroup))
	for _, oneDBInstanceDBSubnetGroup := range arrayOfDBInstanceDBSubnetGroup {
		vpc, err := loadVPCByVPCId(db, oneDBInstanceDBSubnetGroup.VpcId)
		if err == nil && vpc != nil {
			updateDBInstanceDBSubnetGroupVpcId(db, oneDBInstanceDBSubnetGroup, vpc.Id)
			fmt.Println("	", oneDBInstanceDBSubnetGroup.VpcId, "Loaded")
		} else {
			fmt.Println("	ERROR", oneDBInstanceDBSubnetGroup.VpcId, "Not Loaded")
		}
	}

	fmt.Println("\nBuild relation for db_instance_db_subnet_group_subnet to subnet")
	arrayOfDBInstanceDBSubnetGroupSubnet := loadAllDBInstanceDBSubnetGroupSubnet(db)
	fmt.Println("  #Found:", len(arrayOfDBInstanceDBSubnetGroupSubnet))
	for _, oneDBInstanceDBSubnetGroupSubnet := range arrayOfDBInstanceDBSubnetGroupSubnet {
		subnet, err := loadSubnetByAWSReferenceId(db, oneDBInstanceDBSubnetGroupSubnet.SubnetIdentifier)
		if err == nil && subnet != nil {
			updateDBInstanceDBSubnetGroupSubnetSubnetId(db, oneDBInstanceDBSubnetGroupSubnet, subnet.Id)
			fmt.Println("	", oneDBInstanceDBSubnetGroupSubnet.SubnetIdentifier, "Loaded")
		} else {
			fmt.Println("	ERROR", oneDBInstanceDBSubnetGroupSubnet.SubnetIdentifier, "Not Loaded")
		}
	}

	fmt.Println("\nBuild relation for db_instance_db_security_groups to security_groups")
	arrayOfDBInstanceDBSecurityGroup := loadAllDBInstanceDBSecurityGroup(db)
	fmt.Println("  #Found:", len(arrayOfDBInstanceDBSecurityGroup))
	for _, oneDBInstanceDBSecurityGroup := range arrayOfDBInstanceDBSecurityGroup {
		secGroup, err := loadSecurityGroupByAWSReferenceId(db, oneDBInstanceDBSecurityGroup.DBSecurityGroupName)
		if err == nil && secGroup != nil {
			updateDBInstanceDBSecurityGroupSecGroup(db, oneDBInstanceDBSecurityGroup, secGroup.Id)
			fmt.Println("	", oneDBInstanceDBSecurityGroup.DBSecurityGroupName, "Loaded")
		} else {
			fmt.Println("	ERROR", oneDBInstanceDBSecurityGroup.DBSecurityGroupName, "Not Loaded")
		}
	}

	fmt.Println("\nBuild relation for instance to vpc,ami,subnet")
	arrayOfInstance := loadAllInstance(db)
	fmt.Println("  #Found:", len(arrayOfInstance))
	for _, oneInstance := range arrayOfInstance {
		image, err := loadAmiByAWSReferenceId(db, oneInstance.ImageId)
		if err == nil && image != nil {
			subnet, err := loadSubnetByAWSReferenceId(db, oneInstance.SubnetId)
			if err == nil && subnet != nil {
				vpc, err := loadVPCByVPCId(db, oneInstance.VpcId)
				if err == nil && vpc != nil {
					updateInstanceId(db, oneInstance, image.Id, vpc.Id, subnet.Id)
					fmt.Println("	", oneInstance.ImageId, oneInstance.VpcId, oneInstance.SubnetId, "Loaded")
				} else {
					fmt.Println("	ERROR failed on", oneInstance.VpcId, "All Not Loaded")
				}
			} else {
				fmt.Println("	ERROR failed on", oneInstance.SubnetId, "All Not Loaded")
			}
		} else {
			fmt.Println("	ERROR failed on", oneInstance.ImageId, "All Not Loaded")
		}

	}

	fmt.Println("\nBuild relation for instance_network_interface to subnet,vpc")
	arrayOfInstanceNetworkInterface, _ := loadAllInstanceNetworkInterface(db)
	fmt.Println("  #Found:", len(arrayOfInstanceNetworkInterface))
	for _, oneInstanceNetworkInterface := range arrayOfInstanceNetworkInterface {
		subnet, err := loadSubnetByAWSReferenceId(db, oneInstanceNetworkInterface.SubnetId)
		if err == nil && subnet != nil {
			vpc, err := loadVPCByVPCId(db, oneInstanceNetworkInterface.VpcId)
			if err == nil && vpc != nil {
				updateInstanceNetworkInterface(db, oneInstanceNetworkInterface, subnet.Id, vpc.Id)
				fmt.Println("	", oneInstanceNetworkInterface.VpcId, oneInstanceNetworkInterface.SubnetId, "All Loaded")
			} else {
				fmt.Println("	ERROR failed on", oneInstanceNetworkInterface.VpcId, "All Loaded")
			}
		} else {
			fmt.Println("	ERROR failed on", oneInstanceNetworkInterface.SubnetId, "All Not Loaded")
		}
	}

	fmt.Println("\nBuild relation for instance_network_interface_group to security_groups")
	arrayOfInstanceNetworkInterfaceGroup, _ := loadAllInstanceNetworkInterfaceGroup(db)
	fmt.Println("  #Found:", len(arrayOfInstanceNetworkInterfaceGroup))
	for _, oneInstanceNetworkInterfaceGroup := range arrayOfInstanceNetworkInterfaceGroup {
		secGroup, err := loadSecurityGroupByAWSReferenceId(db, oneInstanceNetworkInterfaceGroup.GroupId)
		if err == nil && secGroup != nil {
			updateInstanceNetworkInterfaceGroup(db, oneInstanceNetworkInterfaceGroup, secGroup.Id)
			fmt.Println("	", oneInstanceNetworkInterfaceGroup.GroupId, "Loaded")
		} else {
			fmt.Println("	ERROR", oneInstanceNetworkInterfaceGroup.GroupId, "Not Loaded")
		}
	}

	fmt.Println("\nBuild relation for instance_security_group to security_groups")
	arrayOfInstanceSecurityGroup, _ := loadAllInstanceSecurityGroup(db)
	fmt.Println("  #Found:", len(arrayOfInstanceSecurityGroup))
	for _, oneInstanceSecurityGroup := range arrayOfInstanceSecurityGroup {
		secGroup, err := loadSecurityGroupByAWSReferenceId(db, oneInstanceSecurityGroup.GroupId)
		if err == nil && secGroup != nil {
			updateInstanceSecurityGroup(db, oneInstanceSecurityGroup, secGroup.Id)
			fmt.Println("	", oneInstanceSecurityGroup.GroupId, "Loaded")
		} else {
			fmt.Println("	ERROR", oneInstanceSecurityGroup.GroupId, "Not Loaded")
		}
	}

	fmt.Println("\nBuild relation for instance_block_device_mapping to volumes")
	arrayOfInstanceBlockDeviceMapping, _ := loadAllInstanceBlockDeviceMapping(db)
	fmt.Println("  #Found:", len(arrayOfInstanceBlockDeviceMapping))
	for _, oneInstanceBlockDeviceMapping := range arrayOfInstanceBlockDeviceMapping {
		volume, err := loadVolumebyAwsReferenceId(db, oneInstanceBlockDeviceMapping.EbsVolumeId)
		if err == nil && volume != nil {
			updateInstanceBlockDeviceMapping(db, oneInstanceBlockDeviceMapping, volume.Id)
			fmt.Println("	", oneInstanceBlockDeviceMapping.EbsVolumeId, "Loaded")
		} else {
			fmt.Println("	ERROR", oneInstanceBlockDeviceMapping.EbsVolumeId, "Not Loaded")
		}
	}

	fmt.Println("\nBuild relation for load_balancer to vpc")
	arrayOfLoadBalancer, _ := loadAllLoadBalancer(db)
	fmt.Println("  #Found:", len(arrayOfLoadBalancer))
	for _, oneLoadBalancer := range arrayOfLoadBalancer {
		vpc, err := loadVPCByVPCId(db, oneLoadBalancer.VpcId)
		if err == nil && vpc != nil {
			updateLoadBalancer(db, oneLoadBalancer, vpc.Id)
			fmt.Println("	", oneLoadBalancer.VpcId, "Loaded")
		} else {
			fmt.Println("	ERROR", oneLoadBalancer.VpcId, "Not Loaded")
		}
	}

	fmt.Println("\nBuild relation for load_balancer_availability_zone to subnet")
	arrayOfLoadBalancerAvailabilityZone := loadAllLoadBalancerAvailabilityZone(db)
	fmt.Println("  #Found:", len(arrayOfLoadBalancerAvailabilityZone))
	for _, oneLoadBalancerAvailabilityZone := range arrayOfLoadBalancerAvailabilityZone {
		subnet, err := loadSubnetByAWSReferenceId(db, oneLoadBalancerAvailabilityZone.SubnetId)
		if err == nil && subnet != nil {
			updateLoadBalancerAvailabilityZoneId(db, oneLoadBalancerAvailabilityZone, subnet.Id)
			fmt.Println("	", oneLoadBalancerAvailabilityZone.SubnetId, "Loaded")
		} else {
			fmt.Println("	ERROR", oneLoadBalancerAvailabilityZone.SubnetId, "Not Loaded")
		}
	}

	fmt.Println("\nBuild relation for auto_scaling_group to subnet")
	arrayOfAutoScalingGroup, _ := loadAllAutoScalingGroup(db)
	fmt.Println("  #Found:", len(arrayOfAutoScalingGroup))
	for _, oneAutoScalingGroup := range arrayOfAutoScalingGroup {
		subnet, err := loadSubnetByAWSReferenceId(db, oneAutoScalingGroup.VPCZoneIdentifier)
		if err == nil && subnet != nil {
			updateAutoScalingGroup(db, oneAutoScalingGroup, subnet.Id)
			fmt.Println("	", oneAutoScalingGroup.VPCZoneIdentifier, "Loaded")
		} else {
			fmt.Println("	ERROR", oneAutoScalingGroup.VPCZoneIdentifier, "Not Loaded")
		}
	}

	fmt.Println("\nBuild relation for load_balancer_target_group_health to instance")
	arrayOfLoadBalancerTargetGroupHealth, _ := loadAllLoadBalancerTargetGroupHealthByOrderUnique(db)
	fmt.Println("  #Found:", len(arrayOfLoadBalancerTargetGroupHealth))
	for _, oneLoadBalancerTargetGroupHealth := range arrayOfLoadBalancerTargetGroupHealth {
		instance, err := loadInstanceByAWSReferenceId(db, oneLoadBalancerTargetGroupHealth.TargetId)
		if err == nil && instance != nil {
			updateLoadBalancerTargetGroupHealthRefInstanceId(db, oneLoadBalancerTargetGroupHealth, instance.Id)
			fmt.Println("	", oneLoadBalancerTargetGroupHealth.TargetId, "Loaded")
		} else {
			fmt.Println("	", oneLoadBalancerTargetGroupHealth.TargetId, "Loaded")
		}
	}

	fmt.Println("\nBuild relation for load_balancer_listener_default_actions to load_balancer_target_group")
	arrayOfLoadBalancerDefaultActions, _ := loadAllLoadBalancerListenerDefaultActions(db)
	fmt.Println("  #Found:", len(arrayOfLoadBalancerDefaultActions))
	for _, oneLoadBalancerDefaultActions := range arrayOfLoadBalancerDefaultActions {
		loadBalancerTargetGroup, err := loadLoadBalancerTargetGroupByTargetGroupARN(db, oneLoadBalancerDefaultActions.TargetGroupArn)
		if err == nil && loadBalancerTargetGroup != nil {
			updateLoadBalancerListenerDefaultActionsTargetGroupId(db, oneLoadBalancerDefaultActions, loadBalancerTargetGroup.Id)
			fmt.Println("	", oneLoadBalancerDefaultActions.TargetGroupArn, "Loaded")
		} else {
			fmt.Println("	ERROR", oneLoadBalancerDefaultActions.TargetGroupArn, "Not Loaded")
		}
	}

	fmt.Println("\nBuild relation for auto_scaling_group_instance to instance")
	arrayOfAutoScalingGroupInstance := loadAllAutoScalingGroupInstance(db)
	fmt.Println("  #Found:", len(arrayOfAutoScalingGroupInstance))
	for _, oneAutoScalingGroupInstance := range arrayOfAutoScalingGroupInstance {
		instance, err := loadInstanceByAWSReferenceId(db, oneAutoScalingGroupInstance.InstanceId)
		if err == nil && instance != nil {
			updateAutoScalingGroupInstance(db, oneAutoScalingGroupInstance, instance.Id)
			fmt.Println("	", oneAutoScalingGroupInstance.InstanceId, "Loaded")
		} else {
			fmt.Println("	ERROR", oneAutoScalingGroupInstance.InstanceId, "Not Loaded")
		}
	}

	fmt.Println("\nBuild relation for volume_attachment to instance")
	arrayOfVolumeAttachment, _ := loadAllVolumeAttachment(db)
	fmt.Println("  #Found:", len(arrayOfVolumeAttachment))
	for _, oneVolumeAttachment := range arrayOfVolumeAttachment {
		instance, err := loadInstanceByAWSReferenceId(db, oneVolumeAttachment.InstanceId)
		if err == nil && instance != nil {
			updateVolumeAttachmentId(db, oneVolumeAttachment, instance.Id)
			fmt.Println("	", oneVolumeAttachment.InstanceId, "Loaded")
		} else {
			fmt.Println(err)
			fmt.Println("	ERROR", oneVolumeAttachment.InstanceId, "Not Loaded")
		}
	}

	fmt.Println("\nBuild relation for snapshot to volume")
	arrayOfSnapshots, _ := loadAllSnapshot(db)
	fmt.Println("  #Found:", len(arrayOfSnapshots))
	for _, oneSnapshot := range arrayOfSnapshots {
		volume, err := loadVolumebyAwsReferenceId(db, oneSnapshot.VolumeId)
		if err == nil && volume != nil {
			updateSnapshotVolumeId(db, oneSnapshot, volume.Id)
			fmt.Println("	", oneSnapshot.VolumeId, "Loaded")
		} else {
			fmt.Println(err)
			fmt.Println("	ERROR", oneSnapshot.VolumeId, "Not Loaded")
		}
	}

	fmt.Println("\nBuild relation for AMI to Snapshot")
	arrayOfAmis, _ := loadAllAmi(db)
	fmt.Println("  #Found:", len(arrayOfAmis))
	for _, oneAmi := range arrayOfAmis {
		snapshot, err := loadSnapshotbyAwsReferenceId(db, oneAmi.SnapShotId)
		if err == nil && snapshot != nil {
			updateAmiSnapshotId(db, oneAmi, snapshot.Id)
			fmt.Println("	", oneAmi.SnapShotId, "Loaded")
		} else {
			fmt.Println(err)
			fmt.Println("	ERROR", oneAmi.SnapShotId, "Not Loaded")
		}
	}

	//
}
