package main

import (
	"database/sql"
	"fmt"
)

type LoadBalancerListenerCertificat struct {
	Id                     int64
	LoadBalancerListenerId int64
	CertificatArn          string
}

func newLoadBalancerListenerCertificat(Id int64, LoadBalancerListenerId int64, CertificatArn string) *LoadBalancerListenerCertificat {
	return &LoadBalancerListenerCertificat{Id: Id,
		LoadBalancerListenerId: LoadBalancerListenerId,
		CertificatArn:          CertificatArn}
}

func (d *LoadBalancerListenerCertificat) Dump(db *sql.DB) string {
	loadDump := fmt.Sprintf("LoadBalancerListenerCertificat CertificatArn(%s)", d.CertificatArn)
	return loadDump
}
