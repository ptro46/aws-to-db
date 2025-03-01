package main

import "fmt"

type StackParameter struct {
	Id               int64
	StackId          int64
	ParameterKey     string
	ParameterValue   string
	UsePreviousValue bool
	ResolvedValue    string
}

func NewStackParameter(id int64, stackId int64, parameterKey string, parameterValue string, usePreviousValue bool, resolvedValue string) *StackParameter {
	return &StackParameter{
		Id:               id,
		StackId:          stackId,
		ParameterKey:     parameterKey,
		ParameterValue:   parameterValue,
		UsePreviousValue: usePreviousValue,
		ResolvedValue:    resolvedValue}

}

func (d *StackParameter) String() string {
	return fmt.Sprintf("StackParameter Id(%d) Key Value(%s)", d.Id, d.ParameterKey, d.ParameterValue)
}
