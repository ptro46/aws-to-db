package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type AwsDBInstance struct {
	DBInstanceIdentifier                  string                      `json:"DBInstanceIdentifier,omitempty"`
	DBInstanceClass                       string                      `json:"DBInstanceClass,omitempty"`
	Engine                                string                      `json:"Engine,omitempty"`
	DBInstanceStatus                      string                      `json:"DBInstanceStatus,omitempty"`
	MasterUsername                        string                      `json:"MasterUsername,omitempty"`
	DBName                                string                      `json:"DBName,omitempty"`
	Endpoint                              AwsEndpoint                 `json:"Endpoint,omitempty"`
	AllocatedStorage                      int16                       `json:"AllocatedStorage,omitempty"`
	InstanceCreateTime                    string                      `json:"InstanceCreateTime,omitempty"`
	PreferredBackupWindow                 string                      `json:"PreferredBackupWindow,omitempty"`
	BackupRetentionPeriod                 int16                       `json:"BackupRetentionPeriod,omitempty"`
	DBSecurityGroups                      []AwsDBSecurityGroup        `json:"DBSecurityGroups,omitempty"`
	VpcSecurityGroups                     []AwsVpcSecurityGroup       `json:"VpcSecurityGroups,omitempty"`
	DBParameterGroups                     []AwsDBParameterGroup       `json:"DBParameterGroups,omitempty"`
	AvailabilityZone                      string                      `json:"AvailabilityZone,omitempty"`
	DBSubnetGroup                         AwsDBSubnetGroup            `json:"DBSubnetGroup,omitempty"`
	PreferredMaintenanceWindow            string                      `json:"PreferredMaintenanceWindow,omitempty"`
	PendingModifiedValues                 AwsPendingModifiedValues    `json:"PendingModifiedValues,omitempty"`
	LatestRestorableTime                  string                      `json:"LatestRestorableTime,omitempty"`
	MultiAZ                               bool                        `json:"MultiAZ,omitempty"`
	EngineVersion                         string                      `json:"EngineVersion,omitempty"`
	AutoMinorVersionUpgrade               bool                        `json:"AutoMinorVersionUpgrade,omitempty"`
	ReadReplicaSourceDBInstanceIdentifier string                      `json:"ReadReplicaSourceDBInstanceIdentifier,omitempty"`
	ReadReplicaDBInstanceIdentifiers      []string                    `json:"ReadReplicaDBInstanceIdentifiers,omitempty"`
	ReadReplicaDBClusterIdentifiers       []string                    `json:"ReadReplicaDBClusterIdentifiers,omitempty"`
	LicenseModel                          string                      `json:"LicenseModel,omitempty"`
	Iops                                  int16                       `json:"Iops,omitempty"`
	OptionGroupMemberships                []AwsOptionGroupMemberships `json:"OptionGroupMemberships,omitempty"`
	CharacterSetName                      string                      `json:"CharacterSetName,omitempty"`
	SecondaryAvailabilityZone             string                      `json:"SecondaryAvailabilityZone,omitempty"`
	PubliclyAccessible                    bool                        `json:"PubliclyAccessible,omitempty"`
	StatusInfos                           []AwsStatusInfo             `json:"StatusInfos,omitempty"`
	StorageType                           string                      `json:"StorageType,omitempty"`
	TdeCredentialArn                      string                      `json:"TdeCredentialArn,omitempty"`
	DbInstancePort                        int16                       `json:"DbInstancePort,omitempty"`
	DBClusterIdentifier                   string                      `json:"DBClusterIdentifier,omitempty"`
	StorageEncrypted                      bool                        `json:"StorageEncrypted,omitempty"`
	KmsKeyId                              string                      `json:"KmsKeyId,omitempty"`
	DbiResourceId                         string                      `json:"DbiResourceId,omitempty"`
	CACertificateIdentifier               string                      `json:"CACertificateIdentifier,omitempty"`
	DomainMemberships                     []AwsDomainMemberships      `json:"DomainMemberships,omitempty"`
	CopyTagsToSnapshot                    bool                        `json:"CopyTagsToSnapshot,omitempty"`
	MonitoringInterval                    int16                       `json:"MonitoringInterval,omitempty"`
	EnhancedMonitoringResourceArn         string                      `json:"EnhancedMonitoringResourceArn,omitempty"`
	MonitoringRoleArn                     string                      `json:"MonitoringRoleArn,omitempty"`
	PromotionTier                         int16                       `json:"PromotionTier,omitempty"`
	DBInstanceArn                         string                      `json:"DBInstanceArn,omitempty"`
	Timezone                              string                      `json:"Timezone,omitempty"`
	IAMDatabaseAuthenticationEnabled      bool                        `json:"IAMDatabaseAuthenticationEnabled,omitempty"`
	PerformanceInsightsEnabled            bool                        `json:"PerformanceInsightsEnabled,omitempty"`
	PerformanceInsightsKMSKeyId           string                      `json:"PerformanceInsightsKMSKeyId,omitempty"`
	PerformanceInsightsRetentionPeriod    int16                       `json:"PerformanceInsightsRetentionPeriod,omitempty"`
	EnabledCloudwatchLogsExports          []string                    `json:"EnabledCloudwatchLogsExports,omitempty"`
	ProcessorFeatures                     []AwsProcessorFeatures      `json:"ProcessorFeatures,omitempty"`
	DeletionProtection                    bool                        `json:"DeletionProtection,omitempty"`
	AssociatedRoles                       []AwsAssociatedRoles        `json:"AssociatedRoles,omitempty"`
	ListenerEndpoint                      AwsEndpoint                 `json:"ListenerEndpoint,omitempty"`
	MaxAllocatedStorage                   int16                       `json:"MaxAllocatedStorage,omitempty"`
}

