package main

import (
	"database/sql"
	"fmt"
)

type AutoScalingGroupLauchTemplateOverride struct {
	Id                 int64
	AutoScalingGroupId int64
	InstanceType       string
}

func (d *AutoScalingGroupLauchTemplateOverride) Dump() string {
	return fmt.Sprintf("InstanceType(%s)", d.InstanceType)
}

func NewAutoScalingGroupLauchTemplateOverride(Id int64, AutoScalingGroupId int64, InstanceType string) *AutoScalingGroupLauchTemplateOverride {
	return &AutoScalingGroupLauchTemplateOverride{Id,
		AutoScalingGroupId,
		InstanceType}
}

type AutoScalingGroupAvailabilityZone struct {
	Id                 int64
	AutoScalingGroupId int64
	AvailabilityZone   string
}

func (d *AutoScalingGroupAvailabilityZone) Dump() string {
	return fmt.Sprintf("AvailabilityZone(%s)", d.AvailabilityZone)
}

func NewAutoScalingGroupAvailabilityZone(Id int64, AutoScalingGroupId int64, AvailabilityZone string) *AutoScalingGroupAvailabilityZone {
	return &AutoScalingGroupAvailabilityZone{Id,
		AutoScalingGroupId,
		AvailabilityZone}
}

type AutoScalingGroupLoadBalancerName struct {
	Id                 int64
	AutoScalingGroupId int64
	LoadBalancerName   string
}

func (d *AutoScalingGroupLoadBalancerName) Dump() string {
	return fmt.Sprintf("LoadBalancerName(%s)", d.LoadBalancerName)
}

func NewAutoScalingGroupLoadBalancerName(Id int64, AutoScalingGroupId int64, LoadBalancerName string) *AutoScalingGroupLoadBalancerName {
	return &AutoScalingGroupLoadBalancerName{Id,
		AutoScalingGroupId,
		LoadBalancerName}
}

type AutoScalingGroupTargetGroupARN struct {
	Id                 int64
	AutoScalingGroupId int64
	ArnName            string
}

func (d *AutoScalingGroupTargetGroupARN) Dump() string {
	return fmt.Sprintf("ArnName(%s)", d.ArnName)
}

func NewAutoScalingGroupTargetGroupARN(Id int64, AutoScalingGroupId int64, ArnName string) *AutoScalingGroupTargetGroupARN {
	return &AutoScalingGroupTargetGroupARN{Id,
		AutoScalingGroupId,
		ArnName}
}

type AutoScalingInstance struct {
	Id                      int64
	AutoScalingGroupId      int64
	InstanceId              string
	RefInstanceId           sql.NullFloat64
	AvailabilityZone        string
	LifecycleState          string
	HealthStatus            string
	LaunchConfigurationName string
	LaunchTemplateId        string
	LaunchTemplateName      string
	LaunchTemplateVersion   string
	ProtectedFromScaleIn    bool
}

func (d *AutoScalingInstance) Dump() string {
	return fmt.Sprintf("InstanceId(%s) AvailabilityZone(%s) LifecycleState(%s) HealthStatus(%s) LaunchConfigurationName(%s) LaunchTemplateId(%s) LaunchTemplateName(%s) LaunchTemplateVersion(%s) ProtectedFromScaleIn(%t)",
		d.InstanceId,
		d.AvailabilityZone,
		d.LifecycleState,
		d.HealthStatus,
		d.LaunchConfigurationName,
		d.LaunchTemplateId,
		d.LaunchTemplateName,
		d.LaunchTemplateVersion,
		d.ProtectedFromScaleIn)
}

func NewAutoScalingInstance(Id int64, AutoScalingGroupId int64, InstanceId string, RefInstanceId sql.NullFloat64, AvailabilityZone string, LifecycleState string, HealthStatus string, LaunchConfigurationName string, LaunchTemplateId string, LaunchTemplateName string, LaunchTemplateVersion string, ProtectedFromScaleIn bool) *AutoScalingInstance {
	return &AutoScalingInstance{Id,
		AutoScalingGroupId,
		InstanceId,
		RefInstanceId,
		AvailabilityZone,
		LifecycleState,
		HealthStatus,
		LaunchConfigurationName,
		LaunchTemplateId,
		LaunchTemplateName,
		LaunchTemplateVersion,
		ProtectedFromScaleIn}
}

type AutoScalingGroupsSuspendedProcesses struct {
	Id                 int64
	AutoScalingGroupId int64
	ProcessName        string
	SuspensionReason   string
}

func (d *AutoScalingGroupsSuspendedProcesses) Dump() string {
	return fmt.Sprintf("ProcessName(%s) SuspensionReason(%s)", d.ProcessName, d.SuspensionReason)
}

func NewAutoScalingGroupsSuspendedProcesses(Id int64, AutoScalingGroupId int64, ProcessName string, SuspensionReason string) *AutoScalingGroupsSuspendedProcesses {
	return &AutoScalingGroupsSuspendedProcesses{Id,
		AutoScalingGroupId,
		ProcessName,
		SuspensionReason}
}

type AutoScalingGroupsEnabledMetric struct {
	Id                 int64
	AutoScalingGroupId int64
	Metric             string
	Granularity        string
}

func (d *AutoScalingGroupsEnabledMetric) Dump() string {
	return fmt.Sprintf("Metric(%s) Granularity(%s)", d.Metric, d.Granularity)
}

func NewAutoScalingGroupsEnabledMetric(Id int64, AutoScalingGroupId int64, Metric string, Granularity string) *AutoScalingGroupsEnabledMetric {
	return &AutoScalingGroupsEnabledMetric{Id,
		AutoScalingGroupId,
		Metric,
		Granularity}
}

type AutoScalingGroupsTag struct {
	Id                 int64
	AutoScalingGroupId int64
	ResourceId         string
	ResourceType       string
	Key                string
	Value              string
	PropagateAtLaunch  bool
}

func (d *AutoScalingGroupsTag) Dump() string {
	return fmt.Sprintf("ResourceId(%s) ResourceType(%s) Key(%s) Value(%s) PropagateAtLaunch(%t)", d.ResourceId, d.ResourceType, d.Key, d.Value, d.PropagateAtLaunch)
}

func NewAutoScalingGroupsTag(Id int64, AutoScalingGroupId int64, ResourceId string, ResourceType string, Key string, Value string, PropagateAtLaunch bool) *AutoScalingGroupsTag {
	return &AutoScalingGroupsTag{Id,
		AutoScalingGroupId,
		ResourceId,
		ResourceType,
		Key,
		Value,
		PropagateAtLaunch}
}

type AutoScalingGroupsTerminationPolicy struct {
	Id                 int64
	AutoScalingGroupId int64
	TerminationPolicy  string
}

func (d *AutoScalingGroupsTerminationPolicy) Dump() string {
	return fmt.Sprintf("TerminationPolicy(%s)", d.TerminationPolicy)
}

func NewAutoScalingGroupsTerminationPolicy(Id int64, AutoScalingGroupId int64, TerminationPolicy string) *AutoScalingGroupsTerminationPolicy {
	return &AutoScalingGroupsTerminationPolicy{Id,
		AutoScalingGroupId,
		TerminationPolicy}
}
