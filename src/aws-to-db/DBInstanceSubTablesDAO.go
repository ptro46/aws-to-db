package main

import (
	"database/sql" // package SQL
	"fmt"

	_ "github.com/lib/pq" // driver Postgres
)

//////////////////////////////////////////////////////////
//row result
//////////////////////////////////////////////////////////

func rowResultSetToDBInstanceGroupType(row *sql.Row) (*DBInstanceGroupType, error) {
	var err error
	var Id int64
	var DbInstanceId int64
	var Name string
	var Status string

	err = row.Scan(&Id, &DbInstanceId, &Name, &Status)
	if err != nil {
		return nil, err
	}
	return NewDBInstanceGroupType(Id, DbInstanceId, Name, Status), nil
}

func rowResultSetToDBInstanceDBSecurityGroup(row *sql.Row) (*DBInstanceDBSecurityGroup, error) {
	var err error
	var Id int64
	var DbInstanceId int64
	var DBSecurityGroupName string
	var RefSecurityGroupId sql.NullInt64
	var Status string

	err = row.Scan(&Id, &DbInstanceId, &DBSecurityGroupName, &RefSecurityGroupId, &Status)
	if err != nil {
		return nil, err
	}
	return NewDBInstanceDBSecurityGroup(Id, DbInstanceId, DBSecurityGroupName, RefSecurityGroupId, Status), nil
}

func rowResultSetToDBInstanceVpcSecurityGroups(row *sql.Row) (*DBInstanceVpcSecurityGroups, error) {
	var err error
	var Id int64
	var DbInstanceId int64
	var VpcSecurityGroupId string
	var RefSecurityGroupId sql.NullInt64
	var Status string

	err = row.Scan(&Id, &DbInstanceId, &VpcSecurityGroupId, &RefSecurityGroupId, &Status)
	if err != nil {
		return nil, err
	}
	return NewDBInstanceVpcSecurityGroups(Id, DbInstanceId, VpcSecurityGroupId, RefSecurityGroupId, Status), nil
}

func rowResultSetToDBInstanceAssociatedRoles(row *sql.Row) (*DBInstanceAssociatedRoles, error) {
	var err error
	var Id int64
	var DbInstanceId int64
	var RoleArn string
	var FeatureName string
	var Status string

	err = row.Scan(&Id, &DbInstanceId, &RoleArn, &FeatureName, &Status)
	if err != nil {
		return nil, err
	}
	return NewDBInstanceAssociatedRoles(Id, DbInstanceId, RoleArn, FeatureName, Status), nil
}

func rowResultSetToDBInstanceProcessorFeatures(row *sql.Row) (*DBInstanceProcessorFeatures, error) {
	var err error
	var Id int64
	var DbInstanceId int64
	var Name string
	var Value string

	err = row.Scan(&Id, &DbInstanceId, &Name, &Value)
	if err != nil {
		return nil, err
	}
	return NewDBInstanceProcessorFeatures(Id, DbInstanceId, Name, Value), nil
}

func rowResultSetToDBInstanceDomainMemberships(row *sql.Row) (*DBInstanceDomainMemberships, error) {
	var err error
	var Id int64
	var DbInstanceId int64
	var Domain string
	var Status string
	var FQDN string
	var IamRoleName string

	err = row.Scan(&Id, &DbInstanceId, &Domain, &Status, &FQDN, &IamRoleName)
	if err != nil {
		return nil, err
	}
	return NewDBInstanceDomainMemberships(Id, DbInstanceId, Domain, Status, FQDN, IamRoleName), nil
}

func rowResultSetToDBInstanceStatusInfos(row *sql.Row) (*DBInstanceStatusInfos, error) {
	var err error
	var Id int64
	var DbInstanceId int64
	var StatusType string
	var Normal bool
	var Status string
	var Message string

	err = row.Scan(&Id, &DbInstanceId, &StatusType, &Normal, &Status, &Message)
	if err != nil {
		return nil, err
	}
	return NewDBInstanceStatusInfos(Id, DbInstanceId, StatusType, Normal, Status, Message), nil
}

func rowResultSetToDBInstanceDBSubnetGroup(row *sql.Row) (*DBInstanceDBSubnetGroup, error) {
	var err error
	var Id int64
	var DbInstanceId int64
	var DBSubnetGroupName string
	var DBSubnetGroupDescription string
	var VpcId string
	var RefVpcId sql.NullInt64
	var SubnetGroupStatus string
	var DBSubnetGroupArn string

	err = row.Scan(&Id, &DbInstanceId, &DBSubnetGroupName, &DBSubnetGroupDescription, &VpcId, &RefVpcId, &SubnetGroupStatus, &DBSubnetGroupArn)
	if err != nil {
		return nil, err
	}
	return NewDBInstanceDBSubnetGroup(Id, DbInstanceId, DBSubnetGroupName, DBSubnetGroupDescription, VpcId, RefVpcId, SubnetGroupStatus, DBSubnetGroupArn), nil
}

