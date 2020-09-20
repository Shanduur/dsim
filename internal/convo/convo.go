package convo

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"os"

	"github.com/Sheerley/pluggabl/pkg/plog"
)

// Config struct holds all the informations necessary to configure
// pluggabl apps (client, manager and worker)
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

// Configuration is a default instance of Config struct holding the data
// from loaded from the configuraion file.
var Configuration Config

func (cc *Config) jsonToConfig(cj configJSON) error {
	cc.GarbageCollectionTimeout = cj.GcTimeout
	cc.MaxThreads = cj.MaxThreads
	cc.IsPrimary = cj.IsPrimary
	cc.PrimaryNodeAddress = net.ParseIP(cj.PnAddr)
	cc.SecondaryNodeAddress = net.ParseIP(cj.SnAddr)

	return nil
}

// LoadConfiguration takes string with location of config in JSON format
// and then reads it's contents into the Configuration (Config struct instance)
func LoadConfiguration(location string) error {
	if location == "default" {
		location = defaultFileLocation
	}

	jsonFile, err := os.Open(location)
	if err != nil {
		plog.Fatalf(1, "Failed to read the config file %v: %v\n", location, err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)

	var c configJSON

	err = json.Unmarshal(byteValue, &c)
	if err != nil {
		plog.Fatalf(1, "Failed to unmarshal the config file: %v\n", err)
	}

	err = Configuration.jsonToConfig(c)
	if err != nil {
		plog.Fatalf(1, "Failed to load the Configuration: %v\n", err)
	}

	return nil
}
