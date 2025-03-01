package main

import (
	"database/sql"
	"fmt"
)

type LoadBalancerDefaultAction struct {
	Id                     int64
	LoadBalancerListenerId int64
	TargetGroupArn         string
	TargetGroupId          sql.NullInt64
	Type                   string
}

func newLoadBalancerDefaultAction(Id int64, LoadBalancerListenerId int64, TargetGroupArn string, TargetGroupId sql.NullInt64, Type string) *LoadBalancerDefaultAction {
	return &LoadBalancerDefaultAction{Id: Id,
		LoadBalancerListenerId: LoadBalancerListenerId,
		TargetGroupArn:         TargetGroupArn,
		TargetGroupId:          TargetGroupId,
		Type:                   Type}
}

func (d *LoadBalancerDefaultAction) Dump(db *sql.DB) string {
	loadDump := fmt.Sprintf("LoadBalancerDefaultAction TargetGroupArn(%s) Type(%s) CanonicalHostedZoneId(%s) CreatedTime(%s) LoadBalancerName(%s) Scheme(%s) VpcId(%s) RefVpcId(%d) StateCode(%s) StateReason(%s) Type(%s) IpAddressType(%s)",
		d.TargetGroupArn,
		d.Type)

	return loadDump
}
