package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type AwsBucket struct {
	CreationDate string `json:"CreationDate,omitempty"`
	Name         string `json:"Name,omitempty"`
}

type AwsBuckets struct {
	Buckets []AwsBucket `json:"Buckets,omitempty"`
}

func awsBucketParseAndPersist(db *sql.DB, jsonString string) error {
	fmt.Printf("awsBucketParseAndPersist with json lenght %d\n", len(jsonString))
	awsBuckets := new(AwsBuckets)
	errUnmarshal := json.Unmarshal([]byte(jsonString), awsBuckets)
	if errUnmarshal == nil {
		for _, awsBucket := range awsBuckets.Buckets {
			fmt.Printf("CreationDate(%s) Name(%s)\n", awsBucket.CreationDate, awsBucket.Name)
			bucket := createBucket(db, awsBucket.CreationDate, awsBucket.Name)
			if bucket != nil {
				fmt.Printf("    Bucket %s::%s loaded\n", bucket.CreationDate, bucket.Name)
			} else {
				fmt.Printf("  ERROR  Bucket %s::%s not loaded\n", bucket.CreationDate, bucket.Name)
			}
		}
	} else {
		return errUnmarshal
	}
	return nil
}