func rowResultSetToDBInstanceDBSubnetGroupSubnet(row *sql.Row) (*DBInstanceDBSubnetGroupSubnet, error) {
	var err error
	var Id int64
	var DbInstanceSubnetGroupId int64
	var SubnetIdentifier string
	var RefSubnetId sql.NullInt64
	var Name string
	var SubnetStatus string

	err = row.Scan(&Id, &DbInstanceSubnetGroupId, &SubnetIdentifier, &RefSubnetId, &Name, &SubnetStatus)
	if err != nil {
		return nil, err
	}
	return NewDBInstanceDBSubnetGroupSubnet(Id, DbInstanceSubnetGroupId, SubnetIdentifier, RefSubnetId, Name, SubnetStatus), nil
}

//////////////////////////////////////////////////////////
//rows no fetch
//////////////////////////////////////////////////////////

func rowsNoFetchResultSetToDBInstanceGroupType(rows *sql.Rows) (*DBInstanceGroupType, error) {
	var err error
	var Id int64
	var DbInstanceId int64
	var Name string
	var Status string

	err = rows.Scan(&Id, &DbInstanceId, &Name, &Status)
	if err != nil {
		return nil, err
	}
	return NewDBInstanceGroupType(Id, DbInstanceId, Name, Status), nil
}

func rowsNoFetchResultSetToDBInstanceDBSecurityGroup(rows *sql.Rows) (*DBInstanceDBSecurityGroup, error) {
	var err error
	var Id int64
	var DbInstanceId int64
	var DBSecurityGroupName string
	var RefSecurityGroupId sql.NullInt64
	var Status string

	err = rows.Scan(&Id, &DbInstanceId, &DBSecurityGroupName, &RefSecurityGroupId, &Status)
	if err != nil {
		return nil, err
	}
	return NewDBInstanceDBSecurityGroup(Id, DbInstanceId, DBSecurityGroupName, RefSecurityGroupId, Status), nil
}

func rowsNoFetchResultSetToDBInstanceVpcSecurityGroups(rows *sql.Rows) (*DBInstanceVpcSecurityGroups, error) {
	var err error
	var Id int64
	var DbInstanceId int64
	var VpcSecurityGroupId string
	var RefSecurityGroupId sql.NullInt64
	var Status string

	err = rows.Scan(&Id, &DbInstanceId, &VpcSecurityGroupId, &RefSecurityGroupId, &Status)
	if err != nil {
		return nil, err
	}
	return NewDBInstanceVpcSecurityGroups(Id, DbInstanceId, VpcSecurityGroupId, RefSecurityGroupId, Status), nil
}

func rowsNoFetchResultSetToDBInstanceAssociatedRoles(rows *sql.Rows) (*DBInstanceAssociatedRoles, error) {
	var err error
	var Id int64
	var DbInstanceId int64
	var RoleArn string
	var FeatureName string
	var Status string

	err = rows.Scan(&Id, &DbInstanceId, &RoleArn, &FeatureName, &Status)
	if err != nil {
		return nil, err
	}
	return NewDBInstanceAssociatedRoles(Id, DbInstanceId, RoleArn, FeatureName, Status), nil
}
func rowsNoFetchResultSetToDBInstanceProcessorFeatures(rows *sql.Rows) (*DBInstanceProcessorFeatures, error) {
	var err error
	var Id int64
	var DbInstanceId int64
	var Name string
	var Value string

	err = rows.Scan(&Id, &DbInstanceId, &Name, &Value)
	if err != nil {
		return nil, err
	}
	return NewDBInstanceProcessorFeatures(Id, DbInstanceId, Name, Value), nil
}
func rowsNoFetchResultSetToDBInstanceDomainMemberships(rows *sql.Rows) (*DBInstanceDomainMemberships, error) {
	var err error
	var Id int64
	var DbInstanceId int64
	var Domain string
	var Status string
	var FQDN string
	var IamRoleName string

	err = rows.Scan(&Id, &DbInstanceId, &Domain, &Status, &FQDN, &IamRoleName)
	if err != nil {
		return nil, err
	}
	return NewDBInstanceDomainMemberships(Id, DbInstanceId, Domain, Status, FQDN, IamRoleName), nil
}
func rowsNoFetchResultSetToDBInstanceStatusInfos(rows *sql.Rows) (*DBInstanceStatusInfos, error) {
	var err error
	var Id int64
	var DbInstanceId int64
	var StatusType string
	var Normal bool
	var Status string
	var Message string

	err = rows.Scan(&Id, &DbInstanceId, &StatusType, &Normal, &Status, &Message)
	if err != nil {
		return nil, err
	}
	return NewDBInstanceStatusInfos(Id, DbInstanceId, StatusType, Normal, Status, Message), nil
}
func rowsNoFetchResultSetToDBInstanceDBSubnetGroup(rows *sql.Rows) (*DBInstanceDBSubnetGroup, error) {
	var err error
	var Id int64
	var DbInstanceId int64
	var DBSubnetGroupName string
	var DBSubnetGroupDescription string
	var VpcId string
	var RefVpcId sql.NullInt64
	var SubnetGroupStatus string
	var DBSubnetGroupArn string

	err = rows.Scan(&Id, &DbInstanceId, &DBSubnetGroupName, &DBSubnetGroupDescription, &VpcId, &RefVpcId, &SubnetGroupStatus, &DBSubnetGroupArn)
	if err != nil {
		return nil, err
	}
	return NewDBInstanceDBSubnetGroup(Id, DbInstanceId, DBSubnetGroupName, DBSubnetGroupDescription, VpcId, RefVpcId, SubnetGroupStatus, DBSubnetGroupArn), nil
}

