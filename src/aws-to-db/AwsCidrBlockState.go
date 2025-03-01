package main

import (
	"fmt"
)

type AwsCidrBlockState struct {
	State string `json:"State,omitempty"`
}

func (d AwsCidrBlockState) String() string {
	return fmt.Sprintf("AwsCidrBlockState State(%s)", d.State)
}
