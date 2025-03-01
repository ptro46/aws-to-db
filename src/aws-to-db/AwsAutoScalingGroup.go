package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type AwsAutoScalingGroup struct {
	AutoScalingGroupName             string                        `json:"AutoScalingGroupName,omitempty"`
	AutoScalingGroupARN              string                        `json:"AutoScalingGroupARN,omitempty"`
	LaunchConfigurationName          string                        `json:"LaunchConfigurationName,omitempty"`
	LaunchTemplate                   AwsLaunchTemplate             `json:"LaunchTemplate,omitempty"`
	MixedInstancesPolicy             AwsMixedInstancesPolicy       `json:"MixedInstancesPolicy,omitempty"`
	MinSize                          int16                         `json:"MinSize,omitempty"`
	MaxSize                          int16                         `json:"MaxSize,omitempty"`
	DesiredCapacity                  int16                         `json:"DesiredCapacity,omitempty"`
	DefaultCooldown                  int16                         `json:"DefaultCooldown,omitempty"`
	AvailabilityZones                []string                      `json:"AvailabilityZones,omitempty"`
	LoadBalancerNames                []string                      `json:"LoadBalancerNames,omitempty"`
	TargetGroupARNs                  []string                      `json:"TargetGroupARNs,omitempty"`
	HealthCheckType                  string                        `json:"HealthCheckType,omitempty"`
	HealthCheckGracePeriod           int16                         `json:"HealthCheckGracePeriod,omitempty"`
	Instances                        []AwsAutoScalingGroupInstance `json:"Instances,omitempty"`
	CreatedTime                      string                        `json:"CreatedTime,omitempty"`
	SuspendedProcesses               []AwsSuspendedProcesse        `json:"SuspendedProcesses,omitempty"`
	PlacementGroup                   string                        `json:"PlacementGroup,omitempty"`
	VPCZoneIdentifier                string                        `json:"VPCZoneIdentifier,omitempty"`
	EnabledMetrics                   []AwsEnabledMetric            `json:"EnabledMetrics,omitempty"`
	Status                           string                        `json:"Status,omitempty"`
	Tags                             []AwsAutoScalingGroupTag      `json:"Tags,omitempty"`
	TerminationPolicies              []string                      `json:"TerminationPolicies,omitempty"`
	NewInstancesProtectedFromScaleIn bool                          `json:"NewInstancesProtectedFromScaleIn,omitempty"`
	ServiceLinkedRoleARN             string                        `json:"ServiceLinkedRoleARN,omitempty"`
}

type AwsAutoScalingGroups struct {
	AutoScalingGroups []AwsAutoScalingGroup `json:"AutoScalingGroups,omitempty"`
}

func awsAutoScalingGroupParseAndPersist(db *sql.DB, jsonString string) error {
	fmt.Printf("awsAutoScalingGroupParseAndPersist with json lenght %d\n", len(jsonString))
	awsAutoScalings := new(AwsAutoScalingGroups)
	errUnmarshal := json.Unmarshal([]byte(jsonString), awsAutoScalings)
	if errUnmarshal == nil {
		for _, awsAutoGroup := range awsAutoScalings.AutoScalingGroups {
			fmt.Printf("AutoScalingGroup %s\n", awsAutoGroup.AutoScalingGroupName)
			autoGroup := createAutoScalingGroup(db, awsAutoGroup.AutoScalingGroupName, awsAutoGroup.AutoScalingGroupARN, awsAutoGroup.LaunchConfigurationName, awsAutoGroup.LaunchTemplate.LaunchTemplateId, awsAutoGroup.LaunchTemplate.LaunchTemplateName, awsAutoGroup.LaunchTemplate.Version, awsAutoGroup.MixedInstancesPolicy.LaunchTemplate.LaunchTemplateSpecification.LaunchTemplateId, awsAutoGroup.MixedInstancesPolicy.LaunchTemplate.LaunchTemplateSpecification.LaunchTemplateName, awsAutoGroup.MixedInstancesPolicy.LaunchTemplate.LaunchTemplateSpecification.Version, awsAutoGroup.MixedInstancesPolicy.InstancesDistribution.OnDemandAllocationStrategy, awsAutoGroup.MixedInstancesPolicy.InstancesDistribution.OnDemandPercentageAboveBaseCapacity, awsAutoGroup.MixedInstancesPolicy.InstancesDistribution.OnDemandPercentageAboveBaseCapacity, awsAutoGroup.MixedInstancesPolicy.InstancesDistribution.SpotAllocationStrategy, awsAutoGroup.MixedInstancesPolicy.InstancesDistribution.SpotInstancePools, awsAutoGroup.MixedInstancesPolicy.InstancesDistribution.SpotMaxPrice, awsAutoGroup.MinSize, awsAutoGroup.MaxSize, awsAutoGroup.DesiredCapacity, awsAutoGroup.DefaultCooldown, awsAutoGroup.HealthCheckType, awsAutoGroup.HealthCheckGracePeriod, awsAutoGroup.CreatedTime, awsAutoGroup.PlacementGroup, awsAutoGroup.VPCZoneIdentifier, awsAutoGroup.Status, awsAutoGroup.NewInstancesProtectedFromScaleIn, awsAutoGroup.ServiceLinkedRoleARN)
			if autoGroup != nil {
				fmt.Printf("    AutoScalingGroup %s loaded\n", autoGroup.AutoScalingGroupName)
				createRestOfAutoScalingGroup(db, autoGroup, awsAutoGroup)
			} else {
				fmt.Printf("  ERROR  AutoScalingGroup %s not loaded\n", awsAutoGroup.AutoScalingGroupName)
			}
		}
	} else {
		return errUnmarshal
	}

	return nil
}

