package main

type ViewDBInstanceInputCidr struct {
	DBInstanceId    int64
	DBInstanceName  string
	SubnetCidrBlock string
}

func NewViewDBInstanceInputCidr(DBInstanceId int64, DBInstanceName string, SubnetCidrBlock string) *ViewDBInstanceInputCidr {
	return &ViewDBInstanceInputCidr{DBInstanceId,
		DBInstanceName,
		SubnetCidrBlock}
}
