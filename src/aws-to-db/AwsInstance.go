package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type AwsReservation struct {
	Groups        []AwsGroup    `json:"Groups,omitempty"`
	Instances     []AwsInstance `json:"Instances,omitempty"`
	OwnerId       string        `json:"OwnerId,omitempty"`
	RequesterId   string        `json:"RequesterId,omitempty"`
	ReservationId string        `json:"ReservationId,omitempty"`
}

type AwsReservations struct {
	Reservations []AwsReservation `json:"Reservations,omitempty"`
}

type AwsGroup struct {
	GroupName string `json:"GroupName,omitempty"`
	GroupId   string `json:"GroupId,omitempty"`
}

type AwsInstance struct {
	AmiLaunchIndex                          int16                                       `json:"AmiLaunchIndex,omitempty"`
	ImageId                                 string                                      `json:"ImageId,omitempty"`
	InstanceId                              string                                      `json:"InstanceId,omitempty"`
	InstanceType                            string                                      `json:"InstanceType,omitempty"`
	KernelId                                string                                      `json:"KernelId,omitempty"`
	KeyName                                 string                                      `json:"KeyName,omitempty"`
	LaunchTime                              string                                      `json:"LaunchTime,omitempty"`
	Monitoring                              AwsMonitoring                               `json:"Monitoring,omitempty"`
	Placement                               AwsPlacement                                `json:"Placement,omitempty"`
	Platform                                string                                      `json:"Platform,omitempty"`
	PrivateDnsName                          string                                      `json:"PrivateDnsName,omitempty"`
	PrivateIpAddress                        string                                      `json:"PrivateIpAddress,omitempty"`
	ProductCodes                            []AwsProductCode                            `json:"ProductCodes,omitempty"`
	PublicDnsName                           string                                      `json:"PublicDnsName,omitempty"`
	PublicIpAddress                         string                                      `json:"PublicIpAddress,omitempty"`
	RamdiskId                               string                                      `json:"RamdiskId,omitempty"`
	State                                   AwsState                                    `json:"State,omitempty"`
	StateTransitionReason                   string                                      `json:"StateTransitionReason,omitempty"`
	SubnetId                                string                                      `json:"SubnetId,omitempty"`
	VpcId                                   string                                      `json:"VpcId,omitempty"`
	Architecture                            string                                      `json:"Architecture,omitempty"`
	BlockDeviceMappings                     []AwsBlockDeviceMapping                     `json:"BlockDeviceMappings,omitempty"`
	ClientToken                             string                                      `json:"ClientToken,omitempty"`
	EbsOptimized                            bool                                        `json:"EbsOptimized,omitempty"`
	EnaSupport                              bool                                        `json:"EnaSupport,omitempty"`
	Hypervisor                              string                                      `json:"Hypervisor,omitempty"`
	IamInstanceProfile                      AwsIamInstanceProfile                       `json:"IamInstanceProfile,omitempty"`
	InstanceLifecycle                       string                                      `json:"InstanceLifecycle,omitempty"`
	ElasticGpuAssociations                  []AwsElasticGpuAssociation                  `json:"ElasticGpuAssociations,omitempty"`
	ElasticInferenceAcceleratorAssociations []AwsElasticInferenceAcceleratorAssociation `json:"ElasticInferenceAcceleratorAssociations,omitempty"`
	NetworkInterfaces                       []AwsNetworkInterface                       `json:"NetworkInterfaces,omitempty"`
	RootDeviceName                          string                                      `json:"RootDeviceName,omitempty"`
	RootDeviceType                          string                                      `json:"RootDeviceType,omitempty"`
	SecurityGroups                          []AwsSecurityGroup                          `json:"SecurityGroups,omitempty"`
	SourceDestCheck                         bool                                        `json:"SourceDestCheck,omitempty"`
	SpotInstanceRequestId                   string                                      `json:"SpotInstanceRequestId,omitempty"`
	SriovNetSupport                         string                                      `json:"SriovNetSupport,omitempty"`
	StateReason                             AwsStateReason                              `json:"StateReason,omitempty"`
	Tags                                    []AwsTag                                    `json:"Tags,omitempty"`
	VirtualizationType                      string                                      `json:"VirtualizationType,omitempty"`
	CpuOption                               AwsCpuOption                                `json:"CpuOption,omitempty"`
	CapacityReservationId                   string                                      `json:"CapacityReservationId,omitempty"`
	CapacityReservationSpecification        AwsCapacityReservationSpecification         `json:"CapacityReservationSpecification,omitempty"`
	HibernationOptions                      AwsHibernationOption                        `json:"HibernationOptions,omitempty"`
	Licenses                                []AwsLicense                                `json:"Licenses,omitempty"`
}

