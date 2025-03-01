package main

import "fmt"

type Bucket struct {
	Id           int64
	CreationDate string
	Name         string
}

func NewBucket(id int64, creationDate string, name string) *Bucket {
	return &Bucket{Id: id, CreationDate: creationDate, Name: name}
}

func (d *Bucket) string() string {
	return fmt.Sprintf("CreationDate(%s) Name(%s)", d.CreationDate, d.Name)
}

func (d *Bucket) Dump() string {
	return fmt.Sprintf("Bucket CreationDate(%s) Name(%s)", d.CreationDate, d.Name)
}
