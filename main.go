package main

import (
	"encoding/json"
	"github.com/golang/glog"
	"io/ioutil"
)

// Configuration
const ConfigFileName = "config.json"

// Default values
const DefaultPort = 8000
const DefaultRecordExpiryTimeInSec = 15 * 60
const DefaultMaxRecordCount = 10000
const DefaultRedisMasterUrl = ""

type Config struct {
	RecordExpiryTimeInSec int    `json:"recordTimeToExpireInSec"`
	RecordMaxCount        int    `json:"maxRecordCount"`
	PortNumber            int    `json:"port"`
	RedisMasterURL        string `json:"redisMasterUrl"`
}

func getConfig() Config {

	// Setup config with default parameters
	c := Config{
		PortNumber:            DefaultPort,
		RecordExpiryTimeInSec: DefaultRecordExpiryTimeInSec,
		RecordMaxCount:        DefaultMaxRecordCount,
		RedisMasterURL:        DefaultRedisMasterUrl}

	// Check for config file
	fileData, err := ioutil.ReadFile(ConfigFileName)
	if err != nil {
		glog.Errorf("Failed to open config file: %s", ConfigFileName)
		glog.Info("Using default values for config")
	}

	glog.Info("Parsing config file")
	err = json.Unmarshal(fileData, &c)
	if err != nil {
		glog.Errorf("Failed to parse config file: %s", ConfigFileName)
		glog.Info("Using default values for config")
	}

	return c
}

func main() {

	c := getConfig()

	bc, _ := json.Marshal(c)
	print("Config: ", string(bc))

}
