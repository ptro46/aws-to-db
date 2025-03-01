package main

import (
	"database/sql"
	"fmt"
)

type DBInstance struct {
	Id                                           int64
	DBInstanceIdentifier                         string
	DBInstanceClass                              string
	Engine                                       string
	DBInstanceStatus                             string
	MasterUsername                               string
	DBName                                       string
	EndpointAddress                              string
	EndpointPort                                 int16
	EndpointHostedZoneId                         string
	EllocatedStorage                             int16
	InstanceCreateTime                           string
	PreferredBackupWindow                        string
	BackupRetentionPeriod                        int16
	DBSecurityGroups                             []*DBInstanceDBSecurityGroup
	VpcSecurityGroups                            []*DBInstanceGroupType
	DBParameterGroups                            []*DBInstanceGroupType
	AvailabilityZone                             string
	PreferredMaintenanceWindow                   string
	PendingModifiedValuesDBInstanceClass         string
	PendingModifiedValuesAllocatedStorage        int16
	PendingModifiedValuesMasterUserPassword      string
	PendingModifiedValuesPort                    int16
	PendingModifiedValuesBackupRetentionPeriod   int16
	PendingModifiedValuesMultiAz                 bool
	PendingModifiedValuesEngineVersion           string
	PendingModifiedValuesLicenseModel            string
	PendingModifiedValuesIops                    int16
	PendingModifiedValuesDBInstanceIdentifier    string
	PendingModifiedValuesStorageType             string
	PendingModifiedValuesCaCertificateIdentifier string
	PendingModifiedValuesDBSubnetGroupName       string
	SubnetGroups                                 []*DBInstanceDBSubnetGroup
	LatestRestorableTime                         string
	MultiAz                                      bool
	EngineVersion                                string
	AutoMinorVersionUpgrade                      bool
	ReadReplicaSourceDBInstanceIdentifier        string
	LicenseModel                                 string
	Iops                                         int16
	OptionGroupMemberships                       []*DBInstanceGroupType
	CharacterSetName                             string
	SecondaryAvailabilityZone                    string
	PubliclyAccessible                           bool
	StatusInfos                                  []*DBInstanceStatusInfos
	StorageType                                  string
	TDECredentialArn                             string
	DBInstancePort                               int16
	DBClusterIdentifier                          string
	StorageEncrypted                             bool
	KMSKeyId                                     string
	DBiResourceId                                string
	CaCertificateIdentifier                      string
	DomainMemberships                            []*DBInstanceDomainMemberships
	CopyTagsToSnapshot                           bool
	MonitoringInterval                           int16
	EnhancedMonitoringResourceArn                string
	MonitoringRoleArn                            string
	PromotionTier                                int16
	DBInstanceArn                                string
	Timezone                                     string
	AamDatabaseAuthenticationEnabled             bool
	PerformanceInsightsEnabled                   bool
	PerformanceInsightsKMSKeyId                  string
	PerformanceInsightsRetentionPeriod           int16
	ProcessorFeatures                            []*DBInstanceProcessorFeatures
	DeletionProtection                           bool
	AssociatedRoles                              []*DBInstanceAssociatedRoles
	ListenerEndpointAddress                      string
	ListenerEndpointPort                         int16
	ListenerEndpointHostedZoneId                 string
	MaxAllocatedStorage                          int16
}

