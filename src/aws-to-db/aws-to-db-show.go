package main

import (
	"database/sql"
	"fmt"
	"github.com/fatih/color"
	"strings"
	"time"
)

var blue = color.New(color.FgBlue).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()
var white = color.New(color.FgWhite).SprintFunc()

func awsShow(db *sql.DB, subset string, filterName string) int {
	retcode := 0

	if subset == "ELB" {
		if len(filterName) == 0 {
			arrayOfViewElbTargetsHealth, err := loadAllViewElbTargetsHealth(db)
			if err == nil {
				displayDistinctInstanceName(db, arrayOfViewElbTargetsHealth)
			} else {
				fmt.Printf("%+v\n", err)
			}
		} else {
			arrayOfViewElbTargetsHealth, err := loadViewElbTargetsHealthByElbName(db, "%%"+filterName+"%%")
			if err == nil {
				displayDistinctInstanceName(db, arrayOfViewElbTargetsHealth)
			} else {
				fmt.Printf("%+v\n", err)
			}
		}

	} else if subset == "ASG" {
		if len(filterName) == 0 {
			// []*AutoScalingGroupActivities
			arrayOfAutoScalingGroupActivities, err := loadAllAutoScalingGroupActivitiesOrderByASGName(db)
			if err == nil {
				displayAutoScalingHistory(db, arrayOfAutoScalingGroupActivities)
			} else {
				fmt.Printf("%+v\n", err)
			}
		} else {
			arrayOfAutoScalingGroupActivities, err := loadAllAutoScalingGroupActivitiesWithASGName(db, filterName)
			if err == nil {
				displayAutoScalingHistory(db, arrayOfAutoScalingGroupActivities)
			} else {
				fmt.Printf("%+v\n", err)
			}
		}
	} else if subset == "LaunchConfig" {
		if len(filterName) == 0 {
			arrayOfAutoScalingGroup, err := loadAllAutoScalingGroupOrderByASGName(db)
			if err == nil {
				displayAutoScaling(db, arrayOfAutoScalingGroup)
			} else {
				fmt.Printf("%+v\n", err)
			}
		} else {
			arrayOfAutoScalingGroup, err := loadAllAutoScalingGroupWithASGName(db, filterName)
			if err == nil {
				displayAutoScaling(db, arrayOfAutoScalingGroup)
			} else {
				fmt.Printf("%+v\n", err)
			}
		}

	} else if subset == "LogsLinks" {
		arrayOfViewElbTargetsHealth, err := loadAllViewElbTargetsHealth(db)
		if err == nil {
			displayDistinctInstanceSymlinkLogs(db, arrayOfViewElbTargetsHealth)
		} else {
			fmt.Printf("%+v\n", err)
		}
	} else {
		retcode = -1
		fmt.Println("Invalid show option entered\n Acceptable Inputs:\n		elb")

	}

	return retcode
}

//---
//   ELB
func filterDuplicatedInstances(arrayOfViewElbTargetsHealth []*ViewElbTargetsHealth) []*ViewElbTargetsHealth {
	newArray := make([]*ViewElbTargetsHealth, 0, 0)
	previousInstanceId := ""
	for _, oneTarget := range arrayOfViewElbTargetsHealth {
		if oneTarget.InstanceId != previousInstanceId {
			newArray = append(newArray, oneTarget)
		}
		previousInstanceId = oneTarget.InstanceId
	}
	return newArray
}

