package main

import "fmt"

type StackOutput struct {
	Id          int64
	StackId     int64
	OutputKey   string
	OutputValue string
	Description string
	ExportName  string
}

func NewStackOutput(id int64, stackId int64, outputKey string, outputValue string, description string, exportName string) *StackOutput {
	return &StackOutput{
		Id:          id,
		StackId:     stackId,
		OutputKey:   outputKey,
		OutputValue: outputValue,
		Description: description,
		ExportName:  exportName}
}

func (d *StackOutput) String() string {
	return fmt.Sprintf("StackOutput Id(%d) Description(%s)", d.Id, d.Description)
}
