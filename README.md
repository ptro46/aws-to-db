Procédure à suivre pour ajouter un nouveau service
==================================================

Dans un premier temps je te ferai les étapes 1 à 7, ensuite des que tu as compris le principe tu seras autonome.

1. Identifier la commande AWS à utiliser, par exemple la commande suivante renvoie la liste des VPC.
```shell
aws ec2 describe-vpcs --filters Name=tag:Name,Values=PledgerProdVPC
```

2. Placer cette commande dans un shell script
```shell
$ cat command-vpc.sh 
#!/bin/bash

aws ec2 describe-vpcs --filters Name=tag:Name,Values=PledgerProdVPC 
```

3. Editer le fichier de configuration de l'application et ajouter la commande à utiliser.
```json
$ cat pledger-aws-to-db-config.json 
{
	"db" : {
	} ,
	"commands" : [
		{ 
			"name" : "vpc",
			"cmd" : "command-vpc.sh"
		}
	]
}
```

4. Lancer la commande pour voir ce qu'elle retourne
```json
{
    "Vpcs": [
        {
            "VpcId": "vpc-08b4e6965f71d3be2", 
            "InstanceTenancy": "default", 
            "Tags": [
                {
                    "Value": "PledgerProdVPC", 
                    "Key": "Name"
                }, 
                {
                    "Value": "pledger-network", 
                    "Key": "aws:cloudformation:stack-name"
                }, 
                {
                    "Value": "PledgerProdVPC", 
                    "Key": "aws:cloudformation:logical-id"
                }, 
                {
                    "Value": "arn:aws:cloudformation:eu-west-1:490192481338:stack/pledger-network/a32e6a20-8939-11e9-b088-02c681d3a380", 
                    "Key": "aws:cloudformation:stack-id"
                }
            ], 
            "CidrBlockAssociationSet": [
                {
                    "AssociationId": "vpc-cidr-assoc-09079efbc5c0ffd97", 
                    "CidrBlock": "10.1.0.0/16", 
                    "CidrBlockState": {
                        "State": "associated"
                    }
                }
            ], 
            "State": "available", 
            "DhcpOptionsId": "dopt-880b6aee", 
            "OwnerId": "490192481338", 
            "CidrBlock": "10.1.0.0/16", 
            "IsDefault": false
        }
    ]
}
```

5. Créer la commande de création de la table dans la base de données en fonction des données que l'on souhaite concerver.
```sql
CREATE TABLE public.vpc (
	id 					bigserial,
	vpc_id 				text,
	vpc_name 			text,
	cidr_association_id	text,
	cidr_block 			text,
	cidr_block_state	text,
	vpc_state			text,
	owner_id 			text
) ;
```

6. Créer le code source correspondant aux données de la table (Fichier VPC.go)
```GoLang
package main

import (
	"fmt"
)

type VPC struct {
	Id                int64
	VpcId             string
	VpcName           string
	CidrAssociationId string
	CidrBlock         string
	CidrBlockState    string
	VpcState          string
	OwnerId           string
}

func NewVPC(id int64, vpcId string, vpcName string, cidrAssociationId string, cidrBlock string, cidrBlockState string, vpcState string, ownerId string) *VPC {
	return &VPC{
		Id:                id,
		VpcId:             vpcId,
		VpcName:           vpcName,
		CidrAssociationId: cidrAssociationId,
		CidrBlock:         cidrBlock,
		CidrBlockState:    cidrBlockState,
		VpcState:          vpcState,
		OwnerId:           ownerId}
}

func (d *VPC) String() string {
	return fmt.Sprintf("VPC id(%d) vpcId(%s) vpcName(%s) cidrBlock(%s)",
		d.Id,
		d.VpcId,
		d.VpcName,
		d.CidrBlock)
}
```

