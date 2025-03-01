package main

import "database/sql"
import "fmt"

type InstanceProductCode struct {
	Id              int64
	InstanceId      int64
	ProductCodeId   string
	ProductCodeType string
}

func NewInstanceProductCode(Id int64, InstanceId int64, ProductCodeId string, ProductCodeType string) *InstanceProductCode {
	return &InstanceProductCode{Id, InstanceId, ProductCodeId, ProductCodeType}
}

func (d *InstanceProductCode) Dump() string {
	return fmt.Sprintf("ProductCodeId(%s) ProductCodeType(%s)", d.ProductCodeId, d.ProductCodeId)
}

type InstanceBlockDeviceMapping struct {
	Id                     int64
	InstanceId             int64
	DeviceName             string
	EbsAttachTime          string
	EbsDeleteOnTermination bool
	EbsStatus              string
	EbsVolumeId            string
	RefVolumeId            sql.NullInt64
}

func (d *InstanceBlockDeviceMapping) Dump() string {
	return fmt.Sprintf("DeviceName(%s) EbsAttachTime(%s) EbsDeleteOnTermination(%t) EbsStatus(%s) EbsVolumeId(%s) RefVolumeId(%d)",
		d.DeviceName, d.EbsAttachTime, d.EbsDeleteOnTermination, d.EbsStatus, d.EbsVolumeId, d.RefVolumeId)
}

func NewInstanceBlockDeviceMapping(Id int64, InstanceId int64, DeviceName string, EbsAttachTime string, EbsDeleteOnTermination bool, EbsStatus string, EbsVolumeId string, RefVolumeId sql.NullInt64) *InstanceBlockDeviceMapping {
	return &InstanceBlockDeviceMapping{Id,
		InstanceId,
		DeviceName,
		EbsAttachTime,
		EbsDeleteOnTermination,
		EbsStatus,
		EbsVolumeId,
		RefVolumeId}
}

type InstanceElasticGpuAssociation struct {
	Id                         int64
	InstanceId                 int64
	ElasticGpuId               string
	ElasticGpuAssociationId    string
	ElasticGpuAssociationState string
	ElasticGpuAssociationTime  string
}

func (d *InstanceElasticGpuAssociation) Dump() string {
	return fmt.Sprintf("ElasticGpuId(%s) ElasticGpuAssociationId(%s) ElasticGpuAssociationState(%s) ElasticGpuAssociationTime(%s)",
		d.ElasticGpuId, d.ElasticGpuAssociationId, d.ElasticGpuAssociationState, d.ElasticGpuAssociationTime)
}

func NewInstanceElasticGpuAssociation(Id int64, InstanceId int64, ElasticGpuId string, ElasticGpuAssociationId string, ElasticGpuAssociationState string, ElasticGpuAssociationTime string) *InstanceElasticGpuAssociation {
	return &InstanceElasticGpuAssociation{Id,
		InstanceId,
		ElasticGpuId,
		ElasticGpuAssociationId,
		ElasticGpuAssociationState,
		ElasticGpuAssociationTime}
}

type InstanceElasticInferenceAcceleratorAssociation struct {
	Id                                          int64
	InstanceId                                  int64
	ElasticInferenceAcceleratorArn              string
	ElasticInferenceAcceleratorAssociationId    string
	ElasticInferenceAcceleratorAssociationState string
	ElasticInferenceAcceleratorAssociationTime  string
}

func (d *InstanceElasticInferenceAcceleratorAssociation) Dump() string {
	return fmt.Sprintf("ElasticInferenceAcceleratorArn(%s) ElasticInferenceAcceleratorAssociationId(%s) ElasticInferenceAcceleratorAssociationState(%s) ElasticInferenceAcceleratorAssociationTime(%s)",
		d.ElasticInferenceAcceleratorArn, d.ElasticInferenceAcceleratorAssociationId, d.ElasticInferenceAcceleratorAssociationState, d.ElasticInferenceAcceleratorAssociationTime)
}

func NewInstanceElasticInferenceAcceleratorAssociation(Id int64, InstanceId int64, ElasticInferenceAcceleratorArn string, ElasticInferenceAcceleratorAssociationId string, ElasticInferenceAcceleratorAssociationState string, ElasticInferenceAcceleratorAssociationTime string) *InstanceElasticInferenceAcceleratorAssociation {
	return &InstanceElasticInferenceAcceleratorAssociation{Id,
		InstanceId,
		ElasticInferenceAcceleratorArn,
		ElasticInferenceAcceleratorAssociationId,
		ElasticInferenceAcceleratorAssociationState,
		ElasticInferenceAcceleratorAssociationTime}
}

