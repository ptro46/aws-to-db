package main

import (
	"database/sql"
	"fmt"
)

type LoadBalancerListener struct {
	Id              int64
	LoadBalancerId  int64
	LoadBalancerArn string
	ListenerArn     string
	SslPolicies     string
	Port            int64
	Protocol        string
}

func newLoadBalancerListener(Id int64, LoadBalancerId int64, LoadBalancerArn string, ListenerArn string, SslPolicies string, Port int64, Protocol string) *LoadBalancerListener {
	return &LoadBalancerListener{Id: Id,
		LoadBalancerId:  LoadBalancerId,
		LoadBalancerArn: LoadBalancerArn,
		ListenerArn:     ListenerArn,
		SslPolicies:     SslPolicies,
		Port:            Port,
		Protocol:        Protocol}
}

func (d *LoadBalancerListener) Dump(db *sql.DB) string {
	loadDump := fmt.Sprintf("LoadBalancerListener LoadBalancerId(%s) LoadBalancerArn(%s) ListenerArn(%s) SslPolicies(%s) Port(%d) Protocol(%s)",
		d.LoadBalancerId,
		d.LoadBalancerArn,
		d.ListenerArn,
		d.SslPolicies,
		d.Port,
		d.Protocol)
	return loadDump
}
