package main

import (
	"database/sql" // package SQL
	"fmt"
	"log"

	_ "github.com/lib/pq" // driver Postgres
)

//////////////////////////////////////////////////////////
//RowResult
//////////////////////////////////////////////////////////

func rowResultSetToInstanceProductCode(row *sql.Row) (*InstanceProductCode, error) {
	var err error
	var Id int64
	var InstanceId int64
	var ProductCodeId string
	var ProductCodeType string

	err = row.Scan(&Id, &InstanceId, &ProductCodeId, &ProductCodeType)
	if err != nil {
		return nil, err
	}
	return NewInstanceProductCode(Id, InstanceId, ProductCodeId, ProductCodeType), nil
}

func rowResultSetToInstanceBlockDeviceMapping(row *sql.Row) (*InstanceBlockDeviceMapping, error) {
	var err error
	var Id int64
	var InstanceId int64
	var DeviceName string
	var EbsAttachTime string
	var EbsDeleteOnTermination bool
	var EbsStatus string
	var EbsVolumeId string
	var RefVolumeId sql.NullInt64

	err = row.Scan(&Id, &InstanceId, &DeviceName, &EbsAttachTime, &EbsDeleteOnTermination, &EbsStatus, &EbsVolumeId, &RefVolumeId)
	if err != nil {
		return nil, err
	}
	return NewInstanceBlockDeviceMapping(Id, InstanceId, DeviceName, EbsAttachTime, EbsDeleteOnTermination, EbsStatus, EbsVolumeId, RefVolumeId), nil
}

func rowResultSetToInstanceElasticGpuAssociation(row *sql.Row) (*InstanceElasticGpuAssociation, error) {
	var err error
	var Id int64
	var InstanceId int64
	var ElasticGpuId string
	var ElasticGpuAssociationId string
	var ElasticGpuAssociationState string
	var ElasticGpuAssociationTime string

	err = row.Scan(&Id, &InstanceId, &ElasticGpuId, &ElasticGpuAssociationId, &ElasticGpuAssociationState, &ElasticGpuAssociationTime)
	if err != nil {
		return nil, err
	}
	return NewInstanceElasticGpuAssociation(Id, InstanceId, ElasticGpuId, ElasticGpuAssociationId, ElasticGpuAssociationState, ElasticGpuAssociationTime), nil
}

func rowResultSetToInstanceElasticInferenceAcceleratorAssociation(row *sql.Row) (*InstanceElasticInferenceAcceleratorAssociation, error) {
	var err error
	var Id int64
	var InstanceId int64
	var ElasticInferenceAcceleratorArn string
	var ElasticInferenceAcceleratorAssociationId string
	var ElasticInferenceAcceleratorAssociationState string
	var ElasticInferenceAcceleratorAssociationTime string

	err = row.Scan(&Id, &InstanceId, &ElasticInferenceAcceleratorArn, &ElasticInferenceAcceleratorAssociationId, &ElasticInferenceAcceleratorAssociationState, &ElasticInferenceAcceleratorAssociationTime)
	if err != nil {
		return nil, err
	}
	return NewInstanceElasticInferenceAcceleratorAssociation(Id, InstanceId, ElasticInferenceAcceleratorArn, ElasticInferenceAcceleratorAssociationId, ElasticInferenceAcceleratorAssociationState, ElasticInferenceAcceleratorAssociationTime), nil
}

func rowResultSetToInstanceNetworkInterface(row *sql.Row) (*InstanceNetworkInterface, error) {
	var err error
	var Id int64
	var InstanceId int64
	var AssociationIpOwnerId string
	var AssociationPublicDnsName string
	var AssociationPublicIp string
	var AttachmentAttachTime string
	var AttachmentAttachmentId string
	var AttachmentDeleteOnTermination bool
	var AttachmentDeviceIndex int16
	var AttachmentStatus string
	var Description string
	var MacAddress string
	var NetworkInterfaceId string
	var OwnerId string
	var PrivateDnsName string
	var PrivateIpAddress string
	var SourceDestCheck bool
	var Status string
	var SubnetId string
	var RefSubnetId sql.NullInt64
	var VpcId string
	var RefVpcId sql.NullInt64
	var InterfaceType string

	err = row.Scan(&Id, &InstanceId, &AssociationIpOwnerId, &AssociationPublicDnsName, &AssociationPublicIp, &AttachmentAttachTime, &AttachmentAttachmentId, &AttachmentDeleteOnTermination, &AttachmentDeviceIndex, &AttachmentStatus, &Description, &MacAddress, &NetworkInterfaceId, &OwnerId, &PrivateDnsName, &PrivateIpAddress, &SourceDestCheck, &Status, &SubnetId, &RefSubnetId, &VpcId, &RefVpcId, &InterfaceType)
	if err != nil {
		return nil, err
	}
	return NewInstanceNetworkInterface(Id, InstanceId, AssociationIpOwnerId, AssociationPublicDnsName, AssociationPublicIp, AttachmentAttachTime, AttachmentAttachmentId, AttachmentDeleteOnTermination, AttachmentDeviceIndex, AttachmentStatus, Description, MacAddress, NetworkInterfaceId, OwnerId, PrivateDnsName, PrivateIpAddress, SourceDestCheck, Status, SubnetId, RefSubnetId, VpcId, RefVpcId, InterfaceType), nil
}

func rowResultSetToInstanceNetworkInterfaceGroup(row *sql.Row) (*InstanceNetworkInterfaceGroup, error) {
	var err error
	var Id int64
	var InstanceIdNetworkInterfaceId int64
	var GroupName string
	var GroupId string
	var RefSecurityGroupId sql.NullInt64

	err = row.Scan(&Id, &InstanceIdNetworkInterfaceId, &GroupName, &GroupId, &RefSecurityGroupId)
	if err != nil {
		return nil, err
	}
	return NewInstanceNetworkInterfaceGroup(Id, InstanceIdNetworkInterfaceId, GroupName, GroupId, RefSecurityGroupId), nil
}