type InstanceNetworkInterface struct {
	Id                            int64
	InstanceId                    int64
	AssociationIpOwnerId          string
	AssociationPublicDnsName      string
	AssociationPublicIp           string
	AttachmentAttachTime          string
	AttachmentAttachmentId        string
	AttachmentDeleteOnTermination bool
	AttachmentDeviceIndex         int16
	AttachmentStatus              string
	Description                   string
	Groups                        []*InstanceNetworkInterfaceGroup
	MacAddress                    string
	NetworkInterfaceId            string
	OwnerId                       string
	PrivateDnsName                string
	PrivateIpAddress              string
	PrivateIpAddresses            []*InstanceNetworkInterfacePrivateIpAddress
	SourceDestCheck               bool
	Status                        string
	SubnetId                      string
	RefSubnetId                   sql.NullInt64
	VpcId                         string
	RefVpcId                      sql.NullInt64
	InterfaceType                 string
}

func NewInstanceNetworkInterface(Id int64, InstanceId int64, AssociationIpOwnerId string, AssociationPublicDnsName string, AssociationPublicIp string, AttachmentAttachTime string, AttachmentAttachmentId string, AttachmentDeleteOnTermination bool, AttachmentDeviceIndex int16, AttachmentStatus string, Description string, MacAddress string, NetworkInterfaceId string, OwnerId string, PrivateDnsName string, PrivateIpAddress string, SourceDestCheck bool, Status string, SubnetId string, RefSubnetId sql.NullInt64, VpcId string, RefVpcId sql.NullInt64, InterfaceType string) *InstanceNetworkInterface {
	return &InstanceNetworkInterface{Id: Id,
		InstanceId:                    InstanceId,
		AssociationIpOwnerId:          AssociationIpOwnerId,
		AssociationPublicDnsName:      AssociationPublicDnsName,
		AssociationPublicIp:           AssociationPublicIp,
		AttachmentAttachTime:          AttachmentAttachTime,
		AttachmentAttachmentId:        AttachmentAttachmentId,
		AttachmentDeleteOnTermination: AttachmentDeleteOnTermination,
		AttachmentDeviceIndex:         AttachmentDeviceIndex,
		AttachmentStatus:              AttachmentStatus,
		Description:                   Description,
		MacAddress:                    MacAddress,
		NetworkInterfaceId:            NetworkInterfaceId,
		OwnerId:                       OwnerId,
		PrivateDnsName:                PrivateDnsName,
		PrivateIpAddress:              PrivateIpAddress,
		SourceDestCheck:               SourceDestCheck,
		Status:                        Status,
		SubnetId:                      SubnetId,
		RefSubnetId:                   RefSubnetId,
		VpcId:                         VpcId,
		RefVpcId:                      RefVpcId,
		InterfaceType:                 InterfaceType}
}

func (d *InstanceNetworkInterface) loadDependencies(db *sql.DB) {
	arrayOfGroups, err := loadAllInstanceNetworkInterfaceGroupByOrderUnique(db)
	if err == nil {
		for _, x := range arrayOfGroups {
			if x.InstanceIdNetworkInterfaceId == d.Id {
				d.Groups = append(d.Groups, x)
			}
		}
	} else {
		fmt.Println(err)
	}

	arrayOfPrivateIpAddresses, err := loadAllInstanceNetworkInterfacePrivateIpAddressByOrderUnique(db)
	if err == nil {
		for _, x := range arrayOfPrivateIpAddresses {
			if x.InstanceIdNetworkInterfaceId == d.Id {
				d.PrivateIpAddresses = append(d.PrivateIpAddresses, x)
			}
		}
	} else {
		fmt.Println(err)
	}
}

func (d *InstanceNetworkInterface) Dump() string {
	dumpNet := fmt.Sprintf("AssociationIpOwnerId(%s) AssociationPublicDnsName(%s) AttachmentAttachTime(%s) AttachmentAttachmentId(%s) AttachmentDeleteOnTermination(%t) AttachmentDeviceIndex(%d) AttachmentStatus(%s) Description(%s) MacAddress(%s) NetworkInterfaceId(%s) OwnerId(%s) PrivateDnsName(%s) PrivateIpAddress(%s) SourceDestCheck(%t) Status(%s) SubnetId(%s) RefSubnetId(%d) VpcId(%s) RefVpcId(%d) InterfaceType(%s)",
		d.AssociationIpOwnerId,
		d.AssociationPublicDnsName,
		d.AssociationPublicIp,
		d.AttachmentAttachTime,
		d.AttachmentAttachmentId,
		d.AttachmentDeleteOnTermination,
		d.AttachmentDeviceIndex,
		d.AttachmentStatus,
		d.Description,
		d.Groups,
		d.MacAddress,
		d.NetworkInterfaceId,
		d.OwnerId,
		d.PrivateDnsName,
		d.PrivateIpAddress,
		d.PrivateIpAddresses,
		d.SourceDestCheck,
		d.Status,
		d.SubnetId,
		d.RefSubnetId,
		d.VpcId,
		d.RefVpcId,
		d.InterfaceType)

	dumpNet += "\n 	Groups ["
	for _, x := range d.Groups {
		dumpNet += "{" + x.Dump() + "}"
	}

	dumpNet += "]\n 	PrivateIpAddresses ["
	for _, x := range d.PrivateIpAddresses {
		dumpNet += "{" + x.Dump() + "}"
	}
	dumpNet += "]"
	return dumpNet
}

