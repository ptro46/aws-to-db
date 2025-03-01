package main

import (
	"database/sql"
	"fmt"
)

type AutoScalingGroup struct {
	Id                                                                           int64
	AutoScalingGroupName                                                         string
	AutoScalingGroupARN                                                          string
	LaunchConfigurationName                                                      string
	LaunchTemplateId                                                             string
	LaunchTemplateName                                                           string
	LaunchTemplateVersion                                                        string
	MixedInstancesPolicyLaunchTemplateSpecificationId                            string
	MixedInstancesPolicyLaunchTemplateSpecificationName                          string
	MixedInstancesPolicyLaunchTemplateSpecificationVersion                       string
	MixedInstancesPolicyinstancesDistributionOnDemandAllocationStrategy          string
	MixedInstancesPolicyinstancesDistributionOnDemandBaseCapacity                int16
	MixedInstancesPolicyinstancesDistributionOnDemandPercentageAboveBaseCapacity int16
	MixedInstancesPolicyinstancesDistributionOnDemandSpotAllocationStrategy      string
	MixedInstancesPolicyinstancesDistributionOnDemandSpotInstancePools           int16
	MixedInstancesPolicyinstancesDistributionOnDemandSpotMaxPrice                string
	MinSize                                                                      int16
	MaxSize                                                                      int16
	DesiredCapacity                                                              int16
	DefaultCooldown                                                              int16
	HealthCheckType                                                              string
	HealthCheckGracePeriod                                                       int16
	Instances                                                                    []*AutoScalingInstance
	CreatedTime                                                                  string
	SuspendedProcesses                                                           []*AutoScalingGroupsSuspendedProcesses
	PlacementGroup                                                               string
	VPCZoneIdentifier                                                            string
	EnabledMetrics                                                               []*AutoScalingGroupsEnabledMetric
	RefSubnetId                                                                  sql.NullInt64
	Status                                                                       string
	Tags                                                                         []*AutoScalingGroupsTag
	NewInstancesProtectedFromScaleIn                                             bool
	ServiceLinkedRoleARN                                                         string
	AvailabilityZones                                                            []*AutoScalingGroupAvailabilityZone
	LoadBalancerNames                                                            []*AutoScalingGroupLoadBalancerName
	TargetGroupARNs                                                              []*AutoScalingGroupTargetGroupARN
	TerminationPolicies                                                          []*AutoScalingGroupsTerminationPolicy
	Overrides                                                                    []*AutoScalingGroupLauchTemplateOverride
}

func (d *AutoScalingGroup) loadDependencies(db *sql.DB) {
	arrayOfOveride, err := loadAllAutoScalingGroupLauchTemplateOverrideByOrderUnique(db)
	if err == nil {
		for _, x := range arrayOfOveride {
			if x.AutoScalingGroupId == d.Id {
				d.Overrides = append(d.Overrides, x)
			}
		}
	} else {
		fmt.Println(err, "arrayOfOveride")
	}

	arrayOfInstances, err := loadAllAutoScalingInstanceByOrderUnique(db)
	if err == nil {
		for _, x := range arrayOfInstances {
			if x.AutoScalingGroupId == d.Id {
				d.Instances = append(d.Instances, x)
			}
		}
	} else {
		fmt.Println(err, "arrayOfInstances")
	}

	arrayOfSuspendedprocesses, err := loadAllAutoScalingGroupsSuspendedProcessesByOrderUnique(db)
	if err == nil {
		for _, x := range arrayOfSuspendedprocesses {
			if x.AutoScalingGroupId == d.Id {
				d.SuspendedProcesses = append(d.SuspendedProcesses, x)
			}
		}
	} else {
		fmt.Println(err, "arrayOfSuspendedprocesses")
	}

	arrayOfEnabledMetrics, err := loadAllAutoScalingGroupsEnabledMetricByOrderUnique(db)
	if err == nil {
		for _, x := range arrayOfEnabledMetrics {
			if x.AutoScalingGroupId == d.Id {
				d.EnabledMetrics = append(d.EnabledMetrics, x)
			}
		}
	} else {
		fmt.Println(err, "arrayOfEnabledMetrics")
	}

	arrayOfAvailabilityZones, err := loadAllAutoScalingGroupAvailabilityZoneByOrderUnique(db)
	if err == nil {
		for _, x := range arrayOfAvailabilityZones {
			if x.AutoScalingGroupId == d.Id {
				d.AvailabilityZones = append(d.AvailabilityZones, x)
			}
		}
	} else {
		fmt.Println(err, "arrayOfAvailabilityZones")
	}

	arrayOfLoadBalancerNames, err := loadAllAutoScalingGroupLoadBalancerNameByOrderUnique(db)
	if err == nil {
		for _, x := range arrayOfLoadBalancerNames {
			if x.AutoScalingGroupId == d.Id {
				d.LoadBalancerNames = append(d.LoadBalancerNames, x)
			}
		}
	} else {
		fmt.Println(err, "arrayOfLoadBalancerNames")
	}

	arrayOfTargetGroupARNs, err := loadAllAutoScalingGroupTargetGroupARNByOrderUnique(db)
	if err == nil {
		for _, x := range arrayOfTargetGroupARNs {
			if x.AutoScalingGroupId == d.Id {
				d.TargetGroupARNs = append(d.TargetGroupARNs, x)
			}
		}
	} else {
		fmt.Println(err, "arrayOfTargetGroupARNs")
	}

	arrayOfTerminationPolicies, err := loadAllAutoScalingGroupsTerminationPolicyByOrderUnique(db)
	if err == nil {
		for _, x := range arrayOfTerminationPolicies {
			if x.AutoScalingGroupId == d.Id {
				d.TerminationPolicies = append(d.TerminationPolicies, x)
			}
		}
	} else {
		fmt.Println(err, "arrayOfTerminationPolicies")
	}
}

