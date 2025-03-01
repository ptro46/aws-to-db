package main

import (
	"database/sql" // package SQL
	"fmt"
	"log"

	_ "github.com/lib/pq" // driver Postgres
)

func rowResultSetToViewElbInputCidr(row *sql.Row) (*ViewElbInputCidr, error) {
	var err error
	var LoadBalancerId int64
	var LoadBalancerName string
	var CidrIp string

	err = row.Scan(&LoadBalancerId, &LoadBalancerName, &CidrIp)
	if err != nil {
		return nil, err
	}
	return NewViewElbInputCidr(LoadBalancerId, LoadBalancerName, CidrIp), nil
}

func rowsNoFetchResultSetToViewElbInputCidr(rows *sql.Rows) (*ViewElbInputCidr, error) {
	var err error
	var LoadBalancerId int64
	var LoadBalancerName string
	var CidrIp string

	err = rows.Scan(&LoadBalancerId, &LoadBalancerName, &CidrIp)
	if err != nil {
		return nil, err
	}
	return NewViewElbInputCidr(LoadBalancerId, LoadBalancerName, CidrIp), nil
}

func rowsResultSetToViewElbInputCidr(rows *sql.Rows) (*ViewElbInputCidr, error) {
	var err error
	if rows.Next() {
		var LoadBalancerId int64
		var LoadBalancerName string
		var CidrIp string

		err = rows.Scan(&LoadBalancerId, &LoadBalancerName, &CidrIp)
		if err != nil {
			return nil, err
		}
		return NewViewElbInputCidr(LoadBalancerId, LoadBalancerName, CidrIp), nil
	}
	return nil, err
}

func createViewElbInputCidr(db *sql.DB, LoadBalancerId int64, LoadBalancerName string, CidrIp string) *ViewElbInputCidr {
	rows := db.QueryRow("insert into view_elb_input_cidr(load_balancer_id,load_balancer_name,cidr_ip) values($1,$2,$3) returning load_balancer_id,load_balancer_name,cidr_ip",
		LoadBalancerId, LoadBalancerName, CidrIp)

	viewElbInputCidr, err := rowResultSetToViewElbInputCidr(rows)
	if err != nil {
		return nil
	}
	return viewElbInputCidr
}

func loadAllViewElbInputCidr(db *sql.DB) ([]*ViewElbInputCidr, error) {
	rows, err := db.Query("select load_balancer_id,load_balancer_name,cidr_ip from view_elb_input_cidr order by load_balancer_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfViewElbInputCidr := make([]*ViewElbInputCidr, 0, 0)
	for rows.Next() {
		viewElbInputCidr, err := rowsNoFetchResultSetToViewElbInputCidr(rows)
		if err == nil {
			arrayOfViewElbInputCidr = append(arrayOfViewElbInputCidr, viewElbInputCidr)
		} else {
			log.Println(err)
		}
	}

	return arrayOfViewElbInputCidr, nil
}

func loadAllViewElbInputCidrIps(db *sql.DB, loadBalanceName string) ([]string, error) {
	query := fmt.Sprintf("select cidr_ip from view_elb_input_cidr where load_balancer_name='%s'", loadBalanceName)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfCidrIps := make([]string, 0, 0)
	for rows.Next() {
		var cidrIp string
		err := rows.Scan(&cidrIp)
		if err == nil {
			arrayOfCidrIps = append(arrayOfCidrIps, cidrIp)
		} else {
			log.Println(err)
		}
	}

	return arrayOfCidrIps, nil
}

func loadAllViewElbInputCidrLoadBalancerName(db *sql.DB, cidrIp string) (string, error) {
	query := fmt.Sprintf("select distinct(load_balancer_name) from view_elb_by_cidr where subnet_cidr_block like '%s", cidrIp)
	rows, err := db.Query(query + "%'")
	if err != nil {
		return "", err
	}
	defer rows.Close()
	var name string
	if rows.Next() {
		err = rows.Scan(&name)
		if err != nil {
			return "", err
		}
	}

	return name, nil
}

func deleteAllViewElbInputCidr(db *sql.DB) error {
	_, err := db.Exec("delete from view_elb_input_cidr")

	return err
}
