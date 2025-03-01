package main

type AwsPendingModifiedValues struct {
	DBInstanceClass         string `json:"DBInstanceClass,omitempty"`
	AllocatedStorage        int16  `json:"AllocatedStorage,omitempty"`
	MasterUserPassword      string `json:"MasterUserPassword,omitempty"`
	Port                    int16  `json:"Port,omitempty"`
	BackupRetentionPeriod   int16  `json:"BackupRetentionPeriod,omitempty"`
	MultiAZ                 bool   `json:"MultiAZ,omitempty"`
	EngineVersion           string `json:"EngineVersion,omitempty"`
	LicenseModel            string `json:"LicenseModel,omitempty"`
	Iops                    int16  `json:"Iops,omitempty"`
	DBInstanceIdentifier    string `json:"DBInstanceIdentifier,omitempty"`
	StorageType             string `json:"StorageType,omitempty"`
	CACertificateIdentifier string `json:"CACertificateIdentifier,omitempty"`
	DBSubnetGroupName       string `json:"DBSubnetGroupName,omitempty"`
}
