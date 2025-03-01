package main

type AwsLaunchTemplate struct {
	LaunchTemplateId   string `json:"LaunchTemplateId,omitempty"`
	LaunchTemplateName string `json:"LaunchTemplateName,omitempty"`
	Version            string `json:"Version,omitempty"`
}

type AwsMixedInstancesPolicy struct {
	LaunchTemplate        AwsMixedInstancePoliceLaunchTemplate `json:"LaunchTemplate,omitempty"`
	InstancesDistribution AwsInstancesDistribution             `json:"InstancesDistribution,omitempty"`
}

type AwsInstancesDistribution struct {
	OnDemandAllocationStrategy          string `json:"OnDemandAllocationStrategy,omitempty"`
	OnDemandBaseCapacity                int16  `json:"OnDemandBaseCapacity,omitempty"`
	OnDemandPercentageAboveBaseCapacity int16  `json:"OnDemandPercentageAboveBaseCapacity,omitempty"`
	SpotAllocationStrategy              string `json:"SpotAllocationStrategy,omitempty"`
	SpotInstancePools                   int16  `json:"SpotInstancePools,omitempty"`
	SpotMaxPrice                        string `json:"SpotMaxPrice,omitempty"`
}

type AwsMixedInstancePoliceLaunchTemplate struct {
	LaunchTemplateSpecification AwsLaunchTemplate `json:"LaunchTemplateSpecification,omitempty"`
	Overrides                   []AwsOverride     `json:"Overrides,omitempty"`
}

type AwsOverride struct {
	InstanceType string `json:"InstanceType,omitempty"`
}

type AwsAutoScalingGroupInstance struct {
	InstanceId              string            `json:"InstanceId,omitempty"`
	AvailabilityZone        string            `json:"AvailabilityZone,omitempty"`
	LifecycleState          string            `json:"LifecycleState,omitempty"`
	HealthStatus            string            `json:"HealthStatus,omitempty"`
	LaunchConfigurationName string            `json:"LaunchConfigurationName,omitempty"`
	LaunchTemplate          AwsLaunchTemplate `json:"LaunchTemplate,omitempty"`
	ProtectedFromScaleIn    bool              `json:"ProtectedFromScaleIn,omitempty"`
}

type AwsSuspendedProcesse struct {
	ProcessName      string `json:"ProcessName,omitempty"`
	SuspensionReason string `json:"SuspensionReason,omitempty"`
}

type AwsEnabledMetric struct {
	Metric      string `json:"Metric,omitempty"`
	Granularity string `json:"Granularity,omitempty"`
}

type AwsAutoScalingGroupTag struct {
	ResourceId        string `json:"ResourceId,omitempty"`
	ResourceType      string `json:"ResourceType,omitempty"`
	Key               string `json:"Key,omitempty"`
	Value             string `json:"Value,omitempty"`
	PropagateAtLaunch bool   `json:"PropagateAtLaunch,omitempty"`
}
