package main

import (
	"database/sql"
	"fmt"
)

type Reservation struct {
	Id            int64
	Groups        []*Group
	Instances     []*Instance
	OwnerId       string
	RequesterId   string
	ReservationId string
}

func (d *Reservation) loadDependencies(db *sql.DB) {
	arrayOfGroups, err := loadAllGroupByOrderUnique(db)
	if err == nil {
		for _, group := range arrayOfGroups {
			if group.ReservationId == d.Id {
				d.Groups = append(d.Groups, group)
			}
		}
	} else {
		fmt.Println(err)
	}

	arrayOfInstances, err := loadAllInstanceByOrderUnique(db)
	if err == nil {
		for _, instance := range arrayOfInstances {
			if instance.ReservationId == d.Id {
				d.Instances = append(d.Instances, instance)
			}
		}
	} else {
		fmt.Println(err)
	}
}

func (d *Reservation) Dump(db *sql.DB) string {
	dumpRes := fmt.Sprintf("OwnerId(%s) RequesterId(%s) ReservationId(%s)", d.OwnerId, d.RequesterId, d.ReservationId)
	dumpRes += "\n	 Groups ["
	for _, groups := range d.Groups {
		dumpRes += "{" + groups.Dump() + "}"
	}

	dumpRes += "]\n	 Instances ["
	for _, instances := range d.Instances {
		instances.loadDependencies(db)
		dumpRes += "{" + instances.Dump(db) + "}"
	}

	dumpRes += "]"
	return dumpRes
}

type Group struct {
	Id            int64
	ReservationId int64
	GroupName     string
	GroupId       string
}

func (d *Group) Dump() string {
	return fmt.Sprintf("GroupName(%s) GroupId(%s)", d.GroupName, d.GroupId)
}

type Instance struct {
	Id                                      int64
	ReservationId                           int64
	AmiLaunchIndex                          int16
	ImageId                                 string
	RefImageId                              sql.NullInt64
	InstanceId                              string
	InstanceType                            string
	KernelId                                string
	KeyName                                 string
	LaunchTime                              string
	MonitoringState                         string
	PlacementAvailabilityZone               string
	PlacementAffinity                       string
	PlacementGroupName                      string
	PlacementPartitionNumber                int16
	PlacementHostId                         string
	PlacementTenancy                        string
	PlacementSpreadDomain                   string
	Platform                                string
	PrivateDnsName                          string
	PrivateIpAddress                        string
	ProductCodes                            []*InstanceProductCode
	PublicDnsName                           string
	PublicIpAddress                         string
	RamdiskId                               string
	StateCode                               int16
	StateName                               string
	StateTransitionReason                   string
	SubnetId                                string
	RefSubnetId                             sql.NullInt64
	VpcId                                   string
	RefVpcId                                sql.NullInt64
	Architecture                            string
	BlockDeviceMappings                     []*InstanceBlockDeviceMapping
	ClientToken                             string
	EbsOptimized                            bool
	EnaSupport                              bool
	Hypervisor                              string
	IamInstanceProfileId                    string
	IamInstanceProfileArn                   string
	InstanceLifecycle                       string
	ElasticGpuAssociations                  []*InstanceElasticGpuAssociation
	ElasticInferenceAcceleratorAssociations []*InstanceElasticInferenceAcceleratorAssociation
	NetworkInterfaces                       []*InstanceNetworkInterface
	RootDeviceName                          string
	RootDeviceType                          string
	SecurityGroups                          []*InstanceSecurityGroup
	SourceDestCheck                         bool
	SpotInstanceRequestId                   string
	SriovNetSupport                         string
	StateReasonCode                         string
	StateReasonMessage                      string
	Tags                                    []*InstanceTag
	VirtualizationType                      string
	CpuOptionCoreCount                      int16
	CpuOptionThreadsPerCore                 int16
	CapacityReservationId                   string
	CapacityReservationPreference           string
	CapacityReservationTargetId             string
	HibernationOptionsConfigured            bool
	Licenses                                []*InstanceLicense
}

