package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type AwsStack struct {
	StackId                     string              `json:"StackId,omitempty"`
	StackName                   string              `json:"StackName,omitempty"`
	ChangeSetId                 string              `json:"ChangeSetId,omitempty"`
	Description                 string              `json:"Description,omitempty"`
	Parameters                  []AwsStackParameter `json:"Parameters,omitempty"`
	CreationTime                string              `json:"CreationTime,omitempty"`
	DeletionTime                string              `json:"DeletionTime,omitempty"`
	LastUpdatedTime             string              `json:"LastUpdatedTime,omitempty"`
	StackStatus                 string              `json:"StackStatus,omitempty"`
	StackStatusReason           string              `json:"StackStatusReason,omitempty"`
	DisableRollback             bool                `json:"DisableRollback,omitempty"`
	TimeoutInMinutes            int64               `json:"TimeoutInMinutes,omitempty"`
	Outputs                     []AwsStackOutput    `json:"Outputs,omitempty"`
	RoleARN                     string              `json:"RoleARN,omitempty"`
	Tags                        []AwsTag            `json:"Tags,omitempty"`
	EnableTerminationProtection bool                `json:"EnableTerminationProtection,omitempty"`
	ParentId                    string              `json:"ParentId,omitempty"`
	RootId                      string              `json:"RootId,omitempty"`
	NextToken                   string              `json:"NextToken,omitempty"`
}

type AwsStacks struct {
	Stacks []AwsStack `json:"Stacks,omitempty"`
}

func awsStackParseAndPersist(db *sql.DB, jsonString string) error {
	fmt.Printf("awsStackParseAndPersist with json lenght %d\n", len(jsonString))
	awsStacks := new(AwsStacks)
	errUnmarshal := json.Unmarshal([]byte(jsonString), awsStacks)
	if errUnmarshal == nil {
		fmt.Println(len(awsStacks.Stacks))
		for _, awsStack := range awsStacks.Stacks {
			fmt.Printf("StackName(%s)\n", awsStack.StackName)
			stack := createStack(db,
				awsStack.StackId,
				awsStack.StackName,
				awsStack.ChangeSetId,
				awsStack.Description,
				awsStack.CreationTime,
				awsStack.DeletionTime,
				awsStack.LastUpdatedTime,
				awsStack.StackStatus,
				awsStack.StackStatusReason,
				awsStack.DisableRollback,
				awsStack.TimeoutInMinutes,
				awsStack.RoleARN,
				awsStack.EnableTerminationProtection,
				awsStack.ParentId,
				awsStack.RootId,
				awsStack.NextToken)
			if stack != nil {
				fmt.Printf("    Stack %s loaded\n", stack.StackName)

				for _, stackTag := range awsStack.Tags {
					tag := createStackTag(db, stack.Id, stackTag.Key, stackTag.Value)

					if tag != nil {
						fmt.Printf(" 		Tag %s has loaded\n", tag.Key)
					} else {
						fmt.Printf(" 		ERROR Tag  %s  has not loaded\n", stackTag.Key)
					}
				}

				for _, stackParameter := range awsStack.Parameters {
					paramter := createStackParameter(db, stack.Id, stackParameter.ParameterKey, stackParameter.ParameterValue, stackParameter.UsePreviousValue, stackParameter.ResolvedValue)

					if paramter != nil {
						fmt.Printf(" 	  	 Parameter %s has loaded\n", stackParameter.ParameterKey)
					} else {
						fmt.Printf(" 		ERROR   Parameter %s has not loaded\n", stackParameter.ParameterKey)
					}
				}

				for _, stackOutput := range awsStack.Outputs {
					output := createStackOutput(db, stack.Id, stackOutput.OutputKey, stackOutput.OutputValue, stackOutput.Description, stackOutput.ExportName)

					if output != nil {
						fmt.Printf(" 	  	 Output %s has loaded\n", stackOutput.OutputKey)
					} else {
						fmt.Printf(" 		ERROR   Output %s has not loaded\n", stackOutput.OutputKey)
					}
				}

			} else {
				fmt.Printf("  ERROR  Stack %s::%s not loaded\n", awsStack.Description, awsStack.StackName)
			}
		}
	} else {
		return errUnmarshal
	}
	return nil
}
