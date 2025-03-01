package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type AwsSecurityGroup struct {
	Description         string            `json:"Description,omitempty"`
	GroupName           string            `json:"GroupName,omitempty"`
	IpPermissions       []AwsIpPermission `json:"IpPermissions,omitempty"`
	OwnerId             string            `json:"OwnerId,omitempty"`
	GroupId             string            `json:"GroupId,omitempty"`
	IpPermissionsEgress []AwsIpPermission `json:"IpPermissionsEgress,omitempty"`
	Tags                []AwsTag          `json:"Tags,omitempty"`
	VpcId               string            `json:"VpcId,omitempty"`
}

type AwsSecurityGroups struct {
	SecurityGroups []AwsSecurityGroup `json:"SecurityGroups,omitempty"`
}

func awsSecurityGroupParseAndPersist(db *sql.DB, jsonString string) error {
	fmt.Printf("awsSecurityGroupParseAndPersist with json lenght %d\n", len(jsonString))
	awsSecGroups := new(AwsSecurityGroups)
	errUnmarshal := json.Unmarshal([]byte(jsonString), awsSecGroups)
	if errUnmarshal == nil {
		for _, awsSecGroup := range awsSecGroups.SecurityGroups {
			fmt.Printf("GroupName(%s) OwnerId(%s)\n", awsSecGroup.GroupName, awsSecGroup.OwnerId)
			secGroup := createSecurityGroup(db, awsSecGroup.Description, awsSecGroup.GroupName, awsSecGroup.VpcId, awsSecGroup.OwnerId, awsSecGroup.GroupId)
			if secGroup != nil {
				fmt.Printf("    SecurityGroup %s::%s loaded\n", secGroup.GroupName, secGroup.OwnerId)
				createRestSecurityGroupOfContent(db, secGroup, awsSecGroup)
			} else {
				fmt.Printf("  ERROR  SecurityGroup %s::%s not loaded\n", awsSecGroup.GroupName, awsSecGroup.OwnerId)
			}
		}
	} else {
		return errUnmarshal
	}
	return nil
}

func createRestSecurityGroupOfContent(db *sql.DB, secGroup *SecurityGroup, awsSecGroup AwsSecurityGroup) {
	for _, secGroupTag := range awsSecGroup.Tags {
		tag := createSecurityGroupTag(db, secGroup.Id, secGroupTag.Key, secGroupTag.Value)
		if tag != nil {
			fmt.Printf("		 Tag Key(%s) Loaded\n", tag.Key)
		} else {
			fmt.Printf("		 ERROR Tag Key(%s) Not Loaded\n", secGroupTag.Key)
		}
	}

	for _, secGroupIpPermission := range awsSecGroup.IpPermissions {
		ipPermission := createSecurityGroupIpPermission(db, secGroup.Id, secGroupIpPermission.FromPort, secGroupIpPermission.IpProtocol, secGroupIpPermission.ToPort)
		if ipPermission != nil {
			fmt.Printf("		 IpPermission FromPort(%d) Loaded\n", ipPermission.FromPort)
			for _, awsIpRange := range secGroupIpPermission.IpRanges {
				ipRange := createSecurityGroupIpPermissionIpRange(db,
					ipPermission.Id,
					awsIpRange.CidrIp,
					awsIpRange.Description)
				if ipRange != nil {
					fmt.Printf("		 	IpRange CidrIp(%s) Loaded\n", ipRange.CidrIp)
				} else {
					fmt.Printf("		 	ERROR IpRange CidrIp(%s) Not Loaded\n", awsIpRange.CidrIp)
				}
			}

			for _, awsPrefix := range secGroupIpPermission.PrefixListIds {
				prefix := createSecurityGroupIpPermissionPrefixListId(db,
					ipPermission.Id,
					awsPrefix.Description,
					awsPrefix.PrefixListId)
				if prefix != nil {
					fmt.Printf("		 	PrefixListId description(%s) Loaded\n", prefix.Description)
				} else {
					fmt.Printf("		 	ERROR PrefixListId description(%s) Not Loaded\n", awsPrefix.Description)
				}
			}

			for _, awsUser := range secGroupIpPermission.UserIdGroupPairs {
				user := createSecurityGroupIpPermissionUserIdGroupPair(db,
					ipPermission.Id,
					awsUser.Description,
					awsUser.GroupId,
					awsUser.GroupName,
					awsUser.PeeringStatus,
					awsUser.UserId,
					awsUser.VpcId,
					awsUser.VpcPeeringConnectionId)
				if user != nil {
					fmt.Printf("		 	UserIdGroupPair GroupName(%s) Loaded\n", user.GroupName)
				} else {
					fmt.Printf("		 	ERROR UserIdGroupPair GroupName(%s) Not Loaded\n", awsUser.GroupName)
				}
			}

		} else {
			fmt.Printf("		 ERROR IpPermission FromPort(%s) Not Loaded\n", ipPermission.FromPort)
		}

	}

	for _, secGroupIpPermissionEg := range awsSecGroup.IpPermissionsEgress {
		ipPermissionEg := createSecurityGroupIpPermissionEgress(db, secGroup.Id, secGroupIpPermissionEg.FromPort, secGroupIpPermissionEg.IpProtocol, secGroupIpPermissionEg.ToPort)
		if ipPermissionEg != nil {
			fmt.Printf("		 IpPermissionEgress FromPort(%s) Loaded\n", ipPermissionEg.FromPort)
			for _, awsEgIpRange := range secGroupIpPermissionEg.IpRanges {
				egIpRange := createSecurityGroupIpPermissionEgressIpRange(db,
					ipPermissionEg.Id,
					awsEgIpRange.CidrIp,
					awsEgIpRange.Description)
				if egIpRange != nil {
					fmt.Printf("		 	EgressIpRange CidrIp(%s) Loaded\n", egIpRange.CidrIp)
				} else {
					fmt.Printf("		 	ERROR EgressIpRange CidrIp(%s) Not Loaded\n", awsEgIpRange.CidrIp)
				}
			}

			for _, awsEgPrefix := range secGroupIpPermissionEg.PrefixListIds {
				prefixEg := createSecurityGroupIpPermissionEgressPrefixListId(db,
					ipPermissionEg.Id,
					awsEgPrefix.Description,
					awsEgPrefix.PrefixListId)
				if prefixEg != nil {
					fmt.Printf("		 	EgressPrefixListId description(%s) Loaded\n", prefixEg.Description)
				} else {
					fmt.Printf("		 	ERROR EgressPrefixListId description(%s) Not Loaded\n", awsEgPrefix.Description)
				}
			}

			for _, awsEgUser := range secGroupIpPermissionEg.UserIdGroupPairs {
				userEg := createSecurityGroupIpPermissionEgressUserIdGroupPair(db,
					ipPermissionEg.Id,
					awsEgUser.Description,
					awsEgUser.GroupId,
					awsEgUser.GroupName,
					awsEgUser.PeeringStatus,
					awsEgUser.UserId,
					awsEgUser.VpcId,
					awsEgUser.VpcPeeringConnectionId)
				if userEg != nil {
					fmt.Printf("		 	EgressUserIdGroupPair GroupName(%s) Loaded\n", userEg.GroupName)
				} else {
					fmt.Printf("		 	ERROR EgressUserIdGroupPair GroupName(%s) Not Loaded\n", awsEgUser.GroupName)
				}
			}

		} else {
			fmt.Printf("		 ERROR IpPermissionEgress FromPort(%s) Not Loaded\n", ipPermissionEg.FromPort)
		}
	}

}
