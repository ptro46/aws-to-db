package main

import (
	"database/sql" // package SQL
	"fmt"

	_ "github.com/lib/pq" // driver Postgres
)

func rowResultSetToLoadBalancerListenerCertificat(row *sql.Row) (*LoadBalancerListenerCertificat, error) {
	var err error
	var Id int64
	var LoadBalancerListenerId int64
	var CertificatArn string

	err = row.Scan(&Id, &LoadBalancerListenerId, &CertificatArn)
	if err != nil {
		return nil, err
	}
	return newLoadBalancerListenerCertificat(Id, LoadBalancerListenerId, CertificatArn), nil
}

func rowsNoFetchResultSetToLoadBalancerListenerCertificat(rows *sql.Rows) (*LoadBalancerListenerCertificat, error) {
	var err error
	var Id int64
	var LoadBalancerListenerId int64
	var CertificatArn string

	err = rows.Scan(&Id, &LoadBalancerListenerId, &CertificatArn)
	if err != nil {
		return nil, err
	}
	return newLoadBalancerListenerCertificat(Id, LoadBalancerListenerId, CertificatArn), nil
}

func rowsResultSetToLoadBalancerListenerCertificat(rows *sql.Rows) (*LoadBalancerListenerCertificat, error) {
	var err error
	if rows.Next() {
		var Id int64
		var LoadBalancerListenerId int64
		var CertificatArn string

		err = rows.Scan(&Id, &LoadBalancerListenerId, CertificatArn)
		if err != nil {
			return nil, err
		}
		return newLoadBalancerListenerCertificat(Id, LoadBalancerListenerId, CertificatArn), nil
	}
	return nil, err
}

func createLoadBalancerListenerCertificat(db *sql.DB, LoadBalancerListenerId int64, CertificatArn string) *LoadBalancerListenerCertificat {
	rows := db.QueryRow("insert into load_balancer_listener_certificats(load_balancer_listener_id,certificat_arn) values($1,$2) returning id,load_balancer_listener_id,certificat_arn",
		LoadBalancerListenerId, CertificatArn)

	load, err := rowResultSetToLoadBalancerListenerCertificat(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return load
}

func deleteAllLoadBalancerListenerCertificat(db *sql.DB) error {
	_, err := db.Exec("delete from load_balancer_listener_certificats")

	return err
}
