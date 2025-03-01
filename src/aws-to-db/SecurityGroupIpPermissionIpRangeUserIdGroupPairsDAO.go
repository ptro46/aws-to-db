package main

import (
	"database/sql" // package SQL
	"fmt"
	"log"

	_ "github.com/lib/pq" // driver Postgres
)

func rowResultSetToSecurityGroupIpPermissionUserIdGroupPair(row *sql.Row) (*SecurityGroupIpPermissionsUserIdGroupPair, error) {
	var Id int64
	var SecurityGroupIpPermissionsId int64
	var Description string
	var GroupId string
	var GroupName string
	var PeeringStatus string
	var UserId string
	var VpcId string
	var RefVpcId int64

	var VpcPeeringVonnectionId string

	err := row.Scan(&Id, &SecurityGroupIpPermissionsId, &Description, &GroupId, &GroupName, &PeeringStatus, &UserId, &VpcId, &RefVpcId, &VpcPeeringVonnectionId)
	if err != nil {
		return nil, err
	}
	return NewSecurityGroupIpPermissionsUserIdGroupPair(Id, SecurityGroupIpPermissionsId, Description,
		GroupId, GroupName, PeeringStatus, UserId, VpcId, RefVpcId, VpcPeeringVonnectionId), nil
}

func rowsNoFetchResultSetToSecurityGroupIpPermissionUserIdGroupPair(rows *sql.Rows) (*SecurityGroupIpPermissionsUserIdGroupPair, error) {
	var Id int64
	var SecurityGroupIpPermissionsId int64
	var Description string
	var GroupId string
	var GroupName string
	var PeeringStatus string
	var UserId string
	var VpcId string
	var RefVpcId int64
	var VpcPeeringVonnectionId string

	err := rows.Scan(&Id, &SecurityGroupIpPermissionsId, &Description, &GroupId, &GroupName, &PeeringStatus, &UserId, &VpcId, &RefVpcId, &VpcPeeringVonnectionId)
	if err != nil {
		return nil, err
	}
	return NewSecurityGroupIpPermissionsUserIdGroupPair(Id, SecurityGroupIpPermissionsId, Description,
		GroupId, GroupName, PeeringStatus, UserId, VpcId, RefVpcId, VpcPeeringVonnectionId), nil
}

func rowsResultSetToSecurityGroupIpPermissionUserIdGroupPair(rows *sql.Rows) (*SecurityGroupIpPermissionsUserIdGroupPair, error) {
	var err error
	if rows.Next() {
		var Id int64
		var SecurityGroupIpPermissionsId int64
		var Description string
		var GroupId string
		var GroupName string
		var PeeringStatus string
		var UserId string
		var VpcId string
		var RefVpcId int64
		var VpcPeeringVonnectionId string

		err := rows.Scan(&Id, &SecurityGroupIpPermissionsId, &Description, &GroupId, &GroupName, &PeeringStatus, &UserId, &VpcId, &RefVpcId, &VpcPeeringVonnectionId)
		if err != nil {
			return nil, err
		}
		return NewSecurityGroupIpPermissionsUserIdGroupPair(Id, SecurityGroupIpPermissionsId, Description,
			GroupId, GroupName, PeeringStatus, UserId, VpcId, RefVpcId, VpcPeeringVonnectionId), nil
	}
	return nil, err
}

func createSecurityGroupIpPermissionUserIdGroupPair(db *sql.DB, SecurityGroupIpPermissionsId int64, Description string, GroupId string, GroupName string, PeeringStatus string, UserId string, VpcId string, VpcPeeringVonnectionId string) *SecurityGroupIpPermissionsUserIdGroupPair {
	rows := db.QueryRow("insert into security_group_ip_permissions_user_id_group_pairs(security_group_ip_permissions_id,description,group_id,group_name,peering_status,user_id,vpc_id,ref_vpc_id,vpc_peering_connection_id) values($1,$2,$3,$4,$5,$6,$7,$8) returning id,security_group_ip_permissions_id,description,group_id,group_name,peering_status,user_id,vpc_id,ref_vpc_id,vpc_peering_connection_id",
		SecurityGroupIpPermissionsId, Description, GroupId, GroupName, PeeringStatus, UserId, VpcId, VpcPeeringVonnectionId)

	secGroupIpPermissionUserIdGroupPair, err := rowResultSetToSecurityGroupIpPermissionUserIdGroupPair(rows)
	if err != nil {
		return nil
	}
	return secGroupIpPermissionUserIdGroupPair
}

