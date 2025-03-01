package main

import (
	"database/sql" // package SQL
	"fmt"
	"log"

	_ "github.com/lib/pq" // driver Postgres
)

func rowResultSetToDBInstance(row *sql.Row) (*DBInstance, error) {
	var err error
	var Id int64
	var DBInstanceIdentifier string
	var DBInstanceClass string
	var Engine string
	var DBInstanceStatus string
	var MasterUsername string
	var DBName string
	var EndpointAddress string
	var EndpointPort int16
	var EndpointHostedZoneId string
	var EllocatedStorage int16
	var InstanceCreateTime string
	var PreferredBackupWindow string
	var BackupRetentionPeriod int16
	var AvailabilityZone string
	var PreferredMaintenanceWindow string
	var PendingModifiedValuesDBInstanceClass string
	var PendingModifiedValuesAllocatedStorage int16
	var PendingModifiedValuesMasterUserPassword string
	var PendingModifiedValuesPort int16
	var PendingModifiedValuesBackupRetentionPeriod int16
	var PendingModifiedValuesMultiAz bool
	var PendingModifiedValuesEngineVersion string
	var PendingModifiedValuesLicenseModel string
	var PendingModifiedValuesIops int16
	var PendingModifiedValuesDBInstanceIdentifier string
	var PendingModifiedValuesStorageType string
	var PendingModifiedValuesCaCertificateIdentifier string
	var PendingModifiedValuesDBSubnetGroupName string
	var LatestRestorableTime string
	var MultiAz bool
	var EngineVersion string
	var AutoMinorVersionUpgrade bool
	var ReadReplicaSourceDBInstanceIdentifier string
	var LicenseModel string
	var Iops int16
	var CharacterSetName string
	var SecondaryAvailabilityZone string
	var PubliclyAccessible bool
	var StorageType string
	var TDECredentialArn string
	var DBInstancePort int16
	var DBClusterIdentifier string
	var StorageEncrypted bool
	var KMSKeyId string
	var DBiResourceId string
	var CaCertificateIdentifier string
	var CopyTagsToSnapshot bool
	var MonitoringInterval int16
	var EnhancedMonitoringResourceArn string
	var MonitoringRoleArn string
	var PromotionTier int16
	var DBInstanceArn string
	var Timezone string
	var AamDatabaseAuthenticationEnabled bool
	var PerformanceInsightsEnabled bool
	var PerformanceInsightsKMSKeyId string
	var PerformanceInsightsRetentionPeriod int16
	var DeletionProtection bool
	var ListenerEndpointAddress string
	var ListenerEndpointPort int16
	var ListenerEndpointHostedZoneId string
	var MaxAllocatedStorage int16

	err = row.Scan(&Id, &DBInstanceIdentifier, &DBInstanceClass, &Engine, &DBInstanceStatus, &MasterUsername, &DBName, &EndpointAddress, &EndpointPort,
		&EndpointHostedZoneId, &EllocatedStorage, &InstanceCreateTime, &PreferredBackupWindow, &BackupRetentionPeriod, &AvailabilityZone, &PreferredMaintenanceWindow,
		&PendingModifiedValuesDBInstanceClass, &PendingModifiedValuesAllocatedStorage, &PendingModifiedValuesMasterUserPassword, &PendingModifiedValuesPort,
		&PendingModifiedValuesBackupRetentionPeriod, &PendingModifiedValuesMultiAz, &PendingModifiedValuesEngineVersion, &PendingModifiedValuesLicenseModel,
		&PendingModifiedValuesIops, &PendingModifiedValuesDBInstanceIdentifier, &PendingModifiedValuesStorageType, &PendingModifiedValuesCaCertificateIdentifier,
		&PendingModifiedValuesDBSubnetGroupName, &LatestRestorableTime, &MultiAz, &EngineVersion, &AutoMinorVersionUpgrade, &ReadReplicaSourceDBInstanceIdentifier,
		&LicenseModel, &Iops, &CharacterSetName, &SecondaryAvailabilityZone, &PubliclyAccessible, &StorageType, &TDECredentialArn, &DBInstancePort, &DBClusterIdentifier,
		&StorageEncrypted, &KMSKeyId, &DBiResourceId, &CaCertificateIdentifier, &CopyTagsToSnapshot, &MonitoringInterval, &EnhancedMonitoringResourceArn, &MonitoringRoleArn,
		&PromotionTier, &DBInstanceArn, &Timezone, &AamDatabaseAuthenticationEnabled, &PerformanceInsightsEnabled, &PerformanceInsightsKMSKeyId, &PerformanceInsightsRetentionPeriod,
		&DeletionProtection, &ListenerEndpointAddress, &ListenerEndpointPort, &ListenerEndpointHostedZoneId, &MaxAllocatedStorage)
	if err != nil {
		return nil, err
	}
	return NewDBInstance(Id, DBInstanceIdentifier, DBInstanceClass, Engine, DBInstanceStatus, MasterUsername, DBName, EndpointAddress,
		EndpointPort, EndpointHostedZoneId, EllocatedStorage, InstanceCreateTime, PreferredBackupWindow, BackupRetentionPeriod, AvailabilityZone,
		PreferredMaintenanceWindow, PendingModifiedValuesDBInstanceClass, PendingModifiedValuesAllocatedStorage, PendingModifiedValuesMasterUserPassword,
		PendingModifiedValuesPort, PendingModifiedValuesBackupRetentionPeriod, PendingModifiedValuesMultiAz, PendingModifiedValuesEngineVersion,
		PendingModifiedValuesLicenseModel, PendingModifiedValuesIops, PendingModifiedValuesDBInstanceIdentifier, PendingModifiedValuesStorageType,
		PendingModifiedValuesCaCertificateIdentifier, PendingModifiedValuesDBSubnetGroupName, LatestRestorableTime, MultiAz,
		EngineVersion, AutoMinorVersionUpgrade, ReadReplicaSourceDBInstanceIdentifier, LicenseModel, Iops, CharacterSetName, SecondaryAvailabilityZone,
		PubliclyAccessible, StorageType, TDECredentialArn, DBInstancePort, DBClusterIdentifier, StorageEncrypted, KMSKeyId, DBiResourceId,
		CaCertificateIdentifier, CopyTagsToSnapshot, MonitoringInterval, EnhancedMonitoringResourceArn, MonitoringRoleArn, PromotionTier,
		DBInstanceArn, Timezone, AamDatabaseAuthenticationEnabled, PerformanceInsightsEnabled, PerformanceInsightsKMSKeyId, PerformanceInsightsRetentionPeriod,
		DeletionProtection, ListenerEndpointAddress, ListenerEndpointPort, ListenerEndpointHostedZoneId, MaxAllocatedStorage), nil
}

