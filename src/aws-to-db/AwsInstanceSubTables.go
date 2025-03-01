package main

type AwsMonitoring struct {
	State string `json:"State,omitempty"`
}

type AwsPlacement struct {
	AvailabilityZone string `json:"AvailabilityZone,omitempty"`
	Affinity         string `json:"Affinity,omitempty"`
	GroupName        string `json:"GroupName,omitempty"`
	PartitionNumber  int16  `json:"PartitionNumber,omitempty"`
	HostId           string `json:"HostId,omitempty"`
	Tenancy          string `json:"Tenancy,omitempty"`
	SpreadDomain     string `json:"SpreadDomain,omitempty"`
}

type AwsProductCode struct {
	ProductCodeId   string `json:"ProductCodeId,omitempty"`
	ProductCodeType string `json:"ProductCodeType,omitempty"`
}

type AwsState struct {
	Code int16  `json:"Code,omitempty"`
	Name string `json:"Name,omitempty"`
}

type AwsBlockDeviceMapping struct {
	DeviceName string `json:"DeviceName,omitempty"`
	Ebs        AwsEb  `json:"Ebs,omitempty"`
}

type AwsEb struct {
	AttachTime          string `json:"AttachTime,omitempty"`
	DeleteOnTermination bool   `json:"DeleteOnTermination,omitempty"`
	Status              string `json:"Status,omitempty"`
	VolumeId            string `json:"VolumeId,omitempty"`
}

type AwsIamInstanceProfile struct {
	Arn string `json:"Arn,omitempty"`
	Id  string `json:"Id,omitempty"`
}

type AwsElasticGpuAssociation struct {
	ElasticGpuId               string `json:"ElasticGpuId,omitempty"`
	ElasticGpuAssociationId    string `json:"ElasticGpuAssociationId,omitempty"`
	ElasticGpuAssociationState string `json:"ElasticGpuAssociationState,omitempty"`
	ElasticGpuAssociationTime  string `json:"ElasticGpuAssociationTime,omitempty"`
}

type AwsElasticInferenceAcceleratorAssociation struct {
	ElasticInferenceAcceleratorArn              string `json:"ElasticInferenceAcceleratorArn,omitempty"`
	ElasticInferenceAcceleratorAssociationId    string `json:"ElasticInferenceAcceleratorAssociationId,omitempty"`
	ElasticInferenceAcceleratorAssociationState string `json:"ElasticInferenceAcceleratorAssociationState,omitempty"`
	ElasticInferenceAcceleratorAssociationTime  string `json:"ElasticInferenceAcceleratorAssociationTime,omitempty"`
}

type AwsAssociation struct {
	IpOwnerId     string `json:"IpOwnerId,omitempty"`
	PublicDnsName string `json:"PublicDnsName,omitempty"`
	PublicIp      string `json:"PublicIp,omitempty"`
}

type AwsAttachment struct {
	AttachTime          string `json:"AttachTime,omitempty"`
	AttachmentId        string `json:"AttachmentId,omitempty"`
	DeleteOnTermination bool   `json:"DeleteOnTermination,omitempty"`
	DeviceIndex         int16  `json:"DeviceIndex,omitempty"`
	Status              string `json:"Status,omitempty"`
}

type AwsPrivateIpAddress struct {
	Association      AwsAssociation `json:"Association,omitempty"`
	Primary          bool           `json:"Primary,omitempty"`
	PrivateDnsName   string         `json:"PrivateDnsName,omitempty"`
	PrivateIpAddress string         `json:"PrivateIpAddress,omitempty"`
}

type AwsNetworkInterface struct {
	Association        AwsAssociation        `json:"Association,omitempty"`
	Attachment         AwsAttachment         `json:"Attachment,omitempty"`
	Description        string                `json:"Description,omitempty"`
	Groups             []AwsGroup            `json:"Groups,omitempty"`
	MacAddress         string                `json:"MacAddress,omitempty"`
	NetworkInterfaceId string                `json:"NetworkInterfaceId,omitempty"`
	OwnerId            string                `json:"OwnerId,omitempty"`
	PrivateDnsName     string                `json:"PrivateDnsName,omitempty"`
	PrivateIpAddress   string                `json:"PrivateIpAddress,omitempty"`
	PrivateIpAddresses []AwsPrivateIpAddress `json:"PrivateIpAddresses,omitempty"`
	SourceDestCheck    bool                  `json:"SourceDestCheck,omitempty"`
	Status             string                `json:"Status,omitempty"`
	SubnetId           string                `json:"SubnetId,omitempty"`
	VpcId              string                `json:"VpcId,omitempty"`
	InterfaceType      string                `json:"InterfaceType,omitempty"`
}

type AwsStateReason struct {
	Code    string `json:"Code,omitempty"`
	Message string `json:"Message,omitempty"`
}

type AwsCpuOption struct {
	CoreCount      int16 `json:"CoreCount,omitempty"`
	ThreadsPerCore int16 `json:"ThreadsPerCore,omitempty"`
}

type AwsCapacityReservationSpecification struct {
	CapacityReservationPreference string                       `json:"CapacityReservationPreference,omitempty"`
	CapacityReservationTarget     AwsCapacityReservationTarget `json:"CapacityReservationTarget,omitempty"`
}

type AwsCapacityReservationTarget struct {
	CapacityReservationId string `json:"CapacityReservationId,omitempty"`
}

type AwsHibernationOption struct {
	Configured bool `json:"Configured,omitempty"`
}

type AwsLicense struct {
	LicenseConfigurationArn string `json:"LicenseConfigurationArn,omitempty"`
}
