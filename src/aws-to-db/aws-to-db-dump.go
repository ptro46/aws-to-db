package main

import (
	"database/sql"
	"fmt"
	"os"
)

func awsDump(db *sql.DB) {
	fmt.Println("Dumping Data")
	dumpFile, err := os.Create("DumpFile.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dumpFile.Close()

	arrayOfVpc, err := loadAllByOrderUnique(db)
	if err == nil {
		vpcHeader := fmt.Sprintf("VPC %d\n", len(arrayOfVpc))
		fmt.Printf("%s", vpcHeader)
		dumpFile.WriteString(vpcHeader)
		for _, vpc := range arrayOfVpc {
			dumpInfo := vpc.Dump()
			dumpFile.WriteString("\t" + dumpInfo + "\n")
			fmt.Printf("\t%s\n", dumpInfo)
		}
		fmt.Printf("\n")
		dumpFile.WriteString("\n")
	} else {
		fmt.Println(err)
	}

	arrayOfSubnet, err := loadAllSubnetByOrderUnique(db)
	if err == nil {
		subnetHeader := fmt.Sprintf("Subnets %d\n", len(arrayOfSubnet))
		fmt.Printf("%s", subnetHeader)
		dumpFile.WriteString(subnetHeader)
		for _, subnet := range arrayOfSubnet {
			dumpInfo := subnet.Dump()
			dumpFile.WriteString("\t" + dumpInfo + "\n")
			fmt.Printf("\t%s\n", dumpInfo)
		}
		fmt.Printf("\n")
		dumpFile.WriteString("\n")
	} else {
		fmt.Println(err)
	}

	arrayOfBucket, err := loadAllBucketByOrderUnique(db)
	if err == nil {
		bucketHeader := fmt.Sprintf("Bucket %d\n", len(arrayOfBucket))
		fmt.Printf("%s", bucketHeader)
		dumpFile.WriteString(bucketHeader)
		for _, bucket := range arrayOfBucket {
			dumpInfo := bucket.Dump()
			dumpFile.WriteString("\t" + dumpInfo + "\n")
			fmt.Printf("\t%s\n", dumpInfo)
		}
		fmt.Printf("\n")
		dumpFile.WriteString("\n")
	} else {
		fmt.Println(err)
	}

	arrayOfAMI, err := loadAllAmiByOrderUnique(db)
	if err == nil {
		amiHeader := fmt.Sprintf("Ami %d\n", len(arrayOfAMI))
		fmt.Printf("%s", amiHeader)
		dumpFile.WriteString(amiHeader)
		for _, ami := range arrayOfAMI {
			dumpInfo := ami.Dump()
			dumpFile.WriteString("\t" + dumpInfo + "\n")
			fmt.Printf("\t%s\n", dumpInfo)
		}
		fmt.Printf("\n")
		dumpFile.WriteString("\n")
	} else {
		fmt.Println(err)
	}

	arrayOfVolumes, err := loadAllVolumeByOrderUnique(db)
	if err == nil {
		volumeHeader := fmt.Sprintf("Volume %d\n", len(arrayOfVolumes))
		fmt.Printf("%s", volumeHeader)
		dumpFile.WriteString(volumeHeader)
		for _, volume := range arrayOfVolumes {
			volume.loadDependencies(db)
			dumpInfo := volume.Dump()
			dumpFile.WriteString("\t" + dumpInfo + "\n")
			fmt.Printf("\t%s\n", dumpInfo)
		}
		fmt.Printf("\n")
		dumpFile.WriteString("\n")
	} else {
		fmt.Println(err)
	}

	arrayOfLoadBalancer, err := loadAllLoadBalancerByOrderUnique(db)
	if err == nil {
		loadBalancerHeader := fmt.Sprintf("LoadBalancer %d\n", len(arrayOfLoadBalancer))
		fmt.Printf("%s", loadBalancerHeader)
		dumpFile.WriteString(loadBalancerHeader)
		for _, balacner := range arrayOfLoadBalancer {
			balacner.loadDependencies(db)
			dumpInfo := balacner.Dump(db)
			dumpFile.WriteString("\t" + dumpInfo + "\n")
			fmt.Printf("\t%s\n", dumpInfo)
		}
		fmt.Printf("\n")
		dumpFile.WriteString("\n")
	} else {
		fmt.Println(err)
	}

	arrayOfStack, err := loadAllStackByOrderUnique(db)
	if err == nil {
		stackHeader := fmt.Sprintf("Stack %d\n", len(arrayOfStack))
		fmt.Printf("%s", stackHeader)
		dumpFile.WriteString(stackHeader)
		for _, stack := range arrayOfStack {
			stack.loadDependencies(db)
			dumpInfo := stack.Dump()
			dumpFile.WriteString("\t" + dumpInfo + "\n")
			fmt.Printf("\t%s\n", dumpInfo)
		}
		fmt.Printf("\n")
		dumpFile.WriteString("\n")
	} else {
		fmt.Println(err)
	}

	arrayOfSecGroup, err := loadAllSecurityGroupByOrderUnique(db)
	if err == nil {
		secGroupHeader := fmt.Sprintf("Security Group %d\n", len(arrayOfSecGroup))
		fmt.Printf("%s", secGroupHeader)
		dumpFile.WriteString(secGroupHeader)
		for _, secGroup := range arrayOfSecGroup {
			secGroup.loadDependencies(db)
			dumpInfo := secGroup.Dump(db)
			dumpFile.WriteString("\t" + dumpInfo + "\n")
			fmt.Printf("\t%s\n", dumpInfo)
		}
		fmt.Printf("\n")
		dumpFile.WriteString("\n")
	} else {
		fmt.Println(err)
	}

	arrayOfReservations, err := loadAllReservationsByOrderUnique(db)
	if err == nil {
		resHeader := fmt.Sprintf("Reservation %d\n", len(arrayOfReservations))
		fmt.Printf("%s", resHeader)
		dumpFile.WriteString(resHeader)
		for _, res := range arrayOfReservations {
			res.loadDependencies(db)
			dumpInfo := res.Dump(db)
			dumpFile.WriteString("\t" + dumpInfo + "\n")
			fmt.Printf("\t%s\n", dumpInfo)
		}
		fmt.Printf("\n")
		dumpFile.WriteString("\n")
	} else {
		fmt.Println(err)
	}

	arrayOfDBInstance, err := loadAllDBInstanceByOrderUnique(db)
	if err == nil {
		DBHeader := fmt.Sprintf("DBInstance %d\n", len(arrayOfDBInstance))
		fmt.Printf("%s", DBHeader)
		dumpFile.WriteString(DBHeader)
		for _, DB := range arrayOfDBInstance {
			DB.loadDependencies(db)
			dumpInfo := DB.Dump(db)
			dumpFile.WriteString("\t" + dumpInfo + "\n")
			fmt.Printf("\t%s\n", dumpInfo)
		}
		fmt.Printf("\n")
		dumpFile.WriteString("\n")
	} else {
		fmt.Println(err)
	}

	arrayOfASG, err := loadAllAutoScalingGroupByOrderUnique(db)
	if err == nil {
		ASGHeader := fmt.Sprintf("AutoScalingGroup %d\n", len(arrayOfASG))
		fmt.Printf("%s", ASGHeader)
		dumpFile.WriteString(ASGHeader)
		for _, ASG := range arrayOfASG {
			ASG.loadDependencies(db)
			dumpInfo := ASG.Dump()
			dumpFile.WriteString("\t" + dumpInfo + "\n")
			fmt.Printf("\t%s\n", dumpInfo)
		}
		fmt.Printf("\n")
		dumpFile.WriteString("\n")
	} else {
		fmt.Println(err)
	}

}
