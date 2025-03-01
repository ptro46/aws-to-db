package main

import (
	"database/sql"
	"fmt"
)

type DBInstanceGroupType struct {
	Id           int64
	DbInstanceId int64
	Name         string
	Status       string
}

func (d *DBInstanceGroupType) Dump() string {
	return fmt.Sprintf("Name(%s) Status(%s)", d.Name, d.Status)
}

func NewDBInstanceGroupType(Id int64, DbInstanceId int64, Name string, Status string) *DBInstanceGroupType {
	return &DBInstanceGroupType{Id,
		DbInstanceId, Name, Status}
}

/*
type DBInstanceOptionGroupMemberships struct {
	Id              int64
	DbInstanceId    int64
	OptionGroupName string
	Status          string
}

type DBInstanceDBParameterGroups struct {
	Id                   int64
	DbInstanceId         int64
	DBParameterGroupName string
	ParameterApplyStatus string
}*/

type DBInstanceVpcSecurityGroups struct {
	Id                 int64
	DbInstanceId       int64
	VpcSecurityGroupId string
	RefSecurityGroupId sql.NullInt64
	Status             string
}

func NewDBInstanceVpcSecurityGroups(Id int64, DbInstanceId int64, VpcSecurityGroupId string, RefVpcSecurityGroupId sql.NullInt64, Status string) *DBInstanceVpcSecurityGroups {
	return &DBInstanceVpcSecurityGroups{Id,
		DbInstanceId,
		VpcSecurityGroupId,
		RefVpcSecurityGroupId,
		Status}
}

type DBInstanceDBSecurityGroup struct {
	Id                  int64
	DbInstanceId        int64
	DBSecurityGroupName string
	RefSecurityGroupId  sql.NullInt64
	Status              string
}

func (d *DBInstanceDBSecurityGroup) Dump() string {
	return fmt.Sprintf("DBSecurityGroupName(%s) Status(%s)", d.DBSecurityGroupName, d.Status)
}

func NewDBInstanceDBSecurityGroup(Id int64, DbInstanceId int64, DBSecurityGroupName string, RefSecurityGroupId sql.NullInt64, Status string) *DBInstanceDBSecurityGroup {
	return &DBInstanceDBSecurityGroup{Id,
		DbInstanceId,
		DBSecurityGroupName,
		RefSecurityGroupId,
		Status}
}

type DBInstanceAssociatedRoles struct {
	Id           int64
	DbInstanceId int64
	RoleArn      string
	FeatureName  string
	Status       string
}

func (d *DBInstanceAssociatedRoles) Dump() string {
	return fmt.Sprintf("RoleArn(%s) FeatureName(%s) Status(%s)", d.RoleArn, d.FeatureName, d.Status)
}

func NewDBInstanceAssociatedRoles(Id int64, DbInstanceId int64, RoleArn string, FeatureName string, Status string) *DBInstanceAssociatedRoles {
	return &DBInstanceAssociatedRoles{Id,
		DbInstanceId, RoleArn, FeatureName, Status}
}

type DBInstanceProcessorFeatures struct {
	Id           int64
	DbInstanceId int64
	Name         string
	Value        string
}

func (d *DBInstanceProcessorFeatures) Dump() string {
	return fmt.Sprintf("Name(%s) Value(%s)", d.Name, d.Value)
}

func NewDBInstanceProcessorFeatures(Id int64, DbInstanceId int64, Name string, Value string) *DBInstanceProcessorFeatures {
	return &DBInstanceProcessorFeatures{Id,
		DbInstanceId, Name, Value}
}

type DBInstanceDomainMemberships struct {
	Id           int64
	DbInstanceId int64
	Domain       string
	Status       string
	FQDN         string
	IamRoleName  string
}

func (d *DBInstanceDomainMemberships) Dump() string {
	return fmt.Sprintf("Domain(%s) Status(%s) FQDN(%s) IamRoleName(%s)", d.Domain, d.Status, d.FQDN, d.IamRoleName)
}

