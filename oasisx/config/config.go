package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

//ConfigList DB
type ConfigList struct {
	DbDriverName   string
	DbName         string
	DbUserName     string
	DbUserPassword string
	DbHost         string
	DbPort         string
	//ServerPort     int
}

//Config ConfigList
var Config ConfigList

//init test
func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		DbDriverName:   cfg.Section("db").Key("db_driver_name").String(),
		DbName:         cfg.Section("db").Key("db_name").String(),
		DbUserName:     cfg.Section("db").Key("db_user_name").String(),
		DbUserPassword: cfg.Section("db").Key("db_user_password").String(),
		DbHost:         cfg.Section("db").Key("db_host").String(),
		DbPort:         cfg.Section("db").Key("db_port").String(),
		//ServerPort:     cfg.Section("api").Key("server_port").MustInt(),
	}
}