func rowsNoFetchResultSetToDBInstanceDBSubnetGroupSubnet(rows *sql.Rows) (*DBInstanceDBSubnetGroupSubnet, error) {
	var err error
	var Id int64
	var DbInstanceSubnetGroupId int64
	var SubnetIdentifier string
	var RefSubnetId sql.NullInt64
	var Name string
	var SubnetStatus string

	err = rows.Scan(&Id, &DbInstanceSubnetGroupId, &SubnetIdentifier, &RefSubnetId, &Name, &SubnetStatus)
	if err != nil {
		return nil, err
	}
	return NewDBInstanceDBSubnetGroupSubnet(Id, DbInstanceSubnetGroupId, SubnetIdentifier, RefSubnetId, Name, SubnetStatus), nil
}

//////////////////////////////////////////////////////////
//rows result
//////////////////////////////////////////////////////////

func rowsResultSetToDBInstanceGroupType(rows *sql.Rows) (*DBInstanceGroupType, error) {
	var err error
	if rows.Next() {
		var Id int64
		var DbInstanceId int64
		var Name string
		var Status string

		err = rows.Scan(&Id, &DbInstanceId, &Name, &Status)
		if err != nil {
			return nil, err
		}
		return NewDBInstanceGroupType(Id, DbInstanceId, Name, Status), nil
	}

	return nil, err

}

func rowsResultSetToDBInstanceDBSecurityGroup(rows *sql.Rows) (*DBInstanceDBSecurityGroup, error) {
	var err error
	if rows.Next() {
		var Id int64
		var DbInstanceId int64
		var DBSecurityGroupName string
		var RefSecurityGroupId sql.NullInt64
		var Status string

		err = rows.Scan(&Id, &DbInstanceId, &DBSecurityGroupName, &RefSecurityGroupId, &Status)
		if err != nil {
			return nil, err
		}
		return NewDBInstanceDBSecurityGroup(Id, DbInstanceId, DBSecurityGroupName, RefSecurityGroupId, Status), nil
	}
	return nil, err
}

func rowsResultSetToDBInstanceVpcSecurityGroups(rows *sql.Rows) (*DBInstanceVpcSecurityGroups, error) {
	var err error
	if rows.Next() {
		var Id int64
		var DbInstanceId int64
		var VpcSecurityGroupId string
		var RefSecurityGroupId sql.NullInt64
		var Status string

		err = rows.Scan(&Id, &DbInstanceId, &VpcSecurityGroupId, &RefSecurityGroupId, &Status)
		if err != nil {
			return nil, err
		}
		return NewDBInstanceVpcSecurityGroups(Id, DbInstanceId, VpcSecurityGroupId, RefSecurityGroupId, Status), nil
	}
	return nil, err
}

