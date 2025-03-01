package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type pLoadFunc func(*sql.DB, string) error

func main() {
	if len(os.Args) == 1 {
		fmt.Println("No ConfigFileName was entered")
		return
	}
	var cmdByName = map[string]pLoadFunc{
		"vpc":                          awsVpcParseAndPersist,
		"subnets":                      awsSubnetParseAndPersist,
		"buckets":                      awsBucketParseAndPersist,
		"amis":                         awsAmiParseAndPersist,
		"stacks":                       awsStackParseAndPersist,
		"volumes":                      awsVolumeParseAndPersist,
		"snapshots":                    awsSnapshotParseAndPersist,
		"security_groups":              awsSecurityGroupParseAndPersist,
		"db_instances":                 awsDBInstanceParseAndPersist,
		"instances":                    awsInstanceParseAndPersist,
		"load_balancers":               awsLoadBalancerParseAndPersist,
		"load_balancers_target_groups": awsLoadBalancerTargetGroupParseAndPersist,
		"auto_scaling_groups":          awsAutoScalingGroupParseAndPersist,
		"auto_scaling_activities":      awsAutoScalingGroupActivitiesParseAndPersist,
	}

	var updateCmdByName = map[string]pLoadFunc{
		"instances":                    awsInstanceParseAndPersist,
		"load_balancers":               awsLoadBalancerParseAndPersist,
		"load_balancers_target_groups": awsLoadBalancerTargetGroupParseAndPersist,
		"auto_scaling_groups":          awsAutoScalingGroupParseAndPersist,
		"auto_scaling_activities":      awsAutoScalingGroupActivitiesParseAndPersist,
	}

	configFullFileName := os.Args[1]
	option := "all"
	if len(os.Args) >= 3 {
		option = os.Args[2]
	}

	exitCode := 0

	dbinfo, err := readFile(configFullFileName)
	if err == nil {
		configAwsToDBConfig := new(ConfigAwsToDBConfig)
		errUnmarshal := json.Unmarshal([]byte(dbinfo), configAwsToDBConfig)
		if errUnmarshal == nil {
			// fmt.Printf("%+v\n", configAwsToDBConfig)

			db, err := connectDB(&configAwsToDBConfig.DB)
			defer db.Close()
			if err == nil {

				switch option {
				case "all":
					err := awsPurge(db)
					if err == nil {
						awsLoadAll(db, cmdByName, configAwsToDBConfig.Commands)
					} else {
						fmt.Println(err)
					}
				case "purge":
					awsPurge(db)

				case "correlate":
					awsCorrelation(db)

				case "dump":
					awsDump(db)

				case "show":
					subset := ""
					filterName := ""
					if len(os.Args) == 4 {
						subset = os.Args[3]
						filterName = ""

					} else if len(os.Args) == 5 {
						subset = os.Args[3]
						filterName = os.Args[4]

					} else {
						fmt.Printf("aws-to-db aws-to-db-config.json show ELB [elbName]\n")
						os.Exit(exitCode)
					}
					awsShow(db, subset, filterName)

				case "update":
					awsPrepareUpdate(db)
					awsLoadAll(db, updateCmdByName, configAwsToDBConfig.Commands)

				case "graph":
					buildGraphvizDiGraph(db, configAwsToDBConfig.Graph)

				case "get":
					group := os.Args[3]  // vpc, subnets, ...
					filter := os.Args[4] // vpn_name, subnet_name, ....
					env := ""
					if len(os.Args) >= 6 {
						env = os.Args[5] // PREPROD, PROD, DEV2 ....
					}
					exitCode = awsGetId(db, group, filter, env)

				default:
					awsLoadOption(db, configAwsToDBConfig.Commands, cmdByName, option)

				}
			} else {
				fmt.Println("Error connectDB")
				fmt.Println(err)
			}

		} else {
			fmt.Printf("can not parse config file aws-to-db-config.json %s", errUnmarshal.Error())
		}
	} else {
		fmt.Println("can not read config file aws-to-db-config.json")
	}

	os.Exit(exitCode)
}

func cidrIpCut(cidrIp string) string {
	var cut string
	var i int16

	for _, c := range cidrIp {
		cut += string(c)
		if c == '.' {
			i++
		}
		if i == 3 {
			break
		}
	}

	return cut
}

func awsVoid(db *sql.DB, jsonString string) error {
	fmt.Printf("awsVoid with json lenght %d\n", len(jsonString))
	return nil
}

func readFile(configFilename string) (string, error) {
	contentOfFile, err := ioutil.ReadFile(configFilename)
	if err != nil {
		return "", err
	}
	return string(contentOfFile), nil
}