func (d *Instance) loadDependencies(db *sql.DB) {
	arrayOfProductCodes, err := loadAllInstanceProductCodeByOrderUnique(db)
	if err == nil {
		for _, x := range arrayOfProductCodes {
			if x.InstanceId == d.Id {
				d.ProductCodes = append(d.ProductCodes, x)
			}
		}
	} else {
		fmt.Println(err, "arrayOfProductCodes")
	}

	arrayOfBlockDeviceMappings, err := loadAllInstanceBlockDeviceMappingByOrderUnique(db)
	if err == nil {
		for _, x := range arrayOfBlockDeviceMappings {
			if x.InstanceId == d.Id {
				d.BlockDeviceMappings = append(d.BlockDeviceMappings, x)
			}
		}
	} else {
		fmt.Println(err, "arrayOfBlockDeviceMappings")
	}
	arrayOfElasticGpuAssociations, err := loadAllInstanceElasticGpuAssociationByOrderUnique(db)
	if err == nil {
		for _, x := range arrayOfElasticGpuAssociations {
			if x.InstanceId == d.Id {
				d.ElasticGpuAssociations = append(d.ElasticGpuAssociations, x)
			}
		}
	} else {
		fmt.Println(err, "arrayOfElasticGpuAssociations")
	}
	arrayOfElasticInferenceAcceleratorAssociations, err := loadAllInstanceElasticInferenceAcceleratorAssociationByOrderUnique(db)
	if err == nil {
		for _, x := range arrayOfElasticInferenceAcceleratorAssociations {
			if x.InstanceId == d.Id {
				d.ElasticInferenceAcceleratorAssociations = append(d.ElasticInferenceAcceleratorAssociations, x)
			}
		}
	} else {
		fmt.Println(err, "arrayOfElasticInferenceAcceleratorAssociations")
	}
	arrayOfNetworkInterfaces, err := loadAllInstanceNetworkInterfaceByOrderUnique(db)
	if err == nil {
		for _, x := range arrayOfNetworkInterfaces {
			if x.InstanceId == d.Id {
				d.NetworkInterfaces = append(d.NetworkInterfaces, x)
			}
		}
	} else {
		fmt.Println(err, "arrayOfNetworkInterfaces")
	}
	arrayOfSecurityGroups, err := loadAllInstanceSecurityGroupByOrderUnique(db)
	if err == nil {
		for _, x := range arrayOfSecurityGroups {
			if x.InstanceId == d.Id {
				d.SecurityGroups = append(d.SecurityGroups, x)
			}
		}
	} else {
		fmt.Println(err, "arrayOfSecurityGroups")
	}
	arrayOfTags, err := loadAllInstanceTagsByOrderUnique(db)
	if err == nil {
		for _, x := range arrayOfTags {
			if x.InstanceId == d.Id {
				d.Tags = append(d.Tags, x)
			}
		}
	} else {
		fmt.Println(err, "arrayOfTags")
	}
	arrayOfLicenses, err := loadAllInstanceLicenseByOrderUnique(db)
	if err == nil {
		for _, x := range arrayOfLicenses {
			if x.InstanceId == d.Id {
				d.Licenses = append(d.Licenses, x)
			}
		}
	} else {
		fmt.Println(err, "arrayOfLicenses")
	}
}

