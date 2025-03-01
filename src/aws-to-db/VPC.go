package main

import (
	"fmt"
)

type VPC struct {
	Id                int64
	VpcId             string
	VpcName           string
	CidrAssociationId string
	CidrBlock         string
	CidrBlockState    string
	VpcState          string
	OwnerId           string
}

func NewVPC(id int64, vpcId string, vpcName string, cidrAssociationId string, cidrBlock string, cidrBlockState string, vpcState string, ownerId string) *VPC {
	return &VPC{
		Id:                id,
		VpcId:             vpcId,
		VpcName:           vpcName,
		CidrAssociationId: cidrAssociationId,
		CidrBlock:         cidrBlock,
		CidrBlockState:    cidrBlockState,
		VpcState:          vpcState,
		OwnerId:           ownerId}
}

func (d *VPC) String() string {
	return fmt.Sprintf("VPC id(%d) vpcId(%s) vpcName(%s) cidrBlock(%s)",
		d.Id,
		d.VpcId,
		d.VpcName,
		d.CidrBlock)
}

func (d *VPC) Dump() string {
	return fmt.Sprintf("VPC vpcId(%s) vpcName(%s) cidrAssociationId(%s) CidrBlock(%s) CidrBlockState(%s) VpcState(%s) OwnerId(%s)",
		d.VpcId, d.VpcName, d.CidrAssociationId, d.CidrBlock, d.CidrBlockState, d.VpcState, d.OwnerId)
}
