package convo

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"os"
)

type Config struct {
	PrimaryNodeAddress       net.IP
	SecondaryNodeAddress     net.IP
	GarbageCollectionTimeout int
	MaxThreads               int
	IsPrimary                bool
}

type configJSON struct {
	PnAddr     string `json:"primary-node-address"`
	SnAddr     string `json:"secondary-node-address"`
	GcTimeout  int    `json:"garbage-collection-timeout"`
	MaxThreads int    `json:"max-threads"`
	IsPrimary  bool   `json:"primary"`
}

var defaultFileLocation = "config/config.json"

var Configuration Config

func (cc *Config) jsonToConfig(cj configJSON) error {
	cc.GarbageCollectionTimeout = cj.GcTimeout
	cc.MaxThreads = cj.MaxThreads
	cc.IsPrimary = cj.IsPrimary
	cc.PrimaryNodeAddress = net.ParseIP(cj.PnAddr)
	cc.SecondaryNodeAddress = net.ParseIP(cj.SnAddr)

	return nil
}

func LoadConfiguration(location string) error {
	if location == "default" {
		location = defaultFileLocation
	}

	jsonFile, err := os.Open(location)
	if err != nil {
		log.Fatalf("Failed to read the config file %v: %v\n", location, err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)

	var c configJSON

	err = json.Unmarshal(byteValue, &c)
	if err != nil {
		log.Fatalf("Failed to unmarshal the config file: %v\n", err)
	}

	err = Configuration.jsonToConfig(c)
	if err != nil {
		log.Fatalf("Failed to load the Configuration: %v\n", err)
	}

	return nil
}
