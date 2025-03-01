package main

import (
	"database/sql" // package SQL
	"fmt"

	_ "github.com/lib/pq" // driver Postgres
)

func rowResultSetToLoadBalancerListener(row *sql.Row) (*LoadBalancerListener, error) {
	var err error
	var Id int64
	var LoadBalancerId int64
	var LoadBalancerArn string
	var ListenerArn string
	var SslPolicies string
	var Port int64
	var Protocol string

	err = row.Scan(&Id, &LoadBalancerId, &LoadBalancerArn, &ListenerArn, &SslPolicies, &Port, &Protocol)
	if err != nil {
		return nil, err
	}
	return newLoadBalancerListener(Id, LoadBalancerId, LoadBalancerArn, ListenerArn, SslPolicies, Port, Protocol), nil
}

func rowsNoFetchResultSetToLoadBalancerListener(rows *sql.Rows) (*LoadBalancerListener, error) {
	var err error
	var Id int64
	var LoadBalancerId int64
	var LoadBalancerArn string
	var ListenerArn string
	var SslPolicies string
	var Port int64
	var Protocol string

	err = rows.Scan(&Id, &LoadBalancerId, &LoadBalancerArn, &ListenerArn, &SslPolicies, &Port, &Protocol)
	if err != nil {
		return nil, err
	}
	return newLoadBalancerListener(Id, LoadBalancerId, LoadBalancerArn, ListenerArn, SslPolicies, Port, Protocol), nil
}

func rowsResultSetToLoadBalancerListener(rows *sql.Rows) (*LoadBalancerListener, error) {
	var err error
	if rows.Next() {
		var Id int64
		var LoadBalancerId int64
		var LoadBalancerArn string
		var ListenerArn string
		var SslPolicies string
		var Port int64
		var Protocol string

		err = rows.Scan(&Id, &LoadBalancerId, &LoadBalancerArn, &ListenerArn, &SslPolicies, &Port, &Protocol)
		if err != nil {
			return nil, err
		}
		return newLoadBalancerListener(Id, LoadBalancerId, LoadBalancerArn, ListenerArn, SslPolicies, Port, Protocol), nil
	}
	return nil, err
}

func createLoadBalancerListener(db *sql.DB, LoadBalancerId int64, LoadBalancerArn string, ListenerArn string, SslPolicies string, Port int64, Protocol string) *LoadBalancerListener {
	rows := db.QueryRow("insert into load_balancer_listener(load_balancer_id,load_balancer_arn,listener_arn,ssl_policies,port,protocol) values($1,$2,$3,$4,$5,$6) returning id,load_balancer_id,load_balancer_arn,listener_arn,ssl_policies,port,protocol",
		LoadBalancerId, LoadBalancerArn, ListenerArn, SslPolicies, Port, Protocol)

	load, err := rowResultSetToLoadBalancerListener(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return load
}

func deleteAllLoadBalancerListener(db *sql.DB) error {
	_, err := db.Exec("delete from load_balancer_listener")

	return err
}