func rowsResultSetToDBInstanceAssociatedRoles(rows *sql.Rows) (*DBInstanceAssociatedRoles, error) {
	var err error
	if rows.Next() {
		var Id int64
		var DbInstanceId int64
		var RoleArn string
		var FeatureName string
		var Status string

		err = rows.Scan(&Id, &DbInstanceId, &RoleArn, &FeatureName, &Status)
		if err != nil {
			return nil, err
		}
		return NewDBInstanceAssociatedRoles(Id, DbInstanceId, RoleArn, FeatureName, Status), nil
	}
	return nil, err

}
func rowsResultSetToDBInstanceProcessorFeatures(rows *sql.Rows) (*DBInstanceProcessorFeatures, error) {
	var err error
	if rows.Next() {
		var Id int64
		var DbInstanceId int64
		var Name string
		var Value string

		err = rows.Scan(&Id, &DbInstanceId, &Name, &Value)
		if err != nil {
			return nil, err
		}
		return NewDBInstanceProcessorFeatures(Id, DbInstanceId, Name, Value), nil
	}
	return nil, err

}
func rowsResultSetToDBInstanceDomainMemberships(rows *sql.Rows) (*DBInstanceDomainMemberships, error) {
	var err error
	if rows.Next() {
		var Id int64
		var DbInstanceId int64
		var Domain string
		var Status string
		var FQDN string
		var IamRoleName string

		err = rows.Scan(&Id, &DbInstanceId, &Domain, &Status, &FQDN, &IamRoleName)
		if err != nil {
			return nil, err
		}
		return NewDBInstanceDomainMemberships(Id, DbInstanceId, Domain, Status, FQDN, IamRoleName), nil
	}
	return nil, err

}
func rowsResultSetToDBInstanceStatusInfos(rows *sql.Rows) (*DBInstanceStatusInfos, error) {
	var err error
	if rows.Next() {
		var Id int64
		var DbInstanceId int64
		var StatusType string
		var Normal bool
		var Status string
		var Message string

		err = rows.Scan(&Id, &DbInstanceId, &StatusType, &Normal, &Status, &Message)
		if err != nil {
			return nil, err
		}
		return NewDBInstanceStatusInfos(Id, DbInstanceId, StatusType, Normal, Status, Message), nil
	}
	return nil, err

}
func rowsResultSetToDBInstanceDBSubnetGroup(rows *sql.Rows) (*DBInstanceDBSubnetGroup, error) {
	var err error
	if rows.Next() {
		var Id int64
		var DbInstanceId int64
		var DBSubnetGroupName string
		var DBSubnetGroupDescription string
		var VpcId string
		var RefVpcId sql.NullInt64
		var SubnetGroupStatus string
		var DBSubnetGroupArn string

		err = rows.Scan(&Id, &DbInstanceId, &DBSubnetGroupName, &DBSubnetGroupDescription, &VpcId, &RefVpcId, &SubnetGroupStatus, &DBSubnetGroupArn)
		if err != nil {
			return nil, err
		}
		return NewDBInstanceDBSubnetGroup(Id, DbInstanceId, DBSubnetGroupName, DBSubnetGroupDescription, VpcId, RefVpcId, SubnetGroupStatus, DBSubnetGroupArn), nil
	}
	return nil, err

}
func rowsResultSetToDBInstanceDBSubnetGroupSubnet(rows *sql.Rows) (*DBInstanceDBSubnetGroupSubnet, error) {
	var err error
	if rows.Next() {
		var Id int64
		var DbInstanceSubnetGroupId int64
		var SubnetIdentifier string
		var RefSubnetId sql.NullInt64
		var Name string
		var SubnetStatus string

		err = rows.Scan(&Id, &DbInstanceSubnetGroupId, &SubnetIdentifier, &RefSubnetId, &Name, &SubnetStatus)
		if err != nil {
			return nil, err
		}
		return NewDBInstanceDBSubnetGroupSubnet(Id, DbInstanceSubnetGroupId, SubnetIdentifier, RefSubnetId, Name, SubnetStatus), nil
	}
	return nil, err

}

//////////////////////////////////////////////////////////
//create
//////////////////////////////////////////////////////////

