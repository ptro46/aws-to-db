package main

import (
	"database/sql" // package SQL
	"log"

	_ "github.com/lib/pq" // driver Postgres
)

func rowResultSetToVPC(row *sql.Row) (*VPC, error) {
	var err error
	var Id int64
	var VpcId string
	var VpcName string
	var CidrAssociationId string
	var CidrBlock string
	var CidrBlockState string
	var VpcState string
	var OwnerId string

	err = row.Scan(&Id, &VpcId, &VpcName, &CidrAssociationId, &CidrBlock, &CidrBlockState, &VpcState, &OwnerId)
	if err != nil {
		return nil, err
	}
	return NewVPC(Id, VpcId, VpcName, CidrAssociationId, CidrBlock, CidrBlockState, VpcState, OwnerId), nil
}

func rowsNoFetchResultSetToVPC(rows *sql.Rows) (*VPC, error) {
	var err error
	var Id int64
	var VpcId string
	var VpcName string
	var CidrAssociationId string
	var CidrBlock string
	var CidrBlockState string
	var VpcState string
	var OwnerId string

	err = rows.Scan(&Id, &VpcId, &VpcName, &CidrAssociationId, &CidrBlock, &CidrBlockState, &VpcState, &OwnerId)
	if err != nil {
		return nil, err
	}
	return NewVPC(Id, VpcId, VpcName, CidrAssociationId, CidrBlock, CidrBlockState, VpcState, OwnerId), nil
}

func rowsResultSetToVPC(rows *sql.Rows) (*VPC, error) {
	var err error
	if rows.Next() {
		var err error
		var Id int64
		var VpcId string
		var VpcName string
		var CidrAssociationId string
		var CidrBlock string
		var CidrBlockState string
		var VpcState string
		var OwnerId string

		err = rows.Scan(&Id, &VpcId, &VpcName, &CidrAssociationId, &CidrBlock, &CidrBlockState, &VpcState, &OwnerId)
		if err != nil {
			return nil, err
		}
		return NewVPC(Id, VpcId, VpcName, CidrAssociationId, CidrBlock, CidrBlockState, VpcState, OwnerId), nil
	}
	return nil, err
}

func createVPC(db *sql.DB, VpcId string, VpcName string, CidrAssociationId string, CidrBlock string, CidrBlockState string, VpcState string, OwnerId string) *VPC {
	rows := db.QueryRow("insert into vpc(vpc_id,vpc_name,cidr_association_id,cidr_block,cidr_block_state,vpc_state,owner_id) values($1,$2,$3,$4,$5,$6,$7) returning id,vpc_id,vpc_name,cidr_association_id,cidr_block,cidr_block_state,vpc_state,owner_id", VpcId, VpcName, CidrAssociationId, CidrBlock, CidrBlockState, VpcState, OwnerId)

	vpc, err := rowResultSetToVPC(rows)
	if err != nil {
		return nil
	}
	return vpc
}

func loadVPCByVPCId(db *sql.DB, vpc_id string) (*VPC, error) {
	rows, err := db.Query("select id,vpc_id,vpc_name,cidr_association_id,cidr_block,cidr_block_state,vpc_state,owner_id from vpc where vpc_id=$1", vpc_id)
	if err != nil {
		return nil, err
	}

	vpc, err := rowsResultSetToVPC(rows)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	return vpc, nil
}

func loadVPCById(db *sql.DB, id int64) (*VPC, error) {
	rows, err := db.Query("select id,vpc_id,vpc_name,cidr_association_id,cidr_block,cidr_block_state,vpc_state,owner_id from vpc where id=$1", id)
	if err != nil {
		return nil, err
	}

	vpc, err := rowsResultSetToVPC(rows)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	return vpc, nil
}

func loadVPCByName(db *sql.DB, vpcName string) (*VPC, error) {
	rows, err := db.Query("select id,vpc_id,vpc_name,cidr_association_id,cidr_block,cidr_block_state,vpc_state,owner_id from vpc where vpc_name=$1", vpcName)
	if err != nil {
		return nil, err
	}

	vpc, err := rowsResultSetToVPC(rows)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	return vpc, nil
}

func loadAllVPC(db *sql.DB) ([]*VPC, error) {
	rows, err := db.Query("select id,vpc_id,vpc_name,cidr_association_id,cidr_block,cidr_block_state,vpc_state,owner_id from vpc order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfVPC := make([]*VPC, 0, 0)
	for rows.Next() {
		vpc, err := rowsNoFetchResultSetToVPC(rows)
		if err == nil {
			arrayOfVPC = append(arrayOfVPC, vpc)
		} else {
			log.Println(err)
		}
	}

	return arrayOfVPC, nil
}

func loadAllByOrderUnique(db *sql.DB) ([]*VPC, error) {
	rows, err := db.Query("select id,vpc_id,vpc_name,cidr_association_id,cidr_block,cidr_block_state,vpc_state,owner_id from vpc order by vpc_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfVPC := make([]*VPC, 0, 0)
	for rows.Next() {
		vpc, err := rowsNoFetchResultSetToVPC(rows)
		if err == nil {
			arrayOfVPC = append(arrayOfVPC, vpc)
		} else {
			log.Println(err)
		}
	}

	return arrayOfVPC, nil
}

func deleteAllVPC(db *sql.DB) error {
	_, err := db.Exec("delete from vpc")

	return err
}
