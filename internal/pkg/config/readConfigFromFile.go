package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//readConfigFromFile reads configuration from a configuration file and unmarshals it into the provided config object
func readConfigFromFile(configFileName string, config interface{}) error {
	//Open the file
	configFile, err := os.Open(configFileName)
	if err != nil {
		fmt.Println("Config file error: ", err.Error())
		return err
	}
	defer configFile.Close()

	//Get the configuration
	configDataBytes, err := ioutil.ReadAll(configFile)
	if err != nil {
		fmt.Println("Unable to read config file data. Error: ", err.Error())
		return err
	}

	err = json.Unmarshal(configDataBytes, config)
	if err != nil {
		fmt.Println("Unable to read config data bytes to config object. Error:", err.Error())
		return err
	}

	return nil
}
