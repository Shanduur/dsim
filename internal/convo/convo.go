// Package convo provides functions for reading configuratiion files
// and provides structure holding all the configuration.
package convo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

// Config struct holds all the informations necessary to configure
// pluggabl apps (client, manager and worker)
type Config struct {
	Type                     string
	ManagerAddress           net.IP
	ManagerPort              int
	WorkerAddress            net.IP
	WorkerPort               int
	GarbageCollectionTimeout int
	MaxThreads               int
}

type configJSON struct {
	Type       string `json:"type"`
	MAddr      string `json:"manager-address"`
	MPort      int    `json:"manager-port"`
	WAddr      string `json:"worker-address"`
	WPort      int    `json:"worker-port"`
	GcTimeout  int    `json:"garbage-collection-timeout"`
	MaxThreads int    `json:"max-threads"`
}

var defaultFileLocation = "config/config.json" // TODO: set default config location in ~/.config/

// Configuration is a default instance of Config struct holding the data
// from loaded from the configuraion file.
var Configuration Config

func (cc *Config) jsonToConfig(cj configJSON) {
	cc.GarbageCollectionTimeout = cj.GcTimeout
	cc.MaxThreads = cj.MaxThreads
	cc.Type = cj.Type
	cc.ManagerAddress = net.ParseIP(cj.MAddr)
	cc.ManagerPort = cj.WPort
	cc.ManagerAddress = net.ParseIP(cj.MAddr)
	cc.ManagerPort = cj.WPort
}

// LoadConfiguration takes string with location of config in JSON format
// and then reads it's contents into the Configuration (Config struct instance)
func LoadConfiguration(location string) error {
	if location == "default" {
		location = defaultFileLocation
	}

	jsonFile, err := os.Open(location)
	if err != nil {
		return fmt.Errorf("failed to read the config file %v: %v", location, err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return fmt.Errorf("failed to read config file: %v", err)
	}

	var c configJSON

	err = json.Unmarshal(byteValue, &c)
	if err != nil {
		return fmt.Errorf("failed to unmarshal the config file: %v", err)
	}

	Configuration.jsonToConfig(c)

	return nil
}