type AwsDBInstances struct {
	DBInstances []AwsDBInstance `json:"DBInstances,omitempty"`
}

func awsDBInstanceParseAndPersist(db *sql.DB, jsonString string) error {
	fmt.Printf("awsDBInstanceParseAndPersist with json lenght %d\n", len(jsonString))
	awsDBInstances := new(AwsDBInstances)
	errUnmarshal := json.Unmarshal([]byte(jsonString), awsDBInstances)
	if errUnmarshal == nil {
		for _, awsDBInstance := range awsDBInstances.DBInstances {
			fmt.Printf("DBName(%s) DBInstanceStatus(%s)\n", awsDBInstance.DBName, awsDBInstance.DBInstanceStatus)
			dbInstance := createDBInstance(db, awsDBInstance.DBInstanceIdentifier, awsDBInstance.DBInstanceClass, awsDBInstance.Engine, awsDBInstance.DBInstanceStatus, awsDBInstance.MasterUsername, awsDBInstance.DBName, awsDBInstance.Endpoint.Address, awsDBInstance.Endpoint.Port, awsDBInstance.Endpoint.HostedZoneId, awsDBInstance.AllocatedStorage, awsDBInstance.InstanceCreateTime, awsDBInstance.PreferredBackupWindow,
				awsDBInstance.BackupRetentionPeriod, awsDBInstance.AvailabilityZone, awsDBInstance.PreferredMaintenanceWindow, awsDBInstance.PendingModifiedValues.DBInstanceClass, awsDBInstance.PendingModifiedValues.AllocatedStorage, awsDBInstance.PendingModifiedValues.MasterUserPassword, awsDBInstance.PendingModifiedValues.Port, awsDBInstance.PendingModifiedValues.BackupRetentionPeriod,
				awsDBInstance.PendingModifiedValues.MultiAZ, awsDBInstance.PendingModifiedValues.EngineVersion, awsDBInstance.PendingModifiedValues.LicenseModel, awsDBInstance.PendingModifiedValues.Iops, awsDBInstance.PendingModifiedValues.DBInstanceIdentifier, awsDBInstance.PendingModifiedValues.StorageType, awsDBInstance.PendingModifiedValues.CACertificateIdentifier, awsDBInstance.PendingModifiedValues.DBSubnetGroupName,
				awsDBInstance.LatestRestorableTime, awsDBInstance.MultiAZ, awsDBInstance.EngineVersion, awsDBInstance.AutoMinorVersionUpgrade, awsDBInstance.ReadReplicaSourceDBInstanceIdentifier, awsDBInstance.LicenseModel, awsDBInstance.Iops, awsDBInstance.CharacterSetName, awsDBInstance.SecondaryAvailabilityZone, awsDBInstance.PubliclyAccessible, awsDBInstance.StorageType, awsDBInstance.TdeCredentialArn, awsDBInstance.DbInstancePort,
				awsDBInstance.DBClusterIdentifier, awsDBInstance.StorageEncrypted, awsDBInstance.KmsKeyId, awsDBInstance.DbiResourceId, awsDBInstance.CACertificateIdentifier, awsDBInstance.CopyTagsToSnapshot, awsDBInstance.MonitoringInterval, awsDBInstance.EnhancedMonitoringResourceArn, awsDBInstance.MonitoringRoleArn, awsDBInstance.PromotionTier, awsDBInstance.DBInstanceArn, awsDBInstance.Timezone,
				awsDBInstance.IAMDatabaseAuthenticationEnabled, awsDBInstance.PerformanceInsightsEnabled, awsDBInstance.PerformanceInsightsKMSKeyId, awsDBInstance.PerformanceInsightsRetentionPeriod, awsDBInstance.DeletionProtection, awsDBInstance.ListenerEndpoint.Address, awsDBInstance.ListenerEndpoint.Port, awsDBInstance.ListenerEndpoint.HostedZoneId, awsDBInstance.MaxAllocatedStorage)
			if dbInstance != nil {
				fmt.Printf("    DBInstance %s loaded\n", dbInstance.DBName)
				createRestOfDBInstanceContent(db, dbInstance, awsDBInstance)
			} else {
				fmt.Printf("  ERROR  DBInstance %s not loaded\n", awsDBInstance.DBName)
			}
		}
	} else {
		return errUnmarshal
	}
	return nil
}

