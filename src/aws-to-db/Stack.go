package main

import (
	"database/sql"
	"fmt"
)

type Stack struct {
	Id                          int64
	StackId                     string
	StackName                   string
	ChangeSetId                 string
	Description                 string
	CreationTime                string
	DeletionTime                string
	LastUpdateTime              string
	StackStatus                 string
	StackStatusRaison           string
	DisableRollback             bool
	TimeoutInMinutes            int64
	RoleArn                     string
	Tags                        []*StackTag
	EnableTerminationProtection bool
	ParentId                    string
	RootId                      string
	NextToken                   string
}

func NewStack(id int64, stackId string, stackName string, changeSetId string, description string, creationTime string,
	deletionTime string, lastUpdateTime string, stackStatus string, stackStatusRaison string, disableRollback bool,
	timeoutInMinutes int64, roleArn string, enableTerminationProtection bool, parentId string, rootId string,
	nextToken string) *Stack {
	return &Stack{
		Id:                          id,
		StackId:                     stackId,
		StackName:                   stackName,
		ChangeSetId:                 changeSetId,
		Description:                 description,
		CreationTime:                creationTime,
		DeletionTime:                deletionTime,
		LastUpdateTime:              lastUpdateTime,
		StackStatus:                 stackStatus,
		StackStatusRaison:           stackStatusRaison,
		DisableRollback:             disableRollback,
		TimeoutInMinutes:            timeoutInMinutes,
		RoleArn:                     roleArn,
		EnableTerminationProtection: enableTerminationProtection,
		ParentId:                    parentId,
		RootId:                      rootId,
		NextToken:                   nextToken}
}

func (d Stack) String() string {
	return fmt.Sprintf("Stack stackId(%d) stackName(%s) Description(%s)", d.StackId, d.StackName, d.DeletionTime)
}

func (d *Stack) loadDependencies(db *sql.DB) {
	arrayOfStackTag, err := loadAllStackTagByOrderUnique(db)
	if err == nil {
		for _, stackTag := range arrayOfStackTag {
			if stackTag.StackId == d.Id {
				d.Tags = append(d.Tags, stackTag)
			}
		}
	} else {
		fmt.Println(err)
	}
}

func (d *Stack) Dump() string {
	dumpStack := fmt.Sprintf("StackId(%s) StackName(%s) ChangeSetId(%s) Description(%s) CreationTime(%s) DeletionTime(%s) LastUpdateTime(%s) StackStatus(%s) StackStatusRaison(%s) DisableRollback(%t) TimeoutInMinutes(%d) RoleArn(%s) EnableTerminationProtection(%t) ParentId(%s) RootId(%s) NextToken(%s)",
		d.StackId,
		d.StackName,
		d.ChangeSetId,
		d.Description,
		d.CreationTime,
		d.DeletionTime,
		d.LastUpdateTime,
		d.StackStatus,
		d.StackStatusRaison,
		d.DisableRollback,
		d.TimeoutInMinutes,
		d.RoleArn,
		d.EnableTerminationProtection,
		d.ParentId,
		d.RootId,
		d.NextToken)

	dumpStack += "\n	 Stack Tags ["
	for _, tag := range d.Tags {
		dumpStack += "{" + tag.Dump() + "{"
	}
	dumpStack += "]"
	return dumpStack
}