func rowResultSetToInstanceNetworkInterfacePrivateIpAddress(row *sql.Row) (*InstanceNetworkInterfacePrivateIpAddress, error) {
	var err error
	var Id int64
	var InstanceIdNetworkInterfaceId int64
	var AssociationIpOwnerId string
	var AssociationPublicDnsName string
	var AssociationPublicIp string
	var Primary bool
	var PrivateDnsName string
	var PrivateIpAddress string

	err = row.Scan(&Id, &InstanceIdNetworkInterfaceId, &AssociationIpOwnerId, &AssociationPublicDnsName, &AssociationPublicIp, &Primary, &PrivateDnsName, &PrivateIpAddress)
	if err != nil {
		return nil, err
	}
	return NewInstanceNetworkInterfacePrivateIpAddress(Id, InstanceIdNetworkInterfaceId, AssociationIpOwnerId, AssociationPublicDnsName, AssociationPublicIp, Primary, PrivateDnsName, PrivateIpAddress), nil
}

func rowResultSetToInstanceSecurityGroup(row *sql.Row) (*InstanceSecurityGroup, error) {
	var err error
	var Id int64
	var InstanceId int64
	var GroupName string
	var GroupId string
	var RefSecurityGroupId sql.NullInt64

	err = row.Scan(&Id, &InstanceId, &GroupName, &GroupId, &RefSecurityGroupId)
	if err != nil {
		return nil, err
	}
	return NewInstanceSecurityGroup(Id, InstanceId, GroupName, GroupId, RefSecurityGroupId), nil
}

func rowResultSetToInstanceTag(row *sql.Row) (*InstanceTag, error) {
	var err error
	var Id int64
	var InstanceId int64
	var Key string
	var Value string

	err = row.Scan(&Id, &InstanceId, &Key, &Value)
	if err != nil {
		return nil, err
	}
	return NewInstanceTag(Id, InstanceId, Key, Value), nil
}

func rowResultSetToInstanceLicense(row *sql.Row) (*InstanceLicense, error) {
	var err error
	var Id int64
	var InstanceId int64
	var LicenseConfigurationArn string

	err = row.Scan(&Id, &InstanceId, &LicenseConfigurationArn)
	if err != nil {
		return nil, err
	}
	return NewInstanceLicense(Id, InstanceId, LicenseConfigurationArn), nil
}

//////////////////////////////////////////////////////////
//RowsNoFetch
//////////////////////////////////////////////////////////

func rowsNoFetchResultSetToInstanceProductCode(rows *sql.Rows) (*InstanceProductCode, error) {
	var err error
	var Id int64
	var InstanceId int64
	var ProductCodeId string
	var ProductCodeType string

	err = rows.Scan(&Id, &InstanceId, &ProductCodeId, &ProductCodeType)
	if err != nil {
		return nil, err
	}
	return NewInstanceProductCode(Id, InstanceId, ProductCodeId, ProductCodeType), nil
}

func rowsNoFetchResultSetToInstanceBlockDeviceMapping(rows *sql.Rows) (*InstanceBlockDeviceMapping, error) {
	var err error
	var Id int64
	var InstanceId int64
	var DeviceName string
	var EbsAttachTime string
	var EbsDeleteOnTermination bool
	var EbsStatus string
	var EbsVolumeId string
	var RefVolumeId sql.NullInt64

	err = rows.Scan(&Id, &InstanceId, &DeviceName, &EbsAttachTime, &EbsDeleteOnTermination, &EbsStatus, &EbsVolumeId, &RefVolumeId)
	if err != nil {
		return nil, err
	}
	return NewInstanceBlockDeviceMapping(Id, InstanceId, DeviceName, EbsAttachTime, EbsDeleteOnTermination, EbsStatus, EbsVolumeId, RefVolumeId), nil
}

func rowsNoFetchResultSetToInstanceElasticGpuAssociation(rows *sql.Rows) (*InstanceElasticGpuAssociation, error) {
	var err error
	var Id int64
	var InstanceId int64
	var ElasticGpuId string
	var ElasticGpuAssociationId string
	var ElasticGpuAssociationState string
	var ElasticGpuAssociationTime string

	err = rows.Scan(&Id, &InstanceId, &ElasticGpuId, &ElasticGpuAssociationId, &ElasticGpuAssociationState, &ElasticGpuAssociationTime)
	if err != nil {
		return nil, err
	}
	return NewInstanceElasticGpuAssociation(Id, InstanceId, ElasticGpuId, ElasticGpuAssociationId, ElasticGpuAssociationState, ElasticGpuAssociationTime), nil
}

func rowsNoFetchResultSetToInstanceElasticInferenceAcceleratorAssociation(rows *sql.Rows) (*InstanceElasticInferenceAcceleratorAssociation, error) {
	var err error
	var Id int64
	var InstanceId int64
	var ElasticInferenceAcceleratorArn string
	var ElasticInferenceAcceleratorAssociationId string
	var ElasticInferenceAcceleratorAssociationState string
	var ElasticInferenceAcceleratorAssociationTime string

	err = rows.Scan(&Id, &InstanceId, &ElasticInferenceAcceleratorArn, &ElasticInferenceAcceleratorAssociationId, &ElasticInferenceAcceleratorAssociationState, &ElasticInferenceAcceleratorAssociationTime)
	if err != nil {
		return nil, err
	}
	return NewInstanceElasticInferenceAcceleratorAssociation(Id, InstanceId, ElasticInferenceAcceleratorArn, ElasticInferenceAcceleratorAssociationId, ElasticInferenceAcceleratorAssociationState, ElasticInferenceAcceleratorAssociationTime), nil
}