func createRestOfDBInstanceContent(db *sql.DB, dbInstance *DBInstance, awsDBInstance AwsDBInstance) {

	//db_instance_associated_roles
	for _, assRole := range awsDBInstance.AssociatedRoles {
		role := createDBInstanceAssociatedRoles(db, dbInstance.Id, assRole.RoleArn, assRole.FeatureName, assRole.Status)
		if role != nil {
			fmt.Printf("    	AssociatedRole %s loaded\n", role.FeatureName)
		} else {
			fmt.Printf("    	ERROR AssociatedRole %s not loaded\n", assRole.FeatureName)
		}
	}
	//db_instance_processor_features
	for _, procFeat := range awsDBInstance.ProcessorFeatures {
		feat := createDBInstanceProcessorFeatures(db, dbInstance.Id, procFeat.Name, procFeat.Value)
		if feat != nil {
			fmt.Printf("    	ProcessorFeature %s loaded\n", feat.Name)
		} else {
			fmt.Printf("    	ERROR AssociProcessorFeatureatedRole %s not loaded\n", procFeat.Name)
		}
	}
	//db_instance_domain_memberships
	for _, domainMem := range awsDBInstance.DomainMemberships {
		member := createDBInstanceDomainMemberships(db, dbInstance.Id, domainMem.Domain, domainMem.Status, domainMem.FQDN, domainMem.IAMRoleName)
		if member != nil {
			fmt.Printf("    	DomainMember %s loaded\n", member.Domain)
		} else {
			fmt.Printf("    	ERROR DomainMember %s not loaded\n", domainMem.Domain)
		}
	}
	//db_instance_status_infos
	for _, statInfo := range awsDBInstance.StatusInfos {
		stat := createDBInstanceStatusInfos(db, dbInstance.Id, statInfo.StatusType, statInfo.Normal, statInfo.Status, statInfo.Message)
		if stat != nil {
			fmt.Printf("    	StatInfo %s loaded\n", stat.Normal)
		} else {
			fmt.Printf("    	ERROR StatInfo %s not loaded\n", statInfo.Normal)
		}
	}
	//db_instance_option_group_memberships
	for _, groupMem := range awsDBInstance.OptionGroupMemberships {
		group := createDBInstanceOptionGroupMemberships(db, dbInstance.Id, groupMem.OptionGroupName, groupMem.Status)
		if group != nil {
			fmt.Printf("    	OptionGroupMember %s loaded\n", group.Name)
		} else {
			fmt.Printf("    	ERROR OptionGroupMember %s not loaded\n", groupMem.OptionGroupName)
		}
	}
	//db_instance_db_parameter_groups
	for _, paramGroup := range awsDBInstance.DBParameterGroups {
		group := createDBInstanceDBParameterGroups(db, dbInstance.Id, paramGroup.DBParameterGroupName, paramGroup.ParameterApplyStatus)
		if group != nil {
			fmt.Printf("    	DBParameterGroup %s loaded\n", group.Name)
		} else {
			fmt.Printf("    	ERROR DBParameterGroup %s not loaded\n", paramGroup.DBParameterGroupName)
		}
	}
	//db_instance_vpc_security_groups
	for _, vpcGroup := range awsDBInstance.VpcSecurityGroups {
		vpc := createDBInstanceVpcSecurityGroups(db, dbInstance.Id, vpcGroup.VpcSecurityGroupId, vpcGroup.Status)
		if vpc != nil {
			fmt.Printf("    	VpcSecurityGroup %s loaded\n", vpc.VpcSecurityGroupId)
		} else {
			fmt.Printf("    	ERROR VpcSecurityGroup %s not loaded\n", vpcGroup.VpcSecurityGroupId)
		}
	}
	//db_instance_db_security_groups
	for _, secGroup := range awsDBInstance.DBSecurityGroups {
		sec := createDBInstanceDBSecurityGroups(db, dbInstance.Id, secGroup.DBSecurityGroupName, secGroup.Status)
		if sec != nil {
			fmt.Printf("    	DBSecurityGroup %s loaded\n", sec.Name)
		} else {
			fmt.Printf("    	ERROR VpcSecurityGroup %s not loaded\n", secGroup.DBSecurityGroupName)
		}
	}
	//db_instance_db_subnet_group
	subGroup := createDBInstanceDBSubnetGroup(db, dbInstance.Id, awsDBInstance.DBSubnetGroup.DBSubnetGroupName,
		awsDBInstance.DBSubnetGroup.DBSubnetGroupDescription, awsDBInstance.DBSubnetGroup.VpcId,
		awsDBInstance.DBSubnetGroup.SubnetGroupStatus, awsDBInstance.DBSubnetGroup.DBSubnetGroupArn)

	if subGroup != nil {
		fmt.Printf("    	SubnetGroup %s loaded\n", subGroup.DBSubnetGroupName)
		for _, subnet := range awsDBInstance.DBSubnetGroup.Subnets {
			sub := createDBInstanceDBSubnetGroupSubnet(db, subGroup.Id, subnet.SubnetIdentifier, subnet.SubnetAvailabilityZone.Name, subnet.SubnetStatus)
			if sub != nil {
				fmt.Printf("    		Subnet %s loaded\n", sub.Name)
			} else {
				fmt.Printf("    		Subnet %s loaded\n", subnet.SubnetAvailabilityZone.Name)
			}
		}

	} else {
		fmt.Printf("    	ERROR SubnetGroup %s not loaded\n", awsDBInstance.DBSubnetGroup.DBSubnetGroupName)

	}

}
