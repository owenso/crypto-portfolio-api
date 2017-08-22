package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//Config : config file
type Config struct {
	Database struct {
		URI string `json:"uri,omitempty"`
	} `json:"database"`
}

// LoadConfiguration : Loads config from file
func LoadConfiguration(file string) (Config, error) {
	var config Config
	configFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return config, err
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err
}
