package main

import "database/sql"

type LoadBalancerTargetGroup struct {
	Id                         int64
	TargetGroupArn             string
	TargetGroupName            string
	Protocol                   string
	Port                       int
	VpcId                      string
	RefVpcId                   sql.NullInt64
	HealthCheckProtocol        string
	HealthCheckPort            string
	HealthCheckEnabled         bool
	HealthCheckIntervalSeconds int
	HealthCheckTimeoutSeconds  int
	HealthyThresholdCount      int
	UnhealthyThresholdCount    int
	HealthCheckPath            string
	HttpCode                   string
	TargetType                 string
}

func NewLoadBalancerTargetGroup(Id int64, TargetGroupArn string, TargetGroupName string, Protocol string, Port int, VpcId string, RefVpcId sql.NullInt64, HealthCheckProtocol string, HealthCheckPort string, HealthCheckEnabled bool, HealthCheckIntervalSeconds int, HealthCheckTimeoutSeconds int, HealthyThresholdCount int, UnhealthyThresholdCount int, HealthCheckPath string, HttpCode string, TargetType string) *LoadBalancerTargetGroup {
	return &LoadBalancerTargetGroup{Id: Id,
		TargetGroupArn:             TargetGroupArn,
		TargetGroupName:            TargetGroupName,
		Protocol:                   Protocol,
		Port:                       Port,
		VpcId:                      VpcId,
		RefVpcId:                   RefVpcId,
		HealthCheckProtocol:        HealthCheckProtocol,
		HealthCheckPort:            HealthCheckPort,
		HealthCheckEnabled:         HealthCheckEnabled,
		HealthCheckIntervalSeconds: HealthCheckIntervalSeconds,
		HealthCheckTimeoutSeconds:  HealthCheckTimeoutSeconds,
		HealthyThresholdCount:      HealthyThresholdCount,
		UnhealthyThresholdCount:    UnhealthyThresholdCount,
		HealthCheckPath:            HealthCheckPath,
		HttpCode:                   HttpCode,
		TargetType:                 TargetType}

}

type LoadBalancerTargetGroupLoadBalancerARN struct {
	Id                        int64
	LoadBalancerTargetGroupId int64
	LoadBalancerArn           string
	RefLoadBalancerId         sql.NullInt64
}

func NewLoadBalancerTargetGroupLoadBalancerARN(Id int64, LoadBalancerTargetGroupId int64, LoadBalancerArn string, RefLoadBalancerId sql.NullInt64) *LoadBalancerTargetGroupLoadBalancerARN {
	return &LoadBalancerTargetGroupLoadBalancerARN{Id,
		LoadBalancerTargetGroupId,
		LoadBalancerArn,
		RefLoadBalancerId}
}

type LoadBalancerTargetGroupHealth struct {
	Id                        int64
	LoadBalancerTargetGroupId int64
	HealthCheckPort           string
	TargetId                  string
	RefInstanceId             sql.NullInt64
	Port                      int
	TargetHealthState         string
	TargetHealthReason        string
	TargetHealthDescription   string
}

func NewLoadBalancerTargetGroupHealth(Id int64, LoadBalancerTargetGroupId int64, HealthCheckPort string, TargetId string, RefInstanceId sql.NullInt64, Port int, TargetHealthState string, TargetHealthReason string, TargetHealthDescription string) *LoadBalancerTargetGroupHealth {
	return &LoadBalancerTargetGroupHealth{Id,
		LoadBalancerTargetGroupId,
		HealthCheckPort,
		TargetId,
		RefInstanceId,
		Port,
		TargetHealthState,
		TargetHealthReason,
		TargetHealthDescription}
}
