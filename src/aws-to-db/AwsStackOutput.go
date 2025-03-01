package main

import "fmt"

type AwsStackOutput struct {
	OutputKey   string `json:"OutputKey,omitempty"`
	OutputValue string `json:"OutputValue,omitempty"`
	Description string `json:"Description,omitempty"`
	ExportName  string `json:"ExportName,omitempty"`
}

func (d AwsStackOutput) String() string {
	return fmt.Sprintf("AwsStackOutput OutputKey(%s) OutputValue(%s)", d.OutputKey, d.OutputValue)
}