7. Créer le code source du DAO (Data Access Object) correspondant à la table (Fichier VPCDAO.go)
Dans ce code source nous allons retrouver les méthodes nécessaires : 
```GoLang
func createVPC(db *sql.DB, 
				VpcId string, 
				VpcName string, 
				CidrAssociationId string, 
				CidrBlock string, 
				CidrBlockState string, 
				VpcState string, 
				OwnerId string) *VPC { }
func loadVPCById(db *sql.DB, id int64) (*VPC, error) { }
func loadAllVPC(db *sql.DB) ([]*VPC, error) { }
func deleteAllVPC(db *sql.DB) error { }
```

Source complet.
```GoLang
package main

import (
	"database/sql"        // package SQL
	_ "github.com/lib/pq" // driver Postgres
	"log"
)

func rowResultSetToVPC(row *sql.Row) (*VPC, error) {
	var err error
	var Id int64
	var VpcId string
	var VpcName string
	var CidrAssociationId string
	var CidrBlock string
	var CidrBlockState string
	var VpcState string
	var OwnerId string

	err = row.Scan(&Id, &VpcId, &VpcName, &CidrAssociationId, &CidrBlock, &CidrBlockState, &VpcState, &OwnerId)
	if err != nil {
		return nil, err
	}
	return NewVPC(Id, VpcId, VpcName, CidrAssociationId, CidrBlock, CidrBlockState, VpcState, OwnerId), nil
}

func rowsNoFetchResultSetToVPC(rows *sql.Rows) (*VPC, error) {
	var err error
	var Id int64
	var VpcId string
	var VpcName string
	var CidrAssociationId string
	var CidrBlock string
	var CidrBlockState string
	var VpcState string
	var OwnerId string

	err = rows.Scan(&Id, &VpcId, &VpcName, &CidrAssociationId, &CidrBlock, &CidrBlockState, &VpcState, &OwnerId)
	if err != nil {
		return nil, err
	}
	return NewVPC(Id, VpcId, VpcName, CidrAssociationId, CidrBlock, CidrBlockState, VpcState, OwnerId), nil
}

func rowsResultSetToVPC(rows *sql.Rows) (*VPC, error) {
	var err error
	if rows.Next() {
		var err error
		var Id int64
		var VpcId string
		var VpcName string
		var CidrAssociationId string
		var CidrBlock string
		var CidrBlockState string
		var VpcState string
		var OwnerId string

		err = rows.Scan(&Id, &VpcId, &VpcName, &CidrAssociationId, &CidrBlock, &CidrBlockState, &VpcState, &OwnerId)
		if err != nil {
			return nil, err
		}
		return NewVPC(Id, VpcId, VpcName, CidrAssociationId, CidrBlock, CidrBlockState, VpcState, OwnerId), nil
	}
	return nil, err
}

func createVPC(db *sql.DB, VpcId string, VpcName string, CidrAssociationId string, CidrBlock string, CidrBlockState string, VpcState string, OwnerId string) *VPC {
	rows := db.QueryRow("insert into vpc(vpc_id,vpc_name,cidr_association_id,cidr_block,cidr_block_state,vpc_state,owner_id) values($1,$2,$3,$4,$5,$6,$7) returning id,vpc_id,vpc_name,cidr_association_id,cidr_block,cidr_block_state,vpc_state,owner_id", VpcId, VpcName, CidrAssociationId, CidrBlock, CidrBlockState, VpcState, OwnerId)

	vpc, err := rowResultSetToVPC(rows)
	if err != nil {
		return nil
	}
	return vpc
}

func loadVPCById(db *sql.DB, id int64) (*VPC, error) {
	rows, err := db.Query("select id,vpc_id,vpc_name,cidr_association_id,cidr_block,cidr_block_state,vpc_state,owner_id from vpc where id=$1")
	if err != nil {
		return nil, err
	}

	vpc, err := rowsResultSetToVPC(rows)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	return vpc, nil
}

func loadAllVPC(db *sql.DB) ([]*VPC, error) {
	rows, err := db.Query("select id,vpc_id,vpc_name,cidr_association_id,cidr_block,cidr_block_state,vpc_state,owner_id from vpc order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayOfVPC := make([]*VPC, 0, 0)
	for rows.Next() {
		vpc, err := rowsNoFetchResultSetToVPC(rows)
		if err == nil {
			arrayOfVPC = append(arrayOfVPC, vpc)
		} else {
			log.Println(err)
		}
	}

	return arrayOfVPC, nil
}

func deleteAllVPC(db *sql.DB) error {
	_, err := db.Exec("delete from vpc")

	return err
}
```

