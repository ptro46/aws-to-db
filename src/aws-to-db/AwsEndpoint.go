package main

type AwsEndpoint struct {
	Address      string `json:"Address,omitempty"`
	Port         int16  `json:"Port,omitempty"`
	HostedZoneId string `json:"HostedZoneId,omitempty"`
}
