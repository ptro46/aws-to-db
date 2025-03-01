package main

import (
	"fmt"
)

type DbConfig struct {
	Login      string `json:"login,omitempty"`
	Password   string `json:"password,omitempty"`
	Host       string `json:"host,omitempty"`
	Port       int64  `json:"port,omitempty"`
	Parameters string `json:"parameters,omitempty"`
	Db         string `json:"db,omitempty"`
	Identifier string
	ShortName  string
}

func (config *DbConfig) dbPostgresConnectString(shortName string) string {
	config.ShortName = shortName
	config.Identifier = fmt.Sprintf("%s@%s:%d::%s", config.Login, config.Host, config.Port, config.Db)
	return fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%d %s", config.Login, config.Password, config.Host, config.Db, config.Port, config.Parameters)
}

func (d DbConfig) String() string {
	return fmt.Sprintf("DbConfig %s", fmt.Sprintf("%s@%s:%d::%s", d.Login, d.Host, d.Port, d.Db))
}