8. Ensuite il faut repartir du JSON et construire les structures nécessaires dans le code pour que le Unmarshalling puisse se faire, l'objectif est que tout le json puisse etre chargé par une seule ligne de code et que l'on retrouve toutes ses données dans des structures du language.


**Par convention nous allons nommer tous les sources qui correspondent a des structures JSON de AWS par AwsXXX.go en remplacant le XXX par le nom de l'objet tel qu'il est nommé dans les structures JSON renvoyées par AWS.**

Pour cela aller voir les codes sources suivants :
* AwsVPC.go
* AwsTag.go
* AwsCidrBlockAssociationSet.go
* AwsCidrBlockState.go

Si on regarde un morceau du JSON, par exemple celui de la liste des Tags on retrouve cette structure : 
```json
"Tags": [
                {
                    "Value": "PledgerProdVPC", 
                    "Key": "Name"
                }, 
                {
                    "Value": "pledger-network", 
                    "Key": "aws:cloudformation:stack-name"
                }, 
                {
                    "Value": "PledgerProdVPC", 
                    "Key": "aws:cloudformation:logical-id"
                }, 
                {
                    "Value": "arn:aws:cloudformation:eu-west-1:490192481338:stack/pledger-network/a32e6a20-8939-11e9-b088-02c681d3a380", 
                    "Key": "aws:cloudformation:stack-id"
                }
            ]
```

Il faut donc créer un objet AwsTag dans lequel on va retrouver les deux entrées du JSON : Key,Value
```GoLang
package main

import (
	"fmt"
)

type AwsTag struct {
	Key   string `json:"Key,omitempty"`
	Value string `json:"Value,omitempty"`
}

func (d AwsTag) String() string {
	return fmt.Sprintf("AwsTag Key(%s) Value(%s)", d.Key, d.Value)
}
```

Et comme c'est un tableau dans l'objet qui contient ce tableau de tags (AwsVPC.go) on va retrouver un tableau :
```GoLang
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type AwsVpc struct {
// cut
	Tags                    []AwsTag                     `json:"Tags,omitempty"`
// cut
}
```

9. Il nous reste à ajouter une fonction qui va effectuer le parsing du json et le chargement des éléments en base de données, ce code est dans l'objet racine correspondant au JSON chargé, ici AwsVpc.go
```GoLang
func awsVpcParseAndPersist(db *sql.DB, jsonString string) error {
	fmt.Printf("awsVpcParseAndPersist with json lenght %d\n", len(jsonString))
	awsVpcs := new(AwsVpcs)
	errUnmarshal := json.Unmarshal([]byte(jsonString), awsVpcs)
	if errUnmarshal == nil {
		for _, awsVpc := range awsVpcs.Vpcs {
			fmt.Printf("VpcId(%s) CidrBlock(%s)\n", awsVpc.VpcId, awsVpc.CidrBlock)
			vpcName := findVpcName(&awsVpc)
			vpc := createVPC(db,
				awsVpc.VpcId,
				vpcName,
				awsVpc.CidrBlockAssociationSet[0].AssociationId,
				awsVpc.CidrBlockAssociationSet[0].CidrBlock,
				awsVpc.CidrBlockAssociationSet[0].CidrBlockState.State,
				awsVpc.State,
				awsVpc.OwnerId)
			if vpc != nil {
				fmt.Printf("    VPC %s::%s loaded\n", awsVpc.VpcId, vpcName)
			} else {
				fmt.Printf("  ERROR  VPC %s::%s not loaded\n", awsVpc.VpcId, vpcName)
			}
		}
	} else {
		return errUnmarshal
	}
	return nil
}

func findVpcName(vpc *AwsVpc) string {
	for _, tag := range vpc.Tags {
		if tag.Key == "Name" {
			return tag.Value
		}
	}
	return ""
}
```

