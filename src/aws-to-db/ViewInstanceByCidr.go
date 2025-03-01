package main

type ViewInstanceByCidr struct {
	InstanceId   int64
	InstanceKey  string
	InstanceName string
	CidrBlock    string
}

func NewViewInstanceByCidr(InstanceId int64, InstanceKey string, InstanceName string, CidrBlock string) *ViewInstanceByCidr {
	return &ViewInstanceByCidr{InstanceId,
		InstanceKey,
		InstanceName,
		CidrBlock}
}
