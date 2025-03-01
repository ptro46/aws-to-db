package main

import "fmt"

type SecurityGroupTag struct {
	Id              int64
	SecurityGroupId int64
	Key             string
	Value           string
}

func NewSecurityGroupTag(id int64, securityGroupId int64, key string, value string) *SecurityGroupTag {
	return &SecurityGroupTag{
		Id:              id,
		SecurityGroupId: securityGroupId,
		Key:             key,
		Value:           value}
}

func (d *SecurityGroupTag) String() string {
	return fmt.Sprintf("SecurityGroupTag Key(%s) Value(%s)", d.Key, d.Value)
}

func (d *SecurityGroupTag) Dump() string {
	return fmt.Sprintf("Key(%s) Value(%s)", d.Key, d.Value)
}