type InstanceNetworkInterfaceGroup struct {
	Id                           int64
	InstanceIdNetworkInterfaceId int64
	GroupName                    string
	GroupId                      string
	RefSecurityGroupId           sql.NullInt64
}

func (d *InstanceNetworkInterfaceGroup) Dump() string {
	return fmt.Sprintf("GroupName(%s) GroupId(%s) RefSecurityGroupId(%d)",
		d.GroupName, d.GroupId, d.RefSecurityGroupId)
}

func NewInstanceNetworkInterfaceGroup(Id int64, InstanceIdNetworkInterfaceId int64, GroupName string, GroupId string, RefSecurityGroupId sql.NullInt64) *InstanceNetworkInterfaceGroup {
	return &InstanceNetworkInterfaceGroup{Id,
		InstanceIdNetworkInterfaceId,
		GroupName,
		GroupId,
		RefSecurityGroupId}
}

type InstanceNetworkInterfacePrivateIpAddress struct {
	Id                           int64
	InstanceIdNetworkInterfaceId int64
	AssociationIpOwnerId         string
	AssociationPublicDnsName     string
	AssociationPublicIp          string
	Primary                      bool
	PrivateDnsName               string
	PrivateIpAddress             string
}

func (d *InstanceNetworkInterfacePrivateIpAddress) Dump() string {
	return fmt.Sprintf("AssociationIpOwnerId(%s) AssociationPublicDnsName(%s) AssociationPublicIp(%s) Primary(%t) PrivateDnsName(%s) PrivateIpAddress(%s)",
		d.AssociationIpOwnerId, d.AssociationPublicDnsName, d.AssociationPublicIp, d.Primary, d.PrivateDnsName, d.PrivateIpAddress)
}

func NewInstanceNetworkInterfacePrivateIpAddress(Id int64, InstanceIdNetworkInterfaceId int64, AssociationIpOwnerId string, AssociationPublicDnsName string, AssociationPublicIp string, Primary bool, PrivateDnsName string, PrivateIpAddress string) *InstanceNetworkInterfacePrivateIpAddress {
	return &InstanceNetworkInterfacePrivateIpAddress{Id,
		InstanceIdNetworkInterfaceId,
		AssociationIpOwnerId,
		AssociationPublicDnsName,
		AssociationPublicIp,
		Primary,
		PrivateDnsName,
		PrivateIpAddress}
}

type InstanceSecurityGroup struct {
	Id                 int64
	InstanceId         int64
	GroupName          string
	GroupId            string
	RefSecurityGroupId sql.NullInt64
}

func (d *InstanceSecurityGroup) Dump() string {
	return fmt.Sprintf("GroupName(%s) GroupId(%s) RefSecurityGroupId(%d)",
		d.GroupName, d.GroupId, d.RefSecurityGroupId)
}

func NewInstanceSecurityGroup(Id int64, InstanceId int64, GroupName string, GroupId string, RefSecurityGroupId sql.NullInt64) *InstanceSecurityGroup {
	return &InstanceSecurityGroup{Id,
		InstanceId,
		GroupName,
		GroupId,
		RefSecurityGroupId}
}

type InstanceTag struct {
	Id         int64
	InstanceId int64
	Key        string
	Value      string
}

func (d *InstanceTag) Dump() string {
	return fmt.Sprintf("Key(%s) Value(%s)",
		d.Key, d.Value)
}

func NewInstanceTag(Id int64, InstanceId int64, Key string, Value string) *InstanceTag {
	return &InstanceTag{Id,
		InstanceId,
		Key,
		Value}
}

type InstanceLicense struct {
	Id                      int64
	InstanceId              int64
	LicenseConfigurationArn string
}

func (d *InstanceLicense) Dump() string {
	return fmt.Sprintf("LicenseConfigurationArn(%s)",
		d.LicenseConfigurationArn)
}

func NewInstanceLicense(Id int64, InstanceId int64, LicenseConfigurationArn string) *InstanceLicense {
	return &InstanceLicense{Id,
		InstanceId,
		LicenseConfigurationArn}
}
