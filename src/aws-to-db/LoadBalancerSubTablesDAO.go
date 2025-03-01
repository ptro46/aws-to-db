package main

import (
	"database/sql" // package SQL
	"fmt"
	"log"

	_ "github.com/lib/pq" // driver Postgres
)

//////////////////////////////////////////////////////////
//RowResult
//////////////////////////////////////////////////////////

func rowResultSetToLoadBalancerAvailabilityZone(row *sql.Row) (*LoadBalancerAvailabilityZone, error) {
	var err error
	var Id int64
	var LoadBalancerId int64
	var ZoneName string
	var SubnetId string
	var RefSubnetId sql.NullInt64

	err = row.Scan(&Id, &LoadBalancerId, &ZoneName, &SubnetId, &RefSubnetId)
	if err != nil {
		return nil, err
	}
	return newLoadBalancerAvailabilityZone(Id, LoadBalancerId, ZoneName, SubnetId, RefSubnetId), nil
}

func rowResultSetToLoadBalancerAddress(row *sql.Row) (*LoadBalancerAddress, error) {
	var err error
	var Id int64
	var LoadBalancerAvailabilityZoneId int64
	var IpAddress string
	var AllocationId string

	err = row.Scan(&Id, &LoadBalancerAvailabilityZoneId, &IpAddress, &AllocationId)
	if err != nil {
		return nil, err
	}
	return newLoadBalancerAddress(Id, LoadBalancerAvailabilityZoneId, IpAddress, AllocationId), nil
}

func rowResultSetToLoadBalancerSecurityGroup(row *sql.Row) (*LoadBalancerSecurityGroup, error) {
	var err error
	var Id int64
	var LoadBalancerId int64
	var SecurityGroupId string
	var RefSecurityGroupId sql.NullInt64

	err = row.Scan(&Id, &LoadBalancerId, &SecurityGroupId, &RefSecurityGroupId)
	if err != nil {
		return nil, err
	}
	return newLoadBalancerSecurityGroup(Id, LoadBalancerId, SecurityGroupId, RefSecurityGroupId), nil
}

//////////////////////////////////////////////////////////
//RowsNoFetch
//////////////////////////////////////////////////////////

func rowsNoFetchResultSetToLoadBalancerAvailabilityZone(rows *sql.Rows) (*LoadBalancerAvailabilityZone, error) {
	var err error
	var Id int64
	var LoadBalancerId int64
	var ZoneName string
	var SubnetId string
	var RefSubnetId sql.NullInt64

	err = rows.Scan(&Id, &LoadBalancerId, &ZoneName, &SubnetId, &RefSubnetId)
	if err != nil {
		return nil, err
	}
	return newLoadBalancerAvailabilityZone(Id, LoadBalancerId, ZoneName, SubnetId, RefSubnetId), nil
}

func rowsNoFetchResultSetToLoadBalancerAddress(rows *sql.Rows) (*LoadBalancerAddress, error) {
	var err error
	var Id int64
	var LoadBalancerAvailabilityZoneId int64
	var IpAddress string
	var AllocationId string

	err = rows.Scan(&Id, &LoadBalancerAvailabilityZoneId, &IpAddress, &AllocationId)
	if err != nil {
		return nil, err
	}
	return newLoadBalancerAddress(Id, LoadBalancerAvailabilityZoneId, IpAddress, AllocationId), nil
}

func rowsNoFetchResultSetToLoadBalancerSecurityGroup(rows *sql.Rows) (*LoadBalancerSecurityGroup, error) {
	var err error
	var Id int64
	var LoadBalancerId int64
	var SecurityGroupId string
	var RefSecurityGroupId sql.NullInt64

	err = rows.Scan(&Id, &LoadBalancerId, &SecurityGroupId, &RefSecurityGroupId)
	if err != nil {
		return nil, err
	}
	return newLoadBalancerSecurityGroup(Id, LoadBalancerId, SecurityGroupId, RefSecurityGroupId), nil
}

//////////////////////////////////////////////////////////
//RowsResult
//////////////////////////////////////////////////////////

func rowsResultSetToLoadBalancerAvailabilityZone(rows *sql.Rows) (*LoadBalancerAvailabilityZone, error) {
	var err error
	if rows.Next() {
		var Id int64
		var LoadBalancerId int64
		var ZoneName string
		var SubnetId string
		var RefSubnetId sql.NullInt64

		err = rows.Scan(&Id, &LoadBalancerId, &ZoneName, &SubnetId, &RefSubnetId)
		if err != nil {
			return nil, err
		}
		return newLoadBalancerAvailabilityZone(Id, LoadBalancerId, ZoneName, SubnetId, RefSubnetId), nil
	}

	return nil, err
}

