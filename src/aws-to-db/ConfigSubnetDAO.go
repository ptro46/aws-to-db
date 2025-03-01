package main

import (
	"database/sql" // package SQL
	"log"

	_ "github.com/lib/pq" // driver Postgres
)

func rowResultSetToConfigSubnet(row *sql.Row) (*ConfigSubnet, error) {
	var err error
	var Id int64
	var Name string
	var Description string
	var Env string
	var CidrBlock string

	err = row.Scan(&Id, &Name, &Description, &Env, &CidrBlock)
	if err != nil {
		return nil, err
	}
	return NewConfigSubnet(Id, Name, Description, Env, CidrBlock), nil
}

func rowsNoFetchResultSetToConfigSubnet(rows *sql.Rows) (*ConfigSubnet, error) {
	var err error
	var Id int64
	var Name string
	var Description string
	var Env string
	var CidrBlock string

	err = rows.Scan(&Id, &Name, &Description, &Env, &CidrBlock)
	if err != nil {
		return nil, err
	}
	return NewConfigSubnet(Id, Name, Description, Env, CidrBlock), nil
}

func rowsResultSetToConfigSubnet(rows *sql.Rows) (*ConfigSubnet, error) {
	var err error
	if rows.Next() {
		var Id int64
		var Name string
		var Description string
		var Env string
		var CidrBlock string

		err = rows.Scan(&Id, &Name, &Description, &Env, &CidrBlock)
		if err != nil {
			return nil, err
		}
		return NewConfigSubnet(Id, Name, Description, Env, CidrBlock), nil
	}
	return nil, err
}

func loadConfigSubnetById(db *sql.DB, id int64) (*ConfigSubnet, error) {
	rows, err := db.Query("select id,name,description,env,cidr_block from config_subnet where id=$1", id)
	if err != nil {
		return nil, err
	}

	subnet, err := rowsResultSetToConfigSubnet(rows)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	return subnet, nil
}

func loadConfigSubnetByNameAndEnv(db *sql.DB, subnetName string, env string) (*ConfigSubnet, error) {
	rows, err := db.Query("select id,name,description,env,cidr_block from config_subnet where name=$1 and env=$2", subnetName, env)
	if err != nil {
		return nil, err
	}

	subnet, err := rowsResultSetToConfigSubnet(rows)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	return subnet, nil
}

func loadAllConfigSubnetByEnv(db *sql.DB, env string) ([]*ConfigSubnet, error) {
	rows, err := db.Query("select id,name,description,env,cidr_block from config_subnet where env=$1 order by id", env)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfSubnet := make([]*ConfigSubnet, 0, 0)
	for rows.Next() {
		vpc, err := rowsNoFetchResultSetToConfigSubnet(rows)
		if err == nil {
			arrayOfSubnet = append(arrayOfSubnet, vpc)
		} else {
			log.Println(err)
		}
	}

	return arrayOfSubnet, nil
}
