package main

import (
	"database/sql" // package SQL
	"fmt"
	"log"

	_ "github.com/lib/pq" // driver Postgres
)

func rowResultSetToViewELBByCidr(row *sql.Row) (*ViewELBByCidr, error) {
	var err error
	var SubnetId int64
	var SubnetCidrBlock string
	var LoadBalancerId int64
	var LoadBalancerName string

	err = row.Scan(&SubnetId, &SubnetCidrBlock, &LoadBalancerId, &LoadBalancerName)
	if err != nil {
		return nil, err
	}
	return NewViewELBByCidr(SubnetId, SubnetCidrBlock, LoadBalancerId, LoadBalancerName), nil
}

func rowsNoFetchResultSetToViewELBByCidr(rows *sql.Rows) (*ViewELBByCidr, error) {
	var err error
	var SubnetId int64
	var SubnetCidrBlock string
	var LoadBalancerId int64
	var LoadBalancerName string

	err = rows.Scan(&SubnetId, &SubnetCidrBlock, &LoadBalancerId, &LoadBalancerName)
	if err != nil {
		return nil, err
	}
	return NewViewELBByCidr(SubnetId, SubnetCidrBlock, LoadBalancerId, LoadBalancerName), nil
}

func rowsResultSetToViewELBByCidr(rows *sql.Rows) (*ViewELBByCidr, error) {
	var err error
	if rows.Next() {
		var SubnetId int64
		var SubnetCidrBlock string
		var LoadBalancerId int64
		var LoadBalancerName string

		err = rows.Scan(&SubnetId, &SubnetCidrBlock, &LoadBalancerId, &LoadBalancerName)
		if err != nil {
			return nil, err
		}
		return NewViewELBByCidr(SubnetId, SubnetCidrBlock, LoadBalancerId, LoadBalancerName), nil
	}
	return nil, err
}

func createViewELBByCidr(db *sql.DB, SubnetId int64, SubnetCidrBlock string, LoadBalancerId int64, LoadBalancerName string) *ViewELBByCidr {
	rows := db.QueryRow("insert into view_elb_by_cidr(subnet_id,subnet_cidr_block,load_balancer_id,load_balancer_name) values($1,$2,$3,$4) returning subnet_id,subnet_cidr_block,load_balancer_id,load_balancer_name",
		SubnetId, SubnetCidrBlock, LoadBalancerId, LoadBalancerName)

	viewELBByCidr, err := rowResultSetToViewELBByCidr(rows)
	if err != nil {
		return nil
	}
	return viewELBByCidr
}

func loadAllViewELBByCidrDistinct(db *sql.DB) ([]string, error) {
	rows, err := db.Query("select distinct(load_balancer_name) from public.view_elb_by_cidr ;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfLoadBalancerName := make([]string, 0, 0)
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err == nil {
			arrayOfLoadBalancerName = append(arrayOfLoadBalancerName, name)
		} else {
			log.Println(err)
		}
	}

	return arrayOfLoadBalancerName, nil
}

func loadAllViewELBByCidr(db *sql.DB) ([]*ViewELBByCidr, error) {
	rows, err := db.Query("select subnet_id,subnet_cidr_block,load_balancer_id,load_balancer_name from view_elb_by_cidr order by subnet_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfViewELBByCidr := make([]*ViewELBByCidr, 0, 0)
	for rows.Next() {
		viewELBByCidr, err := rowsNoFetchResultSetToViewELBByCidr(rows)
		if err == nil {
			arrayOfViewELBByCidr = append(arrayOfViewELBByCidr, viewELBByCidr)
		} else {
			log.Println(err)
		}
	}

	return arrayOfViewELBByCidr, nil
}

func loadAllViewELBByCidrLoadBalancerName(db *sql.DB, cidrIp string) (string, error) {
	var err error
	var name string
	query := fmt.Sprintf("select distinct(load_balancer_name) from view_elb_by_cidr where subnet_cidr_block like '%s", cidrIp)
	rows, err := db.Query(query + "%'")
	if err != nil {
		return "", err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&name)
		if err != nil {
			return "", err
		}
	}

	return name, nil
}

func loadAllViewELBByCidrCidrByName(db *sql.DB, loadBalancerName string) (string, error) {
	var err error
	var name string
	query := fmt.Sprintf("select distinct(subnet_cidr_block) from view_elb_by_cidr where load_balancer_name='%s' limit 1", loadBalancerName)
	rows, err := db.Query(query)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&name)
		if err != nil {
			return "", err
		}
	}

	return name, nil
}

func deleteAllViewELBByCidr(db *sql.DB) error {
	_, err := db.Exec("delete from view_elb_by_cidr")

	return err
}