func rowsNoFetchResultSetToInstanceNetworkInterface(rows *sql.Rows) (*InstanceNetworkInterface, error) {
	var err error
	var Id int64
	var InstanceId int64
	var AssociationIpOwnerId string
	var AssociationPublicDnsName string
	var AssociationPublicIp string
	var AttachmentAttachTime string
	var AttachmentAttachmentId string
	var AttachmentDeleteOnTermination bool
	var AttachmentDeviceIndex int16
	var AttachmentStatus string
	var Description string
	var MacAddress string
	var NetworkInterfaceId string
	var OwnerId string
	var PrivateDnsName string
	var PrivateIpAddress string
	var SourceDestCheck bool
	var Status string
	var SubnetId string
	var RefSubnetId sql.NullInt64
	var VpcId string
	var RefVpcId sql.NullInt64
	var InterfaceType string

	err = rows.Scan(&Id, &InstanceId, &AssociationIpOwnerId, &AssociationPublicDnsName, &AssociationPublicIp, &AttachmentAttachTime, &AttachmentAttachmentId, &AttachmentDeleteOnTermination, &AttachmentDeviceIndex, &AttachmentStatus, &Description, &MacAddress, &NetworkInterfaceId, &OwnerId, &PrivateDnsName, &PrivateIpAddress, &SourceDestCheck, &Status, &SubnetId, &RefSubnetId, &VpcId, &RefVpcId, &InterfaceType)
	if err != nil {
		return nil, err
	}
	return NewInstanceNetworkInterface(Id, InstanceId, AssociationIpOwnerId, AssociationPublicDnsName, AssociationPublicIp, AttachmentAttachTime, AttachmentAttachmentId, AttachmentDeleteOnTermination, AttachmentDeviceIndex, AttachmentStatus, Description, MacAddress, NetworkInterfaceId, OwnerId, PrivateDnsName, PrivateIpAddress, SourceDestCheck, Status, SubnetId, RefSubnetId, VpcId, RefVpcId, InterfaceType), nil
}

func rowsNoFetchResultSetToInstanceNetworkInterfaceGroup(rows *sql.Rows) (*InstanceNetworkInterfaceGroup, error) {
	var err error
	var Id int64
	var InstanceIdNetworkInterfaceId int64
	var GroupName string
	var GroupId string
	var RefSecurityGroupId sql.NullInt64

	err = rows.Scan(&Id, &InstanceIdNetworkInterfaceId, &GroupName, &GroupId, &RefSecurityGroupId)
	if err != nil {
		return nil, err
	}
	return NewInstanceNetworkInterfaceGroup(Id, InstanceIdNetworkInterfaceId, GroupName, GroupId, RefSecurityGroupId), nil
}

func rowsNoFetchResultSetToInstanceNetworkInterfacePrivateIpAddress(rows *sql.Rows) (*InstanceNetworkInterfacePrivateIpAddress, error) {
	var err error
	var Id int64
	var InstanceIdNetworkInterfaceId int64
	var AssociationIpOwnerId string
	var AssociationPublicDnsName string
	var AssociationPublicIp string
	var Primary bool
	var PrivateDnsName string
	var PrivateIpAddress string

	err = rows.Scan(&Id, &InstanceIdNetworkInterfaceId, &AssociationIpOwnerId, &AssociationPublicDnsName, &AssociationPublicIp, &Primary, &PrivateDnsName, &PrivateIpAddress)
	if err != nil {
		return nil, err
	}
	return NewInstanceNetworkInterfacePrivateIpAddress(Id, InstanceIdNetworkInterfaceId, AssociationIpOwnerId, AssociationPublicDnsName, AssociationPublicIp, Primary, PrivateDnsName, PrivateIpAddress), nil
}

func rowsNoFetchResultSetToInstanceSecurityGroup(rows *sql.Rows) (*InstanceSecurityGroup, error) {
	var err error
	var Id int64
	var InstanceId int64
	var GroupName string
	var GroupId string
	var RefSecurityGroupId sql.NullInt64

	err = rows.Scan(&Id, &InstanceId, &GroupName, &GroupId, &RefSecurityGroupId)
	if err != nil {
		return nil, err
	}
	return NewInstanceSecurityGroup(Id, InstanceId, GroupName, GroupId, RefSecurityGroupId), nil
}

func rowsNoFetchResultSetToInstanceTag(rows *sql.Rows) (*InstanceTag, error) {
	var err error
	var Id int64
	var InstanceId int64
	var Key string
	var Value string

	err = rows.Scan(&Id, &InstanceId, &Key, &Value)
	if err != nil {
		return nil, err
	}
	return NewInstanceTag(Id, InstanceId, Key, Value), nil
}

func rowsNoFetchResultSetToInstanceLicense(rows *sql.Rows) (*InstanceLicense, error) {
	var err error
	var Id int64
	var InstanceId int64
	var LicenseConfigurationArn string

	err = rows.Scan(&Id, &InstanceId, &LicenseConfigurationArn)
	if err != nil {
		return nil, err
	}
	return NewInstanceLicense(Id, InstanceId, LicenseConfigurationArn), nil
}

//////////////////////////////////////////////////////////
//RowsResult
//////////////////////////////////////////////////////////

func rowsResultSetToInstanceProductCode(rows *sql.Rows) (*InstanceProductCode, error) {
	var err error
	var Id int64
	var InstanceId int64
	var ProductCodeId string
	var ProductCodeType string

	err = rows.Scan(&Id, &InstanceId, &ProductCodeId, &ProductCodeType)
	if err != nil {
		return nil, err
	}
	return NewInstanceProductCode(Id, InstanceId, ProductCodeId, ProductCodeType), nil
}

func rowsResultSetToInstanceBlockDeviceMapping(rows *sql.Rows) (*InstanceBlockDeviceMapping, error) {
	var err error
	var Id int64
	var InstanceId int64
	var DeviceName string
	var EbsAttachTime string
	var EbsDeleteOnTermination bool
	var EbsStatus string
	var EbsVolumeId string
	var RefVolumeId sql.NullInt64

	err = rows.Scan(&Id, &InstanceId, &DeviceName, &EbsAttachTime, &EbsDeleteOnTermination, &EbsStatus, &EbsVolumeId, &RefVolumeId)
	if err != nil {
		return nil, err
	}
	return NewInstanceBlockDeviceMapping(Id, InstanceId, DeviceName, EbsAttachTime, EbsDeleteOnTermination, EbsStatus, EbsVolumeId, RefVolumeId), nil
}