func createSecurityGroupIpPermissionEgressUserIdGroupPair(db *sql.DB, SecurityGroupIpPermissionsId int64, Description string, GroupId string, GroupName string, PeeringStatus string, UserId string, VpcId string, VpcPeeringVonnectionId string) *SecurityGroupIpPermissionsUserIdGroupPair {
	rows := db.QueryRow("insert into security_group_ip_permissions_egress_user_id_group_pairs(security_group_ip_permissions_egress_id,description,group_id,group_name,peering_status,user_id,vpc_id,vpc_peering_connection_id) values($1,$2,$3,$4,$5,$6,$7,$8) returning id,security_group_ip_permissions_egress_id,description,group_id,group_name,peering_status,user_id,vpc_id,ref_vpc_id,vpc_peering_connection_id",
		SecurityGroupIpPermissionsId, Description, GroupId, GroupName, PeeringStatus, UserId, VpcId, VpcPeeringVonnectionId)

	secGroupIpPermissionEgressUserIdGroupPair, err := rowResultSetToSecurityGroupIpPermissionUserIdGroupPair(rows)
	if err != nil {
		return nil
	}
	return secGroupIpPermissionEgressUserIdGroupPair
}

func loadAllSecurityGroupIpPermissionUserIdGroupPair(db *sql.DB) ([]*SecurityGroupIpPermissionsUserIdGroupPair, error) {
	rows, err := db.Query("select id,security_group_ip_permissions_id,description,group_id,group_name,peering_status,user_id,vpc_id,ref_vpc_id,vpc_peering_connection_id from security_group_ip_permissions_user_id_group_pairs order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfSGIpPermissionUserIdGroupPair := make([]*SecurityGroupIpPermissionsUserIdGroupPair, 0, 0)
	for rows.Next() {
		secGroupIpPermissionUserIdGroupPair, err := rowsNoFetchResultSetToSecurityGroupIpPermissionUserIdGroupPair(rows)
		if err == nil {
			arrayOfSGIpPermissionUserIdGroupPair = append(arrayOfSGIpPermissionUserIdGroupPair, secGroupIpPermissionUserIdGroupPair)
		} else {
			log.Println(err)
		}
	}

	return arrayOfSGIpPermissionUserIdGroupPair, nil
}

func loadAllSecurityGroupIpPermissionEgressUserIdGroupPair(db *sql.DB) ([]*SecurityGroupIpPermissionsUserIdGroupPair, error) {
	rows, err := db.Query("select id,security_group_ip_permissions_egress_id,description,group_id,group_name,peering_status,user_id,vpc_id,ref_vpc_id,vpc_peering_connection_id from security_group_ip_permissions_egress_user_id_group_pairs order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfSGIpPermissionEgUserIdGroupPair := make([]*SecurityGroupIpPermissionsUserIdGroupPair, 0, 0)
	for rows.Next() {
		secGroupIpPermissionEgUserIdGroupPair, err := rowsNoFetchResultSetToSecurityGroupIpPermissionUserIdGroupPair(rows)
		if err == nil {
			arrayOfSGIpPermissionEgUserIdGroupPair = append(arrayOfSGIpPermissionEgUserIdGroupPair, secGroupIpPermissionEgUserIdGroupPair)
		} else {
			log.Println(err)
		}
	}

	return arrayOfSGIpPermissionEgUserIdGroupPair, nil
}

func loadAllSecurityGroupIpPermissionUserIdGroupPairByOrderUnique(db *sql.DB) ([]*SecurityGroupIpPermissionsUserIdGroupPair, error) {
	rows, err := db.Query("select id,security_group_ip_permissions_id,description,group_id,group_name,peering_status,user_id,vpc_id,ref_vpc_id,vpc_peering_connection_id from security_group_ip_permissions_user_id_group_pairs order by group_name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfSGIpPermissionUserIdGroupPair := make([]*SecurityGroupIpPermissionsUserIdGroupPair, 0, 0)
	for rows.Next() {
		secGroupIpPermissionUserIdGroupPair, err := rowsNoFetchResultSetToSecurityGroupIpPermissionUserIdGroupPair(rows)
		if err == nil {
			arrayOfSGIpPermissionUserIdGroupPair = append(arrayOfSGIpPermissionUserIdGroupPair, secGroupIpPermissionUserIdGroupPair)
		} else {
			log.Println(err)
		}
	}

	return arrayOfSGIpPermissionUserIdGroupPair, nil
}

func loadAllSecurityGroupIpPermissionEgressUserIdGroupPairByOrderUnique(db *sql.DB) ([]*SecurityGroupIpPermissionsUserIdGroupPair, error) {
	rows, err := db.Query("select id,security_group_ip_permissions_egress_id,description,group_id,group_name,peering_status,user_id,vpc_id,ref_vpc_id,vpc_peering_connection_id from security_group_ip_permissions_egress_user_id_group_pairs order by group_name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfSGIpPermissionEgUserIdGroupPair := make([]*SecurityGroupIpPermissionsUserIdGroupPair, 0, 0)
	for rows.Next() {
		secGroupIpPermissionEgUserIdGroupPair, err := rowsNoFetchResultSetToSecurityGroupIpPermissionUserIdGroupPair(rows)
		if err == nil {
			arrayOfSGIpPermissionEgUserIdGroupPair = append(arrayOfSGIpPermissionEgUserIdGroupPair, secGroupIpPermissionEgUserIdGroupPair)
		} else {
			log.Println(err)
		}
	}

	return arrayOfSGIpPermissionEgUserIdGroupPair, nil
}

func updateSecurityGroupIpPermissionUserIdGroupPair(db *sql.DB, securityGroupIpPermissionUserIdGroupPair *SecurityGroupIpPermissionsUserIdGroupPair, vpcId int64) *SecurityGroupIpPermissionsUserIdGroupPair {
	rows := db.QueryRow("update security_group_ip_permissions_id set ref_vpc_id=$1 where id=$2 returning id,security_group_ip_permissions_egress_id,description,group_id,group_name,peering_status,user_id,vpc_id,ref_vpc_id,vpc_peering_connection_id", vpcId, securityGroupIpPermissionUserIdGroupPair.Id)

	userIdGroupPair, err := rowResultSetToSecurityGroupIpPermissionUserIdGroupPair(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return userIdGroupPair
}

func updateSecurityGroupIpPermissionEgressUserIdGroupPair(db *sql.DB, securityGroupIpPermissionEgressUserIdGroupPair *SecurityGroupIpPermissionsUserIdGroupPair, vpcId int64) *SecurityGroupIpPermissionsUserIdGroupPair {
	rows := db.QueryRow("update security_group_ip_permissions_egress_id set ref_vpc_id=$1 where id=$2 returning id,security_group_ip_permissions_egress_id,description,group_id,group_name,peering_status,user_id,vpc_id,ref_vpc_id,vpc_peering_connection_id", vpcId, securityGroupIpPermissionEgressUserIdGroupPair.Id)

	egressUserIdGroupPair, err := rowResultSetToSecurityGroupIpPermissionUserIdGroupPair(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return egressUserIdGroupPair
}

func deleteAllSecurityGroupIpPermissionUserIdGroupPair(db *sql.DB) error {
	_, err := db.Exec("delete from security_group_ip_permissions_user_id_group_pairs")

	return err
}
func deleteAllSecurityGroupIpPermissionEgressUserIdGroupPair(db *sql.DB) error {
	_, err := db.Exec("delete from security_group_ip_permissions_egress_user_id_group_pairs")

	return err
}
