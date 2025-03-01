package main

import (
	"database/sql" // package SQL
	"log"

	_ "github.com/lib/pq" // driver Postgres
)

func rowResultSetToSecurityGroupIpPermissionPrefixListId(row *sql.Row) (*SecurityGroupIpPermissionsPrefixListId, error) {
	var Id int64
	var SecurityGroupIpPermissionsId int64
	var Description string
	var PrefixListId string

	err := row.Scan(&Id, &SecurityGroupIpPermissionsId, &Description, &PrefixListId)
	if err != nil {
		return nil, err
	}
	return NewSecurityGroupIpPermissionsPrefixListId(Id, SecurityGroupIpPermissionsId, Description, PrefixListId), nil
}

func rowsNoFetchResultSetToSecurityGroupIpPermissionPrefixListId(rows *sql.Rows) (*SecurityGroupIpPermissionsPrefixListId, error) {
	var Id int64
	var SecurityGroupIpPermissionsId int64
	var Description string
	var PrefixListId string

	err := rows.Scan(&Id, &SecurityGroupIpPermissionsId, &Description, &PrefixListId)
	if err != nil {
		return nil, err
	}
	return NewSecurityGroupIpPermissionsPrefixListId(Id, SecurityGroupIpPermissionsId, Description, PrefixListId), nil
}

func rowsResultSetToSecurityGroupIpPermissionPrefixListId(rows *sql.Rows) (*SecurityGroupIpPermissionsPrefixListId, error) {
	var err error
	if rows.Next() {
		var Id int64
		var SecurityGroupIpPermissionsId int64
		var Description string
		var PrefixListId string

		err := rows.Scan(&Id, &SecurityGroupIpPermissionsId, &Description, &PrefixListId)
		if err != nil {
			return nil, err
		}
		return NewSecurityGroupIpPermissionsPrefixListId(Id, SecurityGroupIpPermissionsId, Description, PrefixListId), nil
	}
	return nil, err
}

func createSecurityGroupIpPermissionPrefixListId(db *sql.DB, SecurityGroupIpPermissionsId int64, Description string, PrefixListId string) *SecurityGroupIpPermissionsPrefixListId {
	rows := db.QueryRow("insert into security_group_ip_permissions_prefix_list_ids(security_group_ip_permissions_id,description,prefix_list_id) values($1,$2,$3) returning id,security_group_ip_permissions_id,description,prefix_list_id", SecurityGroupIpPermissionsId, Description, PrefixListId)

	secGroupIpPermissionPrefixListId, err := rowResultSetToSecurityGroupIpPermissionPrefixListId(rows)
	if err != nil {
		return nil
	}
	return secGroupIpPermissionPrefixListId
}

func createSecurityGroupIpPermissionEgressPrefixListId(db *sql.DB, SecurityGroupIpPermissionsId int64, Description string, PrefixListId string) *SecurityGroupIpPermissionsPrefixListId {
	rows := db.QueryRow("insert into security_group_ip_permissions_egress_prefix_list_ids(security_group_ip_permissions_egress_id,description,prefix_list_id) values($1,$2,$3) returning id,security_group_ip_permissions_egress_id,description,prefix_list_id", SecurityGroupIpPermissionsId, Description, PrefixListId)

	secGroupIpPermissionEgressPrefixListId, err := rowResultSetToSecurityGroupIpPermissionPrefixListId(rows)
	if err != nil {
		return nil
	}
	return secGroupIpPermissionEgressPrefixListId
}

func loadAllSecurityGroupIpPermissionPrefixListId(db *sql.DB) ([]*SecurityGroupIpPermissionsPrefixListId, error) {
	rows, err := db.Query("select id,security_group_ip_permissions_id,description,prefix_list_id from security_group_ip_permissions_prefix_list_ids order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfSGIpPermissionPrefixListId := make([]*SecurityGroupIpPermissionsPrefixListId, 0, 0)
	for rows.Next() {
		secGroupIpPermissionPrefixListId, err := rowsNoFetchResultSetToSecurityGroupIpPermissionPrefixListId(rows)
		if err == nil {
			arrayOfSGIpPermissionPrefixListId = append(arrayOfSGIpPermissionPrefixListId, secGroupIpPermissionPrefixListId)
		} else {
			log.Println(err)
		}
	}

	return arrayOfSGIpPermissionPrefixListId, nil
}

func loadAllSecurityGroupIpPermissionEgressPrefixListId(db *sql.DB) ([]*SecurityGroupIpPermissionsPrefixListId, error) {
	rows, err := db.Query("select id,security_group_ip_permissions_egress_id,description,prefix_list_id from security_group_ip_permissions_egress_prefix_list_ids order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfSGIpPermissionEgPrefixListId := make([]*SecurityGroupIpPermissionsPrefixListId, 0, 0)
	for rows.Next() {
		secGroupIpPermissionEgPrefixListId, err := rowsNoFetchResultSetToSecurityGroupIpPermissionPrefixListId(rows)
		if err == nil {
			arrayOfSGIpPermissionEgPrefixListId = append(arrayOfSGIpPermissionEgPrefixListId, secGroupIpPermissionEgPrefixListId)
		} else {
			log.Println(err)
		}
	}

	return arrayOfSGIpPermissionEgPrefixListId, nil
}

func loadAllSecurityGroupIpPermissionPrefixListIdByOrderUnique(db *sql.DB) ([]*SecurityGroupIpPermissionsPrefixListId, error) {
	rows, err := db.Query("select id,security_group_ip_permissions_id,description,prefix_list_id from security_group_ip_permissions_prefix_list_ids order by prefix_list_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfSGIpPermissionPrefixListId := make([]*SecurityGroupIpPermissionsPrefixListId, 0, 0)
	for rows.Next() {
		secGroupIpPermissionPrefixListId, err := rowsNoFetchResultSetToSecurityGroupIpPermissionPrefixListId(rows)
		if err == nil {
			arrayOfSGIpPermissionPrefixListId = append(arrayOfSGIpPermissionPrefixListId, secGroupIpPermissionPrefixListId)
		} else {
			log.Println(err)
		}
	}

	return arrayOfSGIpPermissionPrefixListId, nil
}

func loadAllSecurityGroupIpPermissionEgressPrefixListIdByOrderUnique(db *sql.DB) ([]*SecurityGroupIpPermissionsPrefixListId, error) {
	rows, err := db.Query("select id,security_group_ip_permissions_egress_id,description,prefix_list_id from security_group_ip_permissions_egress_prefix_list_ids order by prefix_list_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfSGIpPermissionEgPrefixListId := make([]*SecurityGroupIpPermissionsPrefixListId, 0, 0)
	for rows.Next() {
		secGroupIpPermissionEgPrefixListId, err := rowsNoFetchResultSetToSecurityGroupIpPermissionPrefixListId(rows)
		if err == nil {
			arrayOfSGIpPermissionEgPrefixListId = append(arrayOfSGIpPermissionEgPrefixListId, secGroupIpPermissionEgPrefixListId)
		} else {
			log.Println(err)
		}
	}

	return arrayOfSGIpPermissionEgPrefixListId, nil
}

func deleteAllSecurityGroupIpPermissionPrefixListId(db *sql.DB) error {
	_, err := db.Exec("delete from security_group_ip_permissions_prefix_list_ids")

	return err
}
func deleteAllSecurityGroupIpPermissionEgressPrefixListId(db *sql.DB) error {
	_, err := db.Exec("delete from security_group_ip_permissions_egress_prefix_list_ids")

	return err
}
