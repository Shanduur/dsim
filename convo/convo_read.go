package convo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"runtime/debug"
)

// LoadConfiguration takes string with location of config in JSON format
// and then reads it's contents into the Configuration (Config struct instance)
func LoadConfiguration(location string) (conf Config, err error) {
	jsonFile, err := os.Open(location)
	if err != nil {
		err = fmt.Errorf("failed to read the config file %v: %v", location, err)
		return
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		err = fmt.Errorf("failed to read config file: %v", err)
		return
	}

	var c configJSON

	err = json.Unmarshal(byteValue, &c)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal the config file: %v", err)
		return
	}

	conf.jsonToConfig(c)

	if conf.Type == "primary" || conf.Type == "secondary" {
		debug.SetGCPercent(conf.GarbageCollectionTimeout)
	}

	SavedConfig = conf

	return
}
