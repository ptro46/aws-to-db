package main

type AwsStatusInfo struct {
	StatusType string `json:"StatusType,omitempty"`
	Normal     bool   `json:"Normal,omitempty"`
	Status     string `json:"Status,omitempty"`
	Message    string `json:"Message,omitempty"`
}