func (d *AutoScalingGroup) Dump() string {
	dumpAuto := fmt.Sprintf("AutoScalingGroupName(%s) AutoScalingGroupARN(%s) LaunchConfigurationName(%s) LaunchTemplateId(%s) LaunchTemplateName(%s) LaunchTemplateVersion(%s) MixedInstancesPolicyLaunchTemplateSpecificationId(%s) MixedInstancesPolicyLaunchTemplateSpecificationName(%s) MixedInstancesPolicyLaunchTemplateSpecificationVersion(%s) MixedInstancesPolicyinstancesDistributionOnDemandAllocationStrategy(%s) MixedInstancesPolicyinstancesDistributionOnDemandBaseCapacity(%d) MixedInstancesPolicyinstancesDistributionOnDemandPercentageAboveBaseCapacity(%d) MixedInstancesPolicyinstancesDistributionOnDemandSpotAllocationStrategy(%s) MixedInstancesPolicyinstancesDistributionOnDemandSpotInstancePools(%d) MixedInstancesPolicyinstancesDistributionOnDemandSpotMaxPrice(%s) MinSize(%d) MaxSize(%d) DesiredCapacity(%d) DefaultCooldown(%d) HealthCheckType(%s) HealthCheckGracePeriod(%d) CreatedTime(%s) PlacementGroup(%s) VPCZoneIdentifier(%s) Status(%s) NewInstancesProtectedFromScaleIn(%t) ServiceLinkedRoleARN(%s)",
		d.AutoScalingGroupName,
		d.AutoScalingGroupARN,
		d.LaunchConfigurationName,
		d.LaunchTemplateId,
		d.LaunchTemplateName,
		d.LaunchTemplateVersion,
		d.MixedInstancesPolicyLaunchTemplateSpecificationId,
		d.MixedInstancesPolicyLaunchTemplateSpecificationName,
		d.MixedInstancesPolicyLaunchTemplateSpecificationVersion,
		d.MixedInstancesPolicyinstancesDistributionOnDemandAllocationStrategy,
		d.MixedInstancesPolicyinstancesDistributionOnDemandBaseCapacity,
		d.MixedInstancesPolicyinstancesDistributionOnDemandPercentageAboveBaseCapacity,
		d.MixedInstancesPolicyinstancesDistributionOnDemandSpotAllocationStrategy,
		d.MixedInstancesPolicyinstancesDistributionOnDemandSpotInstancePools,
		d.MixedInstancesPolicyinstancesDistributionOnDemandSpotMaxPrice,
		d.MinSize,
		d.MaxSize,
		d.DesiredCapacity,
		d.DefaultCooldown,
		d.HealthCheckType,
		d.HealthCheckGracePeriod,
		d.CreatedTime,
		d.PlacementGroup,
		d.VPCZoneIdentifier,
		d.Status,
		d.NewInstancesProtectedFromScaleIn,
		d.ServiceLinkedRoleARN)

	dumpAuto += "\n		Instances ["
	for _, x := range d.Instances {
		dumpAuto += "{" + x.Dump() + "}"
	}

	dumpAuto += "]\n		Overides ["
	for _, x := range d.Overrides {
		dumpAuto += "{" + x.Dump() + "}"
	}

	dumpAuto += "]\n		SuspendedProcesses ["
	for _, x := range d.SuspendedProcesses {
		dumpAuto += "{" + x.Dump() + "}"
	}

	dumpAuto += "]\n		EnabledMetrics ["
	for _, x := range d.EnabledMetrics {
		dumpAuto += "{" + x.Dump() + "}"
	}
	dumpAuto += "]\n		Tags ["
	for _, x := range d.Tags {
		dumpAuto += "{" + x.Dump() + "}"
	}
	dumpAuto += "]\n		AvailabilityZones ["
	for _, x := range d.AvailabilityZones {
		dumpAuto += "{" + x.Dump() + "}"
	}
	dumpAuto += "]\n		LoadBalancerNames ["
	for _, x := range d.LoadBalancerNames {
		dumpAuto += "{" + x.Dump() + "}"
	}
	dumpAuto += "]\n		TargetGroupARNs ["
	for _, x := range d.TargetGroupARNs {
		dumpAuto += "{" + x.Dump() + "}"
	}
	dumpAuto += "]\n		TerminationPolicies ["
	for _, x := range d.TerminationPolicies {
		dumpAuto += "{" + x.Dump() + "}"
	}
	dumpAuto += "]"
	return dumpAuto

}

