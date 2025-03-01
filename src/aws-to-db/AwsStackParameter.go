package main

import "fmt"

type AwsStackParameter struct {
	ParameterKey     string `json:"ParameterKey,omitempty"`
	ParameterValue   string `json:"ParameterValue,omitempty"`
	UsePreviousValue bool   `json:"UsePreviousValue,omitempty"`
	ResolvedValue    string `json:"ResolvedValue,omitempty"`
}

func (d AwsStackParameter) String() string {
	return fmt.Sprintf("AwsStackParameter ParameterKey(%s) ParamterValue(%s)", d.ParameterKey, d.ParameterValue)
}
