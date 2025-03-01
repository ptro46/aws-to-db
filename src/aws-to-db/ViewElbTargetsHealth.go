package main

import (
	"fmt"
)

type ViewElbTargetsHealth struct {
	Id                         int64
	LoadBalancerArn            string
	LoadBalancerName           string
	LoadBalancerPort           int64
	LoadBalancerProtocol       string
	Type                       string
	HealthCheckIntervalSeconds int
	HealthCheckTimeoutSeconds  int
	HealthyThresholdCount      int
	UnhealthyThresholdCount    int
	HealthCheckPath            string
	Httpcode                   string
	TargetType                 string
	TargetHealthState          string
	InstanceId                 string
	InstanceName               string
}

func NewViewElbTargetsHealth(id int64, loadBalancerArn string, loadBalancerName string, loadBalancerPort int64, loadBalancerProtocol string, typee string, healthCheckIntervalSeconds int, healthCheckTimeoutSeconds int, healthyThresholdCount int, unhealthyThresholdCount int, healthCheckPath string, httpcode string, targetType string, targetHealthState string, instanceId string, instanceName string) *ViewElbTargetsHealth {
	return &ViewElbTargetsHealth{
		Id:                         id,
		LoadBalancerArn:            loadBalancerArn,
		LoadBalancerName:           loadBalancerName,
		LoadBalancerPort:           loadBalancerPort,
		LoadBalancerProtocol:       loadBalancerProtocol,
		Type:                       typee,
		HealthCheckIntervalSeconds: healthCheckIntervalSeconds,
		HealthCheckTimeoutSeconds:  healthCheckTimeoutSeconds,
		HealthyThresholdCount:      healthyThresholdCount,
		UnhealthyThresholdCount:    unhealthyThresholdCount,
		HealthCheckPath:            healthCheckPath,
		Httpcode:                   httpcode,
		TargetType:                 targetType,
		TargetHealthState:          targetHealthState,
		InstanceId:                 instanceId,
		InstanceName:               instanceName}
}

func (d *ViewElbTargetsHealth) String() string {
	return fmt.Sprintf("ViewElbTargetsHealth loadBalancerName(%s) loadBalancerArn(%s)", d.LoadBalancerName, d.LoadBalancerArn)
}

type DisplayElbTargetsHealth struct {
	Type              string
	TargetType        string
	TargetHealthState string
	InstanceId        string
	InstanceName      string
	HealthCheckPath   string
	PrivateDnsName    string
	PrivateIpAddress  string
	PublicDnsName     string
	PublicIpAddress   string
}

func NewDisplayElbTargetsHealth(viewElbTargetsHealth *ViewElbTargetsHealth) *DisplayElbTargetsHealth {
	return &DisplayElbTargetsHealth{
		Type:              viewElbTargetsHealth.Type,
		TargetType:        viewElbTargetsHealth.TargetType,
		TargetHealthState: viewElbTargetsHealth.TargetHealthState,
		InstanceId:        viewElbTargetsHealth.InstanceId,
		InstanceName:      viewElbTargetsHealth.InstanceName,
		HealthCheckPath:   viewElbTargetsHealth.HealthCheckPath,
		PrivateDnsName:    "",
		PrivateIpAddress:  "",
		PublicDnsName:     "",
		PublicIpAddress:   ""}
}

func (d *DisplayElbTargetsHealth) getIps() string {
	ret := ""
	if d.PrivateIpAddress != "" && d.PublicIpAddress != "" {
		ret = fmt.Sprintf("%-15v / %-15v", d.PrivateIpAddress, d.PublicIpAddress)

	} else if d.PrivateIpAddress == "" && d.PublicIpAddress != "" {
		ret = fmt.Sprintf("%-15v / %-15v", " ", d.PublicIpAddress)

	} else if d.PrivateIpAddress != "" && d.PublicIpAddress == "" {
		ret = fmt.Sprintf("%-15v / %-15v", d.PrivateIpAddress, " ")

	} else if d.PrivateIpAddress == "" && d.PublicIpAddress == "" {
		ret = fmt.Sprintf("%-15v / %-15v", " ", " ")

	}
	return ret
}