func (d *DBInstance) loadDependencies(db *sql.DB) {
	arrayOfDBSecurityGroups, err := loadAllDBInstanceDBSecurityGroupByOrderUnique(db)
	if err == nil {
		for _, x := range arrayOfDBSecurityGroups {
			if x.DbInstanceId == d.Id {
				d.DBSecurityGroups = append(d.DBSecurityGroups, x)
			}
		}
	} else {
		fmt.Println(err, "arrayOfDBSecurityGroups")
	}

	arrayOfDBInstanceDBSubnetGroup, err := loadAllDBInstanceDBSubnetGroupByOrderUnique(db)
	if err == nil {
		for _, x := range arrayOfDBInstanceDBSubnetGroup {
			if x.DbInstanceId == d.Id {
				d.SubnetGroups = append(d.SubnetGroups, x)
			}
		}
	} else {
		fmt.Println(err, "arrayOfDBInstanceDBSubnetGroup")
	}

	arrayOfVpcSecurityGroups, err := loadAllDBInstanceDBVpcSecurityGroupsByOrderUnique(db)
	if err == nil {
		for _, x := range arrayOfVpcSecurityGroups {
			if x.DbInstanceId == d.Id {
				d.VpcSecurityGroups = append(d.VpcSecurityGroups, x)
			}
		}
	} else {
		fmt.Println(err, "arrayOfVpcSecurityGroups")
	}
	arrayOfDBParameterGroups, err := loadAllDBInstanceDBParameterGroupsByOrderUnique(db)
	if err == nil {
		for _, x := range arrayOfDBParameterGroups {
			if x.DbInstanceId == d.Id {
				d.DBParameterGroups = append(d.DBParameterGroups, x)
			}
		}
	} else {
		fmt.Println(err, "arrayOfDBParameterGroups")
	}

	arrayOfOptionGroupMemberships, err := loadAllDBInstanceDBOptionGroupMembershipsByOrderUnique(db)
	if err == nil {
		for _, x := range arrayOfOptionGroupMemberships {
			if x.DbInstanceId == d.Id {
				d.OptionGroupMemberships = append(d.OptionGroupMemberships, x)
			}
		}
	} else {
		fmt.Println(err, "arrayOfOptionGroupMemberships")
	}

	arrayOfStatusInfos, err := loadAllDBInstanceStatusInfosByOrderUnique(db)
	if err == nil {
		for _, x := range arrayOfStatusInfos {
			if x.DbInstanceId == d.Id {
				d.StatusInfos = append(d.StatusInfos, x)
			}
		}
	} else {
		fmt.Println(err, "arrayOfStatusInfos")
	}

	arrayOfDomainMemberships, err := loadAllDBInstanceDomainMembershipsByOrderUnique(db)
	if err == nil {
		for _, x := range arrayOfDomainMemberships {
			if x.DbInstanceId == d.Id {
				d.DomainMemberships = append(d.DomainMemberships, x)
			}
		}
	} else {
		fmt.Println(err, "arrayOfDomainMemberships")
	}

	arrayOfProcessorFeatures, err := loadAllDBInstanceProcessorFeaturesByOrderUnique(db)
	if err == nil {
		for _, x := range arrayOfProcessorFeatures {
			if x.DbInstanceId == d.Id {
				d.ProcessorFeatures = append(d.ProcessorFeatures, x)
			}
		}
	} else {
		fmt.Println(err, "arrayOfProcessorFeatures")
	}

	arrayOfAssociatedRoles, err := loadAllDBInstanceAssociatedRolesByOrderUnique(db)
	if err == nil {
		for _, x := range arrayOfAssociatedRoles {
			if x.DbInstanceId == d.Id {
				d.AssociatedRoles = append(d.AssociatedRoles, x)
			}
		}
	} else {
		fmt.Println(err, "arrayOfAssociatedRoles")
	}

}