func displayDistinctInstanceSymlinkLogs(db *sql.DB, arrayOfViewElbTargetsHealth []*ViewElbTargetsHealth) {
	orderArray := [19]string{
		"ELBDealsFrontend HTTPS:443",
		"ELBPartFrontend HTTPS:443",
		"ELBBackOfficeFrontend TCP:1443",
		"ELBMiddleOfficeFrontend TCP:1443",
		"RC",
		"ELBDealsBackend HTTPS:9443",
		"ELBPartenairesBackend HTTPS:9443",
		"ELBBackOfficeBackend HTTPS:9443",
		"ELBMiddleOfficeBackend HTTPS:9443",
		"RC",
		"ELBCypheringBackend HTTPS:9443",
		"ELBGatewayBackend HTTPS:9443",
		"ELBFilerBackend HTTPS:9443",
		"RC",
		"ELBHsmVPNBackend TCP:443",
		"ELBHsmBackendBackend HTTPS:9443",
		"RC",
		"ELBLogsBackendA UDP:514",
		"ELBLogsBackendB UDP:514"}
	arrayOfTargets := filterDuplicatedInstances(arrayOfViewElbTargetsHealth)
	targets := make(map[string][]*DisplayElbTargetsHealth)
	for _, oneTarget := range arrayOfTargets {
		key := fmt.Sprintf("%s %s:%d", oneTarget.LoadBalancerName, oneTarget.LoadBalancerProtocol, oneTarget.LoadBalancerPort)

		display := NewDisplayElbTargetsHealth(oneTarget)
		instanceDetails, err := loadInstanceByAWSReferenceId(db, oneTarget.InstanceId)
		if err == nil {
			display.PrivateDnsName = instanceDetails.PrivateDnsName
			display.PrivateIpAddress = instanceDetails.PrivateIpAddress
			display.PublicDnsName = instanceDetails.PublicDnsName
			display.PublicIpAddress = instanceDetails.PublicIpAddress
		}
		_, ok := targets[key]
		if ok {
			targets[key] = append(targets[key], display)
		} else {
			targets[key] = make([]*DisplayElbTargetsHealth, 0, 0)
			targets[key] = append(targets[key], display)
		}
	}
	for _, key := range orderArray {
		if key == "RC" {
			fmt.Printf("\n")
		} else {
			fmt.Printf("%s\n", key)
			elbTargets, ok := targets[key]
			if ok {
				for _, oneTarget := range elbTargets {
					if oneTarget.TargetHealthState == "healthy" {
						fmt.Printf("\t%s : ln -s ip-%s %s\n", blue(oneTarget.TargetHealthState), strings.ReplaceAll(oneTarget.PrivateIpAddress, ".", "-"), oneTarget.InstanceName)
					} else {
						fmt.Printf("\t%s : ln -s ip-%s %s\n", red(oneTarget.TargetHealthState), strings.ReplaceAll(oneTarget.PrivateIpAddress, ".", "-"), oneTarget.InstanceName)
					}
				}
			} else {
				fmt.Printf("\t%s\n", red("NOT FOUND"))
			}
		}
	}
}

func displayDistinctInstanceName(db *sql.DB, arrayOfViewElbTargetsHealth []*ViewElbTargetsHealth) {
	orderArray := []string{
		"ELBDealsFrontend HTTPS:443",
		"ELBPartFrontend HTTPS:443",
		"ELBBackOfficeFrontend TCP:1443",
		"ELBMiddleOfficeFrontend TCP:1443",
		"RC",
		"ELBDealsBackend HTTPS:9443",
		"ELBPartenairesBackend HTTPS:9443",
		"ELBBackOfficeBackend HTTPS:9443",
		"ELBMiddleOfficeBackend HTTPS:9443",
		"RC",
		"ELBCypheringBackend HTTPS:9443",
		"ELBGatewayBackend HTTPS:9443",
		"ELBFilerBackend HTTPS:9443",
		"RC",
		"ELBHsmVPNBackend TCP:443",
		"ELBHsmBackendBackend HTTPS:9443",
		"RC",
		"ELBLogsBackendA UDP:514",
		"ELBLogsBackendB UDP:514",
		"RC",
		"ELBPtroFrontend HTTP:80"}
	arrayOfTargets := filterDuplicatedInstances(arrayOfViewElbTargetsHealth)
	targets := make(map[string][]*DisplayElbTargetsHealth)
	for _, oneTarget := range arrayOfTargets {
		key := fmt.Sprintf("%s %s:%d", oneTarget.LoadBalancerName, oneTarget.LoadBalancerProtocol, oneTarget.LoadBalancerPort)

		display := NewDisplayElbTargetsHealth(oneTarget)
		instanceDetails, err := loadInstanceByAWSReferenceId(db, oneTarget.InstanceId)
		if err == nil {
			display.PrivateDnsName = instanceDetails.PrivateDnsName
			display.PrivateIpAddress = instanceDetails.PrivateIpAddress
			display.PublicDnsName = instanceDetails.PublicDnsName
			display.PublicIpAddress = instanceDetails.PublicIpAddress
		}
		_, ok := targets[key]
		if ok {
			targets[key] = append(targets[key], display)
		} else {
			targets[key] = make([]*DisplayElbTargetsHealth, 0, 0)
			targets[key] = append(targets[key], display)
		}
	}
	for _, key := range orderArray {
		if key == "RC" {
			fmt.Printf("\n")
		} else {
			fmt.Printf("%s\n", key)
			elbTargets, ok := targets[key]
			if ok {
				for _, oneTarget := range elbTargets {
					if oneTarget.TargetHealthState == "healthy" {
						fmt.Printf("\t%s %s %s %s  %s %s %s\n", oneTarget.Type, oneTarget.TargetType, blue(oneTarget.TargetHealthState), white(oneTarget.getIps()), oneTarget.InstanceId, oneTarget.InstanceName, oneTarget.HealthCheckPath)
					} else {
						fmt.Printf("\t%s %s %s %s %s %s %s\n", oneTarget.Type, oneTarget.TargetType, red(oneTarget.TargetHealthState), white(oneTarget.getIps()), oneTarget.InstanceId, oneTarget.InstanceName, oneTarget.HealthCheckPath)
					}
				}
			} else {
				fmt.Printf("\t%s\n", red("NOT FOUND"))
			}
		}
	}
}

