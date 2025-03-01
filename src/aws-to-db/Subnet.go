package main

import "fmt"

type Subnet struct {
	Id                      int64
	RefVpcId                int64
	SubnetId                string
	MapPublicAtLaunch       bool
	AvailabilityZoneId      string
	Name                    string
	CloudformationStackName string
	AvailableIpAddressCount int16
	State                   string
	AvailabilityZone        string
	OwnerId                 string
	CidrBlock               string
}

func NewSubnet(id int64, refVpcId int64, subnetId string, mapPublicAtLaunch bool, availabilityZoneId string, name string, cloudformationStackName string, availableIpAddressCount int16, state string, availabilityZone string, ownerId string, cidrBlock string) *Subnet {
	return &Subnet{
		Id:                      id,
		RefVpcId:                refVpcId,
		SubnetId:                subnetId,
		MapPublicAtLaunch:       mapPublicAtLaunch,
		AvailabilityZoneId:      availabilityZoneId,
		Name:                    name,
		CloudformationStackName: cloudformationStackName,
		AvailableIpAddressCount: availableIpAddressCount,
		State:                   state,
		AvailabilityZone:        availabilityZone,
		OwnerId:                 ownerId,
		CidrBlock:               cidrBlock}
}

func (d *Subnet) String() string {
	return fmt.Sprintf("Subnet Id(%d) ref_vpc_id(%d) subnetId(%s) name(%s) state(%s) CidrBlock(%s)\n",
		d.Id, d.RefVpcId, d.SubnetId, d.Name, d.State, d.CidrBlock)
}

func (d *Subnet) Dump() string {
	return fmt.Sprintf("Subent SubnetId(%s) MapPublicAtLaunch(%b) AvailabilityZoneId(%s) Name(%s) CloudformationStackName(%s) AvailableIpAddressCount(%d) State(%s) AvailabilityZone(%s) OwnerId(%s) CidrBlock(%s)",
		d.SubnetId,
		d.MapPublicAtLaunch,
		d.AvailabilityZoneId,
		d.Name,
		d.CloudformationStackName,
		d.AvailableIpAddressCount,
		d.State,
		d.AvailabilityZone,
		d.OwnerId,
		d.CidrBlock)
}
