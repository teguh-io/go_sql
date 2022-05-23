package configs

import (
	"encoding/json"
	"io/ioutil"
)

type DBConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

const (
	FILENAME = "configs/config.json"
)

func GetDBConfig() *DBConfig {
	configFile, err := ioutil.ReadFile(FILENAME)
	if err != nil {
		panic("Failed to open config file")
	}

	var dbConfig DBConfig
	err = json.Unmarshal(configFile, &dbConfig)
	if err != nil {
		panic("Failed to read config file")
	}

	return &dbConfig
}