func NewAutoScalingGroup(Id int64, AutoScalingGroupName string, AutoScalingGroupARN string, LaunchConfigurationName string, LaunchTemplateId string, LaunchTemplateName string, LaunchTemplateVersion string, MixedInstancesPolicyLaunchTemplateSpecificationId string, MixedInstancesPolicyLaunchTemplateSpecificationName string, MixedInstancesPolicyLaunchTemplateSpecificationVersion string, MixedInstancesPolicyinstancesDistributionOnDemandAllocationStrategy string, MixedInstancesPolicyinstancesDistributionOnDemandBaseCapacity int16, MixedInstancesPolicyinstancesDistributionOnDemandPercentageAboveBaseCapacity int16, MixedInstancesPolicyinstancesDistributionOnDemandSpotAllocationStrategy string, MixedInstancesPolicyinstancesDistributionOnDemandSpotInstancePools int16, MixedInstancesPolicyinstancesDistributionOnDemandSpotMaxPrice string, MinSize int16, MaxSize int16, DesiredCapacity int16, DefaultCooldown int16, HealthCheckType string, HealthCheckGracePeriod int16, CreatedTime string, PlacementGroup string, VPCZoneIdentifier string, RefSubnetId sql.NullInt64, Status string, NewInstancesProtectedFromScaleIn bool, ServiceLinkedRoleARN string) *AutoScalingGroup {
	return &AutoScalingGroup{Id: Id,
		AutoScalingGroupName:    AutoScalingGroupName,
		AutoScalingGroupARN:     AutoScalingGroupARN,
		LaunchConfigurationName: LaunchConfigurationName,
		LaunchTemplateId:        LaunchTemplateId,
		LaunchTemplateName:      LaunchTemplateName,
		LaunchTemplateVersion:   LaunchTemplateVersion,
		MixedInstancesPolicyLaunchTemplateSpecificationId:                            MixedInstancesPolicyLaunchTemplateSpecificationId,
		MixedInstancesPolicyLaunchTemplateSpecificationName:                          MixedInstancesPolicyLaunchTemplateSpecificationName,
		MixedInstancesPolicyLaunchTemplateSpecificationVersion:                       MixedInstancesPolicyLaunchTemplateSpecificationVersion,
		MixedInstancesPolicyinstancesDistributionOnDemandAllocationStrategy:          MixedInstancesPolicyinstancesDistributionOnDemandAllocationStrategy,
		MixedInstancesPolicyinstancesDistributionOnDemandBaseCapacity:                MixedInstancesPolicyinstancesDistributionOnDemandBaseCapacity,
		MixedInstancesPolicyinstancesDistributionOnDemandPercentageAboveBaseCapacity: MixedInstancesPolicyinstancesDistributionOnDemandPercentageAboveBaseCapacity,
		MixedInstancesPolicyinstancesDistributionOnDemandSpotAllocationStrategy:      MixedInstancesPolicyinstancesDistributionOnDemandSpotAllocationStrategy,
		MixedInstancesPolicyinstancesDistributionOnDemandSpotInstancePools:           MixedInstancesPolicyinstancesDistributionOnDemandSpotInstancePools,
		MixedInstancesPolicyinstancesDistributionOnDemandSpotMaxPrice:                MixedInstancesPolicyinstancesDistributionOnDemandSpotMaxPrice,
		MinSize:                          MinSize,
		MaxSize:                          MaxSize,
		DesiredCapacity:                  DesiredCapacity,
		DefaultCooldown:                  DefaultCooldown,
		HealthCheckType:                  HealthCheckType,
		HealthCheckGracePeriod:           HealthCheckGracePeriod,
		CreatedTime:                      CreatedTime,
		PlacementGroup:                   PlacementGroup,
		VPCZoneIdentifier:                VPCZoneIdentifier,
		RefSubnetId:                      RefSubnetId,
		Status:                           Status,
		NewInstancesProtectedFromScaleIn: NewInstancesProtectedFromScaleIn,
		ServiceLinkedRoleARN:             ServiceLinkedRoleARN}
}