func createDBInstanceOptionGroupMemberships(db *sql.DB, DbInstanceId int64, OptionGroupName string, Status string) *DBInstanceGroupType {
	rows := db.QueryRow("insert into db_instance_option_group_memberships(db_instance_id,option_group_name,status) values($1,$2,$3) returning id,db_instance_id,option_group_name,status", DbInstanceId, OptionGroupName, Status)

	dbInstanceOptionGroupMembership, err := rowResultSetToDBInstanceGroupType(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return dbInstanceOptionGroupMembership
}

func createDBInstanceDBSecurityGroup(db *sql.DB, DbInstanceId int64, DBSecurityGroupName string, Status string) *DBInstanceDBSecurityGroup {
	rows := db.QueryRow("insert into db_instance_db_security_groups(db_instance_id,db_security_group_name,status) values($1,$2,$3) returning id,db_instance_id,db_security_group_name,ref_security_group_id,status", DbInstanceId, DBSecurityGroupName, Status)

	dbInstanceSecGroup, err := rowResultSetToDBInstanceDBSecurityGroup(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return dbInstanceSecGroup
}

func createDBInstanceDBParameterGroups(db *sql.DB, DbInstanceId int64, DBParameterGroupName string, ParameterApplyStatus string) *DBInstanceGroupType {
	rows := db.QueryRow("insert into db_instance_db_parameter_groups(db_instance_id,db_parameter_group_name,parameter_apply_status) values($1,$2,$3) returning id,db_instance_id,db_parameter_group_name,parameter_apply_status", DbInstanceId, DBParameterGroupName, ParameterApplyStatus)

	dbInstanceDBParameterGroups, err := rowResultSetToDBInstanceGroupType(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return dbInstanceDBParameterGroups
}

func createDBInstanceVpcSecurityGroups(db *sql.DB, DbInstanceId int64, VpcSecurityGroupId string, Status string) *DBInstanceVpcSecurityGroups {
	rows := db.QueryRow("insert into db_instance_vpc_security_groups(db_instance_id,vpc_security_group_id,status) values($1,$2,$3) returning id,db_instance_id,vpc_security_group_id,ref_vpc_security_group_id,status", DbInstanceId, VpcSecurityGroupId, Status)

	dbInstanceVpcSecurityGroups, err := rowResultSetToDBInstanceVpcSecurityGroups(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return dbInstanceVpcSecurityGroups
}

func createDBInstanceDBSecurityGroups(db *sql.DB, DbInstanceId int64, DBSecurityGroupName string, Status string) *DBInstanceGroupType {
	rows := db.QueryRow("insert into db_instance_db_security_groups(db_instance_id,db_security_group_name,satus) values($1,$2,$3) returning id,db_instance_id,db_security_group_name,ref_security_group_id,satus", DbInstanceId, DBSecurityGroupName, Status)

	dbInstanceDBSecurityGroups, err := rowResultSetToDBInstanceGroupType(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return dbInstanceDBSecurityGroups
}

func createDBInstanceAssociatedRoles(db *sql.DB, DbInstanceId int64, RoleArn string, FeatureName string, Status string) *DBInstanceAssociatedRoles {
	rows := db.QueryRow("insert into db_instance_associated_roles(db_instance_id,role_arn,feature_name,status) values($1,$2,$3,$4) returning id,db_instance_id,role_arn,feature_name,status", DbInstanceId, RoleArn, FeatureName, Status)

	dbInstanceAssociatedRoles, err := rowResultSetToDBInstanceAssociatedRoles(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return dbInstanceAssociatedRoles
}

func createDBInstanceProcessorFeatures(db *sql.DB, DbInstanceId int64, Name string, Value string) *DBInstanceProcessorFeatures {
	rows := db.QueryRow("insert into db_instance_processor_features(db_instance_id,name,value) values($1,$2,$3) returning id,db_instance_id,name,value", DbInstanceId, Name, Value)

	dbInstanceProcessorFeatures, err := rowResultSetToDBInstanceProcessorFeatures(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return dbInstanceProcessorFeatures
}

func createDBInstanceDomainMemberships(db *sql.DB, DbInstanceId int64, Domain string, Status string, FQDN string, IamRoleName string) *DBInstanceDomainMemberships {
	rows := db.QueryRow("insert into db_instance_domain_memberships(db_instance_id,domain,status,fqdn,iam_role_name) values($1,$2,$3,$4,$5) returning id,db_instance_id,domain,status,fqdn,iam_role_name", DbInstanceId, Domain, Status, FQDN, IamRoleName)

	dbInstanceDomainMemberships, err := rowResultSetToDBInstanceDomainMemberships(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return dbInstanceDomainMemberships
}

func createDBInstanceStatusInfos(db *sql.DB, DbInstanceId int64, StatusType string, Normal bool, Status string, Message string) *DBInstanceStatusInfos {
	rows := db.QueryRow("insert into db_instance_status_infos(db_instance_id,status_type,normal,status,message) values($1,$2,$3,$4,$5) returning id,db_instance_id,status_type,normal,status,message", DbInstanceId, StatusType, Normal, Status, Message)

	dbInstanceStatusInfos, err := rowResultSetToDBInstanceStatusInfos(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return dbInstanceStatusInfos
}

func createDBInstanceDBSubnetGroup(db *sql.DB, DbInstanceId int64, DBSubnetGroupName string, DBSubnetGroupDescription string, VpcId string, SubnetGroupStatus string, DBSubnetGroupArn string) *DBInstanceDBSubnetGroup {
	rows := db.QueryRow("insert into db_instance_db_subnet_group(db_instance_id,db_subnet_group_name,db_subnet_group_description,vpc_id,subnet_group_status,db_subnet_group_arn) values($1,$2,$3,$4,$5,$6) returning id,db_instance_id,db_subnet_group_name,db_subnet_group_description,vpc_id,ref_vpc_id,subnet_group_status,db_subnet_group_arn", DbInstanceId, DBSubnetGroupName, DBSubnetGroupDescription, VpcId, SubnetGroupStatus, DBSubnetGroupArn)

	dbInstanceDBSubnetGroup, err := rowResultSetToDBInstanceDBSubnetGroup(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return dbInstanceDBSubnetGroup
}

func createDBInstanceDBSubnetGroupSubnet(db *sql.DB, DbInstanceSubnetGroupId int64, SubnetIdentifier string, Name string, SubnetStatus string) *DBInstanceDBSubnetGroupSubnet {
	rows := db.QueryRow("insert into db_instance_db_subnet_group_subnet(db_instance_db_subnet_group_id,subnet_identifier,name,subnet_status) values($1,$2,$3,$4) returning id,db_instance_db_subnet_group_id,subnet_identifier,ref_subnet_id,name,subnet_status", DbInstanceSubnetGroupId, SubnetIdentifier, Name, SubnetStatus)

	dbInstanceDBSubnetGroupSubnet, err := rowResultSetToDBInstanceDBSubnetGroupSubnet(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return dbInstanceDBSubnetGroupSubnet
}

//////////////////////////////////////////////////////////
//load all unique
//////////////////////////////////////////////////////////

func loadAllDBInstanceAssociatedRolesByOrderUnique(db *sql.DB) ([]*DBInstanceAssociatedRoles, error) {
	results := make([]*DBInstanceAssociatedRoles, 0, 0)
	sqlSelect := "SELECT id,db_instance_id,role_arn,feature_name,status from db_instance_associated_roles order by feature_name"

	rows, err := db.Query(sqlSelect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		dBInstanceAssociatedRoles, err := rowsNoFetchResultSetToDBInstanceAssociatedRoles(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, dBInstanceAssociatedRoles)
	}
	return results, nil
}
func loadAllDBInstanceProcessorFeaturesByOrderUnique(db *sql.DB) ([]*DBInstanceProcessorFeatures, error) {
	results := make([]*DBInstanceProcessorFeatures, 0, 0)
	sqlSelect := "SELECT id,db_instance_id,name,value from db_instance_processor_features order by value"

	rows, err := db.Query(sqlSelect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		dBInstanceProcessorFeatures, err := rowsNoFetchResultSetToDBInstanceProcessorFeatures(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, dBInstanceProcessorFeatures)
	}
	return results, nil
}
func loadAllDBInstanceDomainMembershipsByOrderUnique(db *sql.DB) ([]*DBInstanceDomainMemberships, error) {
	results := make([]*DBInstanceDomainMemberships, 0, 0)
	sqlSelect := "SELECT id,db_instance_id,domain,status,fqdn,iam_role_name from db_instance_domain_memberships order by iam_role_name"

	rows, err := db.Query(sqlSelect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		dBInstanceDomainMemberships, err := rowsNoFetchResultSetToDBInstanceDomainMemberships(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, dBInstanceDomainMemberships)
	}
	return results, nil
}
func loadAllDBInstanceStatusInfosByOrderUnique(db *sql.DB) ([]*DBInstanceStatusInfos, error) {
	results := make([]*DBInstanceStatusInfos, 0, 0)
	sqlSelect := "SELECT id,db_instance_id,status_type,normal,status,message from db_instance_status_infos order by status_type"

	rows, err := db.Query(sqlSelect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		dBInstanceStatusInfos, err := rowsNoFetchResultSetToDBInstanceStatusInfos(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, dBInstanceStatusInfos)
	}
	return results, nil
}

func loadAllDBInstanceDBOptionGroupMembershipsByOrderUnique(db *sql.DB) ([]*DBInstanceGroupType, error) {
	results := make([]*DBInstanceGroupType, 0, 0)
	sqlSelect := "SELECT id,db_instance_id,option_group_name,status from db_instance_option_group_memberships order by option_group_name"

	rows, err := db.Query(sqlSelect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		dBInstanceOptionGroupMemberships, err := rowsNoFetchResultSetToDBInstanceGroupType(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, dBInstanceOptionGroupMemberships)
	}
	return results, nil
}
func loadAllDBInstanceDBParameterGroupsByOrderUnique(db *sql.DB) ([]*DBInstanceGroupType, error) {
	results := make([]*DBInstanceGroupType, 0, 0)
	sqlSelect := "SELECT id,db_instance_id,db_parameter_group_name,parameter_apply_status from db_instance_db_parameter_groups order by db_parameter_group_name"

	rows, err := db.Query(sqlSelect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		dBInstanceDBParameterGroups, err := rowsNoFetchResultSetToDBInstanceGroupType(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, dBInstanceDBParameterGroups)
	}
	return results, nil
}

func loadAllDBInstanceDBVpcSecurityGroupsByOrderUnique(db *sql.DB) ([]*DBInstanceGroupType, error) {
	results := make([]*DBInstanceGroupType, 0, 0)
	sqlSelect := "SELECT id,db_instance_id,vpc_security_group_id,ref_vpc_security_group_id,status from db_instance_vpc_security_groups order by status"

	rows, err := db.Query(sqlSelect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		dBInstanceVpcSecurityGroups, err := rowsNoFetchResultSetToDBInstanceGroupType(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, dBInstanceVpcSecurityGroups)
	}
	return results, nil
}
func loadAllDBInstanceDBSubnetGroupByOrderUnique(db *sql.DB) ([]*DBInstanceDBSubnetGroup, error) {
	results := make([]*DBInstanceDBSubnetGroup, 0, 0)
	sqlSelect := "SELECT id,db_instance_id,db_subnet_group_name,db_subnet_group_description,vpc_id,ref_vpc_id,subnet_group_status,db_subnet_group_arn from db_instance_db_subnet_group order by db_subnet_group_arn"

	rows, err := db.Query(sqlSelect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		dbInstanceDBSubnetGroup, err := rowsNoFetchResultSetToDBInstanceDBSubnetGroup(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, dbInstanceDBSubnetGroup)
	}
	return results, nil
}

func loadAllDBInstanceDBSubnetGroupSubnetByOrderUnique(db *sql.DB) ([]*DBInstanceDBSubnetGroupSubnet, error) {
	results := make([]*DBInstanceDBSubnetGroupSubnet, 0, 0)
	sqlSelect := "SELECT id,db_instance_db_subnet_group_id,subnet_identifier,ref_subnet_id,name,subnet_status from db_instance_db_subnet_group_subnet order by subnet_identifier"

	rows, err := db.Query(sqlSelect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		dbInstanceDBSubnetGroupSubnet, err := rowsNoFetchResultSetToDBInstanceDBSubnetGroupSubnet(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, dbInstanceDBSubnetGroupSubnet)
	}
	return results, nil
}

func loadAllDBInstanceDBSecurityGroupByOrderUnique(db *sql.DB) ([]*DBInstanceDBSecurityGroup, error) {
	results := make([]*DBInstanceDBSecurityGroup, 0, 0)
	sqlSelect := "SELECT id,db_instance_id,db_security_group_name,ref_security_group_id,status from db_instance_db_security_groups order by db_security_group_name"

	rows, err := db.Query(sqlSelect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		dBInstanceDBSecurityGroup, err := rowsNoFetchResultSetToDBInstanceDBSecurityGroup(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, dBInstanceDBSecurityGroup)
	}
	return results, nil
}

//////////////////////////////////////////////////////////
//load all with Update
//////////////////////////////////////////////////////////

func loadAllDBInstanceDBSubnetGroup(db *sql.DB) []*DBInstanceDBSubnetGroup {
	results := make([]*DBInstanceDBSubnetGroup, 0, 0)
	sqlSelect := "SELECT id,db_instance_id,db_subnet_group_name,db_subnet_group_description,vpc_id,ref_vpc_id,subnet_group_status,db_subnet_group_arn from db_instance_db_subnet_group"

	rows, err := db.Query(sqlSelect)
	if err != nil {
		return results
	}
	defer rows.Close()

	for rows.Next() {
		dbInstanceDBSubnetGroup, err := rowsNoFetchResultSetToDBInstanceDBSubnetGroup(rows)
		if err != nil {
			return make([]*DBInstanceDBSubnetGroup, 0, 0)
		}
		results = append(results, dbInstanceDBSubnetGroup)
	}
	return results
}

func loadAllDBInstanceDBSubnetGroupSubnet(db *sql.DB) []*DBInstanceDBSubnetGroupSubnet {
	results := make([]*DBInstanceDBSubnetGroupSubnet, 0, 0)
	sqlSelect := "SELECT id,db_instance_db_subnet_group_id,subnet_identifier,ref_subnet_id,name,subnet_status from db_instance_db_subnet_group_subnet"

	rows, err := db.Query(sqlSelect)
	if err != nil {
		return results
	}
	defer rows.Close()

	for rows.Next() {
		dbInstanceDBSubnetGroupSubnet, err := rowsNoFetchResultSetToDBInstanceDBSubnetGroupSubnet(rows)
		if err != nil {
			return make([]*DBInstanceDBSubnetGroupSubnet, 0, 0)
		}
		results = append(results, dbInstanceDBSubnetGroupSubnet)
	}
	return results
}

func loadAllDBInstanceDBSecurityGroup(db *sql.DB) []*DBInstanceDBSecurityGroup {
	results := make([]*DBInstanceDBSecurityGroup, 0, 0)
	sqlSelect := "SELECT id,db_instance_id,db_security_group_name,ref_security_group_id,status from db_instance_db_security_groups"

	rows, err := db.Query(sqlSelect)
	if err != nil {
		return results
	}
	defer rows.Close()

	for rows.Next() {
		dBInstanceDBSecurityGroup, err := rowsNoFetchResultSetToDBInstanceDBSecurityGroup(rows)
		if err != nil {
			return make([]*DBInstanceDBSecurityGroup, 0, 0)
		}
		results = append(results, dBInstanceDBSecurityGroup)
	}
	return results
}

func loadAllDBInstanceVpcSecurityGroup(db *sql.DB) ([]*DBInstanceVpcSecurityGroups, error) {
	results := make([]*DBInstanceVpcSecurityGroups, 0, 0)
	sqlSelect := "SELECT id,db_instance_id,vpc_security_group_id,ref_vpc_security_group_id,status from db_instance_vpc_security_groups"

	rows, err := db.Query(sqlSelect)
	if err != nil {

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		dBInstanceVpcSecurityGroups, err := rowsNoFetchResultSetToDBInstanceVpcSecurityGroups(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, dBInstanceVpcSecurityGroups)
	}
	return results, nil
}

func updateDBInstanceVpcSecurityGroup(db *sql.DB, dBInstanceVpcSecurityGroup *DBInstanceVpcSecurityGroups, secGroupID int64) *DBInstanceVpcSecurityGroups {
	rows := db.QueryRow("update db_instance_vpc_security_groups set ref_vpc_security_group_id=$1 where id=$2 returning id,db_instance_id,vpc_security_group_id,ref_vpc_security_group_id,status", secGroupID, dBInstanceVpcSecurityGroup.Id)

	dBInstanceVpcSecurityGroup, err := rowResultSetToDBInstanceVpcSecurityGroups(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return dBInstanceVpcSecurityGroup
}

func updateDBInstanceDBSubnetGroupVpcId(db *sql.DB, dbInstanceDBSubnetGroup *DBInstanceDBSubnetGroup, vpcId int64) *DBInstanceDBSubnetGroup {
	rows := db.QueryRow("update db_instance_db_subnet_group set ref_vpc_id=$1 where id=$2 returning id,db_instance_id,db_subnet_group_name,db_subnet_group_description,vpc_id,ref_vpc_id,subnet_group_status,db_subnet_group_arn", vpcId, dbInstanceDBSubnetGroup.Id)

	dbInstanceDBSubnetGroup, err := rowResultSetToDBInstanceDBSubnetGroup(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return dbInstanceDBSubnetGroup
}

func updateDBInstanceDBSubnetGroupSubnetSubnetId(db *sql.DB, dbInstanceDBSubnetGroupSubnet *DBInstanceDBSubnetGroupSubnet, subnetId int64) *DBInstanceDBSubnetGroupSubnet {
	rows := db.QueryRow("update db_instance_db_subnet_group_subnet set ref_subnet_id=$1 where id=$2 returning id,db_instance_db_subnet_group_id,subnet_identifier,ref_subnet_id,name,subnet_status", subnetId, dbInstanceDBSubnetGroupSubnet.Id)

	dbInstanceDBSubnetGroupSubnet, err := rowResultSetToDBInstanceDBSubnetGroupSubnet(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return dbInstanceDBSubnetGroupSubnet
}

func updateDBInstanceDBSecurityGroupSecGroup(db *sql.DB, dBInstanceDBSecurityGroup *DBInstanceDBSecurityGroup, secGroupId int64) *DBInstanceDBSecurityGroup {
	rows := db.QueryRow("update db_instance_db_security_groups set ref_security_group_id=$1 where id=$2 returning id,db_instance_id,db_security_group_name,ref_security_group_id,status", secGroupId, dBInstanceDBSecurityGroup.Id)

	dBInstanceDBSecurityGroup, err := rowResultSetToDBInstanceDBSecurityGroup(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return dBInstanceDBSecurityGroup
}

//////////////////////////////////////////////////////////
//delete all
//////////////////////////////////////////////////////////

func deleteAllDBInstanceOptionGroupMemberships(db *sql.DB) error {
	_, err := db.Exec("delete from db_instance_option_group_memberships")

	return err
}

func deleteAllDBInstanceDBParameterGroups(db *sql.DB) error {
	_, err := db.Exec("delete from db_instance_db_parameter_groups")

	return err
}

func deleteAllDBInstanceVpcSecurityGroups(db *sql.DB) error {
	_, err := db.Exec("delete from db_instance_vpc_security_groups")

	return err
}

func deleteAllDBInstanceDBSecurityGroups(db *sql.DB) error {
	_, err := db.Exec("delete from db_instance_db_security_groups")

	return err
}

func deleteAllDBInstanceAssociatedRoles(db *sql.DB) error {
	_, err := db.Exec("delete from db_instance_associated_roles")

	return err
}

func deleteAllDBInstanceProcessorFeatures(db *sql.DB) error {
	_, err := db.Exec("delete from db_instance_processor_features")

	return err
}

func deleteAllDBInstanceDBSubnetGroupSubnet(db *sql.DB) error {
	_, err := db.Exec("delete from db_instance_db_subnet_group_subnet")

	return err
}

func deleteAllDBInstanceDomainMemberships(db *sql.DB) error {
	_, err := db.Exec("delete from db_instance_domain_memberships")

	return err
}

func deleteAllDBInstanceStatusInfos(db *sql.DB) error {
	_, err := db.Exec("delete from db_instance_status_infos")

	return err
}

func deleteAllDBInstanceDBSubnetGroup(db *sql.DB) error {
	_, err := db.Exec("delete from db_instance_db_subnet_group")

	return err
}