func (d *DBInstance) Dump(db *sql.DB) string {
	dumpDB := fmt.Sprintf("DBInstanceIdentifier(%s) DBInstanceClass(%s) Engine(%s) DBInstanceStatus(%s) MasterUsername(%s) DBName(%s) EndpointAddress(%s) EndpointPort(%d) EndpointHostedZoneId(%s) EllocatedStorage(%d) InstanceCreateTime(%s) PreferredBackupWindow(%s) BackupRetentionPeriod(%d) AvailabilityZone(%s) PreferredMaintenanceWindow(%s) PendingModifiedValuesDBInstanceClass(%s) PendingModifiedValuesAllocatedStorage(%s) PendingModifiedValuesMasterUserPassword(%s) PendingModifiedValuesPort(%d) PendingModifiedValuesBackupRetentionPeriod(%d) PendingModifiedValuesMultiAz(%t) PendingModifiedValuesEngineVersion(%s) PendingModifiedValuesLicenseModel(%s) PendingModifiedValuesIops(%d) PendingModifiedValuesDBInstanceIdentifier(%s) PendingModifiedValuesStorageType(%s) PendingModifiedValuesCaCertificateIdentifier(%s) PendingModifiedValuesDBSubnetGroupName(%s) LatestRestorableTime(%s) MultiAz(%t) EngineVersion(%s) AutoMinorVersionUpgrade(%t) ReadReplicaSourceDBInstanceIdentifier(%s) LicenseModel(%s) Iops(%d) CharacterSetName(%s) SecondaryAvailabilityZone(%s) PubliclyAccessible(%t) StorageType(%s) TDECredentialArn(%s) DBInstancePort(%d) DBClusterIdentifier(%s) StorageEncrypted(%t) KMSKeyId(%s) DBiResourceId(%s) CaCertificateIdentifier(%s) CopyTagsToSnapshot(%t) MonitoringInterval(%d) EnhancedMonitoringResourceArn(%s) MonitoringRoleArn(%s) PromotionTier(%d) DBInstanceArn(%s) Timezone(%s) AamDatabaseAuthenticationEnabled(%t) PerformanceInsightsEnabled(%s) PerformanceInsightsKMSKeyId(%s) PerformanceInsightsRetentionPeriod(%d) DeletionProtection(%t) ListenerEndpointAddress(%s) ListenerEndpointPort(%d) ListenerEndpointHostedZoneId(%s) MaxAllocatedStorage(%d)",
		d.DBInstanceIdentifier,
		d.DBInstanceClass,
		d.Engine,
		d.DBInstanceStatus,
		d.MasterUsername,
		d.DBName,
		d.EndpointAddress,
		d.EndpointPort,
		d.EndpointHostedZoneId,
		d.EllocatedStorage,
		d.InstanceCreateTime,
		d.PreferredBackupWindow,
		d.BackupRetentionPeriod,
		d.AvailabilityZone,
		d.PreferredMaintenanceWindow,
		d.PendingModifiedValuesDBInstanceClass,
		d.PendingModifiedValuesAllocatedStorage,
		d.PendingModifiedValuesMasterUserPassword,
		d.PendingModifiedValuesPort,
		d.PendingModifiedValuesBackupRetentionPeriod,
		d.PendingModifiedValuesMultiAz,
		d.PendingModifiedValuesEngineVersion,
		d.PendingModifiedValuesLicenseModel,
		d.PendingModifiedValuesIops,
		d.PendingModifiedValuesDBInstanceIdentifier,
		d.PendingModifiedValuesStorageType,
		d.PendingModifiedValuesCaCertificateIdentifier,
		d.PendingModifiedValuesDBSubnetGroupName,
		d.LatestRestorableTime,
		d.MultiAz,
		d.EngineVersion,
		d.AutoMinorVersionUpgrade,
		d.ReadReplicaSourceDBInstanceIdentifier,
		d.LicenseModel,
		d.Iops,
		d.CharacterSetName,
		d.SecondaryAvailabilityZone,
		d.PubliclyAccessible,
		d.StorageType,
		d.TDECredentialArn,
		d.DBInstancePort,
		d.DBClusterIdentifier,
		d.StorageEncrypted,
		d.KMSKeyId,
		d.DBiResourceId,
		d.CaCertificateIdentifier,
		d.CopyTagsToSnapshot,
		d.MonitoringInterval,
		d.EnhancedMonitoringResourceArn,
		d.MonitoringRoleArn,
		d.PromotionTier,
		d.DBInstanceArn,
		d.Timezone,
		d.AamDatabaseAuthenticationEnabled,
		d.PerformanceInsightsEnabled,
		d.PerformanceInsightsKMSKeyId,
		d.PerformanceInsightsRetentionPeriod,
		d.DeletionProtection,
		d.ListenerEndpointAddress,
		d.ListenerEndpointPort,
		d.ListenerEndpointHostedZoneId,
		d.MaxAllocatedStorage)

	dumpDB += "\n        DBSecurityGroups ["
	for _, x := range d.DBSecurityGroups {
		dumpDB += "{" + x.Dump() + "}"
	}
	dumpDB += "]\n        VpcSecurityGroups ["
	for _, x := range d.VpcSecurityGroups {
		dumpDB += "{" + x.Dump() + "}"
	}
	dumpDB += "]\n        DBParameterGroups ["
	for _, x := range d.DBParameterGroups {
		dumpDB += "{" + x.Dump() + "}"
	}
	dumpDB += "]\n        OptionGroupMemberships ["
	for _, x := range d.OptionGroupMemberships {
		dumpDB += "{" + x.Dump() + "}"
	}
	dumpDB += "]\n        StatusInfos ["
	for _, x := range d.StatusInfos {
		dumpDB += "{" + x.Dump() + "}"
	}
	dumpDB += "]\n        DomainMemberships ["
	for _, x := range d.DomainMemberships {
		dumpDB += "{" + x.Dump() + "}"
	}
	dumpDB += "]\n        ProcessorFeatures ["
	for _, x := range d.ProcessorFeatures {
		dumpDB += "{" + x.Dump() + "}"
	}
	dumpDB += "]\n        AssociatedRoles ["
	for _, x := range d.AssociatedRoles {
		dumpDB += "{" + x.Dump() + "}"
	}
	dumpDB += "]\n        SubnetGroups ["
	for _, x := range d.SubnetGroups {
		x.loadDependencies(db)
		dumpDB += "{" + x.Dump() + "}"
	}
	dumpDB += "]"
	return dumpDB

}

