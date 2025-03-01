package main

import "fmt"

type ConfigSubnet struct {
	Id          int64
	Name        string
	Description string
	Env         string
	CidrBlock   string
}

func NewConfigSubnet(id int64, name string, description string, env string, cidrBlock string) *ConfigSubnet {
	return &ConfigSubnet{
		Id:          id,
		Name:        name,
		Description: description,
		Env:         env,
		CidrBlock:   cidrBlock}
}

func (d *ConfigSubnet) String() string {
	return fmt.Sprintf("ConfigSubnet Id(%d) name(%s) description(%s) env(%s) CidrBlock(%s)\n",
		d.Id, d.Name, d.Description, d.Env, d.CidrBlock)
}
