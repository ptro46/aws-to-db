package main

import (
	"database/sql"
	"fmt"
)

type SecurityGroup struct {
	Id                  int64
	Description         string
	GroupName           string
	IpPermissions       []*SecurityGroupIpPermission
	VpcId               string
	RefVpcId            sql.NullInt64
	OwnerId             string
	GroupId             string
	IpPermissionsEgress []*SecurityGroupIpPermission
	Tags                []*SecurityGroupTag
}

func NewSecurityGroup(id int64, description string, groupName string, vpcId string, refVpcId sql.NullInt64, ownerId string, groupId string) *SecurityGroup {
	return &SecurityGroup{
		Id:          id,
		Description: description,
		GroupName:   groupName,
		VpcId:       vpcId,
		RefVpcId:    refVpcId,
		OwnerId:     ownerId,
		GroupId:     groupId}
}

func (d *SecurityGroup) String() string {
	return fmt.Sprintf("SecurityGroup groupName(%s) vpcId(%s)", d.GroupName, d.VpcId)
}

func (d *SecurityGroup) loadDependencies(db *sql.DB) {
	arrayOfIpPermissions, err := loadAllSecurityGroupIpPermissionByOrderUnique(db)
	if err == nil {
		for _, ipPermission := range arrayOfIpPermissions {
			if ipPermission.SecurityGroupId == d.Id {
				d.IpPermissions = append(d.IpPermissions, ipPermission)
			}
		}
	} else {
		fmt.Println(err)
	}

	arrayOfIpPermissionsEgress, err := loadAllSecurityGroupIpPermissionEgress(db)
	if err == nil {
		for _, ipPermissionEgress := range arrayOfIpPermissionsEgress {
			if ipPermissionEgress.SecurityGroupId == d.Id {
				d.IpPermissionsEgress = append(d.IpPermissionsEgress, ipPermissionEgress)
			}
		}
	} else {
		fmt.Println(err)
	}

	arrayOfTags, err := loadAllSecurityGroupTag(db)
	if err == nil {
		for _, tag := range arrayOfTags {
			if tag.SecurityGroupId == d.Id {
				d.Tags = append(d.Tags, tag)
			}
		}
	} else {
		fmt.Println(err)
	}
}

func (d *SecurityGroup) Dump(db *sql.DB) string {
	dumpSecGroup := fmt.Sprintf("Security Group Description(%s) GroupName(%s) VpcId(%s) RefVpcId(%d) OwnerId(%s) GroupId(%s)",
		d.Description,
		d.GroupName,
		d.VpcId,
		d.RefVpcId,
		d.OwnerId,
		d.GroupId)
	dumpSecGroup += "\n 	IpPermissions ["
	for _, permission := range d.IpPermissions {
		permission.loadDependencies(db)
		dumpSecGroup += "{" + permission.Dump() + "}"
	}

	dumpSecGroup += "]\n 	IpPermissionsEgress ["
	for _, permissionEgress := range d.IpPermissionsEgress {
		permissionEgress.loadDependenciesEgress(db)
		dumpSecGroup += "{" + permissionEgress.Dump() + "}"
	}

	dumpSecGroup += "]\n 	Tag ["
	for _, tag := range d.Tags {
		dumpSecGroup += "{" + tag.Dump() + "}"
	}
	dumpSecGroup += "]"
	return dumpSecGroup
}