func createRestOfAutoScalingGroup(db *sql.DB, autoGroup *AutoScalingGroup, awsAutoGroup AwsAutoScalingGroup) {

	for _, awsTempOveride := range awsAutoGroup.MixedInstancesPolicy.LaunchTemplate.Overrides {
		tempOveride := createAutoScalingGroupLauchTemplateOverride(db, autoGroup.Id, awsTempOveride.InstanceType)
		if tempOveride != nil {
			fmt.Printf("   	 GroupLauchTemplateOverride %s loaded\n", tempOveride.InstanceType)
		} else {
			fmt.Printf("  	ERROR  GroupLauchTemplateOverride %s not loaded\n", awsTempOveride.InstanceType)
		}
	}

	for _, awsAvailZone := range awsAutoGroup.AvailabilityZones {
		availZone := createAutoScalingGroupAvailabilityZone(db, autoGroup.Id, awsAvailZone)
		if availZone != nil {
			fmt.Printf("   	 AvailabilityZone %s loaded\n", availZone.AvailabilityZone)
		} else {
			fmt.Printf("  	ERROR  AvailabilityZone %s not loaded\n", awsAvailZone)
		}
	}

	for _, awsBalancerName := range awsAutoGroup.LoadBalancerNames {
		balanceName := createAutoScalingGroupLoadBalancerName(db, autoGroup.Id, awsBalancerName)
		if balanceName != nil {
			fmt.Printf("   	 BalancerName %s loaded\n", balanceName)
		} else {
			fmt.Printf("  	ERROR  BalancerName %s not loaded\n", awsBalancerName)
		}
	}

	for _, awsGroupARN := range awsAutoGroup.TargetGroupARNs {
		groupARN := createAutoScalingGroupTargetGroupARN(db, autoGroup.Id, awsGroupARN)
		if groupARN != nil {
			fmt.Printf("   	 TargetGroupARN %s loaded\n", groupARN.ArnName)
		} else {
			fmt.Printf("  	ERROR  TargetGroupARN %s not loaded\n", awsGroupARN)
		}
	}

	for _, awsAutoInstance := range awsAutoGroup.Instances {
		autoInstance := createAutoScalingInstance(db, autoGroup.Id, awsAutoInstance.InstanceId, awsAutoInstance.AvailabilityZone, awsAutoInstance.LifecycleState, awsAutoInstance.HealthStatus, awsAutoInstance.LaunchConfigurationName, awsAutoInstance.LaunchTemplate.LaunchTemplateId, awsAutoInstance.LaunchTemplate.LaunchTemplateName, awsAutoInstance.LaunchTemplate.Version, awsAutoInstance.ProtectedFromScaleIn)
		if autoInstance != nil {
			fmt.Printf("   	 Instance %s loaded\n", autoInstance.LaunchTemplateName)
		} else {
			fmt.Printf("  	ERROR  Instance %s not loaded\n", awsAutoInstance.LaunchTemplate.LaunchTemplateName)
		}
	}

	for _, awsSuspendProcess := range awsAutoGroup.SuspendedProcesses {
		suspendProcess := createAutoScalingGroupsSuspendedProcesses(db, autoGroup.Id, awsSuspendProcess.ProcessName, awsSuspendProcess.SuspensionReason)
		if suspendProcess != nil {
			fmt.Printf("   	 SuspendProcess %s loaded\n", suspendProcess.ProcessName)
		} else {
			fmt.Printf("  	ERROR  SuspendProcess %s not loaded\n", awsSuspendProcess.ProcessName)
		}
	}

	for _, awsEnabledMetric := range awsAutoGroup.EnabledMetrics {
		enabledMetric := createAutoScalingGroupsEnabledMetric(db, autoGroup.Id, awsEnabledMetric.Metric, awsEnabledMetric.Granularity)
		if enabledMetric != nil {
			fmt.Printf("   	 EnabledMetric %s loaded\n", enabledMetric.Metric)
		} else {
			fmt.Printf("  	ERROR  EnabledMetric %s not loaded\n", awsEnabledMetric.Metric)
		}
	}

	for _, awsTag := range awsAutoGroup.Tags {
		tag := createAutoScalingGroupsTag(db, autoGroup.Id, awsTag.ResourceId, awsTag.ResourceType, awsTag.Key, awsTag.Value, awsTag.PropagateAtLaunch)
		if tag != nil {
			fmt.Printf("   	 Tag %s loaded\n", tag.Key)
		} else {
			fmt.Printf("  	ERROR  Tag %s not loaded\n", awsTag.Key)
		}
	}

	for _, awsTerminPolicy := range awsAutoGroup.TerminationPolicies {
		terminPolicy := createAutoScalingGroupsTerminationPolicy(db, autoGroup.Id, awsTerminPolicy)
		if terminPolicy != nil {
			fmt.Printf("   	 TerminationPolicy %s loaded\n", terminPolicy.TerminationPolicy)
		} else {
			fmt.Printf("  	ERROR  TerminationPolicy %s not loaded\n", awsTerminPolicy)
		}
	}

}
