package main

import (
	"database/sql"
	"fmt"
)

type LoadBalancerAvailabilityZone struct {
	Id                    int64
	LoadBalancerId        int64
	ZoneName              string
	SubnetId              string
	RefSubnetId           sql.NullInt64
	LoadBalancerAddresses []*LoadBalancerAddress
}

func newLoadBalancerAvailabilityZone(Id int64, LoadBalancerId int64, ZoneName string, SubnetId string, RefSubnetId sql.NullInt64) *LoadBalancerAvailabilityZone {
	return &LoadBalancerAvailabilityZone{Id: Id,
		LoadBalancerId: LoadBalancerId,
		ZoneName:       ZoneName,
		SubnetId:       SubnetId,
		RefSubnetId:    RefSubnetId}
}

func (d *LoadBalancerAvailabilityZone) loadDependencies(db *sql.DB) {
	arrayofLoadBalancerAddress, err := loadAllLoadBalancerAddressByOrderUnique(db)
	if err == nil {
		for _, address := range arrayofLoadBalancerAddress {
			if address.LoadBalancerAvailabilityZoneId == d.Id {
				d.LoadBalancerAddresses = append(d.LoadBalancerAddresses, address)
			}
		}
	} else {
		fmt.Println(err)
	}
}

func (d *LoadBalancerAvailabilityZone) Dump() string {
	dumpZone := fmt.Sprintf("ZoneName(%s) SubnetId(%s) RefSubnetId(%d)",
		d.ZoneName,
		d.SubnetId,
		d.RefSubnetId)

	dumpZone += "\n 	LoadBalancerAddress ["
	for _, address := range d.LoadBalancerAddresses {
		dumpZone += "{" + address.Dump() + "}"
	}
	dumpZone += "]"
	return dumpZone
}

type LoadBalancerAddress struct {
	Id                             int64
	LoadBalancerAvailabilityZoneId int64
	IpAddress                      string
	AllocationId                   string
}

func newLoadBalancerAddress(Id int64, LoadBalancerAvailabilityZoneId int64, IpAddress string, AllocationId string) *LoadBalancerAddress {
	return &LoadBalancerAddress{
		Id,
		LoadBalancerAvailabilityZoneId,
		IpAddress,
		AllocationId}
}

func (d *LoadBalancerAddress) Dump() string {
	return fmt.Sprintf("IpAddress(%s) AllocationId(%s)",
		d.IpAddress,
		d.AllocationId)
}

type LoadBalancerSecurityGroup struct {
	Id                 int64
	LoadBalancerId     int64
	SecurityGroupId    string
	RefSecurityGroupId sql.NullInt64
}

func newLoadBalancerSecurityGroup(Id int64, LoadBalancerId int64, SecurityGroupId string, RefSecurityGroupId sql.NullInt64) *LoadBalancerSecurityGroup {
	return &LoadBalancerSecurityGroup{Id, LoadBalancerId, SecurityGroupId, RefSecurityGroupId}
}

func (d *LoadBalancerSecurityGroup) Dump() string {
	return fmt.Sprintf("SecurityGroupId(%s) RefSecurityGroupId(%d)",
		d.SecurityGroupId,
		d.RefSecurityGroupId)
}