func rowsResultSetToLoadBalancerAddress(rows *sql.Rows) (*LoadBalancerAddress, error) {
	var err error
	if rows.Next() {
		var Id int64
		var LoadBalancerAvailabilityZoneId int64
		var IpAddress string
		var AllocationId string

		err = rows.Scan(&Id, &LoadBalancerAvailabilityZoneId, &IpAddress, &AllocationId)
		if err != nil {
			return nil, err
		}
		return newLoadBalancerAddress(Id, LoadBalancerAvailabilityZoneId, IpAddress, AllocationId), nil
	}

	return nil, err

}

func rowsResultSetToLoadBalancerSecurityGroup(rows *sql.Rows) (*LoadBalancerSecurityGroup, error) {
	var err error
	if rows.Next() {
		var Id int64
		var LoadBalancerId int64
		var SecurityGroupId string
		var RefSecurityGroupId sql.NullInt64

		err = rows.Scan(&Id, &LoadBalancerId, &SecurityGroupId, &RefSecurityGroupId)
		if err != nil {
			return nil, err
		}
		return newLoadBalancerSecurityGroup(Id, LoadBalancerId, SecurityGroupId, RefSecurityGroupId), nil
	}

	return nil, err

}

//////////////////////////////////////////////////////////
//Create
//////////////////////////////////////////////////////////

func createLoadBalancerAvailabilityZone(db *sql.DB, LoadBalancerId int64, ZoneName string, SubnetId string) *LoadBalancerAvailabilityZone {
	rows := db.QueryRow("insert into load_balancer_availability_zone(load_balancer_id,zone_name,subnet_id) values($1,$2,$3) returning id,load_balancer_id,zone_name,subnet_id,ref_subnet_id",
		LoadBalancerId, ZoneName, SubnetId)

	loadZone, err := rowResultSetToLoadBalancerAvailabilityZone(rows)
	if err != nil {
		return nil
	}
	return loadZone
}

func createLoadBalancerAddress(db *sql.DB, LoadBalancerAvailabilityZoneId int64, IpAddress string, AllocationId string) *LoadBalancerAddress {
	rows := db.QueryRow("insert into load_balancer_availability_zone_load_balancer_address(load_balancer_availability_zone_id,ip_address,allocation_id) values($1,$2,$3) returning id,load_balancer_availability_zone_id,ip_address,allocation_id",
		LoadBalancerAvailabilityZoneId, IpAddress, AllocationId)

	loadAddress, err := rowResultSetToLoadBalancerAddress(rows)
	if err != nil {
		return nil
	}
	return loadAddress
}

func createLoadBalancerSecurityGroup(db *sql.DB, LoadBalancerId int64, SecurityGroupId string) *LoadBalancerSecurityGroup {
	rows := db.QueryRow("insert into load_balancer_security_group(load_balancer_id,security_group_id) values($1,$2) returning id,load_balancer_id,security_group_id,ref_security_group_id",
		LoadBalancerId, SecurityGroupId)

	loadSecGroup, err := rowResultSetToLoadBalancerSecurityGroup(rows)
	if err != nil {
		return nil
	}
	return loadSecGroup
}

//////////////////////////////////////////////////////////
//LoadAll
//////////////////////////////////////////////////////////

