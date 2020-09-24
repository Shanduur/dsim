package convo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

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
