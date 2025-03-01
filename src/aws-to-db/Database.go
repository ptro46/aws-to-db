package main

import (
	"database/sql" // package SQL

	_ "github.com/lib/pq" // driver Postgres
)

func connectDB(dbConfig *DbConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbConfig.dbPostgresConnectString("aws-to-db-identifier"))
	return db, err
}
