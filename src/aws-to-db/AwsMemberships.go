package main

type AwsOptionGroupMemberships struct {
	OptionGroupName string `json:"OptionGroupName,omitempty"`
	Status          string `json:"Status,omitempty"`
}

type AwsDomainMemberships struct {
	Domain      string `json:"Domain,omitempty"`
	Status      string `json:"Status,omitempty"`
	FQDN        string `json:"FQDN,omitempty"`
	IAMRoleName string `json:"IAMRoleName,omitempty"`
}
