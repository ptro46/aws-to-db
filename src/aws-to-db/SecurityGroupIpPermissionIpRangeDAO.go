package main

import (
	"database/sql" // package SQL
	"fmt"
	"log"

	_ "github.com/lib/pq" // driver Postgres
)

func rowResultSetToSecurityGroupIpPermissionIpRange(row *sql.Row) (*SecurityGroupIpPermissionsIpRange, error) {
	var Id int64
	var SecurityGroupIpPermissionsId int64
	var CidrIp string
	var Description string

	err := row.Scan(&Id, &SecurityGroupIpPermissionsId, &CidrIp, &Description)
	if err != nil {
		return nil, err
	}
	return NewSecurityGroupIpPermissionsIpRange(Id, SecurityGroupIpPermissionsId, CidrIp, Description), nil
}

func rowsNoFetchResultSetToSecurityGroupIpPermissionIpRange(rows *sql.Rows) (*SecurityGroupIpPermissionsIpRange, error) {
	var Id int64
	var SecurityGroupIpPermissionsId int64
	var CidrIp string
	var Description string

	err := rows.Scan(&Id, &SecurityGroupIpPermissionsId, &CidrIp, &Description)
	if err != nil {
		return nil, err
	}
	return NewSecurityGroupIpPermissionsIpRange(Id, SecurityGroupIpPermissionsId, CidrIp, Description), nil
}

func rowsResultSetToSecurityGroupIpPermissionIpRange(rows *sql.Rows) (*SecurityGroupIpPermissionsIpRange, error) {
	var err error
	if rows.Next() {
		var Id int64
		var SecurityGroupIpPermissionsId int64
		var CidrIp string
		var Description string

		err := rows.Scan(&Id, &SecurityGroupIpPermissionsId, &CidrIp, &Description)
		if err != nil {
			return nil, err
		}
		return NewSecurityGroupIpPermissionsIpRange(Id, SecurityGroupIpPermissionsId, CidrIp, Description), nil
	}
	return nil, err
}

func createSecurityGroupIpPermissionIpRange(db *sql.DB, SecurityGroupIpPermissionsId int64, CidrIp string, Description string) *SecurityGroupIpPermissionsIpRange {
	rows := db.QueryRow("insert into security_group_ip_permissions_ip_ranges(security_group_ip_permissions_id,cidr_ip,description) values($1,$2,$3) returning id,security_group_ip_permissions_id,cidr_ip,description ", SecurityGroupIpPermissionsId, CidrIp, Description)

	secGroupIpPermissionIpRange, err := rowResultSetToSecurityGroupIpPermissionIpRange(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return secGroupIpPermissionIpRange
}

func createSecurityGroupIpPermissionEgressIpRange(db *sql.DB, SecurityGroupIpPermissionsId int64, CidrIp string, Description string) *SecurityGroupIpPermissionsIpRange {
	rows := db.QueryRow("insert into security_group_ip_permissions_egress_ip_ranges(security_group_ip_permissions_egress_id,cidr_ip,description) values($1,$2,$3) returning id,security_group_ip_permissions_egress_id,cidr_ip,description", SecurityGroupIpPermissionsId, CidrIp, Description)

	secGroupIpPermissionEgressIpRange, err := rowResultSetToSecurityGroupIpPermissionIpRange(rows)
	if err != nil {
		fmt.Println(err)
		fmt.Println("DSDS")

		return nil
	}
	return secGroupIpPermissionEgressIpRange
}

func loadAllSecurityGroupIpPermissionIpRange(db *sql.DB) ([]*SecurityGroupIpPermissionsIpRange, error) {
	rows, err := db.Query("select id,security_group_ip_permissions_id,cidr_ip,description from security_group_ip_permissions_ip_ranges order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfSGIpPermissionIpRange := make([]*SecurityGroupIpPermissionsIpRange, 0, 0)
	for rows.Next() {
		secGroupIpPermissionIpRange, err := rowsNoFetchResultSetToSecurityGroupIpPermissionIpRange(rows)
		if err == nil {
			arrayOfSGIpPermissionIpRange = append(arrayOfSGIpPermissionIpRange, secGroupIpPermissionIpRange)
		} else {
			log.Println(err)
		}
	}

	return arrayOfSGIpPermissionIpRange, nil
}

func loadAllSecurityGroupIpPermissionEgressIpRange(db *sql.DB) ([]*SecurityGroupIpPermissionsIpRange, error) {
	rows, err := db.Query("select id,security_group_ip_permissions_egress_id,cidr_ip,description from security_group_ip_permissions_egress_ip_ranges order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfSGIpPermissionEgIpRange := make([]*SecurityGroupIpPermissionsIpRange, 0, 0)
	for rows.Next() {
		secGroupIpPermissionEgIpRange, err := rowsNoFetchResultSetToSecurityGroupIpPermissionIpRange(rows)
		if err == nil {
			arrayOfSGIpPermissionEgIpRange = append(arrayOfSGIpPermissionEgIpRange, secGroupIpPermissionEgIpRange)
		} else {
			log.Println(err)
		}
	}

	return arrayOfSGIpPermissionEgIpRange, nil
}

func loadAllSecurityGroupIpPermissionIpRangeByOrderUnique(db *sql.DB) ([]*SecurityGroupIpPermissionsIpRange, error) {
	rows, err := db.Query("select id,security_group_ip_permissions_id,cidr_ip,description from security_group_ip_permissions_ip_ranges order by cidr_ip")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfSGIpPermissionIpRange := make([]*SecurityGroupIpPermissionsIpRange, 0, 0)
	for rows.Next() {
		secGroupIpPermissionIpRange, err := rowsNoFetchResultSetToSecurityGroupIpPermissionIpRange(rows)
		if err == nil {
			arrayOfSGIpPermissionIpRange = append(arrayOfSGIpPermissionIpRange, secGroupIpPermissionIpRange)
		} else {
			log.Println(err)
		}
	}

	return arrayOfSGIpPermissionIpRange, nil
}

func loadAllSecurityGroupIpPermissionEgressIpRangeByOrderUnique(db *sql.DB) ([]*SecurityGroupIpPermissionsIpRange, error) {
	rows, err := db.Query("select id,security_group_ip_permissions_egress_id,cidr_ip,description from security_group_ip_permissions_egress_ip_ranges order by cidr_ip")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfSGIpPermissionEgIpRange := make([]*SecurityGroupIpPermissionsIpRange, 0, 0)
	for rows.Next() {
		secGroupIpPermissionEgIpRange, err := rowsNoFetchResultSetToSecurityGroupIpPermissionIpRange(rows)
		if err == nil {
			arrayOfSGIpPermissionEgIpRange = append(arrayOfSGIpPermissionEgIpRange, secGroupIpPermissionEgIpRange)
		} else {
			log.Println(err)
		}
	}

	return arrayOfSGIpPermissionEgIpRange, nil
}

func deleteAllSecurityGroupIpPermissionIpRange(db *sql.DB) error {
	_, err := db.Exec("delete from security_group_ip_permissions_ip_ranges")

	return err
}
func deleteAllSecurityGroupIpPermissionEgressIpRange(db *sql.DB) error {
	_, err := db.Exec("delete from security_group_ip_permissions_egress_ip_ranges")

	return err
}