func NewDBInstance(Id int64,
	DBInstanceIdentifier string,
	DBInstanceClass string,
	Engine string,
	DBInstanceStatus string,
	MasterUsername string,
	DBName string,
	EndpointAddress string,
	EndpointPort int16,
	EndpointHostedZoneId string,
	EllocatedStorage int16,
	InstanceCreateTime string,
	PreferredBackupWindow string,
	BackupRetentionPeriod int16,
	AvailabilityZone string,
	PreferredMaintenanceWindow string,
	PendingModifiedValuesDBInstanceClass string,
	PendingModifiedValuesAllocatedStorage int16,
	PendingModifiedValuesMasterUserPassword string,
	PendingModifiedValuesPort int16,
	PendingModifiedValuesBackupRetentionPeriod int16,
	PendingModifiedValuesMultiAz bool,
	PendingModifiedValuesEngineVersion string,
	PendingModifiedValuesLicenseModel string,
	PendingModifiedValuesIops int16,
	PendingModifiedValuesDBInstanceIdentifier string,
	PendingModifiedValuesStorageType string,
	PendingModifiedValuesCaCertificateIdentifier string,
	PendingModifiedValuesDBSubnetGroupName string,
	LatestRestorableTime string,
	MultiAz bool,
	EngineVersion string,
	AutoMinorVersionUpgrade bool,
	ReadReplicaSourceDBInstanceIdentifier string,
	LicenseModel string,
	Iops int16,
	CharacterSetName string,
	SecondaryAvailabilityZone string,
	PubliclyAccessible bool,
	StorageType string,
	TDECredentialArn string,
	DBInstancePort int16,
	DBClusterIdentifier string,
	StorageEncrypted bool,
	KMSKeyId string,
	DBiResourceId string,
	CaCertificateIdentifier string,
	CopyTagsToSnapshot bool,
	MonitoringInterval int16,
	EnhancedMonitoringResourceArn string,
	MonitoringRoleArn string,
	PromotionTier int16,
	DBInstanceArn string,
	Timezone string,
	AamDatabaseAuthenticationEnabled bool,
	PerformanceInsightsEnabled bool,
	PerformanceInsightsKMSKeyId string,
	PerformanceInsightsRetentionPeriod int16,
	DeletionProtection bool,
	ListenerEndpointAddress string,
	ListenerEndpointPort int16,
	ListenerEndpointHostedZoneId string,
	MaxAllocatedStorage int16) *DBInstance {

	return &DBInstance{Id: Id,
		DBInstanceIdentifier:                         DBInstanceIdentifier,
		DBInstanceClass:                              DBInstanceClass,
		Engine:                                       Engine,
		DBInstanceStatus:                             DBInstanceStatus,
		MasterUsername:                               MasterUsername,
		DBName:                                       DBName,
		EndpointAddress:                              EndpointAddress,
		EndpointPort:                                 EndpointPort,
		EndpointHostedZoneId:                         EndpointHostedZoneId,
		EllocatedStorage:                             EllocatedStorage,
		InstanceCreateTime:                           InstanceCreateTime,
		PreferredBackupWindow:                        PreferredBackupWindow,
		BackupRetentionPeriod:                        BackupRetentionPeriod,
		AvailabilityZone:                             AvailabilityZone,
		PreferredMaintenanceWindow:                   PreferredMaintenanceWindow,
		PendingModifiedValuesDBInstanceClass:         PendingModifiedValuesDBInstanceClass,
		PendingModifiedValuesAllocatedStorage:        PendingModifiedValuesAllocatedStorage,
		PendingModifiedValuesMasterUserPassword:      PendingModifiedValuesMasterUserPassword,
		PendingModifiedValuesPort:                    PendingModifiedValuesPort,
		PendingModifiedValuesBackupRetentionPeriod:   PendingModifiedValuesBackupRetentionPeriod,
		PendingModifiedValuesMultiAz:                 PendingModifiedValuesMultiAz,
		PendingModifiedValuesEngineVersion:           PendingModifiedValuesEngineVersion,
		PendingModifiedValuesLicenseModel:            PendingModifiedValuesLicenseModel,
		PendingModifiedValuesIops:                    PendingModifiedValuesIops,
		PendingModifiedValuesDBInstanceIdentifier:    PendingModifiedValuesDBInstanceIdentifier,
		PendingModifiedValuesStorageType:             PendingModifiedValuesStorageType,
		PendingModifiedValuesCaCertificateIdentifier: PendingModifiedValuesCaCertificateIdentifier,
		PendingModifiedValuesDBSubnetGroupName:       PendingModifiedValuesDBSubnetGroupName,
		LatestRestorableTime:                         LatestRestorableTime,
		MultiAz:                                      MultiAz,
		EngineVersion:                                EngineVersion,
		AutoMinorVersionUpgrade:                      AutoMinorVersionUpgrade,
		ReadReplicaSourceDBInstanceIdentifier:        ReadReplicaSourceDBInstanceIdentifier,
		LicenseModel:                                 LicenseModel,
		Iops:                                         Iops,
		CharacterSetName:                             CharacterSetName,
		SecondaryAvailabilityZone:                    SecondaryAvailabilityZone,
		PubliclyAccessible:                           PubliclyAccessible,
		StorageType:                                  StorageType,
		TDECredentialArn:                             TDECredentialArn,
		DBInstancePort:                               DBInstancePort,
		DBClusterIdentifier:                          DBClusterIdentifier,
		StorageEncrypted:                             StorageEncrypted,
		KMSKeyId:                                     KMSKeyId,
		DBiResourceId:                                DBiResourceId,
		CaCertificateIdentifier:                      CaCertificateIdentifier,
		CopyTagsToSnapshot:                           CopyTagsToSnapshot,
		MonitoringInterval:                           MonitoringInterval,
		EnhancedMonitoringResourceArn:                EnhancedMonitoringResourceArn,
		MonitoringRoleArn:                            MonitoringRoleArn,
		PromotionTier:                                PromotionTier,
		DBInstanceArn:                                DBInstanceArn,
		Timezone:                                     Timezone,
		AamDatabaseAuthenticationEnabled:             AamDatabaseAuthenticationEnabled,
		PerformanceInsightsEnabled:                   PerformanceInsightsEnabled,
		PerformanceInsightsKMSKeyId:                  PerformanceInsightsKMSKeyId,
		PerformanceInsightsRetentionPeriod:           PerformanceInsightsRetentionPeriod,
		DeletionProtection:                           DeletionProtection,
		ListenerEndpointAddress:                      ListenerEndpointAddress,
		ListenerEndpointPort:                         ListenerEndpointPort,
		ListenerEndpointHostedZoneId:                 ListenerEndpointHostedZoneId,
		MaxAllocatedStorage:                          MaxAllocatedStorage}

}
