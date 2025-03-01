package main

import (
	"database/sql" // package SQL
	"log"

	_ "github.com/lib/pq" // driver Postgres
)

func rowResultSetToSecurityGroupIpPermission(row *sql.Row) (*SecurityGroupIpPermission, error) {
	var Id int64
	var SecurityGroupId int64
	var FromPort int16
	var IpProtocol string
	var ToPort int16

	err := row.Scan(&Id, &SecurityGroupId, &FromPort, &IpProtocol, &ToPort)
	if err != nil {
		return nil, err
	}
	return NewSecurityGroupIpPermission(Id, SecurityGroupId, FromPort, IpProtocol, ToPort), nil
}

func rowsNoFetchResultSetToSecurityGroupIpPermission(rows *sql.Rows) (*SecurityGroupIpPermission, error) {
	var Id int64
	var SecurityGroupId int64
	var FromPort int16
	var IpProtocol string
	var ToPort int16

	err := rows.Scan(&Id, &SecurityGroupId, &FromPort, &IpProtocol, &ToPort)
	if err != nil {
		return nil, err
	}
	return NewSecurityGroupIpPermission(Id, SecurityGroupId, FromPort, IpProtocol, ToPort), nil
}

func rowsResultSetToSecurityGroupIpPermission(rows *sql.Rows) (*SecurityGroupIpPermission, error) {
	var err error
	if rows.Next() {
		var Id int64
		var SecurityGroupId int64
		var FromPort int16
		var IpProtocol string
		var ToPort int16

		err = rows.Scan(&Id, &SecurityGroupId, &FromPort, &IpProtocol, &ToPort)
		if err != nil {
			return nil, err
		}
		return NewSecurityGroupIpPermission(Id, SecurityGroupId, FromPort, IpProtocol, ToPort), nil
	}
	return nil, err
}

func createSecurityGroupIpPermission(db *sql.DB, SecurityGroupId int64, FromPort int16, IpProtocol string, ToPort int16) *SecurityGroupIpPermission {
	rows := db.QueryRow("insert into security_group_ip_permissions(security_group_id,from_port,ip_protocol,to_port) values($1,$2,$3,$4) returning id,security_group_id,from_port,ip_protocol, to_port", SecurityGroupId, FromPort, IpProtocol, ToPort)

	secGroupIpPermission, err := rowResultSetToSecurityGroupIpPermission(rows)
	if err != nil {
		return nil
	}
	return secGroupIpPermission
}

func createSecurityGroupIpPermissionEgress(db *sql.DB, SecurityGroupId int64, FromPort int16, IpProtocol string, ToPort int16) *SecurityGroupIpPermission {
	rows := db.QueryRow("insert into security_group_ip_permissions_egress(security_group_id,from_port,ip_protocol,to_port) values($1,$2,$3,$4) returning id,security_group_id,from_port,ip_protocol, to_port", SecurityGroupId, FromPort, IpProtocol, ToPort)

	secGroupIpPermissionEgress, err := rowResultSetToSecurityGroupIpPermission(rows)
	if err != nil {
		return nil
	}
	return secGroupIpPermissionEgress
}

func loadAllSecurityGroupIpPermission(db *sql.DB) ([]*SecurityGroupIpPermission, error) {
	rows, err := db.Query("select id,security_group_id,from_port,ip_protocol,to_port from security_group_ip_permissions order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfSGIpPermission := make([]*SecurityGroupIpPermission, 0, 0)
	for rows.Next() {
		secGroupIpPermission, err := rowsNoFetchResultSetToSecurityGroupIpPermission(rows)
		if err == nil {
			arrayOfSGIpPermission = append(arrayOfSGIpPermission, secGroupIpPermission)
		} else {
			log.Println(err)
		}
	}

	return arrayOfSGIpPermission, nil
}

func loadAllSecurityGroupIpPermissionEgress(db *sql.DB) ([]*SecurityGroupIpPermission, error) {
	rows, err := db.Query("select id,security_group_id,from_port,ip_protocol,to_port from security_group_ip_permissions_egress order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfSGIpPermissionEg := make([]*SecurityGroupIpPermission, 0, 0)
	for rows.Next() {
		secGroupIpPermissionEg, err := rowsNoFetchResultSetToSecurityGroupIpPermission(rows)
		if err == nil {
			arrayOfSGIpPermissionEg = append(arrayOfSGIpPermissionEg, secGroupIpPermissionEg)
		} else {
			log.Println(err)
		}
	}

	return arrayOfSGIpPermissionEg, nil
}

func loadAllSecurityGroupIpPermissionByOrderUnique(db *sql.DB) ([]*SecurityGroupIpPermission, error) {
	rows, err := db.Query("select id,security_group_id,from_port,ip_protocol,to_port from security_group_ip_permissions order by from_port")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfSGIpPermission := make([]*SecurityGroupIpPermission, 0, 0)
	for rows.Next() {
		secGroupIpPermission, err := rowsNoFetchResultSetToSecurityGroupIpPermission(rows)
		if err == nil {
			arrayOfSGIpPermission = append(arrayOfSGIpPermission, secGroupIpPermission)
		} else {
			log.Println(err)
		}
	}

	return arrayOfSGIpPermission, nil
}

func loadAllSecurityGroupIpPermissionEgressByOrderUnique(db *sql.DB) ([]*SecurityGroupIpPermission, error) {
	rows, err := db.Query("select id,security_group_id,from_port,ip_protocol,to_port from security_group_ip_permissions_egress order by from_port")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfSGIpPermissionEg := make([]*SecurityGroupIpPermission, 0, 0)
	for rows.Next() {
		secGroupIpPermissionEg, err := rowsNoFetchResultSetToSecurityGroupIpPermission(rows)
		if err == nil {
			arrayOfSGIpPermissionEg = append(arrayOfSGIpPermissionEg, secGroupIpPermissionEg)
		} else {
			log.Println(err)
		}
	}

	return arrayOfSGIpPermissionEg, nil
}

func deleteAllSecurityGroupIpPermission(db *sql.DB) error {
	_, err := db.Exec("delete from security_group_ip_permissions")

	return err
}
func deleteAllSecurityGroupIpPermissionEgress(db *sql.DB) error {
	_, err := db.Exec("delete from security_group_ip_permissions_egress")

	return err
}