func (d *Instance) Dump(db *sql.DB) string {
	dumpInst := fmt.Sprintf("AmiLaunchIndex(%d) ImageId(%s) RefImageId(%d) InstanceId(%s) InstanceType(%s) KernelId(%s) KeyName(%s) LaunchTime(%s) MonitoringState(%s) PlacementAvailabilityZone(%s) PlacementAffinity(%s) PlacementGroupName(%s) PlacementPartitionNumber(%d) PlacementHostId(%s) PlacementTenancy(%s) PlacementSpreadDomain(%s) Platform(%s) PrivateDnsName(%s) PrivateIpAddress(%s) PublicDnsName(%s) PublicIpAddress(%s) RamdiskId(%s) StateCode(%d) StateName(%s) StateTransitionReason(%s) SubnetId(%s) RefSubnetId(%d) VpcId(%s) RefVpcId(%d) Architecture(%s) ClientToken(%s) EbsOptimized(%t) EnaSupport(%t) Hypervisor(%s) IamInstanceProfileId(%s) IamInstanceProfileArn(%s) InstanceLifecycle(%s) RootDeviceName(%s) RootDeviceType(%s) SourceDestCheck(%t) SpotInstanceRequestId(%s) SriovNetSupport(%s) StateReasonCode(%s) StateReasonMessage(%s) VirtualizationType(%s) CpuOptionCoreCount(%d) CpuOptionThreadsPerCore(%d) CapacityReservationId(%s) CapacityReservationPreference(%s) CapacityReservationTargetId(%s) HibernationOptionsConfigured(%t)",
		d.AmiLaunchIndex,
		d.ImageId,
		d.RefImageId,
		d.InstanceId,
		d.InstanceType,
		d.KernelId,
		d.KeyName,
		d.LaunchTime,
		d.MonitoringState,
		d.PlacementAvailabilityZone,
		d.PlacementAffinity,
		d.PlacementGroupName,
		d.PlacementPartitionNumber,
		d.PlacementHostId,
		d.PlacementTenancy,
		d.PlacementSpreadDomain,
		d.Platform,
		d.PrivateDnsName,
		d.PrivateIpAddress,
		d.PublicDnsName,
		d.PublicIpAddress,
		d.RamdiskId,
		d.StateCode,
		d.StateName,
		d.StateTransitionReason,
		d.SubnetId,
		d.RefSubnetId,
		d.VpcId,
		d.RefVpcId,
		d.Architecture,
		d.ClientToken,
		d.EbsOptimized,
		d.EnaSupport,
		d.Hypervisor,
		d.IamInstanceProfileId,
		d.IamInstanceProfileArn,
		d.InstanceLifecycle,
		d.RootDeviceName,
		d.RootDeviceType,
		d.SourceDestCheck,
		d.SpotInstanceRequestId,
		d.SriovNetSupport,
		d.StateReasonCode,
		d.StateReasonMessage,
		d.VirtualizationType,
		d.CpuOptionCoreCount,
		d.CpuOptionThreadsPerCore,
		d.CapacityReservationId,
		d.CapacityReservationPreference,
		d.CapacityReservationTargetId,
		d.HibernationOptionsConfigured)

	dumpInst += "\n	 ProductCodes ["
	for _, x := range d.ProductCodes {
		dumpInst += "{" + x.Dump() + "}"
	}

	dumpInst += "]\n	 BlockDeviceMappings ["
	for _, x := range d.BlockDeviceMappings {
		dumpInst += "{" + x.Dump() + "}"
	}

	dumpInst += "]\n	 ElasticGpuAssociations ["
	for _, x := range d.ElasticGpuAssociations {
		dumpInst += "{" + x.Dump() + "}"
	}

	dumpInst += "]\n	 ElasticInferenceAcceleratorAssociations ["
	for _, x := range d.ElasticInferenceAcceleratorAssociations {
		dumpInst += "{" + x.Dump() + "}"
	}

	dumpInst += "]\n	 NetworkInterfaces ["
	for _, x := range d.NetworkInterfaces {
		x.loadDependencies(db)
		dumpInst += "{" + x.Dump() + "}"
	}

	dumpInst += "]\n	 SecurityGroups ["
	for _, x := range d.SecurityGroups {
		dumpInst += "{" + x.Dump() + "}"
	}

	dumpInst += "]\n	 Tags ["
	for _, x := range d.Tags {
		dumpInst += "{" + x.Dump() + "}"
	}

	dumpInst += "]\n	 Licenses ["
	for _, x := range d.Licenses {
		dumpInst += "{" + x.Dump() + "}"
	}

	dumpInst += "]"
	return dumpInst
}

