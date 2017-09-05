package config

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
	Secret string `json:"secret"`
}

// LoadConfiguration : Loads config from file
func LoadConfiguration() (Config, error) {
	var config Config
	configFile, err := os.Open("./config/config." + os.Getenv("ENV") + ".json")
	if err != nil {
		fmt.Println(err)
		return config, err
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err
}
