package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type AwsLoadBalancerListener struct {
	Protocol        string                                  `json:"Protocol,omitempty"`
	DefaultActions  []AwsLoadBalancerListenerDefaultActions `json:"DefaultActions,omitempty"`
	SslPolicy       string                                  `json:"SslPolicy,omitempty"`
	Certificates    []AwsLoadBalancerListenerCertificats    `json:"Certificates,omitempty"`
	LoadBalancerArn string                                  `json:"LoadBalancerArn,omitempty"`
	Port            int64                                   `json:"Port,omitempty"`
	ListenerArn     string                                  `json:"ListenerArn,omitempty"`
}

type AwsLoadBalancerListenerDefaultActions struct {
	TargetGroupArn string `json:"TargetGroupArn,omitempty"`
	Type           string `json:"Type,omitempty"`
}

type AwsLoadBalancerListenerCertificats struct {
	CertificateArn string `json:"CertificateArn,omitempty"`
	IsDefault      bool   `json:"IsDefault,omitempty"`
}

type AwsListeners struct {
	Listeners []AwsLoadBalancerListener `json:"Listeners,omitempty"`
}

func awsLoadBalancerListenersParseAndPersist(db *sql.DB, jsonString string, loadBalancerId int64) error {
	fmt.Printf("	awsLoadBalancerListenersParseAndPersist with json lenght %d\n", len(jsonString))
	awsListeners := new(AwsListeners)
	errUnmarshal := json.Unmarshal([]byte(jsonString), awsListeners)
	if errUnmarshal == nil {
		for _, awsListener := range awsListeners.Listeners {
			listener := createLoadBalancerListener(db, loadBalancerId, awsListener.LoadBalancerArn, awsListener.ListenerArn, awsListener.SslPolicy, awsListener.Port, awsListener.Protocol)
			if listener != nil {
				fmt.Printf("    		LoadBalancerListener %s loaded\n", awsListener.ListenerArn)
				for _, defaultAction := range awsListener.DefaultActions {
					fmt.Printf("    		    LoadBalancerListenerDefaultAction %s::%s loaded\n", defaultAction.Type, defaultAction.TargetGroupArn)
					createLoadBalancerListenerDefaultAction(db, listener.Id, defaultAction.TargetGroupArn, defaultAction.Type)
				}
				for _, certificat := range awsListener.Certificates {
					fmt.Printf("    		    LoadBalancerListenerCertificat %s loaded\n", certificat.CertificateArn)
					createLoadBalancerListenerCertificat(db, listener.Id, certificat.CertificateArn)
				}
			} else {
				fmt.Printf("  ERROR  		LoadBalancerListener %s loaded\n", awsListener.ListenerArn)
			}
		}
	} else {
		return errUnmarshal
	}
	return nil
}
