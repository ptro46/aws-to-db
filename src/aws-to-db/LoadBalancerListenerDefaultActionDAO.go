package main

import (
	"database/sql" // package SQL
	"fmt"

	_ "github.com/lib/pq" // driver Postgres
)

func rowResultSetToLoadBalancerListenerDefaultAction(row *sql.Row) (*LoadBalancerDefaultAction, error) {
	var err error
	var Id int64
	var LoadBalancerListenerId int64
	var TargetGroupArn string
	var TargetGroupId sql.NullInt64
	var Type string

	err = row.Scan(&Id, &LoadBalancerListenerId, &TargetGroupArn, &TargetGroupId, &Type)
	if err != nil {
		return nil, err
	}
	return newLoadBalancerDefaultAction(Id, LoadBalancerListenerId, TargetGroupArn, TargetGroupId, Type), nil
}

func rowsNoFetchResultSetToLoadBalancerListenerDefaultAction(rows *sql.Rows) (*LoadBalancerDefaultAction, error) {
	var err error
	var Id int64
	var LoadBalancerListenerId int64
	var TargetGroupArn string
	var TargetGroupId sql.NullInt64
	var Type string

	err = rows.Scan(&Id, &LoadBalancerListenerId, &TargetGroupArn, &TargetGroupId, &Type)
	if err != nil {
		return nil, err
	}
	return newLoadBalancerDefaultAction(Id, LoadBalancerListenerId, TargetGroupArn, TargetGroupId, Type), nil
}

func rowsResultSetToLoadBalancerListenerDefaultAction(rows *sql.Rows) (*LoadBalancerDefaultAction, error) {
	var err error
	if rows.Next() {
		var Id int64
		var LoadBalancerListenerId int64
		var TargetGroupArn string
		var TargetGroupId sql.NullInt64
		var Type string

		err = rows.Scan(&Id, &LoadBalancerListenerId, &TargetGroupArn, &TargetGroupId, &Type)
		if err != nil {
			return nil, err
		}
		return newLoadBalancerDefaultAction(Id, LoadBalancerListenerId, TargetGroupArn, TargetGroupId, Type), nil
	}
	return nil, err
}

func createLoadBalancerListenerDefaultAction(db *sql.DB, LoadBalancerListenerId int64, TargetGroupArn string, Type string) *LoadBalancerDefaultAction {
	rows := db.QueryRow("insert into load_balancer_listener_default_actions(load_balancer_listener_id,target_group_arn,type) values($1,$2,$3) returning id,load_balancer_listener_id,target_group_arn,target_group_id,type",
		LoadBalancerListenerId, TargetGroupArn, Type)

	load, err := rowResultSetToLoadBalancerListenerDefaultAction(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return load
}

func loadAllLoadBalancerListenerDefaultActions(db *sql.DB) ([]*LoadBalancerDefaultAction, error) {
	results := make([]*LoadBalancerDefaultAction, 0, 0)
	sqlSelect := "SELECT id,load_balancer_listener_id,target_group_arn,target_group_id,type from load_balancer_listener_default_actions order by id"

	rows, err := db.Query(sqlSelect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		loadBalancerDefaultAction, err := rowsNoFetchResultSetToLoadBalancerListenerDefaultAction(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, loadBalancerDefaultAction)
	}
	return results, nil
}

func updateLoadBalancerListenerDefaultActionsTargetGroupId(db *sql.DB, loadBalancerDefaultAction *LoadBalancerDefaultAction, targetGroupId int64) *LoadBalancerDefaultAction {
	rows := db.QueryRow("update load_balancer_listener_default_actions set target_group_id=$1 where id=$2 returning id,load_balancer_listener_id,target_group_arn,target_group_id,type",
		targetGroupId, loadBalancerDefaultAction.Id)

	load, err := rowResultSetToLoadBalancerListenerDefaultAction(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return load
}

func deleteAllLoadBalancerListenerDefaultAction(db *sql.DB) error {
	_, err := db.Exec("delete from load_balancer_listener_default_actions")

	return err
}