func NewReservation(Id int64, OwnerId string, RequesterId string, ReservationId string) *Reservation {
	return &Reservation{Id: Id,
		OwnerId:       OwnerId,
		RequesterId:   RequesterId,
		ReservationId: ReservationId}
}

func NewGroup(Id int64, ReservationId int64, GroupName string, GroupId string) *Group {
	return &Group{Id: Id,
		ReservationId: ReservationId,
		GroupName:     GroupName,
		GroupId:       GroupId}
}

func NewInstance(Id int64, ReservationId int64, AmiLaunchIndex int16, ImageId string, RefImageId sql.NullInt64, InstanceId string, InstanceType string, KernelId string, KeyName string, LaunchTime string, MonitoringState string, PlacementAvailabilityZone string, PlacementAffinity string, PlacementGroupName string, PlacementPartitionNumber int16, PlacementHostId string, PlacementTenancy string, PlacementSpreadDomain string, Platform string, PrivateDnsName string, PrivateIpAddress string, PublicDnsName string,
	PublicIpAddress string, RamdiskId string, StateCode int16, StateName string, StateTransitionReason string, SubnetId string, RefSubnetId sql.NullInt64, VpcId string, RefVpcId sql.NullInt64, Architecture string, ClientToken string, EbsOptimized bool, EnaSupport bool, Hypervisor string, IamInstanceProfileId string, IamInstanceProfileArn string, InstanceLifecycle string, RootDeviceName string, RootDeviceType string, SourceDestCheck bool, SpotInstanceRequestId string, SriovNetSupport string, StateReasonCode string, StateReasonMessage string,
	VirtualizationType string, CpuOptionCoreCount int16, CpuOptionThreadsPerCore int16, CapacityReservationId string, CapacityReservationPreference string, CapacityReservationTargetId string, HibernationOptionsConfigured bool) *Instance {

	return &Instance{Id: Id, ReservationId: ReservationId, AmiLaunchIndex: AmiLaunchIndex, ImageId: ImageId, RefImageId: RefImageId, InstanceId: InstanceId, InstanceType: InstanceType, KernelId: KernelId, KeyName: KeyName, LaunchTime: LaunchTime, MonitoringState: MonitoringState, PlacementAvailabilityZone: PlacementAvailabilityZone, PlacementAffinity: PlacementAffinity, PlacementGroupName: PlacementGroupName, PlacementPartitionNumber: PlacementPartitionNumber, PlacementHostId: PlacementHostId, PlacementTenancy: PlacementTenancy, PlacementSpreadDomain: PlacementSpreadDomain, Platform: Platform, PrivateDnsName: PrivateDnsName, PrivateIpAddress: PrivateIpAddress, PublicDnsName: PublicDnsName,
		PublicIpAddress: PublicIpAddress, RamdiskId: RamdiskId, StateCode: StateCode, StateName: StateName, StateTransitionReason: StateTransitionReason, SubnetId: SubnetId, RefSubnetId: RefSubnetId, VpcId: VpcId, RefVpcId: RefVpcId, Architecture: Architecture, ClientToken: ClientToken, EbsOptimized: EbsOptimized, EnaSupport: EnaSupport, Hypervisor: Hypervisor, IamInstanceProfileId: IamInstanceProfileId, IamInstanceProfileArn: IamInstanceProfileArn, InstanceLifecycle: InstanceLifecycle, RootDeviceName: RootDeviceName, RootDeviceType: RootDeviceType, SourceDestCheck: SourceDestCheck, SpotInstanceRequestId: SpotInstanceRequestId, SriovNetSupport: SriovNetSupport, StateReasonCode: StateReasonCode, StateReasonMessage: StateReasonMessage,
		VirtualizationType: VirtualizationType, CpuOptionCoreCount: CpuOptionCoreCount, CpuOptionThreadsPerCore: CpuOptionThreadsPerCore, CapacityReservationId: CapacityReservationId, CapacityReservationPreference: CapacityReservationPreference, CapacityReservationTargetId: CapacityReservationTargetId, HibernationOptionsConfigured: HibernationOptionsConfigured}
}
