package main

import (
	"database/sql" // package SQL
	"log"

	_ "github.com/lib/pq" // driver Postgres
)

func rowResultSetToBucket(row *sql.Row) (*Bucket, error) {
	var err error
	var Id int64
	var CreationDate string
	var Name string

	err = row.Scan(&Id, &CreationDate, &Name)
	if err != nil {
		return nil, err
	}
	return NewBucket(Id, CreationDate, Name), nil
}

func rowsNoFetchResultSetToBucket(rows *sql.Rows) (*Bucket, error) {
	var err error
	var Id int64
	var CreationDate string
	var Name string

	err = rows.Scan(&Id, &CreationDate, &Name)
	if err != nil {
		return nil, err
	}
	return NewBucket(Id, CreationDate, Name), nil
}

func rowsResultSetToBucket(rows *sql.Rows) (*Bucket, error) {
	var err error
	if rows.Next() {
		var Id int64
		var CreationDate string
		var Name string

		err = rows.Scan(&Id, &CreationDate, &Name)
		if err != nil {
			return nil, err
		}
		return NewBucket(Id, CreationDate, Name), nil
	}
	return nil, err
}

func createBucket(db *sql.DB, CreationDate string, Name string) *Bucket {
	rows := db.QueryRow("insert into bucket(creation_date,name) values($1,$2) returning id,creation_date,name", CreationDate, Name)

	bucket, err := rowResultSetToBucket(rows)
	if err != nil {
		return nil
	}
	return bucket
}

func loadAllBucket(db *sql.DB) ([]*Bucket, error) {
	rows, err := db.Query("select id,creation_date,name from bucket order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfBucket := make([]*Bucket, 0, 0)
	for rows.Next() {
		bucket, err := rowsNoFetchResultSetToBucket(rows)
		if err == nil {
			arrayOfBucket = append(arrayOfBucket, bucket)
		} else {
			log.Println(err)
		}
	}

	return arrayOfBucket, nil
}

func loadAllBucketByOrderUnique(db *sql.DB) ([]*Bucket, error) {
	rows, err := db.Query("select id,creation_date,name from bucket order by name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfBucket := make([]*Bucket, 0, 0)
	for rows.Next() {
		bucket, err := rowsNoFetchResultSetToBucket(rows)
		if err == nil {
			arrayOfBucket = append(arrayOfBucket, bucket)
		} else {
			log.Println(err)
		}
	}

	return arrayOfBucket, nil
}

func deleteAllBucket(db *sql.DB) error {
	_, err := db.Exec("delete from bucket")

	return err
}