func rowsResultSetToInstanceElasticGpuAssociation(rows *sql.Rows) (*InstanceElasticGpuAssociation, error) {
	var err error
	var Id int64
	var InstanceId int64
	var ElasticGpuId string
	var ElasticGpuAssociationId string
	var ElasticGpuAssociationState string
	var ElasticGpuAssociationTime string

	err = rows.Scan(&Id, &InstanceId, &ElasticGpuId, &ElasticGpuAssociationId, &ElasticGpuAssociationState, &ElasticGpuAssociationTime)
	if err != nil {
		return nil, err
	}
	return NewInstanceElasticGpuAssociation(Id, InstanceId, ElasticGpuId, ElasticGpuAssociationId, ElasticGpuAssociationState, ElasticGpuAssociationTime), nil
}

func rowsResultSetToInstanceElasticInferenceAcceleratorAssociation(rows *sql.Rows) (*InstanceElasticInferenceAcceleratorAssociation, error) {
	var err error
	var Id int64
	var InstanceId int64
	var ElasticInferenceAcceleratorArn string
	var ElasticInferenceAcceleratorAssociationId string
	var ElasticInferenceAcceleratorAssociationState string
	var ElasticInferenceAcceleratorAssociationTime string

	err = rows.Scan(&Id, &InstanceId, &ElasticInferenceAcceleratorArn, &ElasticInferenceAcceleratorAssociationId, &ElasticInferenceAcceleratorAssociationState, &ElasticInferenceAcceleratorAssociationTime)
	if err != nil {
		return nil, err
	}
	return NewInstanceElasticInferenceAcceleratorAssociation(Id, InstanceId, ElasticInferenceAcceleratorArn, ElasticInferenceAcceleratorAssociationId, ElasticInferenceAcceleratorAssociationState, ElasticInferenceAcceleratorAssociationTime), nil
}

func rowsResultSetToInstanceNetworkInterface(rows *sql.Rows) (*InstanceNetworkInterface, error) {
	var err error
	var Id int64
	var InstanceId int64
	var AssociationIpOwnerId string
	var AssociationPublicDnsName string
	var AssociationPublicIp string
	var AttachmentAttachTime string
	var AttachmentAttachmentId string
	var AttachmentDeleteOnTermination bool
	var AttachmentDeviceIndex int16
	var AttachmentStatus string
	var Description string
	var MacAddress string
	var NetworkInterfaceId string
	var OwnerId string
	var PrivateDnsName string
	var PrivateIpAddress string
	var SourceDestCheck bool
	var Status string
	var SubnetId string
	var RefSubnetId sql.NullInt64
	var VpcId string
	var RefVpcId sql.NullInt64
	var InterfaceType string

	err = rows.Scan(&Id, &InstanceId, &AssociationIpOwnerId, &AssociationPublicDnsName, &AssociationPublicIp, &AttachmentAttachTime, &AttachmentAttachmentId, &AttachmentDeleteOnTermination, &AttachmentDeviceIndex, &AttachmentStatus, &Description, &MacAddress, &NetworkInterfaceId, &OwnerId, &PrivateDnsName, &PrivateIpAddress, &SourceDestCheck, &Status, &SubnetId, &RefSubnetId, &VpcId, RefVpcId, &InterfaceType)
	if err != nil {
		return nil, err
	}
	return NewInstanceNetworkInterface(Id, InstanceId, AssociationIpOwnerId, AssociationPublicDnsName, AssociationPublicIp, AttachmentAttachTime, AttachmentAttachmentId, AttachmentDeleteOnTermination, AttachmentDeviceIndex, AttachmentStatus, Description, MacAddress, NetworkInterfaceId, OwnerId, PrivateDnsName, PrivateIpAddress, SourceDestCheck, Status, SubnetId, RefSubnetId, VpcId, RefVpcId, InterfaceType), nil
}

func rowsResultSetToInstanceNetworkInterfaceGroup(rows *sql.Rows) (*InstanceNetworkInterfaceGroup, error) {
	var err error
	var Id int64
	var InstanceIdNetworkInterfaceId int64
	var GroupName string
	var GroupId string
	var RefSecurityGroupId sql.NullInt64

	err = rows.Scan(&Id, &InstanceIdNetworkInterfaceId, &GroupName, &GroupId, &RefSecurityGroupId)
	if err != nil {
		return nil, err
	}
	return NewInstanceNetworkInterfaceGroup(Id, InstanceIdNetworkInterfaceId, GroupName, GroupId, RefSecurityGroupId), nil
}

func rowsResultSetToInstanceNetworkInterfacePrivateIpAddress(rows *sql.Rows) (*InstanceNetworkInterfacePrivateIpAddress, error) {
	var err error
	var Id int64
	var InstanceIdNetworkInterfaceId int64
	var AssociationIpOwnerId string
	var AssociationPublicDnsName string
	var AssociationPublicIp string
	var Primary bool
	var PrivateDnsName string
	var PrivateIpAddress string

	err = rows.Scan(&Id, &InstanceIdNetworkInterfaceId, &AssociationIpOwnerId, &AssociationPublicDnsName, &AssociationPublicIp, &Primary, &PrivateDnsName, &PrivateIpAddress)
	if err != nil {
		return nil, err
	}
	return NewInstanceNetworkInterfacePrivateIpAddress(Id, InstanceIdNetworkInterfaceId, AssociationIpOwnerId, AssociationPublicDnsName, AssociationPublicIp, Primary, PrivateDnsName, PrivateIpAddress), nil
}

func rowsResultSetToInstanceSecurityGroup(rows *sql.Rows) (*InstanceSecurityGroup, error) {
	var err error
	var Id int64
	var InstanceId int64
	var GroupName string
	var GroupId string
	var RefSecurityGroupId sql.NullInt64

	err = rows.Scan(&Id, &InstanceId, &GroupName, &GroupId, &RefSecurityGroupId)
	if err != nil {
		return nil, err
	}
	return NewInstanceSecurityGroup(Id, InstanceId, GroupName, GroupId, RefSecurityGroupId), nil
}

func rowsResultSetToInstanceTag(rows *sql.Rows) (*InstanceTag, error) {
	var err error
	var Id int64
	var InstanceId int64
	var Key string
	var Value string

	err = rows.Scan(&Id, &InstanceId, &Key, &Value)
	if err != nil {
		return nil, err
	}
	return NewInstanceTag(Id, InstanceId, Key, Value), nil
}

