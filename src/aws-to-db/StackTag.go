package main

import "fmt"

type StackTag struct {
	Id      int64
	StackId int64
	Key     string
	Value   string
}

func NewStackTag(id int64, stackId int64, key string, value string) *StackTag {
	return &StackTag{
		Id:      id,
		StackId: stackId,
		Key:     key,
		Value:   value}
}

func (d *StackTag) String() string {
	return fmt.Sprintf("StackTag Id(%d) Key(%s) Value(%s)", d.Id, d.Key, d.Value)
}

func (d *StackTag) Dump() string {
	return fmt.Sprintf("Key(%s) Value(%s)", d.Key, d.Value)
}