func rowsNoFetchResultSetToDBInstance(rows *sql.Rows) (*DBInstance, error) {
	var err error
	var Id int64
	var DBInstanceIdentifier string
	var DBInstanceClass string
	var Engine string
	var DBInstanceStatus string
	var MasterUsername string
	var DBName string
	var EndpointAddress string
	var EndpointPort int16
	var EndpointHostedZoneId string
	var EllocatedStorage int16
	var InstanceCreateTime string
	var PreferredBackupWindow string
	var BackupRetentionPeriod int16
	var AvailabilityZone string
	var PreferredMaintenanceWindow string
	var PendingModifiedValuesDBInstanceClass string
	var PendingModifiedValuesAllocatedStorage int16
	var PendingModifiedValuesMasterUserPassword string
	var PendingModifiedValuesPort int16
	var PendingModifiedValuesBackupRetentionPeriod int16
	var PendingModifiedValuesMultiAz bool
	var PendingModifiedValuesEngineVersion string
	var PendingModifiedValuesLicenseModel string
	var PendingModifiedValuesIops int16
	var PendingModifiedValuesDBInstanceIdentifier string
	var PendingModifiedValuesStorageType string
	var PendingModifiedValuesCaCertificateIdentifier string
	var PendingModifiedValuesDBSubnetGroupName string
	var LatestRestorableTime string
	var MultiAz bool
	var EngineVersion string
	var AutoMinorVersionUpgrade bool
	var ReadReplicaSourceDBInstanceIdentifier string
	var LicenseModel string
	var Iops int16
	var CharacterSetName string
	var SecondaryAvailabilityZone string
	var PubliclyAccessible bool
	var StorageType string
	var TDECredentialArn string
	var DBInstancePort int16
	var DBClusterIdentifier string
	var StorageEncrypted bool
	var KMSKeyId string
	var DBiResourceId string
	var CaCertificateIdentifier string
	var CopyTagsToSnapshot bool
	var MonitoringInterval int16
	var EnhancedMonitoringResourceArn string
	var MonitoringRoleArn string
	var PromotionTier int16
	var DBInstanceArn string
	var Timezone string
	var AamDatabaseAuthenticationEnabled bool
	var PerformanceInsightsEnabled bool
	var PerformanceInsightsKMSKeyId string
	var PerformanceInsightsRetentionPeriod int16
	var DeletionProtection bool
	var ListenerEndpointAddress string
	var ListenerEndpointPort int16
	var ListenerEndpointHostedZoneId string
	var MaxAllocatedStorage int16

	err = rows.Scan(&Id, &DBInstanceIdentifier, &DBInstanceClass, &Engine, &DBInstanceStatus, &MasterUsername, &DBName, &EndpointAddress, &EndpointPort,
		&EndpointHostedZoneId, &EllocatedStorage, &InstanceCreateTime, &PreferredBackupWindow, &BackupRetentionPeriod, &AvailabilityZone, &PreferredMaintenanceWindow,
		&PendingModifiedValuesDBInstanceClass, &PendingModifiedValuesAllocatedStorage, &PendingModifiedValuesMasterUserPassword, &PendingModifiedValuesPort,
		&PendingModifiedValuesBackupRetentionPeriod, &PendingModifiedValuesMultiAz, &PendingModifiedValuesEngineVersion, &PendingModifiedValuesLicenseModel,
		&PendingModifiedValuesIops, &PendingModifiedValuesDBInstanceIdentifier, &PendingModifiedValuesStorageType, &PendingModifiedValuesCaCertificateIdentifier,
		&PendingModifiedValuesDBSubnetGroupName, &LatestRestorableTime, &MultiAz, &EngineVersion, &AutoMinorVersionUpgrade, &ReadReplicaSourceDBInstanceIdentifier,
		&LicenseModel, &Iops, &CharacterSetName, &SecondaryAvailabilityZone, &PubliclyAccessible, &StorageType, &TDECredentialArn, &DBInstancePort, &DBClusterIdentifier,
		&StorageEncrypted, &KMSKeyId, &DBiResourceId, &CaCertificateIdentifier, &CopyTagsToSnapshot, &MonitoringInterval, &EnhancedMonitoringResourceArn, &MonitoringRoleArn,
		&PromotionTier, &DBInstanceArn, &Timezone, &AamDatabaseAuthenticationEnabled, &PerformanceInsightsEnabled, &PerformanceInsightsKMSKeyId, &PerformanceInsightsRetentionPeriod,
		&DeletionProtection, &ListenerEndpointAddress, &ListenerEndpointPort, &ListenerEndpointHostedZoneId, &MaxAllocatedStorage)
	if err != nil {
		return nil, err
	}
	return NewDBInstance(Id, DBInstanceIdentifier, DBInstanceClass, Engine, DBInstanceStatus, MasterUsername, DBName, EndpointAddress,
		EndpointPort, EndpointHostedZoneId, EllocatedStorage, InstanceCreateTime, PreferredBackupWindow, BackupRetentionPeriod, AvailabilityZone,
		PreferredMaintenanceWindow, PendingModifiedValuesDBInstanceClass, PendingModifiedValuesAllocatedStorage, PendingModifiedValuesMasterUserPassword,
		PendingModifiedValuesPort, PendingModifiedValuesBackupRetentionPeriod, PendingModifiedValuesMultiAz, PendingModifiedValuesEngineVersion,
		PendingModifiedValuesLicenseModel, PendingModifiedValuesIops, PendingModifiedValuesDBInstanceIdentifier, PendingModifiedValuesStorageType,
		PendingModifiedValuesCaCertificateIdentifier, PendingModifiedValuesDBSubnetGroupName, LatestRestorableTime, MultiAz,
		EngineVersion, AutoMinorVersionUpgrade, ReadReplicaSourceDBInstanceIdentifier, LicenseModel, Iops, CharacterSetName, SecondaryAvailabilityZone,
		PubliclyAccessible, StorageType, TDECredentialArn, DBInstancePort, DBClusterIdentifier, StorageEncrypted, KMSKeyId, DBiResourceId,
		CaCertificateIdentifier, CopyTagsToSnapshot, MonitoringInterval, EnhancedMonitoringResourceArn, MonitoringRoleArn, PromotionTier,
		DBInstanceArn, Timezone, AamDatabaseAuthenticationEnabled, PerformanceInsightsEnabled, PerformanceInsightsKMSKeyId, PerformanceInsightsRetentionPeriod,
		DeletionProtection, ListenerEndpointAddress, ListenerEndpointPort, ListenerEndpointHostedZoneId, MaxAllocatedStorage), nil
}