func loadAllLoadBalancerSecurityGroup(db *sql.DB) ([]*LoadBalancerSecurityGroup, error) {
	rows, err := db.Query("select id,load_balancer_id,security_group_id,ref_security_group_id from load_balancer_security_group order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfLoadBalancerSecurityGroup := make([]*LoadBalancerSecurityGroup, 0, 0)
	for rows.Next() {
		loadBalancerSecurityGroup, err := rowsNoFetchResultSetToLoadBalancerSecurityGroup(rows)
		if err == nil {
			arrayOfLoadBalancerSecurityGroup = append(arrayOfLoadBalancerSecurityGroup, loadBalancerSecurityGroup)
		} else {
			log.Println(err)
		}
	}
	return arrayOfLoadBalancerSecurityGroup, nil
}

func loadAllLoadBalancerSecurityGroupByOrderUnique(db *sql.DB) ([]*LoadBalancerSecurityGroup, error) {
	rows, err := db.Query("select id,load_balancer_id,security_group_id,ref_security_group_id from load_balancer_security_group order by security_group_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfLoadBalancerSecurityGroup := make([]*LoadBalancerSecurityGroup, 0, 0)
	for rows.Next() {
		loadBalancerSecurityGroup, err := rowsNoFetchResultSetToLoadBalancerSecurityGroup(rows)
		if err == nil {
			arrayOfLoadBalancerSecurityGroup = append(arrayOfLoadBalancerSecurityGroup, loadBalancerSecurityGroup)
		} else {
			log.Println(err)
		}
	}
	return arrayOfLoadBalancerSecurityGroup, nil
}

func loadAllLoadBalancerAvailabilityZone(db *sql.DB) []*LoadBalancerAvailabilityZone {
	results := make([]*LoadBalancerAvailabilityZone, 0, 0)
	sqlSelect := "SELECT id,load_balancer_id,zone_name,subnet_id,ref_subnet_id from load_balancer_availability_zone"

	rows, err := db.Query(sqlSelect)
	if err != nil {
		return results
	}
	defer rows.Close()

	for rows.Next() {
		loadBalancerAvailabilityZone, err := rowsNoFetchResultSetToLoadBalancerAvailabilityZone(rows)
		if err != nil {
			return make([]*LoadBalancerAvailabilityZone, 0, 0)
		}
		results = append(results, loadBalancerAvailabilityZone)
	}
	return results
}

func loadAllLoadBalancerAvailabilityZoneByOrderUnique(db *sql.DB) ([]*LoadBalancerAvailabilityZone, error) {
	results := make([]*LoadBalancerAvailabilityZone, 0, 0)
	sqlSelect := "SELECT id,load_balancer_id,zone_name,subnet_id,ref_subnet_id from load_balancer_availability_zone order by subnet_id"

	rows, err := db.Query(sqlSelect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		loadBalancerAvailabilityZone, err := rowsNoFetchResultSetToLoadBalancerAvailabilityZone(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, loadBalancerAvailabilityZone)
	}
	return results, nil
}

func loadAllLoadBalancerAddressByOrderUnique(db *sql.DB) ([]*LoadBalancerAddress, error) {
	results := make([]*LoadBalancerAddress, 0, 0)
	sqlSelect := "SELECT id,load_balancer_availability_zone_id,ip_address,allocation_id from load_balancer_availability_zone_load_balancer_address order by load_balancer_availability_zone_id"

	rows, err := db.Query(sqlSelect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		loadBalancerAddress, err := rowsNoFetchResultSetToLoadBalancerAddress(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, loadBalancerAddress)
	}
	return results, nil
}

func updateLoadBalancerSecurityGroupId(db *sql.DB, loadBalancerSecurityGroup *LoadBalancerSecurityGroup, secGroup int64) *LoadBalancerSecurityGroup {
	rows := db.QueryRow("update load_balancer_security_group set ref_security_group_id=$1 where id=$2 returning id,load_balancer_id,security_group_id,ref_security_group_id", secGroup, loadBalancerSecurityGroup.Id)

	loadBalancerSecurityGroup, err := rowResultSetToLoadBalancerSecurityGroup(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return loadBalancerSecurityGroup
}

func updateLoadBalancerAvailabilityZoneId(db *sql.DB, loadBalancerAvailabilityZone *LoadBalancerAvailabilityZone, subnetId int64) *LoadBalancerAvailabilityZone {
	rows := db.QueryRow("update load_balancer_availability_zone set ref_subnet_id=$1 where id=$2 returning id,load_balancer_id,zone_name,subnet_id,ref_subnet_id", subnetId, loadBalancerAvailabilityZone.Id)

	loadBalancerAvailabilityZone, err := rowResultSetToLoadBalancerAvailabilityZone(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return loadBalancerAvailabilityZone
}

//////////////////////////////////////////////////////////
//DeleteAll
//////////////////////////////////////////////////////////

func deleteAllLoadBalancerAvailabilityZone(db *sql.DB) error {
	_, err := db.Exec("delete from load_balancer_availability_zone")

	return err
}

func deleteAllLoadBalancerAddress(db *sql.DB) error {
	_, err := db.Exec("delete from load_balancer_availability_zone_load_balancer_address")

	return err
}
func deleteAllLoadBalancerSecurityGroup(db *sql.DB) error {
	_, err := db.Exec("delete from load_balancer_security_group")

	return err
}
