package main

import (
	"database/sql" // package SQL
	"log"

	_ "github.com/lib/pq" // driver Postgres
)

func rowResultSetToSubnet(row *sql.Row) (*Subnet, error) {
	var err error
	var Id int64
	var RefVpcId int64
	var SubnetId string
	var MapPublicAtLaunch bool
	var AvailabilityZoneId string
	var Name string
	var CloudformationStackName string
	var AvailableIpAddressCount int16
	var State string
	var AvailabilityZone string
	var OwnerId string
	var CidrBlock string

	err = row.Scan(&Id, &RefVpcId, &SubnetId, &MapPublicAtLaunch, &AvailabilityZoneId,
		&Name, &CloudformationStackName, &AvailableIpAddressCount, &State, &AvailabilityZone,
		&OwnerId, &CidrBlock)
	if err != nil {
		return nil, err
	}
	return NewSubnet(Id, RefVpcId, SubnetId, MapPublicAtLaunch, AvailabilityZoneId, Name, CloudformationStackName,
		AvailableIpAddressCount, State, AvailabilityZone, OwnerId, CidrBlock), nil
}

func rowsNoFetchResultSetToSubnet(rows *sql.Rows) (*Subnet, error) {
	var err error
	var Id int64
	var RefVpcId int64
	var SubnetId string
	var MapPublicAtLaunch bool
	var AvailabilityZoneId string
	var Name string
	var CloudformationStackName string
	var AvailableIpAddressCount int16
	var State string
	var AvailabilityZone string
	var OwnerId string
	var CidrBlock string

	err = rows.Scan(&Id, &RefVpcId, &SubnetId, &MapPublicAtLaunch, &AvailabilityZoneId,
		&Name, &CloudformationStackName, &AvailableIpAddressCount, &State, &AvailabilityZone,
		&OwnerId, &CidrBlock)
	if err != nil {
		return nil, err
	}
	return NewSubnet(Id, RefVpcId, SubnetId, MapPublicAtLaunch, AvailabilityZoneId, Name, CloudformationStackName,
		AvailableIpAddressCount, State, AvailabilityZone, OwnerId, CidrBlock), nil
}

func rowsResultSetToSubnet(rows *sql.Rows) (*Subnet, error) {
	var err error
	if rows.Next() {
		var Id int64
		var RefVpcId int64
		var SubnetId string
		var MapPublicAtLaunch bool
		var AvailabilityZoneId string
		var Name string
		var CloudformationStackName string
		var AvailableIpAddressCount int16
		var State string
		var AvailabilityZone string
		var OwnerId string
		var CidrBlock string

		err = rows.Scan(&Id, &RefVpcId, &SubnetId, &MapPublicAtLaunch, &AvailabilityZoneId,
			&Name, &CloudformationStackName, &AvailableIpAddressCount, &State, &AvailabilityZone,
			&OwnerId, &CidrBlock)
		if err != nil {
			return nil, err
		}
		return NewSubnet(Id, RefVpcId, SubnetId, MapPublicAtLaunch, AvailabilityZoneId, Name, CloudformationStackName,
			AvailableIpAddressCount, State, AvailabilityZone, OwnerId, CidrBlock), nil
	}
	return nil, err
}

func createSubnet(db *sql.DB, RefVpcId int64, SubnetId string, MapPublicAtLaunch bool, AvailabilityZoneId string, Name string, CloudformationStackName string, AvailableIpAddressCount int16, State string, AvailabilityZone string, OwnerId string, CidrBlock string) *Subnet {
	rows := db.QueryRow("insert into subnet(ref_vpc_id,subnet_id,map_public_at_launch,availability_zone_id,name,cloudformation_stack_name,available_ip_address_count,state,availability_zone,owner_id,cidr_block) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) returning id,ref_vpc_id,subnet_id,map_public_at_launch,availability_zone_id,name,cloudformation_stack_name,available_ip_address_count,state,availability_zone,owner_id,cidr_block", RefVpcId, SubnetId, MapPublicAtLaunch, AvailabilityZoneId, Name, CloudformationStackName, AvailableIpAddressCount, State, AvailabilityZone, OwnerId, CidrBlock)

	subnet, err := rowResultSetToSubnet(rows)
	if err != nil {
		return nil
	}
	return subnet
}

func loadSubnetById(db *sql.DB, id int64) (*Subnet, error) {
	rows, err := db.Query("select id,ref_vpc_id,subnet_id,map_public_at_launch,availability_zone_id,name,cloudformation_stack_name,available_ip_address_count,state,availability_zone,owner_id,cidr_block from subnet where id=$1", id)
	if err != nil {
		return nil, err
	}

	subnet, err := rowsResultSetToSubnet(rows)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	return subnet, nil
}

func loadSubnetByName(db *sql.DB, subnetName string) (*Subnet, error) {
	rows, err := db.Query("select id,ref_vpc_id,subnet_id,map_public_at_launch,availability_zone_id,name,cloudformation_stack_name,available_ip_address_count,state,availability_zone,owner_id,cidr_block from subnet where name=$1", subnetName)
	if err != nil {
		return nil, err
	}

	subnet, err := rowsResultSetToSubnet(rows)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	return subnet, nil
}

func loadSubnetByAWSReferenceId(db *sql.DB, subnetReferenceId string) (*Subnet, error) {
	rows, err := db.Query("select id,ref_vpc_id,subnet_id,map_public_at_launch,availability_zone_id,name,cloudformation_stack_name,available_ip_address_count,state,availability_zone,owner_id,cidr_block from subnet where subnet_id=$1", subnetReferenceId)
	if err != nil {
		return nil, err
	}

	subnet, err := rowsResultSetToSubnet(rows)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	return subnet, nil
}

func loadAllSubnet(db *sql.DB) ([]*Subnet, error) {
	rows, err := db.Query("select id,ref_vpc_id,subnet_id,map_public_at_launch,availability_zone_id,name,cloudformation_stack_name,available_ip_address_count,state,availability_zone,owner_id,cidr_block from subnet order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfSubnet := make([]*Subnet, 0, 0)
	for rows.Next() {
		vpc, err := rowsNoFetchResultSetToSubnet(rows)
		if err == nil {
			arrayOfSubnet = append(arrayOfSubnet, vpc)
		} else {
			log.Println(err)
		}
	}

	return arrayOfSubnet, nil
}

func loadAllSubnetByOrderUnique(db *sql.DB) ([]*Subnet, error) {
	rows, err := db.Query("select id,ref_vpc_id,subnet_id,map_public_at_launch,availability_zone_id,name,cloudformation_stack_name,available_ip_address_count,state,availability_zone,owner_id,cidr_block from subnet order by subnet_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfSubnet := make([]*Subnet, 0, 0)
	for rows.Next() {
		vpc, err := rowsNoFetchResultSetToSubnet(rows)
		if err == nil {
			arrayOfSubnet = append(arrayOfSubnet, vpc)
		} else {
			log.Println(err)
		}
	}

	return arrayOfSubnet, nil
}

func deleteAllSubnet(db *sql.DB) error {
	_, err := db.Exec("delete from subnet")

	return err
}
