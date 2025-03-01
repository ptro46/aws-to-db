package main

import "fmt"

type AwsPrefixListId struct {
	Description  string `json:"Description,omitempty"`
	PrefixListId string `json:"PrefixListId,omitempty"`
}

func (d *AwsPrefixListId) String() string {
	return fmt.Sprintf("AwsPrefixListId description(%s) prefixListId(%s)", d.Description, d.PrefixListId)
}