func rowsResultSetToDBInstance(rows *sql.Rows) (*DBInstance, error) {
	var err error
	if rows.Next() {
		var Id int64
		var DBInstanceIdentifier string
		var DBInstanceClass string
		var Engine string
		var DBInstanceStatus string
		var MasterUsername string
		var DBName string
		var EndpointAddress string
		var EndpointPort int16
		var EndpointHostedZoneId string
		var EllocatedStorage int16
		var InstanceCreateTime string
		var PreferredBackupWindow string
		var BackupRetentionPeriod int16
		var AvailabilityZone string
		var PreferredMaintenanceWindow string
		var PendingModifiedValuesDBInstanceClass string
		var PendingModifiedValuesAllocatedStorage int16
		var PendingModifiedValuesMasterUserPassword string
		var PendingModifiedValuesPort int16
		var PendingModifiedValuesBackupRetentionPeriod int16
		var PendingModifiedValuesMultiAz bool
		var PendingModifiedValuesEngineVersion string
		var PendingModifiedValuesLicenseModel string
		var PendingModifiedValuesIops int16
		var PendingModifiedValuesDBInstanceIdentifier string
		var PendingModifiedValuesStorageType string
		var PendingModifiedValuesCaCertificateIdentifier string
		var PendingModifiedValuesDBSubnetGroupName string
		var LatestRestorableTime string
		var MultiAz bool
		var EngineVersion string
		var AutoMinorVersionUpgrade bool
		var ReadReplicaSourceDBInstanceIdentifier string
		var LicenseModel string
		var Iops int16
		var CharacterSetName string
		var SecondaryAvailabilityZone string
		var PubliclyAccessible bool
		var StorageType string
		var TDECredentialArn string
		var DBInstancePort int16
		var DBClusterIdentifier string
		var StorageEncrypted bool
		var KMSKeyId string
		var DBiResourceId string
		var CaCertificateIdentifier string
		var CopyTagsToSnapshot bool
		var MonitoringInterval int16
		var EnhancedMonitoringResourceArn string
		var MonitoringRoleArn string
		var PromotionTier int16
		var DBInstanceArn string
		var Timezone string
		var AamDatabaseAuthenticationEnabled bool
		var PerformanceInsightsEnabled bool
		var PerformanceInsightsKMSKeyId string
		var PerformanceInsightsRetentionPeriod int16
		var DeletionProtection bool
		var ListenerEndpointAddress string
		var ListenerEndpointPort int16
		var ListenerEndpointHostedZoneId string
		var MaxAllocatedStorage int16

		err = rows.Scan(&Id, &DBInstanceIdentifier, &DBInstanceClass, &Engine, &DBInstanceStatus, &MasterUsername, &DBName, &EndpointAddress, &EndpointPort,
			&EndpointHostedZoneId, &EllocatedStorage, &InstanceCreateTime, &PreferredBackupWindow, &BackupRetentionPeriod, &AvailabilityZone, &PreferredMaintenanceWindow,
			&PendingModifiedValuesDBInstanceClass, &PendingModifiedValuesAllocatedStorage, &PendingModifiedValuesMasterUserPassword, &PendingModifiedValuesPort,
			&PendingModifiedValuesBackupRetentionPeriod, &PendingModifiedValuesMultiAz, &PendingModifiedValuesEngineVersion, &PendingModifiedValuesLicenseModel,
			&PendingModifiedValuesIops, &PendingModifiedValuesDBInstanceIdentifier, &PendingModifiedValuesStorageType, &PendingModifiedValuesCaCertificateIdentifier,
			&PendingModifiedValuesDBSubnetGroupName, &LatestRestorableTime, &MultiAz, &EngineVersion, &AutoMinorVersionUpgrade, &ReadReplicaSourceDBInstanceIdentifier,
			&LicenseModel, &Iops, &CharacterSetName, &SecondaryAvailabilityZone, &PubliclyAccessible, &StorageType, &TDECredentialArn, &DBInstancePort, &DBClusterIdentifier,
			&StorageEncrypted, &KMSKeyId, &DBiResourceId, &CaCertificateIdentifier, &CopyTagsToSnapshot, &MonitoringInterval, &EnhancedMonitoringResourceArn, &MonitoringRoleArn,
			&PromotionTier, &DBInstanceArn, &Timezone, &AamDatabaseAuthenticationEnabled, &PerformanceInsightsEnabled, &PerformanceInsightsKMSKeyId, &PerformanceInsightsRetentionPeriod,
			&DeletionProtection, &ListenerEndpointAddress, &ListenerEndpointPort, &ListenerEndpointHostedZoneId, &MaxAllocatedStorage)
		if err != nil {
			return nil, err
		}
		return NewDBInstance(Id, DBInstanceIdentifier, DBInstanceClass, Engine, DBInstanceStatus, MasterUsername, DBName, EndpointAddress,
			EndpointPort, EndpointHostedZoneId, EllocatedStorage, InstanceCreateTime, PreferredBackupWindow, BackupRetentionPeriod, AvailabilityZone,
			PreferredMaintenanceWindow, PendingModifiedValuesDBInstanceClass, PendingModifiedValuesAllocatedStorage, PendingModifiedValuesMasterUserPassword,
			PendingModifiedValuesPort, PendingModifiedValuesBackupRetentionPeriod, PendingModifiedValuesMultiAz, PendingModifiedValuesEngineVersion,
			PendingModifiedValuesLicenseModel, PendingModifiedValuesIops, PendingModifiedValuesDBInstanceIdentifier, PendingModifiedValuesStorageType,
			PendingModifiedValuesCaCertificateIdentifier, PendingModifiedValuesDBSubnetGroupName, LatestRestorableTime, MultiAz,
			EngineVersion, AutoMinorVersionUpgrade, ReadReplicaSourceDBInstanceIdentifier, LicenseModel, Iops, CharacterSetName, SecondaryAvailabilityZone,
			PubliclyAccessible, StorageType, TDECredentialArn, DBInstancePort, DBClusterIdentifier, StorageEncrypted, KMSKeyId, DBiResourceId,
			CaCertificateIdentifier, CopyTagsToSnapshot, MonitoringInterval, EnhancedMonitoringResourceArn, MonitoringRoleArn, PromotionTier,
			DBInstanceArn, Timezone, AamDatabaseAuthenticationEnabled, PerformanceInsightsEnabled, PerformanceInsightsKMSKeyId, PerformanceInsightsRetentionPeriod,
			DeletionProtection, ListenerEndpointAddress, ListenerEndpointPort, ListenerEndpointHostedZoneId, MaxAllocatedStorage), nil
	}
	return nil, err
}

