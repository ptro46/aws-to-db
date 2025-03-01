package main

import (
	"database/sql"
	"fmt"
)

type LoadBalancer struct {
	Id                    int64
	LoadBalancerArn       string
	DNSName               string
	CanonicalHostedZoneId string
	CreatedTime           string
	LoadBalancerName      string
	Scheme                string
	VpcId                 string
	RefVpcId              sql.NullInt64
	StateCode             string
	StateReason           string
	Type                  string
	AvailabilityZones     []*LoadBalancerAvailabilityZone
	SecurityGroups        []*LoadBalancerSecurityGroup
	IpAddressType         string
}

func newLoadBalancer(Id int64, LoadBalancerArn string, DNSName string, CanonicalHostedZoneId string, CreatedTime string, LoadBalancerName string, Scheme string, VpcId string, RefVpcId sql.NullInt64, StateCode string, StateReason string, Type string, IpAddressType string) *LoadBalancer {
	return &LoadBalancer{Id: Id,
		LoadBalancerArn:       LoadBalancerArn,
		DNSName:               DNSName,
		CanonicalHostedZoneId: CanonicalHostedZoneId,
		CreatedTime:           CreatedTime,
		LoadBalancerName:      LoadBalancerName,
		Scheme:                Scheme,
		VpcId:                 VpcId,
		RefVpcId:              RefVpcId,
		StateCode:             StateCode,
		StateReason:           StateReason,
		Type:                  Type,
		IpAddressType:         IpAddressType}
}

func (d *LoadBalancer) loadDependencies(db *sql.DB) {
	arrayofAvailabilityZones, err := loadAllLoadBalancerAvailabilityZoneByOrderUnique(db)
	if err == nil {
		for _, zone := range arrayofAvailabilityZones {
			if zone.LoadBalancerId == d.Id {
				d.AvailabilityZones = append(d.AvailabilityZones, zone)
			}
		}
	} else {
		fmt.Println(err)
	}

	arrayofSecGroups, err := loadAllLoadBalancerSecurityGroupByOrderUnique(db)
	if err == nil {
		for _, secGroup := range arrayofSecGroups {
			if secGroup.LoadBalancerId == d.Id {
				d.SecurityGroups = append(d.SecurityGroups, secGroup)
			}
		}
	} else {
		fmt.Println(err)
	}

}

func (d *LoadBalancer) Dump(db *sql.DB) string {
	loadDump := fmt.Sprintf("Load Balancer LoadBalancerArn(%s) DNSName(%s) CanonicalHostedZoneId(%s) CreatedTime(%s) LoadBalancerName(%s) Scheme(%s) VpcId(%s) RefVpcId(%d) StateCode(%s) StateReason(%s) Type(%s) IpAddressType(%s)",
		d.LoadBalancerArn,
		d.DNSName,
		d.CanonicalHostedZoneId,
		d.CreatedTime,
		d.LoadBalancerName,
		d.Scheme,
		d.VpcId,
		d.RefVpcId,
		d.StateCode,
		d.StateReason,
		d.Type,
		d.IpAddressType)

	loadDump += "\n 	AvailabilityZones ["
	for _, zone := range d.AvailabilityZones {
		zone.loadDependencies(db)
		loadDump += "{" + zone.Dump() + "}"
	}

	loadDump += "]\n 	AvailabilitySecurityGroup ["
	for _, secGroup := range d.SecurityGroups {
		loadDump += "{" + secGroup.Dump() + "}"
	}
	loadDump += "]"
	return loadDump
}
