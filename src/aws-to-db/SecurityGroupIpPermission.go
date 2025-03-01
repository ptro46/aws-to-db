package main

import (
	"database/sql"
	"fmt"
)

type SecurityGroupIpPermission struct {
	Id               int64
	SecurityGroupId  int64
	FromPort         int16
	IpProtocol       string
	IpRanges         []*SecurityGroupIpPermissionsIpRange
	PrefixListIds    []*SecurityGroupIpPermissionsPrefixListId
	ToPort           int16
	UserIdGroupPairs []*SecurityGroupIpPermissionsUserIdGroupPair
}

func (d *SecurityGroupIpPermission) loadDependencies(db *sql.DB) {
	arrayOfIpRanges, err := loadAllSecurityGroupIpPermissionIpRangeByOrderUnique(db)
	if err == nil {
		for _, ipRange := range arrayOfIpRanges {
			if ipRange.SecurityGroupIpPermissionsId == d.Id {
				d.IpRanges = append(d.IpRanges, ipRange)
			}
		}
	} else {
		fmt.Println(err)
	}

	arrayOfPrefixListIds, err := loadAllSecurityGroupIpPermissionPrefixListIdByOrderUnique(db)
	if err == nil {
		for _, prefix := range arrayOfPrefixListIds {
			if prefix.SecurityGroupIpPermissionsId == d.Id {
				d.PrefixListIds = append(d.PrefixListIds, prefix)
			}
		}
	} else {
		fmt.Println(err)
	}

	arrayOfUserIdGroupPairs, err := loadAllSecurityGroupIpPermissionUserIdGroupPairByOrderUnique(db)
	if err == nil {
		for _, pair := range arrayOfUserIdGroupPairs {
			if pair.SecurityGroupIpPermissionsId == d.Id {
				d.UserIdGroupPairs = append(d.UserIdGroupPairs, pair)
			}
		}
	} else {
		fmt.Println(err)
	}
}

func (d *SecurityGroupIpPermission) loadDependenciesEgress(db *sql.DB) {
	arrayOfIpRanges, err := loadAllSecurityGroupIpPermissionEgressIpRangeByOrderUnique(db)

	if err == nil {
		for _, ipRange := range arrayOfIpRanges {
			if ipRange.SecurityGroupIpPermissionsId == d.Id {
				d.IpRanges = append(d.IpRanges, ipRange)
			}
		}
	} else {
		fmt.Println(err)
	}

	arrayOfPrefixListIds, err := loadAllSecurityGroupIpPermissionEgressPrefixListIdByOrderUnique(db)
	if err == nil {
		for _, prefix := range arrayOfPrefixListIds {
			if prefix.SecurityGroupIpPermissionsId == d.Id {
				d.PrefixListIds = append(d.PrefixListIds, prefix)
			}
		}
	} else {
		fmt.Println(err)
	}

	arrayOfUserIdGroupPairs, err := loadAllSecurityGroupIpPermissionEgressUserIdGroupPairByOrderUnique(db)
	if err == nil {
		for _, pair := range arrayOfUserIdGroupPairs {
			if pair.SecurityGroupIpPermissionsId == d.Id {
				d.UserIdGroupPairs = append(d.UserIdGroupPairs, pair)
			}
		}
	} else {
		fmt.Println(err)
	}
}

func (d *SecurityGroupIpPermission) Dump() string {
	dumpIpPermission := fmt.Sprintf("FromPort(%d) IpProtocol(%s) ToPort(%d)", d.FromPort, d.IpProtocol, d.ToPort)
	dumpIpPermission += "\n 	IpRange ["
	for _, ipRange := range d.IpRanges {
		dumpIpPermission += "{" + ipRange.Dump() + "}"
	}
	dumpIpPermission += "]\n	 PrefixListId ["
	for _, prefix := range d.PrefixListIds {
		dumpIpPermission += "{" + prefix.Dump() + "}"
	}
	dumpIpPermission += "]\n	 UserIdGroupPair  ["
	for _, pair := range d.UserIdGroupPairs {
		dumpIpPermission += "{" + pair.Dump() + "}"
	}
	dumpIpPermission += "]"
	return dumpIpPermission
}