func NewDBInstanceDomainMemberships(Id int64, DbInstanceId int64,
	Domain string, Status string, FQDN string, IamRoleName string) *DBInstanceDomainMemberships {
	return &DBInstanceDomainMemberships{Id,
		DbInstanceId, Domain, Status, FQDN, IamRoleName}
}

type DBInstanceStatusInfos struct {
	Id           int64
	DbInstanceId int64
	StatusType   string
	Normal       bool
	Status       string
	Message      string
}

func (d *DBInstanceStatusInfos) Dump() string {
	return fmt.Sprintf("StatusType(%s) Normal(%t) Status(%s) Message(%s)", d.StatusType, d.Normal, d.Status, d.Message)
}

func NewDBInstanceStatusInfos(Id int64, DbInstanceId int64,
	StatusType string, Normal bool, Status string, Message string) *DBInstanceStatusInfos {
	return &DBInstanceStatusInfos{Id, DbInstanceId,
		StatusType, Normal, Status, Message}
}

type DBInstanceDBSubnetGroup struct {
	Id                       int64
	DbInstanceId             int64
	DBSubnetGroupName        string
	DBSubnetGroupDescription string
	VpcId                    string
	RefVpcId                 sql.NullInt64
	SubnetGroupStatus        string
	Subnets                  []*DBInstanceDBSubnetGroupSubnet
	DBSubnetGroupArn         string
}

func (d *DBInstanceDBSubnetGroup) loadDependencies(db *sql.DB) {
	arrayOfDBInstanceDBSubnetGroupSubnet, err := loadAllDBInstanceDBSubnetGroupSubnetByOrderUnique(db)
	if err == nil {
		for _, x := range arrayOfDBInstanceDBSubnetGroupSubnet {
			if x.DbInstanceSubnetGroupId == d.Id {
				d.Subnets = append(d.Subnets, x)
			}
		}
	} else {
		fmt.Println(err)
	}

}

func (d *DBInstanceDBSubnetGroup) Dump() string {
	return fmt.Sprintf("DBSubnetGroupName(%s) DBSubnetGroupDescription(%s) VpcId(%s) SubnetGroupStatus(%s) DBSubnetGroupArn(%s)",
		d.DBSubnetGroupName, d.DBSubnetGroupDescription, d.VpcId, d.SubnetGroupStatus, d.DBSubnetGroupArn)
}

func NewDBInstanceDBSubnetGroup(Id int64, DbInstanceId int64,
	DBSubnetGroupName string, DBSubnetGroupDescription string, VpcId string, RefVpcId sql.NullInt64, SubnetGroupStatus string, DBSubnetGroupArn string) *DBInstanceDBSubnetGroup {
	return &DBInstanceDBSubnetGroup{Id: Id, DbInstanceId: DbInstanceId,
		DBSubnetGroupName: DBSubnetGroupName, DBSubnetGroupDescription: DBSubnetGroupDescription, VpcId: VpcId, RefVpcId: RefVpcId, SubnetGroupStatus: SubnetGroupStatus, DBSubnetGroupArn: DBSubnetGroupArn}
}

type DBInstanceDBSubnetGroupSubnet struct {
	Id                      int64
	DbInstanceSubnetGroupId int64
	SubnetIdentifier        string
	RefSubnetId             sql.NullInt64
	Name                    string
	SubnetStatus            string
}

func (d *DBInstanceDBSubnetGroupSubnet) Dump() string {
	return fmt.Sprintf("SubnetIdentifier(%s) Name(%s) SubnetStatus(%s)",
		d.SubnetIdentifier, d.Name, d.SubnetStatus)
}

func NewDBInstanceDBSubnetGroupSubnet(Id int64, DbInstanceSubnetGroupId int64,
	SubnetIdentifier string, RefSubnetId sql.NullInt64, Name string, SubnetStatus string) *DBInstanceDBSubnetGroupSubnet {
	return &DBInstanceDBSubnetGroupSubnet{Id, DbInstanceSubnetGroupId,
		SubnetIdentifier, RefSubnetId, Name, SubnetStatus}
}