func awsInstanceParseAndPersist(db *sql.DB, jsonString string) error {
	fmt.Printf("awsInstanceParseAndPersist with json lenght %d\n", len(jsonString))
	awsReservations := new(AwsReservations)
	errUnmarshal := json.Unmarshal([]byte(jsonString), awsReservations)
	if errUnmarshal == nil {
		for _, awsReservation := range awsReservations.Reservations {
			fmt.Printf("OwnerId(%s)\n", awsReservation.OwnerId)
			reservation := createReservation(db, awsReservation.OwnerId, awsReservation.RequesterId, awsReservation.ReservationId)
			if reservation != nil {
				fmt.Printf("    Reservation (%s) loaded\n", reservation.OwnerId)
				createRestOfInstanceContent(db, reservation, awsReservation)
			} else {
				fmt.Printf("  ERROR  reservation (%s) not loaded\n", awsReservation.OwnerId)
			}
		}
	} else {
		return errUnmarshal
	}
	return nil
}

func createRestOfInstanceContent(db *sql.DB, reservation *Reservation, awsReservation AwsReservation) {

	for _, awsGroup := range awsReservation.Groups {
		group := createGroup(db, reservation.Id, awsGroup.GroupName, awsGroup.GroupId)
		if group != nil {
			fmt.Printf("    	Group (%s) loaded\n", group.GroupName)
		} else {
			fmt.Printf("    	ERROR Group (%s) not loaded\n", awsGroup.GroupName)

		}
	}

	for _, awsInstance := range awsReservation.Instances {
		instance := createInstance(db, reservation.Id, awsInstance.AmiLaunchIndex, awsInstance.ImageId, awsInstance.InstanceId, awsInstance.InstanceType, awsInstance.KernelId, awsInstance.KeyName, awsInstance.LaunchTime, awsInstance.Monitoring.State, awsInstance.Placement.AvailabilityZone, awsInstance.Placement.Affinity, awsInstance.Placement.GroupName, awsInstance.Placement.PartitionNumber, awsInstance.Placement.HostId, awsInstance.Placement.Tenancy, awsInstance.Placement.SpreadDomain, awsInstance.Platform, awsInstance.PrivateDnsName, awsInstance.PrivateIpAddress, awsInstance.PublicDnsName, awsInstance.PublicIpAddress, awsInstance.RamdiskId, awsInstance.State.Code, awsInstance.State.Name, awsInstance.StateTransitionReason, awsInstance.SubnetId, awsInstance.VpcId, awsInstance.Architecture, awsInstance.ClientToken, awsInstance.EbsOptimized, awsInstance.EnaSupport, awsInstance.Hypervisor, awsInstance.IamInstanceProfile.Id, awsInstance.IamInstanceProfile.Arn, awsInstance.InstanceLifecycle, awsInstance.RootDeviceName, awsInstance.RootDeviceType, awsInstance.SourceDestCheck, awsInstance.SpotInstanceRequestId, awsInstance.SriovNetSupport, awsInstance.StateReason.Code, awsInstance.StateReason.Message, awsInstance.VirtualizationType, awsInstance.CpuOption.CoreCount, awsInstance.CpuOption.ThreadsPerCore, awsInstance.CapacityReservationId, awsInstance.CapacityReservationSpecification.CapacityReservationPreference, awsInstance.CapacityReservationSpecification.CapacityReservationTarget.CapacityReservationId, awsInstance.HibernationOptions.Configured)
		if instance != nil {
			fmt.Printf("    	Instance (%s) loaded\n", instance.KeyName)

			for _, awsProductCode := range awsInstance.ProductCodes {
				productCode := createInstanceProductCode(db, instance.Id, awsProductCode.ProductCodeId, awsProductCode.ProductCodeType)
				if productCode != nil {
					fmt.Printf("    		ProductCode (%s) loaded\n", productCode.ProductCodeId)
				} else {
					fmt.Printf("    		ERROR ProductCode (%s) not loaded\n", awsProductCode.ProductCodeId)

				}
			}

			for _, awsBlockDeviceMapping := range awsInstance.BlockDeviceMappings {
				block := createInstanceBlockDeviceMapping(db, instance.Id, awsBlockDeviceMapping.DeviceName, awsBlockDeviceMapping.Ebs.AttachTime, awsBlockDeviceMapping.Ebs.DeleteOnTermination, awsBlockDeviceMapping.Ebs.Status, awsBlockDeviceMapping.Ebs.VolumeId)
				if block != nil {
					fmt.Printf("    		BlockDeviceMapping (%s) loaded\n", block.DeviceName)
				} else {
					fmt.Printf("    		ERROR BlockDeviceMapping (%s) not loaded\n", awsBlockDeviceMapping.DeviceName)

				}
			}

			for _, awsEleasticGpu := range awsInstance.ElasticGpuAssociations {
				gpu := createInstanceElasticGpuAssociation(db, instance.Id, awsEleasticGpu.ElasticGpuId, awsEleasticGpu.ElasticGpuAssociationId, awsEleasticGpu.ElasticGpuAssociationState, awsEleasticGpu.ElasticGpuAssociationTime)
				if gpu != nil {
					fmt.Printf("    		ElasticGpuAssociation (%s) loaded\n", gpu.ElasticGpuId)
				} else {
					fmt.Printf("    		ERROR ElasticGpuAssociation (%s) not loaded\n", awsEleasticGpu.ElasticGpuId)

				}
			}

			for _, awsEleasticInterface := range awsInstance.ElasticInferenceAcceleratorAssociations {
				elInterface := createInstanceElasticInferenceAcceleratorAssociation(db, instance.Id, awsEleasticInterface.ElasticInferenceAcceleratorArn, awsEleasticInterface.ElasticInferenceAcceleratorAssociationId, awsEleasticInterface.ElasticInferenceAcceleratorAssociationState, awsEleasticInterface.ElasticInferenceAcceleratorAssociationTime)
				if elInterface != nil {
					fmt.Printf("    		ElasticInferenceAcceleratorAssociation (%s) loaded\n", elInterface.ElasticInferenceAcceleratorAssociationId)
				} else {
					fmt.Printf("    		ERROR ElasticInferenceAcceleratorAssociation (%s) not loaded\n", awsEleasticInterface.ElasticInferenceAcceleratorAssociationId)
				}
			}

			for _, awsNetwork := range awsInstance.NetworkInterfaces {
				network := createInstanceNetworkInterface(db, instance.Id, awsNetwork.Association.IpOwnerId, awsNetwork.Association.PublicDnsName, awsNetwork.Association.PublicIp, awsNetwork.Attachment.AttachTime, awsNetwork.Attachment.AttachmentId, awsNetwork.Attachment.DeleteOnTermination, awsNetwork.Attachment.DeviceIndex, awsNetwork.Attachment.Status, awsNetwork.Description, awsNetwork.MacAddress, awsNetwork.NetworkInterfaceId, awsNetwork.OwnerId, awsNetwork.PrivateDnsName, awsNetwork.PrivateIpAddress, awsNetwork.SourceDestCheck, awsNetwork.Status, awsNetwork.SubnetId, awsNetwork.VpcId, awsNetwork.InterfaceType)
				if network != nil {
					fmt.Printf("    		NetworkInterface (%s) loaded\n", network.MacAddress)
					for _, awsNetworkGroup := range awsNetwork.Groups {
						networkGroup := createInstanceNetworkInterfaceGroup(db, network.Id, awsNetworkGroup.GroupName, awsNetworkGroup.GroupId)
						if networkGroup != nil {
							fmt.Printf("    		NetworkInterfaceGroup (%s) loaded\n", networkGroup.GroupName)
						} else {
							fmt.Printf("    		ERROR NetworkInterfaceGroup (%s) not loaded\n", awsNetworkGroup.GroupName)
						}
					}
					for _, awsNetworkIp := range awsNetwork.PrivateIpAddresses {
						networkIp := createInstanceNetworkInterfacePrivateIpAddress(db, network.Id, awsNetworkIp.Association.IpOwnerId, awsNetworkIp.Association.PublicDnsName, awsNetworkIp.Association.PublicIp, awsNetworkIp.Primary, awsNetworkIp.PrivateDnsName, awsNetworkIp.PrivateIpAddress)
						if networkIp != nil {
							fmt.Printf("    		NetworkInterfacePrivateIpAddress (%s) loaded\n", networkIp.AssociationPublicDnsName)
						} else {
							fmt.Printf("    		ERROR NetworkInterfacePrivateIpAddress (%s) not loaded\n", awsNetworkIp.Association.IpOwnerId)
						}
					}
				} else {
					fmt.Printf("    		ERROR NetworkInterface (%s) not loaded\n", awsNetwork.MacAddress)
				}
			}

			for _, awsSecurityGroup := range awsInstance.SecurityGroups {
				securityGroup := createInstanceSecurityGroup(db, instance.Id, awsSecurityGroup.GroupName, awsSecurityGroup.GroupId)
				if securityGroup != nil {
					fmt.Printf("    		SecurityGroup (%s) loaded\n", securityGroup.GroupName)
				} else {
					fmt.Printf("    		ERROR SecurityGroup (%s) not loaded\n", awsSecurityGroup.GroupName)
				}
			}

			for _, awsTag := range awsInstance.Tags {
				tag := createInstanceTag(db, instance.Id, awsTag.Key, awsTag.Value)
				if tag != nil {
					fmt.Printf("    		Tag (%s) loaded\n", tag.Value)
				} else {
					fmt.Printf("    		Tag SecurityGroup (%s) not loaded\n", awsTag.Value)
				}
			}

			for _, awsTag := range awsInstance.Tags {
				tag := createInstanceTag(db, instance.Id, awsTag.Key, awsTag.Value)
				if tag != nil {
					fmt.Printf("    		Tag (%s) loaded\n", tag.Value)
				} else {
					fmt.Printf("    		Tag SecurityGroup (%s) not loaded\n", awsTag.Value)
				}
			}

			for _, awsLicense := range awsInstance.Licenses {
				license := createInstanceLicense(db, instance.Id, awsLicense.LicenseConfigurationArn)
				if license != nil {
					fmt.Printf("    		License (%s) loaded\n", license.LicenseConfigurationArn)
				} else {
					fmt.Printf("    		Tag SecurityGroup (%s) not loaded\n", awsLicense.LicenseConfigurationArn)
				}
			}

		} else {
			fmt.Printf("    	ERROR Instance (%s) not loaded\n", awsInstance.CapacityReservationId)

		}
	}

}