type SecurityGroupIpPermissionsIpRange struct {
	Id                           int64
	SecurityGroupIpPermissionsId int64
	CidrIp                       string
	Description                  string
}

func (d *SecurityGroupIpPermissionsIpRange) Dump() string {
	return fmt.Sprintf("CidrIp(%s) Description(%s)", d.CidrIp, d.Description)
}

type SecurityGroupIpPermissionsPrefixListId struct {
	Id                           int64
	SecurityGroupIpPermissionsId int64
	Description                  string
	PrefixListId                 string
}

func (d *SecurityGroupIpPermissionsPrefixListId) Dump() string {
	return fmt.Sprintf("Description(%s) PrefixListId(%s)", d.Description, d.PrefixListId)
}

type SecurityGroupIpPermissionsUserIdGroupPair struct {
	Id                           int64
	SecurityGroupIpPermissionsId int64
	Description                  string
	GroupId                      string
	GroupName                    string
	PeeringStatus                string
	UserId                       string
	VpcId                        string
	RefVpcId                     int64
	VpcPeeringVonnectionId       string
}

func (d *SecurityGroupIpPermissionsUserIdGroupPair) Dump() string {
	return fmt.Sprintf("Description(%s) GroupId(%s) GroupName(%s) PeeringStatus(%s) UserId(%s) VpcId (%s) RefVpcId(%d) VpcPeeringVonnectionId(%s)",
		d.Description,
		d.GroupId,
		d.GroupName,
		d.PeeringStatus,
		d.UserId,
		d.VpcId,
		d.RefVpcId,
		d.VpcPeeringVonnectionId)
}

func NewSecurityGroupIpPermissionsUserIdGroupPair(id int64, securityGroupIpPermissionsId int64, description string, groupId string, groupName string, peeringStatus string, userId string, vpcId string, refVpcId int64, vpcPeeringVonnectionId string) *SecurityGroupIpPermissionsUserIdGroupPair {
	return &SecurityGroupIpPermissionsUserIdGroupPair{
		Id:                           id,
		SecurityGroupIpPermissionsId: securityGroupIpPermissionsId,
		Description:                  description,
		GroupId:                      groupId,
		GroupName:                    groupName,
		PeeringStatus:                peeringStatus,
		UserId:                       userId,
		VpcId:                        vpcId,
		RefVpcId:                     refVpcId,
		VpcPeeringVonnectionId:       vpcPeeringVonnectionId}
}

func NewSecurityGroupIpPermissionsPrefixListId(id int64, securityGroupIpPermissionsId int64, description string, prefixListId string) *SecurityGroupIpPermissionsPrefixListId {
	return &SecurityGroupIpPermissionsPrefixListId{
		Id:                           id,
		SecurityGroupIpPermissionsId: securityGroupIpPermissionsId,
		Description:                  description,
		PrefixListId:                 prefixListId}
}

func NewSecurityGroupIpPermission(id int64, securityGroupId int64, fromPort int16, ipProtocol string, toPort int16) *SecurityGroupIpPermission {
	return &SecurityGroupIpPermission{
		Id:              id,
		SecurityGroupId: securityGroupId,
		FromPort:        fromPort,
		IpProtocol:      ipProtocol,
		ToPort:          toPort}
}

func NewSecurityGroupIpPermissionsIpRange(id int64, securityGroupIpPermissionsId int64, cidrIp string, description string) *SecurityGroupIpPermissionsIpRange {
	return &SecurityGroupIpPermissionsIpRange{
		Id:                           id,
		SecurityGroupIpPermissionsId: securityGroupIpPermissionsId,
		CidrIp:                       cidrIp,
		Description:                  description}
}

func (d *SecurityGroupIpPermission) String() string {
	return fmt.Sprintf("SecurityGroupIpPermissionsEgress fromPort(%d)", d.FromPort)
}
