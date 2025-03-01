package main

import (
	"fmt"
)

type AwsCidrBlockAssociationSet struct {
	AssociationId  string            `json:"AssociationId,omitempty"`
	CidrBlock      string            `json:"CidrBlock,omitempty"`
	CidrBlockState AwsCidrBlockState `json:"CidrBlockState,omitempty"`
}

func (d AwsCidrBlockAssociationSet) String() string {
	return fmt.Sprintf("AwsCidrBlockAssociationSet AssociationId(%s) CidrBlock(%s)", d.AssociationId, d.CidrBlock)
}
