package main

type AwsAvailabilityZone struct {
	ZoneName              string                   `json:"ZoneName,omitempty"`
	SubnetId              string                   `json:"SubnetId,omitempty"`
	LoadBalancerAddresses []AwsLoadBalancerAddress `json:"LoadBalancerAddresses,omitempty"`
}
