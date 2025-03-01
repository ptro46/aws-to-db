package main

type GraphConfig struct {
	Name       string                  `json:"name,omitempty"`
	Properties []GraphConfigProperties `json:"properties,omitempty"`
	Rank       int                     `json:"rank,omitempty"`
}
