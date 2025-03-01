package main

import (
	"fmt"
)

type AwsTag struct {
	Key   string `json:"Key,omitempty"`
	Value string `json:"Value,omitempty"`
}

func (d AwsTag) String() string {
	return fmt.Sprintf("AwsTag Key(%s) Value(%s)", d.Key, d.Value)
}
