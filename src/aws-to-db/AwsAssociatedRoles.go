package main

type AwsAssociatedRoles struct {
	RoleArn     string `json:"RoleArn,omitempty"`
	FeatureName string `json:"FeatureName,omitempty"`
	Status      string `json:"Status,omitempty"`
}
