package main

import (
	"database/sql"
	"fmt"
)

func awsGetId(db *sql.DB, option string, value string, env string) int {
	retcode := 0

	if option == "vpc" {
		if value == "all" {
			arrayOfVPCs, _ := loadAllByOrderUnique(db)
			for _, vpc := range arrayOfVPCs {
				fmt.Printf("%s %s\n", vpc.VpcName, vpc.VpcId)
			}
		} else {
			vpc, _ := loadVPCByName(db, value)
			if vpc != nil {
				fmt.Printf("%s\n", vpc.VpcId)
			} else {
				retcode = 1
			}
		}
	} else if option == "subnet" {
		if value == "all" {
			arrayOfSubnets, _ := loadAllSubnetByOrderUnique(db)
			for _, subnet := range arrayOfSubnets {
				fmt.Printf("%s %s\n", subnet.Name, subnet.SubnetId)
			}
		} else {
			subnet, _ := loadSubnetByName(db, value)
			if subnet != nil {
				fmt.Printf("%s\n", subnet.SubnetId)
			} else {
				retcode = 2
			}
		}
	} else if option == "amis" {
		if value == "all" {
			arrayOfAMIs, _ := loadAllAmiByOrderUnique(db)
			for _, ami := range arrayOfAMIs {
				fmt.Printf("%s %s\n", ami.Name, ami.ImageId)
			}
		} else {
			ami, _ := loadAmiByName(db, value)
			if ami != nil {
				fmt.Printf("%s\n", ami.ImageId)
			} else {
				retcode = 3
			}
		}
	} else if option == "instances" {
		if value == "all" {
			arrayOfInstances, _ := loadAllInstanceByOrderUnique(db)
			for _, instance := range arrayOfInstances {
				fmt.Printf("%s %s\n", instance.InstanceId, instance.ImageId)
			}
		} else {
			instance, err := loadInstanceByName(db, value)
			if instance != nil {
				fmt.Printf("%s\n", instance.InstanceId)
			} else {
				fmt.Println(err)
				retcode = 3
			}
		}
	} else if option == "config_subnet" {
		if value == "all" {
			arrayOfConfigSubnets, _ := loadAllConfigSubnetByEnv(db, env)
			for _, configSubnets := range arrayOfConfigSubnets {
				fmt.Printf("%s %s\n", configSubnets.Name, configSubnets.CidrBlock)
			}
		} else {
			configSubnets, _ := loadConfigSubnetByNameAndEnv(db, value, env)
			if configSubnets != nil {
				fmt.Printf("%s\n", configSubnets.CidrBlock)
			} else {
				retcode = 3
			}
		}
	} else {
		retcode = -1
		fmt.Println("Invalid get option entered\n Acceptable Inputs:\n		subnets, vpc, buckets, amis, stacks, volumes, security_groups, db_instances, instances, load_balancers, auto_scaling_groups")
	}

	return retcode
}