10. Le code qui effectue le chargement du JSON et son mapping dans les structures du langage se résume a ces deux lignes :
```GoLang
	awsVpcs := new(AwsVpcs)
	errUnmarshal := json.Unmarshal([]byte(jsonString), awsVpcs)

```

11. Le code qui suit l'exploite juste pour appeler la fonction de creation en base de données :
```GoLang
vpc := createVPC(db,
				awsVpc.VpcId,
				vpcName,
				awsVpc.CidrBlockAssociationSet[0].AssociationId,
				awsVpc.CidrBlockAssociationSet[0].CidrBlock,
				awsVpc.CidrBlockAssociationSet[0].CidrBlockState.State,
				awsVpc.State,
				awsVpc.OwnerId)
```

12. Pour intégrer ce code il ne reste qu'une chose à faire, c'est le référencer dans la fonction principale (pledger-aws-to-db.go) en faisant le lien entre le nom indiqué dans le fichier de configuration et la fonction à utiliser 
```json
{
	"db" : {
	} ,
	"commands" : [
		{ 
			"name" : "vpc",
			"cmd" : "command-vpc.sh"
		}
	]
}
```

13. Donc ici le nom de la commande est "vpc", ce nom se retrouve dans le code et on y associe la fonction de parsing / chargement :
```GoLang
	var cmdByName = map[string]pLoadFunc{
		"vpc":     awsVpcParseAndPersist,
		"subnets": awsVoid,
	}
```
(La fonction awsVoid ne fait rien)

14. Il ne reste qu'une chose a faire c'est appeller la fonction qui fait le ménage dans les données avant de les charger:
```GoLang
log.Println("Go")
deleteAllVPC(db)
```

Ajout du service de chargement des Subnets
==================================================
1. La commande a lancer pour avoir la liste des subnets
```shell
./command-subnets.sh
```

2. La liste des subnets est une répétition de la structure suivante 
```json
{
            "MapPublicIpOnLaunch": true, 
            "AvailabilityZoneId": "euw1-az1", 
            "Tags": [
                {
                    "Value": "SubnetDealsFrontEndZoneA", 
                    "Key": "aws:cloudformation:logical-id"
                }, 
                {
                    "Value": "SubnetDealsFrontEndZoneA", 
                    "Key": "Name"
                }, 
                {
                    "Value": "arn:aws:cloudformation:eu-west-1:490192481338:stack/pledger-network/a32e6a20-8939-11e9-b088-02c681d3a380", 
                    "Key": "aws:cloudformation:stack-id"
                }, 
                {
                    "Value": "pledger-network", 
                    "Key": "aws:cloudformation:stack-name"
                }
            ], 
            "AvailableIpAddressCount": 123, 
            "DefaultForAz": false, 
            "SubnetArn": "arn:aws:ec2:eu-west-1:490192481338:subnet/subnet-0ec92da7ddb7d90c3", 
            "Ipv6CidrBlockAssociationSet": [], 
            "VpcId": "vpc-08b4e6965f71d3be2", 
            "State": "available", 
            "AvailabilityZone": "eu-west-1a", 
            "SubnetId": "subnet-0ec92da7ddb7d90c3", 
            "OwnerId": "490192481338", 
            "CidrBlock": "10.1.35.0/25", 
            "AssignIpv6AddressOnCreation": false
        }
```