func rowsResultSetToInstanceLicense(rows *sql.Rows) (*InstanceLicense, error) {
	var err error
	var Id int64
	var InstanceId int64
	var LicenseConfigurationArn string

	err = rows.Scan(&Id, &InstanceId, &LicenseConfigurationArn)
	if err != nil {
		return nil, err
	}
	return NewInstanceLicense(Id, InstanceId, LicenseConfigurationArn), nil
}

//////////////////////////////////////////////////////////
//Create
//////////////////////////////////////////////////////////

func createInstanceProductCode(db *sql.DB, InstanceId int64, ProductCodeId string, ProductCodeType string) *InstanceProductCode {
	rows := db.QueryRow("insert into instance_product_code(instance_id,product_code_id,product_code_type) values($1,$2,$3) returning id,instance_id,product_code_id,product_code_type",
		InstanceId, ProductCodeId, ProductCodeType)

	instanceProductCode, err := rowResultSetToInstanceProductCode(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return instanceProductCode
}

func createInstanceBlockDeviceMapping(db *sql.DB, InstanceId int64, DeviceName string, EbsAttachTime string, EbsDeleteOnTermination bool, EbsStatus string, EbsVolumeId string) *InstanceBlockDeviceMapping {
	rows := db.QueryRow("insert into instance_block_device_mapping(instance_id,device_name,ebs_attach_time,ebs_delete_on_termination,ebs_status,ebs_volume_id) values($1,$2,$3,$4,$5,$6) returning id,instance_id,device_name,ebs_attach_time,ebs_delete_on_termination,ebs_status,ebs_volume_id,ref_volume_id",
		InstanceId, DeviceName, EbsAttachTime, EbsDeleteOnTermination, EbsStatus, EbsVolumeId)

	instanceBlockDeviceMapping, err := rowResultSetToInstanceBlockDeviceMapping(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return instanceBlockDeviceMapping
}
func createInstanceElasticGpuAssociation(db *sql.DB, InstanceId int64, ElasticGpuId string, ElasticGpuAssociationId string, ElasticGpuAssociationState string, ElasticGpuAssociationTime string) *InstanceElasticGpuAssociation {
	rows := db.QueryRow("insert into instance_elastic_gpu_association(instance_id,elastic_gpu_id,elastic_gpu_association_id,elastic_gpu_association_state,elastic_gpu_association_time) values($1,$2,$3,$4,$5,$6) returning id,instance_id,elastic_gpu_id,elastic_gpu_association_id,elastic_gpu_association_state,elastic_gpu_association_time",
		InstanceId, ElasticGpuId, ElasticGpuAssociationId, ElasticGpuAssociationState, ElasticGpuAssociationTime)

	instanceElasticGpuAssociation, err := rowResultSetToInstanceElasticGpuAssociation(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return instanceElasticGpuAssociation
}
func createInstanceElasticInferenceAcceleratorAssociation(db *sql.DB, InstanceId int64, ElasticInferenceAcceleratorArn string, ElasticInferenceAcceleratorAssociationId string, ElasticInferenceAcceleratorAssociationState string, ElasticInferenceAcceleratorAssociationTime string) *InstanceElasticInferenceAcceleratorAssociation {
	rows := db.QueryRow("insert into instance_elastic_inference_accelerator_association(instance_id,elasticInference_accelerator_arn,elasticInference_accelerator_association_id,elasticInference_accelerator_association_state,elasticInference_accelerator_association_time) values($1,$2,$3,$4,$5) returning id,instance_id,elasticInference_accelerator_arn,elasticInference_accelerator_association_id,elasticInference_accelerator_association_state,elasticInference_accelerator_association_time",
		InstanceId, ElasticInferenceAcceleratorArn, ElasticInferenceAcceleratorAssociationId, ElasticInferenceAcceleratorAssociationState, ElasticInferenceAcceleratorAssociationTime)

	instanceElasticInferenceAcceleratorAssociation, err := rowResultSetToInstanceElasticInferenceAcceleratorAssociation(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return instanceElasticInferenceAcceleratorAssociation
}
func createInstanceNetworkInterface(db *sql.DB, InstanceId int64, AssociationIpOwnerId string, AssociationPublicDnsName string, AssociationPublicIp string, AttachmentAttachTime string, AttachmentAttachmentId string, AttachmentDeleteOnTermination bool, AttachmentDeviceIndex int16, AttachmentStatus string, Description string, MacAddress string, NetworkInterfaceId string, OwnerId string, PrivateDnsName string, PrivateIpAddress string, SourceDestCheck bool, Status string, SubnetId string, VpcId string, InterfaceType string) *InstanceNetworkInterface {
	rows := db.QueryRow("insert into instance_network_interface(instance_id,association_ip_owner_id,association_public_dns_name,association_public_ip,attachment_attach_time,attachment_attachment_id,attachment_delete_on_termination,attachment_device_index,attachment_status,description,mac_address,network_interface_id,owner_id,private_dns_name,private_ip_address,source_dest_check,status,subnet_id,vpc_id,interface_type) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20) returning id,instance_id,association_ip_owner_id,association_public_dns_name,association_public_ip,attachment_attach_time,attachment_attachment_id,attachment_delete_on_termination,attachment_device_index,attachment_status,description,mac_address,network_interface_id,owner_id,private_dns_name,private_ip_address,source_dest_check,status,subnet_id,ref_subnet_id,vpc_id,ref_vpc_id,interface_type",
		InstanceId, AssociationIpOwnerId, AssociationPublicDnsName, AssociationPublicIp, AttachmentAttachTime, AttachmentAttachmentId, AttachmentDeleteOnTermination, AttachmentDeviceIndex, AttachmentStatus, Description, MacAddress, NetworkInterfaceId, OwnerId, PrivateDnsName, PrivateIpAddress, SourceDestCheck, Status, SubnetId, VpcId, InterfaceType)

	instanceNetworkInterface, err := rowResultSetToInstanceNetworkInterface(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return instanceNetworkInterface
}
func createInstanceNetworkInterfaceGroup(db *sql.DB, InstanceIdNetworkInterfaceId int64, GroupName string, GroupId string) *InstanceNetworkInterfaceGroup {
	rows := db.QueryRow("insert into instance_network_interface_group(instance_network_interface_id,group_name,group_id) values($1,$2,$3) returning id,instance_network_interface_id,group_name,group_id,ref_security_group_id",
		InstanceIdNetworkInterfaceId, GroupName, GroupId)

	instanceNetworkInterfaceGroup, err := rowResultSetToInstanceNetworkInterfaceGroup(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return instanceNetworkInterfaceGroup
}
func createInstanceNetworkInterfacePrivateIpAddress(db *sql.DB, InstanceIdNetworkInterfaceId int64, AssociationIpOwnerId string, AssociationPublicDnsName string, AssociationPublicIp string, Primary bool, PrivateDnsName string, PrivateIpAddress string) *InstanceNetworkInterfacePrivateIpAddress {
	rows := db.QueryRow("insert into instance_network_interface_private_ip_address(instance_network_interface_id,association_ip_owner_id,association_public_dns_name,association_public_ip,is_primary,private_dns_name,private_ip_address) values($1,$2,$3,$4,$5,$6,$7) returning id,instance_network_interface_id,association_ip_owner_id,association_public_dns_name,association_public_ip,is_primary,private_dns_name,private_ip_address",
		InstanceIdNetworkInterfaceId, AssociationIpOwnerId, AssociationPublicDnsName, AssociationPublicIp, Primary, PrivateDnsName, PrivateIpAddress)
	instanceNetworkInterfacePrivateIpAddress, err := rowResultSetToInstanceNetworkInterfacePrivateIpAddress(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return instanceNetworkInterfacePrivateIpAddress
}
func createInstanceSecurityGroup(db *sql.DB, InstanceId int64, GroupName string, GroupId string) *InstanceSecurityGroup {
	rows := db.QueryRow("insert into instance_security_group(instance_id,group_name,group_id) values($1,$2,$3) returning id,instance_id,group_name,group_id,ref_security_group_id",
		InstanceId, GroupName, GroupId)

	instanceSecurityGroup, err := rowResultSetToInstanceSecurityGroup(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return instanceSecurityGroup
}
func createInstanceTag(db *sql.DB, InstanceId int64, Key string, Value string) *InstanceTag {
	rows := db.QueryRow("insert into instance_tag(instance_id,key,value) values($1,$2,$3) returning id,instance_id,key,value",
		InstanceId, Key, Value)

	instanceTag, err := rowResultSetToInstanceTag(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return instanceTag
}
func createInstanceLicense(db *sql.DB, InstanceId int64, LicenseConfigurationArn string) *InstanceLicense {
	rows := db.QueryRow("insert into instance_license(instance_id,license_configuration_arn) values($1,$2) returning id,instance_id,license_configuration_arn",
		InstanceId, LicenseConfigurationArn)

	instanceLicense, err := rowResultSetToInstanceLicense(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return instanceLicense
}

//////////////////////////////////////////////////////////
//Load
//////////////////////////////////////////////////////////

func loadAllInstanceBlockDeviceMapping(db *sql.DB) ([]*InstanceBlockDeviceMapping, error) {
	rows, err := db.Query("select id,instance_id,device_name,ebs_attach_time,ebs_delete_on_termination,ebs_status,ebs_volume_id,ref_volume_id from instance_block_device_mapping order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfInstanceBlockDeviceMapping := make([]*InstanceBlockDeviceMapping, 0, 0)
	for rows.Next() {
		instanceBlockDeviceMapping, err := rowsNoFetchResultSetToInstanceBlockDeviceMapping(rows)
		if err == nil {
			arrayOfInstanceBlockDeviceMapping = append(arrayOfInstanceBlockDeviceMapping, instanceBlockDeviceMapping)
		} else {
			log.Println(err)
		}
	}

	return arrayOfInstanceBlockDeviceMapping, nil
}

func loadAllInstanceNetworkInterface(db *sql.DB) ([]*InstanceNetworkInterface, error) {
	rows, err := db.Query("select id,instance_id,association_ip_owner_id,association_public_dns_name,association_public_ip,attachment_attach_time,attachment_attachment_id,attachment_delete_on_termination,attachment_device_index,attachment_status,description,mac_address,network_interface_id,owner_id,private_dns_name,private_ip_address,source_dest_check,status,subnet_id,ref_subnet_id,vpc_id,ref_vpc_id,interface_type from instance_network_interface order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfInstanceNetworkInterface := make([]*InstanceNetworkInterface, 0, 0)
	for rows.Next() {
		instanceNetworkInterface, err := rowsNoFetchResultSetToInstanceNetworkInterface(rows)
		if err == nil {
			arrayOfInstanceNetworkInterface = append(arrayOfInstanceNetworkInterface, instanceNetworkInterface)
		} else {
			log.Println(err)
		}
	}

	return arrayOfInstanceNetworkInterface, nil
}

func loadAllInstanceNetworkInterfaceGroup(db *sql.DB) ([]*InstanceNetworkInterfaceGroup, error) {
	rows, err := db.Query("select id,instance_network_interface_id,group_name,group_id,ref_security_group_id from instance_network_interface_group order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfInstanceNetworkInterfaceGroup := make([]*InstanceNetworkInterfaceGroup, 0, 0)
	for rows.Next() {
		instanceNetworkInterfaceGroup, err := rowsNoFetchResultSetToInstanceNetworkInterfaceGroup(rows)
		if err == nil {
			arrayOfInstanceNetworkInterfaceGroup = append(arrayOfInstanceNetworkInterfaceGroup, instanceNetworkInterfaceGroup)
		} else {
			log.Println(err)
		}
	}

	return arrayOfInstanceNetworkInterfaceGroup, nil
}

func loadAllInstanceSecurityGroup(db *sql.DB) ([]*InstanceSecurityGroup, error) {
	rows, err := db.Query("select id,instance_id,group_name,group_id,ref_security_group_id from instance_security_group order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfInstanceSecurityGroup := make([]*InstanceSecurityGroup, 0, 0)
	for rows.Next() {
		instanceSecurityGroup, err := rowsNoFetchResultSetToInstanceSecurityGroup(rows)
		if err == nil {
			arrayOfInstanceSecurityGroup = append(arrayOfInstanceSecurityGroup, instanceSecurityGroup)
		} else {
			log.Println(err)
		}
	}

	return arrayOfInstanceSecurityGroup, nil
}

func updateInstanceBlockDeviceMapping(db *sql.DB, instanceBlockDeviceMapping *InstanceBlockDeviceMapping, volumeId int64) *InstanceBlockDeviceMapping {
	rows := db.QueryRow("update instance_block_device_mapping set ref_volume_id=$1 where id=$2 returning id,instance_id,device_name,ebs_attach_time,ebs_delete_on_termination,ebs_status,ebs_volume_id,ref_volume_id", volumeId, instanceBlockDeviceMapping.Id)

	instanceBlockDeviceMapping, err := rowResultSetToInstanceBlockDeviceMapping(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return instanceBlockDeviceMapping
}

func updateInstanceNetworkInterface(db *sql.DB, instanceNetworkInterface *InstanceNetworkInterface, subnetId int64, vpcId int64) *InstanceNetworkInterface {
	rows := db.QueryRow("update instance_network_interface set ref_subnet_id=$1,ref_vpc_id=$2 where id=$3 returning id,instance_id,association_ip_owner_id,association_public_dns_name,association_public_ip,attachment_attach_time,attachment_attachment_id,attachment_delete_on_termination,attachment_device_index,attachment_status,description,mac_address,network_interface_id,owner_id,private_dns_name,private_ip_address,source_dest_check,status,subnet_id,ref_subnet_id,vpc_id,ref_vpc_id,interface_type", subnetId, vpcId, instanceNetworkInterface.Id)

	instanceNetworkInterface, err := rowResultSetToInstanceNetworkInterface(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return instanceNetworkInterface
}

func updateInstanceNetworkInterfaceGroup(db *sql.DB, instanceNetworkInterfaceGroup *InstanceNetworkInterfaceGroup, secGroupId int64) *InstanceNetworkInterfaceGroup {
	rows := db.QueryRow("update instance_network_interface_group set ref_security_group_id=$1 where id=$2 returning id,instance_network_interface_id,group_name,group_id,ref_security_group_id", secGroupId, instanceNetworkInterfaceGroup.Id)

	instanceNetworkInterfaceGroup, err := rowResultSetToInstanceNetworkInterfaceGroup(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return instanceNetworkInterfaceGroup
}

func updateInstanceSecurityGroup(db *sql.DB, instanceSecurityGroup *InstanceSecurityGroup, secGroupId int64) *InstanceSecurityGroup {
	rows := db.QueryRow("update instance_security_group set ref_security_group_id=$1 where id=$2 returning id,instance_id,group_name,group_id,ref_security_group_id", secGroupId, instanceSecurityGroup.Id)

	instanceSecurityGroup, err := rowResultSetToInstanceSecurityGroup(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return instanceSecurityGroup
}

//////////////////////////////////////////////////////////
//LoadAll
//////////////////////////////////////////////////////////

func loadAllInstanceNetworkInterfaceGroupByOrderUnique(db *sql.DB) ([]*InstanceNetworkInterfaceGroup, error) {
	rows, err := db.Query("select id,instance_network_interface_id,group_name,group_id,ref_security_group_id from instance_network_interface_group order by group_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfInstanceNetworkInterfaceGroup := make([]*InstanceNetworkInterfaceGroup, 0, 0)
	for rows.Next() {
		instanceNetworkInterfaceGroup, err := rowsNoFetchResultSetToInstanceNetworkInterfaceGroup(rows)
		if err == nil {
			arrayOfInstanceNetworkInterfaceGroup = append(arrayOfInstanceNetworkInterfaceGroup, instanceNetworkInterfaceGroup)
		} else {
			log.Println(err)
		}
	}

	return arrayOfInstanceNetworkInterfaceGroup, nil
}

func loadAllInstanceNetworkInterfacePrivateIpAddressByOrderUnique(db *sql.DB) ([]*InstanceNetworkInterfacePrivateIpAddress, error) {
	rows, err := db.Query("select id,instance_network_interface_id,association_ip_owner_id,association_public_dns_name,association_public_ip,is_primary,private_dns_name,private_ip_address from instance_network_interface_private_ip_address order by association_public_ip")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfInstanceNetworkInterfacePrivateIpAddress := make([]*InstanceNetworkInterfacePrivateIpAddress, 0, 0)
	for rows.Next() {
		instanceNetworkInterfacePrivateIpAddress, err := rowsNoFetchResultSetToInstanceNetworkInterfacePrivateIpAddress(rows)
		if err == nil {
			arrayOfInstanceNetworkInterfacePrivateIpAddress = append(arrayOfInstanceNetworkInterfacePrivateIpAddress, instanceNetworkInterfacePrivateIpAddress)
		} else {
			log.Println(err)
		}
	}

	return arrayOfInstanceNetworkInterfacePrivateIpAddress, nil
}

func loadAllInstanceProductCodeByOrderUnique(db *sql.DB) ([]*InstanceProductCode, error) {
	results := make([]*InstanceProductCode, 0, 0)
	sqlSelect := "SELECT id,instance_id,product_code_id,product_code_type from instance_product_code order by product_code_id"

	rows, err := db.Query(sqlSelect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		instanceProductCode, err := rowsNoFetchResultSetToInstanceProductCode(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, instanceProductCode)
	}
	return results, nil
}

func loadAllInstanceBlockDeviceMappingByOrderUnique(db *sql.DB) ([]*InstanceBlockDeviceMapping, error) {
	rows, err := db.Query("select id,instance_id,device_name,ebs_attach_time,ebs_delete_on_termination,ebs_status,ebs_volume_id,ref_volume_id from instance_block_device_mapping order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfInstanceBlockDeviceMapping := make([]*InstanceBlockDeviceMapping, 0, 0)
	for rows.Next() {
		instanceBlockDeviceMapping, err := rowsNoFetchResultSetToInstanceBlockDeviceMapping(rows)
		if err == nil {
			arrayOfInstanceBlockDeviceMapping = append(arrayOfInstanceBlockDeviceMapping, instanceBlockDeviceMapping)
		} else {
			log.Println(err)
		}
	}

	return arrayOfInstanceBlockDeviceMapping, nil
}

func loadAllInstanceElasticGpuAssociationByOrderUnique(db *sql.DB) ([]*InstanceElasticGpuAssociation, error) {
	results := make([]*InstanceElasticGpuAssociation, 0, 0)
	sqlSelect := "SELECT id,instance_id,elastic_gpu_id,elastic_gpu_association_id,elastic_gpu_association_state,elastic_gpu_association_time from instance_elastic_gpu_association order by elastic_gpu_association_id"

	rows, err := db.Query(sqlSelect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		instanceElasticGpuAssociation, err := rowsNoFetchResultSetToInstanceElasticGpuAssociation(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, instanceElasticGpuAssociation)
	}
	return results, nil
}

func loadAllInstanceElasticInferenceAcceleratorAssociationByOrderUnique(db *sql.DB) ([]*InstanceElasticInferenceAcceleratorAssociation, error) {
	results := make([]*InstanceElasticInferenceAcceleratorAssociation, 0, 0)
	sqlSelect := "SELECT id,instance_id,elasticInference_accelerator_arn,elasticInference_accelerator_association_id,elasticInference_accelerator_association_state,elasticInference_accelerator_association_time from instance_elastic_inference_accelerator_association order by elasticInference_accelerator_association_id"

	rows, err := db.Query(sqlSelect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		instanceElasticInferenceAcceleratorAssociation, err := rowsNoFetchResultSetToInstanceElasticInferenceAcceleratorAssociation(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, instanceElasticInferenceAcceleratorAssociation)
	}
	return results, nil
}

func loadAllInstanceNetworkInterfaceByOrderUnique(db *sql.DB) ([]*InstanceNetworkInterface, error) {
	rows, err := db.Query("select id,instance_id,association_ip_owner_id,association_public_dns_name,association_public_ip,attachment_attach_time,attachment_attachment_id,attachment_delete_on_termination,attachment_device_index,attachment_status,description,mac_address,network_interface_id,owner_id,private_dns_name,private_ip_address,source_dest_check,status,subnet_id,ref_subnet_id,vpc_id,ref_vpc_id,interface_type from instance_network_interface order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfInstanceNetworkInterface := make([]*InstanceNetworkInterface, 0, 0)
	for rows.Next() {
		instanceNetworkInterface, err := rowsNoFetchResultSetToInstanceNetworkInterface(rows)
		if err == nil {
			arrayOfInstanceNetworkInterface = append(arrayOfInstanceNetworkInterface, instanceNetworkInterface)
		} else {
			log.Println(err)
		}
	}

	return arrayOfInstanceNetworkInterface, nil
}

func loadAllInstanceSecurityGroupByOrderUnique(db *sql.DB) ([]*InstanceSecurityGroup, error) {
	rows, err := db.Query("select id,instance_id,group_name,group_id,ref_security_group_id from instance_security_group order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfInstanceSecurityGroup := make([]*InstanceSecurityGroup, 0, 0)
	for rows.Next() {
		instanceSecurityGroup, err := rowsNoFetchResultSetToInstanceSecurityGroup(rows)
		if err == nil {
			arrayOfInstanceSecurityGroup = append(arrayOfInstanceSecurityGroup, instanceSecurityGroup)
		} else {
			log.Println(err)
		}
	}

	return arrayOfInstanceSecurityGroup, nil
}

func loadAllInstanceTagsByOrderUnique(db *sql.DB) ([]*InstanceTag, error) {
	results := make([]*InstanceTag, 0, 0)
	sqlSelect := "SELECT id,instance_id,key,value from instance_tag order by value"

	rows, err := db.Query(sqlSelect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		instanceTag, err := rowsNoFetchResultSetToInstanceTag(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, instanceTag)
	}
	return results, nil
}

func loadAllInstanceLicenseByOrderUnique(db *sql.DB) ([]*InstanceLicense, error) {
	results := make([]*InstanceLicense, 0, 0)
	sqlSelect := "SELECT id,instance_id,license_configuration_arn from instance_license order by license_configuration_arn"

	rows, err := db.Query(sqlSelect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		instanceLicense, err := rowsNoFetchResultSetToInstanceLicense(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, instanceLicense)
	}
	return results, nil
}

//////////////////////////////////////////////////////////
//DeleteAll
//////////////////////////////////////////////////////////

func deleteAllInstanceProductCode(db *sql.DB) error {
	_, err := db.Exec("delete from instance_product_code")

	return err
}

func deleteAllInstanceBlockDeviceMapping(db *sql.DB) error {
	_, err := db.Exec("delete from instance_block_device_mapping")

	return err
}

func deleteAllInstanceElasticGpuAssociation(db *sql.DB) error {
	_, err := db.Exec("delete from instance_elastic_gpu_association")

	return err
}

func deleteAllInstanceElasticInferenceAcceleratorAssociation(db *sql.DB) error {
	_, err := db.Exec("delete from instance_elastic_inference_accelerator_association")

	return err
}

func deleteAllInstanceNetworkInterface(db *sql.DB) error {
	_, err := db.Exec("delete from instance_network_interface")

	return err
}

func deleteAllInstanceNetworkInterfaceGroup(db *sql.DB) error {
	_, err := db.Exec("delete from instance_network_interface_group")

	return err
}

func deleteAllInstanceNetworkInterfacePrivateIpAddress(db *sql.DB) error {
	_, err := db.Exec("delete from instance_network_interface_private_ip_address")

	return err
}

func deleteAllInstanceSecurityGroup(db *sql.DB) error {
	_, err := db.Exec("delete from instance_security_group")

	return err
}

func deleteAllInstanceTag(db *sql.DB) error {
	_, err := db.Exec("delete from instance_tag")

	return err
}

func deleteAllInstanceLicense(db *sql.DB) error {
	_, err := db.Exec("delete from instance_license")

	return err
}