//   ELB
//---

//---
//   ASG
// select id,description,auto_scaling_group_name,start_time,progress,end_time,status_code,details from public.auto_scaling_activities ;
func displayAutoScalingHistory(db *sql.DB, arrayOfAutoScalingGroupActivities []*AutoScalingGroupActivities) {
	orderArray := []string{
		"asg-deals-frontend-1",
		"asg-deals-frontend-2",
		"asg-partenaires-frontend-1",
		"asg-partenaires-frontend-2",
		"asg-backoffice-frontend-1",
		"asg-backoffice-frontend-2",
		"asg-middleoffice-frontend-1",
		"asg-middleoffice-frontend-2",
		"RC",
		"asg-deals-backend-1",
		"asg-deals-backend-2",
		"asg-partenaires-backend-1",
		"asg-partenaires-backend-2",
		"asg-backoffice-backend-1",
		"asg-backoffice-backend-2",
		"asg-middleoffice-backend-1",
		"asg-middleoffice-backend-2",
		"RC",
		"asg-cyphering-backend-1",
		"asg-cyphering-backend-2",
		"asg-gateway-backend-1",
		"asg-gateway-backend-2",
		"asg-filer-backend-1",
		"asg-filer-backend-2",
		"RC",
		"asg-hsm-backend-1",
		"asg-hsm-backend-2",
		"RC",
		"asg-ptro-frontend-1",
		"asg-ptro-frontend-2"}
	targets := make(map[string][]*AutoScalingGroupActivities)
	for _, oneTarget := range arrayOfAutoScalingGroupActivities {
		key := oneTarget.AutoScalingGroupName
		_, ok := targets[key]
		if ok {
			targets[key] = append(targets[key], oneTarget)
		} else {
			targets[key] = make([]*AutoScalingGroupActivities, 0, 0)
			targets[key] = append(targets[key], oneTarget)
		}
	}
	// 2020-09-04T06:43:06.560Z

	for _, key := range orderArray {
		if key == "RC" {
			fmt.Printf("\n")
		} else {
			fmt.Printf("%s\n", key)
			asgTargets, ok := targets[key]
			if ok {
				for _, oneTarget := range asgTargets {
					tStart, errStart := time.Parse(time.RFC3339, oneTarget.StartTime)

					if errStart == nil {
						fmt.Printf("\t%s %s %s\n", tStart.Format("2006-01-02 15:04:05"), oneTarget.Description, oneTarget.StatusCode)
					} else {
						fmt.Printf("\t%s %s %d%% %s %s\n", oneTarget.Description, oneTarget.StartTime, oneTarget.Progress, oneTarget.EndTime, oneTarget.StatusCode)
					}
				}
			} else {
				fmt.Printf("\t%s\n", red("NOT FOUND"))
			}
		}
	}
}

func displayAutoScaling(db *sql.DB, arrayOfAutoScalingGroup []*AutoScalingGroup) {
	for _, oneASG := range arrayOfAutoScalingGroup {
		fmt.Printf("%s %s\n", oneASG.AutoScalingGroupName, oneASG.LaunchConfigurationName)
	}
}

//   ASG
//---
