package main

import (
	"database/sql" // package SQL
	"fmt"
	"log"

	_ "github.com/lib/pq" // driver Postgres
)

func rowResultSetToLoadBalancer(row *sql.Row) (*LoadBalancer, error) {
	var err error
	var Id int64
	var LoadBalancerArn string
	var DNSName string
	var CanonicalHostedZoneId string
	var CreatedTime string
	var LoadBalancerName string
	var Scheme string
	var VpcId string
	var RefVpcId sql.NullInt64
	var StateCode string
	var StateReason string
	var Type string
	var IpAddressType string

	err = row.Scan(&Id, &LoadBalancerArn, &DNSName, &CanonicalHostedZoneId, &CreatedTime, &LoadBalancerName, &Scheme, &VpcId, &RefVpcId, &StateCode, &StateReason, &Type, &IpAddressType)
	if err != nil {
		return nil, err
	}
	return newLoadBalancer(Id, LoadBalancerArn, DNSName, CanonicalHostedZoneId, CreatedTime, LoadBalancerName, Scheme, VpcId, RefVpcId, StateCode, StateReason, Type, IpAddressType), nil
}

func rowsNoFetchResultSetToLoadBalancer(rows *sql.Rows) (*LoadBalancer, error) {
	var err error
	var Id int64
	var LoadBalancerArn string
	var DNSName string
	var CanonicalHostedZoneId string
	var CreatedTime string
	var LoadBalancerName string
	var Scheme string
	var VpcId string
	var RefVpcId sql.NullInt64
	var StateCode string
	var StateReason string
	var Type string
	var IpAddressType string

	err = rows.Scan(&Id, &LoadBalancerArn, &DNSName, &CanonicalHostedZoneId, &CreatedTime, &LoadBalancerName, &Scheme, &VpcId, &RefVpcId, &StateCode, &StateReason, &Type, &IpAddressType)
	if err != nil {
		return nil, err
	}
	return newLoadBalancer(Id, LoadBalancerArn, DNSName, CanonicalHostedZoneId, CreatedTime, LoadBalancerName, Scheme, VpcId, RefVpcId, StateCode, StateReason, Type, IpAddressType), nil
}

func rowsResultSetToLoadBalancer(rows *sql.Rows) (*LoadBalancer, error) {
	var err error
	if rows.Next() {
		var Id int64
		var LoadBalancerArn string
		var DNSName string
		var CanonicalHostedZoneId string
		var CreatedTime string
		var LoadBalancerName string
		var Scheme string
		var VpcId string
		var RefVpcId sql.NullInt64
		var StateCode string
		var StateReason string
		var Type string
		var IpAddressType string

		err = rows.Scan(&Id, &LoadBalancerArn, &DNSName, &CanonicalHostedZoneId, &CreatedTime, &LoadBalancerName, &Scheme, &VpcId, &RefVpcId, &StateCode, &StateReason, &Type, &IpAddressType)
		if err != nil {
			return nil, err
		}
		return newLoadBalancer(Id, LoadBalancerArn, DNSName, CanonicalHostedZoneId, CreatedTime, LoadBalancerName, Scheme, VpcId, RefVpcId, StateCode, StateReason, Type, IpAddressType), nil
	}
	return nil, err
}

func createLoadBalancer(db *sql.DB, LoadBalancerArn string, DNSName string, CanonicalHostedZoneId string, CreatedTime string, LoadBalancerName string, Scheme string, VpcId string, StateCode string, StateReason string, Type string, IpAddressType string) *LoadBalancer {
	rows := db.QueryRow("insert into load_balancer(load_balancer_arn,dns_name,canonical_hosted_zone_id,created_time,load_balancer_name,scheme,vpc_id,state_code,state_reason,type,ip_address_type) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) returning id,load_balancer_arn,dns_name,canonical_hosted_zone_id,created_time,load_balancer_name,scheme,vpc_id,ref_vpc_id,state_code,state_reason,type,ip_address_type",
		LoadBalancerArn, DNSName, CanonicalHostedZoneId, CreatedTime, LoadBalancerName, Scheme, VpcId, StateCode, StateReason, Type, IpAddressType)

	load, err := rowResultSetToLoadBalancer(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return load
}

func loadAllLoadBalancer(db *sql.DB) ([]*LoadBalancer, error) {
	rows, err := db.Query("select id,load_balancer_arn,dns_name,canonical_hosted_zone_id,created_time,load_balancer_name,scheme,vpc_id,ref_vpc_id,state_code,state_reason,type,ip_address_type from load_balancer order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfLoad := make([]*LoadBalancer, 0, 0)
	for rows.Next() {
		load, err := rowsNoFetchResultSetToLoadBalancer(rows)
		if err == nil {
			arrayOfLoad = append(arrayOfLoad, load)
		} else {
			log.Println(err)
		}
	}

	return arrayOfLoad, nil
}

func loadAllLoadBalancerByOrderUnique(db *sql.DB) ([]*LoadBalancer, error) {
	rows, err := db.Query("select id,load_balancer_arn,dns_name,canonical_hosted_zone_id,created_time,load_balancer_name,scheme,vpc_id,ref_vpc_id,state_code,state_reason,type,ip_address_type from load_balancer order by load_balancer_name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfLoad := make([]*LoadBalancer, 0, 0)
	for rows.Next() {
		load, err := rowsNoFetchResultSetToLoadBalancer(rows)
		if err == nil {
			arrayOfLoad = append(arrayOfLoad, load)
		} else {
			log.Println(err)
		}
	}

	return arrayOfLoad, nil
}

func loadAllLoadBalancerByARN(db *sql.DB, arn string) (*LoadBalancer, error) {
	rows, err := db.Query("select id,load_balancer_arn,dns_name,canonical_hosted_zone_id,created_time,load_balancer_name,scheme,vpc_id,ref_vpc_id,state_code,state_reason,type,ip_address_type from load_balancer where load_balancer_arn=$1", arn)
	if err != nil {
		return nil, err
	}

	loadBalancer, err := rowsResultSetToLoadBalancer(rows)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	return loadBalancer, nil
}

func updateLoadBalancer(db *sql.DB, loadBalancer *LoadBalancer, vpcId int64) *LoadBalancer {
	rows := db.QueryRow("update load_balancer set ref_vpc_id=$1 where id=$2 returning id,load_balancer_arn,dns_name,canonical_hosted_zone_id,created_time,load_balancer_name,scheme,vpc_id,ref_vpc_id,state_code,state_reason,type,ip_address_type", vpcId, loadBalancer.Id)

	loadBalancer, err := rowResultSetToLoadBalancer(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return loadBalancer
}

func deleteAllLoadBalancer(db *sql.DB) error {
	_, err := db.Exec("delete from load_balancer")

	return err
}