3. Création des scripts de base de données (je l'ai fait)
```sql
CREATE TABLE public.subnet (
	id 							bigserial,
	ref_vpc_id					bigint,
	subnet_id 					text,
	map_public_at_launch		boolean,
	availability_zone_id		text,
	name 						text,
	cloudformation_stack_name	text,
	available_ip_address_count	int,
	state 						text,
	availability_zone 			text,
	owner_id					text,
	cidr_block 					text 
) ;

ALTER TABLE public.subnet ADD CONSTRAINT pk_subnet PRIMARY KEY(id);

ALTER TABLE ONLY public.subnet ADD CONSTRAINT fk_subnet_vpc FOREIGN KEY (ref_vpc_id) REFERENCES public.vpc(id);

```

4. Création du code base de données dans Go
Je te laisse créer les codes sources qui suivent sur leur modele indiqué à coté
* Créer Subnet.go sur le modèle de VPC.go
* Créer SubnetDAO.go sur le modèle de VPCDAO.go

5. Création du code de mapping du JSON dans Go
Je te laisse créer les codes sources qui suivent sur leur modele indiqué à coté
* Créer AWSSubnet.go sur le modèle de AwsVPC.go

6. Dans AWSSubnet.go ajouter la fonction de parse du JSON et son chargement en base
Je te laisse créer les codes sources qui suivent sur leur modele indiqué à coté
* dans AWSSubnet.go ajouter
```GoLang
func awsSubnetParseAndPersist(db *sql.DB, jsonString string) error {
}
```

7. Il reste juste à tester :)


Plan d'implémentation des commandes aws
==================================================

1. command-vpc.sh
	* Liste des VPC.
	* Status : **Fait.**

2. command-subnets.sh
	* Liste des subnets.
	* Status : **Fait.**

3. command-buckets.sh
	* Liste des buckets.
	* Status : **Fait.**

4. command-ami.sh
	* Liste des images AMI.
	* Status : **Fait.**

5. command-stacks.sh
	* Liste des Stack Cloud Formation.
	* Status : **Fait.**

6. command-volumes.sh
	* Liste des Volumes.
	* Status : **Fait.**

7. command-security-groups.sh
	* Liste des Security Groups.
	* Status : **Fait.**

8. command-rds.sh
	* Liste des base de données RDS.
	* Status : **Fait.**

9. command-instances.sh
	* Liste des instances.
	* Status : **Fait.**

10. command-load-balancers.sh
	* Liste des ELB (Elastic Load Balancer).
	* Status : **Fait.**

11. command-auto-scaling-groups.sh
	* Liste des ASG (Auto Scaling Groups).
	* Status : **Fait.**

Evolutions
=============

Deux évolutions à prévoir, une qui va te demander genre 1h et une autre plus profonde.

1. Pourvoir ne charger qu'une des commandes sans les autres.

Actuellement tu as une map qui te donne la liste des commandes et tu les charges une par une apres avoir delete all de toutes les données 
```GoLang
	var cmdByName = map[string]pLoadFunc{
		"vpc":                 awsVpcParseAndPersist,
		"subnets":             awsSubnetParseAndPersist,
		"buckets":             awsBucketParseAndPersist,
		"amis":                awsAmiParseAndPersist,
		"stacks":              awsStackParseAndPersist,
		"volumes":             awsVolumeParseAndPersist,
		"security_groups":     awsSecurityGroupParseAndPersist,
		"db_instances":        awsDBInstanceParseAndPersist,
		"instances":           awsInstanceParseAndPersist,
		"load_balancers":      awsLoadBalancerParseAndPersist,
		"auto_scaling_groups": awsAutoScalingGroupParseAndPersist,
	}
```

On garde ce mode mais on va en introduire un second, qui est un chargement à la demande et une option de purge.
L'appel de l'outil se fait par 
```shell
$> pledger-aws-to-go fichier_de_configuration
```

il faudrait deja garder cette appel qui continue a faire ce qu'il fait aujourd'hui c'est a dire : 
* tout purger
* tout charger

Il faudrait ajouter un appel qui serait 
```shell
$> pledger-aws-to-go fichier_de_configuration --purge
```
Qui ne ferait que les delete et pas le chargement

et un appel qui serait
```shell
$> pledger-aws-to-go fichier_de_configuration --load vpc # charge que les VPC sans faire la purge
```

```shell
$> pledger-aws-to-go fichier_de_configuration --load subnets # charge que les subnets sans faire la purge
```

etc.

2. La corrélation des données

Ca je regarde ton doc References-aux-Donnees.md et je t'en dit plus dans la matinée, déjà tu peux commencer par le point 1.

