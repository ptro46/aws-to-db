package main

type ViewELBByCidr struct {
	SubnetId         int64
	SubnetCidrBlock  string
	LoadBalancerId   int64
	LoadBalancerName string
}

func NewViewELBByCidr(SubnetId int64, SubnetCidrBlock string, LoadBalancerId int64, LoadBalancerName string) *ViewELBByCidr {
	return &ViewELBByCidr{SubnetId: SubnetId,
		SubnetCidrBlock:  SubnetCidrBlock,
		LoadBalancerId:   LoadBalancerId,
		LoadBalancerName: LoadBalancerName}
}
