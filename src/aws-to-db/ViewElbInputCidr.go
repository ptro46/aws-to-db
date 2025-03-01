package main

type ViewElbInputCidr struct {
	LoadBalancerId   int64
	LoadBalancerName string
	CidrIp           string
}

func NewViewElbInputCidr(LoadBalancerId int64, LoadBalancerName string, CidrIp string) *ViewElbInputCidr {
	return &ViewElbInputCidr{LoadBalancerId,
		LoadBalancerName,
		CidrIp}
}