Alors prenons un exemple : db_instance_db_subnet_group_subnet.subnet_identifier ----> Subnet
si je regarde ta table : 
```sql
CREATE TABLE public.db_instance_db_subnet_group_subnet (
	id bigserial,
	db_instance_db_subnet_group_id bigint,
	subnet_identifier text,
	name text,
	subnet_status text
) ;
```

Elle va donc devenir :
```sql
CREATE TABLE public.db_instance_db_subnet_group_subnet (
	id bigserial,
	db_instance_db_subnet_group_id bigint,
	subnet_identifier text,
	ref_subnet_id bigint,
	name text,
	subnet_status text
) ;

ALTER TABLE ONLY public.db_instance_db_subnet_group_subnet ADD CONSTRAINT fk_db_i_db_subnet_grp_subnet_subnet FOREIGN KEY (ref_subnet_id) REFERENCES public.subnet(id);

```

donc je modifie ce code en conséquence :
/// DBInstanceSubTables.go
```GoLang
type DBInstanceDBSubnetGroupSubnet struct {
	Id                      int64
	DbInstanceSubnetGroupId int64
	SubnetIdentifier        string
	RefSubnetId             int64
	Name                    string
	SubnetStatus            string
}

func NewDBInstanceDBSubnetGroupSubnet(Id int64, DbInstanceSubnetGroupId int64,
	SubnetIdentifier string, RefSubnetId int64, Name string, SubnetStatus string) *DBInstanceDBSubnetGroupSubnet {
	return &DBInstanceDBSubnetGroupSubnet{Id, DbInstanceSubnetGroupId,
		SubnetIdentifier, RefSubnetId, Name, SubnetStatus}
}
```

/// DBInstanceSubTablesDAO.go
```GoLang
func rowResultSetToDBInstanceDBSubnetGroupSubnet(row *sql.Row) (*DBInstanceDBSubnetGroupSubnet, error) {
	var err error
	var Id int64
	var DbInstanceSubnetGroupId int64
	var SubnetIdentifier string
	var RefSubnetId int64
	var Name string
	var SubnetStatus string

	err = row.Scan(&Id, &DbInstanceSubnetGroupId, &SubnetIdentifier, &RefSubnetId, &Name, &SubnetStatus)
	if err != nil {
		return nil, err
	}
	return NewDBInstanceDBSubnetGroupSubnet(Id, DbInstanceSubnetGroupId, SubnetIdentifier, RefSubnetId, Name, SubnetStatus), nil
}

func rowsNoFetchResultSetToDBInstanceDBSubnetGroupSubnet(rows *sql.Rows) (*DBInstanceDBSubnetGroupSubnet, error) {
	var err error
	var Id int64
	var DbInstanceSubnetGroupId int64
	var SubnetIdentifier string
	var RefSubnetId int64
	var Name string
	var SubnetStatus string

	err = rows.Scan(&Id, &DbInstanceSubnetGroupId, &SubnetIdentifier, &RefSubnetId, &Name, &SubnetStatus)
	if err != nil {
		return nil, err
	}
	return NewDBInstanceDBSubnetGroupSubnet(Id, DbInstanceSubnetGroupId, SubnetIdentifier, RefSubnetId, Name, SubnetStatus), nil
}

func rowsResultSetToDBInstanceDBSubnetGroupSubnet(rows *sql.Rows) (*DBInstanceDBSubnetGroupSubnet, error) {
	var err error
	if rows.Next() {
		var Id int64
		var DbInstanceSubnetGroupId int64
		var SubnetIdentifier string
		var RefSubnetId int64
		var Name string
		var SubnetStatus string

		err = rows.Scan(&Id, &DbInstanceSubnetGroupId, &SubnetIdentifier, &RefSubnetId, &Name, &SubnetStatus)
		if err != nil {
			return nil, err
		}
		return NewDBInstanceDBSubnetGroupSubnet(Id, DbInstanceSubnetGroupId, SubnetIdentifier, RefSubnetId, Name, SubnetStatus), nil
	}
	return nil, err

}
// Attention la subtilité ici est dans la clause ; returning il faut ajouter le ref_subnet_id apres le subnet_identifier
func createDBInstanceDBSubnetGroupSubnet(db *sql.DB, DbInstanceSubnetGroupId int64, SubnetIdentifier string, Name string, SubnetStatus string) *DBInstanceDBSubnetGroupSubnet {
	rows := db.QueryRow("insert into db_instance_db_subnet_group_subnet(db_instance_db_subnet_group_id,subnet_identifier,name,subnet_status) values($1,$2,$3,$4) returning id,db_instance_db_subnet_group_id,subnet_identifier,ref_subnet_id,name,subnet_status", DbInstanceSubnetGroupId, SubnetIdentifier, Name, SubnetStatus)

	dbInstanceDBSubnetGroupSubnet, err := rowResultSetToDBInstanceDBSubnetGroupSubnet(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return dbInstanceDBSubnetGroupSubnet
}
```