func createDBInstance(db *sql.DB,
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
	rows := db.QueryRow("insert into db_instance(db_instance_identifier,db_instance_class,engine,db_instance_status,master_username,db_name,endpoint_address,endpoint_port,endpoint_hosted_zone_id,allocated_storage,instance_create_time,preferred_backup_window,backup_retention_period,availability_zone,preferred_maintenance_window,pending_modified_values_db_instance_class,pending_modified_values_allocated_storage,pending_modified_values_master_user_password,pending_modified_values_port,pending_modified_values_backup_retention_period,pending_modified_values_multi_az,pending_modified_values_engine_version,pending_modified_values_license_model,pending_modified_values_iops,pending_modified_values_db_instance_identifier,pending_modified_values_storage_type,pending_modified_values_ca_certificate_identifier,pending_modified_values_db_subnet_group_name,latest_restorable_time,multi_az,engine_version,auto_minor_version_upgrade,read_replica_source_db_instance_identifier,license_model,iops,character_set_name,secondary_availability_zone,publicly_accessible,storage_type,tde_credential_arn,db_instance_port,db_cluster_identifier,storage_encrypted,kms_key_id,db_i_resource_id,ca_certificate_identifier,copy_tags_to_snapshot,monitoring_interval,enhanced_monitoring_resource_arn,monitoring_role_arn,promotion_tier,db_instance_arn,timezone,iam_database_authentication_enabled,performance_insights_enabled,performance_insights_kms_key_id,performance_insights_retention_period,deletion_protection,listener_endpoint_address,listener_endpoint_port,listener_endpoint_hosted_zone_id,max_allocated_storage) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26,$27,$28,$29,$30,$31,$32,$33,$34,$35,$36,$37,$38,$39,$40,$41,$42,$43,$44,$45,$46,$47,$48,$49,$50,$51,$52,$53,$54,$55,$56,$57,$58,$59,$60,$61,$62) returning id,db_instance_identifier,db_instance_class,engine,db_instance_status,master_username,db_name,endpoint_address,endpoint_port,endpoint_hosted_zone_id,allocated_storage,instance_create_time,preferred_backup_window,backup_retention_period,availability_zone,preferred_maintenance_window,pending_modified_values_db_instance_class,pending_modified_values_allocated_storage,pending_modified_values_master_user_password,pending_modified_values_port,pending_modified_values_backup_retention_period,pending_modified_values_multi_az,pending_modified_values_engine_version,pending_modified_values_license_model,pending_modified_values_iops,pending_modified_values_db_instance_identifier,pending_modified_values_storage_type,pending_modified_values_ca_certificate_identifier,pending_modified_values_db_subnet_group_name,latest_restorable_time,multi_az,engine_version,auto_minor_version_upgrade,read_replica_source_db_instance_identifier,license_model,iops,character_set_name,secondary_availability_zone,publicly_accessible,storage_type,tde_credential_arn,db_instance_port,db_cluster_identifier,storage_encrypted,kms_key_id,db_i_resource_id,ca_certificate_identifier,copy_tags_to_snapshot,monitoring_interval,enhanced_monitoring_resource_arn,monitoring_role_arn,promotion_tier,db_instance_arn,timezone,iam_database_authentication_enabled,performance_insights_enabled,performance_insights_kms_key_id,performance_insights_retention_period,deletion_protection,listener_endpoint_address,listener_endpoint_port,listener_endpoint_hosted_zone_id,max_allocated_storage",
		DBInstanceIdentifier, DBInstanceClass, Engine, DBInstanceStatus, MasterUsername, DBName, EndpointAddress, EndpointPort, EndpointHostedZoneId, EllocatedStorage,
		InstanceCreateTime, PreferredBackupWindow, BackupRetentionPeriod, AvailabilityZone, PreferredMaintenanceWindow, PendingModifiedValuesDBInstanceClass, PendingModifiedValuesAllocatedStorage,
		PendingModifiedValuesMasterUserPassword, PendingModifiedValuesPort, PendingModifiedValuesBackupRetentionPeriod, PendingModifiedValuesMultiAz, PendingModifiedValuesEngineVersion,
		PendingModifiedValuesLicenseModel, PendingModifiedValuesIops, PendingModifiedValuesDBInstanceIdentifier, PendingModifiedValuesStorageType, PendingModifiedValuesCaCertificateIdentifier,
		PendingModifiedValuesDBSubnetGroupName, LatestRestorableTime, MultiAz, EngineVersion, AutoMinorVersionUpgrade, ReadReplicaSourceDBInstanceIdentifier, LicenseModel, Iops,
		CharacterSetName, SecondaryAvailabilityZone, PubliclyAccessible, StorageType, TDECredentialArn, DBInstancePort, DBClusterIdentifier, StorageEncrypted, KMSKeyId,
		DBiResourceId, CaCertificateIdentifier, CopyTagsToSnapshot, MonitoringInterval, EnhancedMonitoringResourceArn, MonitoringRoleArn, PromotionTier, DBInstanceArn, Timezone,
		AamDatabaseAuthenticationEnabled, PerformanceInsightsEnabled, PerformanceInsightsKMSKeyId, PerformanceInsightsRetentionPeriod, DeletionProtection, ListenerEndpointAddress, ListenerEndpointPort,
		ListenerEndpointHostedZoneId, MaxAllocatedStorage)

	dbInstance, err := rowResultSetToDBInstance(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return dbInstance
}

func loadAllDBIstance(db *sql.DB) ([]*DBInstance, error) {
	rows, err := db.Query("select id,db_instance_identifier,db_instance_class,engine,db_instance_status,master_username,db_name,endpoint_address,endpoint_port,endpoint_hosted_zone_id,allocated_storage,instance_create_time,preferred_backup_window,backup_retention_period,availability_zone,preferred_maintenance_window,pending_modified_values_db_instance_class,pending_modified_values_allocated_storage,pending_modified_values_master_user_password,pending_modified_values_port,pending_modified_values_backup_retention_period,pending_modified_values_multi_az,pending_modified_values_engine_version,pending_modified_values_license_model,pending_modified_values_iops,pending_modified_values_db_instance_identifier,pending_modified_values_storage_type,pending_modified_values_ca_certificate_identifier,pending_modified_values_db_subnet_group_name,latest_restorable_time,multi_az,engine_version,auto_minor_version_upgrade,read_replica_source_db_instance_identifier,license_model,iops,character_set_name,secondary_availability_zone,publicly_accessible,storage_type,tde_credential_arn,db_instance_port,db_cluster_identifier,storage_encrypted,kms_key_id,db_i_resource_id,ca_certificate_identifier,copy_tags_to_snapshot,monitoring_interval,enhanced_monitoring_resource_arn,monitoring_role_arn,promotion_tier,db_instance_arn,timezone,iam_database_authentication_enabled,performance_insights_enabled,performance_insights_kms_key_id,performance_insights_retention_period,deletion_protection,listener_endpoint_address,listener_endpoint_port,listener_endpoint_hosted_zone_id,max_allocated_storage from db_instance order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfDBInstance := make([]*DBInstance, 0, 0)
	for rows.Next() {
		dbInstance, err := rowsNoFetchResultSetToDBInstance(rows)
		if err == nil {
			arrayOfDBInstance = append(arrayOfDBInstance, dbInstance)
		} else {
			log.Println(err)
		}
	}

	return arrayOfDBInstance, nil
}

func loadAllDBInstanceByOrderUnique(db *sql.DB) ([]*DBInstance, error) {
	rows, err := db.Query("select id,db_instance_identifier,db_instance_class,engine,db_instance_status,master_username,db_name,endpoint_address,endpoint_port,endpoint_hosted_zone_id,allocated_storage,instance_create_time,preferred_backup_window,backup_retention_period,availability_zone,preferred_maintenance_window,pending_modified_values_db_instance_class,pending_modified_values_allocated_storage,pending_modified_values_master_user_password,pending_modified_values_port,pending_modified_values_backup_retention_period,pending_modified_values_multi_az,pending_modified_values_engine_version,pending_modified_values_license_model,pending_modified_values_iops,pending_modified_values_db_instance_identifier,pending_modified_values_storage_type,pending_modified_values_ca_certificate_identifier,pending_modified_values_db_subnet_group_name,latest_restorable_time,multi_az,engine_version,auto_minor_version_upgrade,read_replica_source_db_instance_identifier,license_model,iops,character_set_name,secondary_availability_zone,publicly_accessible,storage_type,tde_credential_arn,db_instance_port,db_cluster_identifier,storage_encrypted,kms_key_id,db_i_resource_id,ca_certificate_identifier,copy_tags_to_snapshot,monitoring_interval,enhanced_monitoring_resource_arn,monitoring_role_arn,promotion_tier,db_instance_arn,timezone,iam_database_authentication_enabled,performance_insights_enabled,performance_insights_kms_key_id,performance_insights_retention_period,deletion_protection,listener_endpoint_address,listener_endpoint_port,listener_endpoint_hosted_zone_id,max_allocated_storage from db_instance order by db_instance_identifier")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfDBInstance := make([]*DBInstance, 0, 0)
	for rows.Next() {
		dbInstance, err := rowsNoFetchResultSetToDBInstance(rows)
		if err == nil {
			arrayOfDBInstance = append(arrayOfDBInstance, dbInstance)
		} else {
			log.Println(err)
		}
	}

	return arrayOfDBInstance, nil
}

func deleteAllDBInstance(db *sql.DB) error {
	_, err := db.Exec("delete from db_instance")

	return err
}
