package main

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
)

func buildGraphvizDiGraph(db *sql.DB, graphNodes []GraphConfig) {
	dotFile, err := os.Create("aws-si.dot")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dotFile.Close()

	graph := NewGraphvizDiGraph("Example")
	for _, nodeJSON := range graphNodes {
		node := NewGraphvizNode(nodeJSON.Name)
		for _, prop := range nodeJSON.Properties {
			node.addProperty(prop.Key, prop.Value)
		}
		graph.addNode(node)
	}

	arrayOfELBByCidrLoadBalancerName, err := loadAllViewELBByCidrDistinct(db)
	if err == nil {
		for _, loadBalancerName := range arrayOfELBByCidrLoadBalancerName {
			arrayOfELBInputCidr, err := loadAllViewElbInputCidrIps(db, loadBalancerName)
			if err == nil {
				fmt.Println("	", loadBalancerName)
				for _, elbInputCidr := range arrayOfELBInputCidr {
					cidrIp := cidrIpCut(elbInputCidr)
					loadBalancerNameTo, err := loadAllViewELBByCidrLoadBalancerName(db, cidrIp)
					if err == nil && loadBalancerNameTo != "" && loadBalancerName != loadBalancerNameTo {
						arrow := NewGraphvizArrow(graph.Nodes[loadBalancerNameTo], graph.Nodes[loadBalancerName])
						arrow.addProperty("label", elbInputCidr)
						arrow.addProperty("color", "blue")
						graph.addArrow(arrow)

					} else if err == nil && loadBalancerNameTo == "" {
						loadBalancerNameTo, err := loadAllViewInstanceByCidrInstanceName(db, cidrIp)
						if err == nil && loadBalancerNameTo != "" {
							arrow := NewGraphvizArrow(graph.Nodes[loadBalancerNameTo], graph.Nodes[loadBalancerName])
							arrow.addProperty("label", elbInputCidr)
							arrow.addProperty("color", "red")
							graph.addArrow(arrow)

						}
					} else {
						if err != nil {
							fmt.Println(err)
						}
					}
				}
			} else {
				if err != nil {
					fmt.Println(err)
				}
			}

		}
	} else {
		fmt.Println(err)
	}

	arrayOfViewDB, err := loadAllViewDBInstanceInputCidr(db)
	if err == nil {
		for _, viewDB := range arrayOfViewDB {
			loadBalancerName, err := loadAllViewELBByCidrLoadBalancerName(db, cidrIpCut(viewDB.SubnetCidrBlock))
			if err == nil && loadBalancerName != "" {
				arrow := NewGraphvizArrow(graph.Nodes[loadBalancerName], graph.Nodes[viewDB.DBInstanceName])
				arrow.addProperty("label", viewDB.SubnetCidrBlock)
				arrow.addProperty("color", "green")
				graph.addArrow(arrow)

			} else {
				if err == nil {
					loadBalancerName, err = loadAllViewInstanceByCidrInstanceName(db, cidrIpCut(viewDB.SubnetCidrBlock))
					if err == nil {
						arrow := NewGraphvizArrow(graph.Nodes[loadBalancerName], graph.Nodes[viewDB.DBInstanceName])
						arrow.addProperty("label", viewDB.SubnetCidrBlock)
						arrow.addProperty("color", "green")
						graph.addArrow(arrow)
					} else {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			}
		}
	}

	fmt.Printf("%s", graph.String(graphNodes))
	dotFile.WriteString(graph.String(graphNodes))

	fmt.Println("/bin/bash command-generate-graph.sh")
	err = exec.Command("/bin/bash", "command-generate-graph.sh").Run()

	if err != nil {
		fmt.Println("ERROR", err)
	} else {
		fmt.Println("PDF Created")
	}

}