/// DBInstanceSubTablesDAO.go
```GoLang
func loadAllDBInstanceDBSubnetGroupSubnet(db *sql.DB) []*DBInstanceDBSubnetGroupSubnet {
	results := make([]*DBInstanceDBSubnetGroupSubnet, 0, 0)
	sqlSelect := "SELECT id,db_instance_db_subnet_group_id,subnet_identifier,ref_subnet_id,name,subnet_status from db_instance_db_subnet_group_subnet"

	rows, err := db.Query(sqlSelect)
	if err != nil {
		return results
	}
	defer rows.Close()

	for rows.Next() {
		dbInstanceDBSubnetGroupSubnet, err := rowsResultSetToDBInstanceDBSubnetGroupSubnet(rows)
		if err != nil {
			return make([]*DBInstanceDBSubnetGroupSubnet, 0, 0)
		}
		results = append(results, dbInstanceDBSubnetGroupSubnet)
	}
	return results
}
```

/// SubnetDAO.go
```GoLang
func loadSubnetByAWSReferenceId(db *sql.DB, subnetReferenceId string) (*Subnet, error) {
	rows, err := db.Query("select id,ref_vpc_id,subnet_id,map_public_at_launch,availability_zone_id,name,cloudformation_stack_name,available_ip_address_count,state,availability_zone,owner_id,cidr_block from subnet where subnet_id=$1", subnetReferenceId)
	if err != nil {
		return nil, err
	}

	subnet, err := rowsResultSetToSubnet(rows)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	return subnet, nil
}
```

/// DBInstanceSubTablesDAO.go
```GoLang
func updateDBInstanceDBSubnetGroupSubnetSubnetId(db *sql.DB, dbInstanceDBSubnetGroupSubnet *DBInstanceDBSubnetGroupSubnet, subnetId int64) *DBInstanceDBSubnetGroupSubnet {
	rows := db.QueryRow("update db_instance_db_subnet_group_subnet set ref_subnet_id=$1 where id=$2 returning id,db_instance_db_subnet_group_id,subnet_identifier,ref_subnet_id,name,subnet_status", subnetId, dbInstanceDBSubnetGroupSubnet.Id)

	dbInstanceDBSubnetGroupSubnet, err := rowResultSetToDBInstanceDBSubnetGroupSubnet(rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return dbInstanceDBSubnetGroupSubnet
}
```

/// pledger-aws-to-db.go
```GoLang
func awsCorrelation(db *sql.DB) {
	log.Println("Build relation for db_instance_db_subnet_group_subnet to subnet")
	arrayOfDBInstanceDBSubnetGroupSubnet := loadAllDBInstanceDBSubnetGroupSubnet(db)
	for _, oneDBInstanceDBSubnetGroupSubnet := range arrayOfDBInstanceDBSubnetGroupSubnet {
		subnet, err := loadSubnetByAWSReferenceId(db, oneDBInstanceDBSubnetGroupSubnet.SubnetIdentifier)
		if err == nil && subnet != nil {
			updateDBInstanceDBSubnetGroupSubnetSubnetId(db, oneDBInstanceDBSubnetGroupSubnet, subnet.Id)
		}
	}
}
```
