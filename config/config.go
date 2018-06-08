package config

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/ghodss/yaml"
)

// ConfigPath is the path of config file
// ConfigFileName is the config file name
var (
	ConfigPath     = filepath.Join(os.Getenv("HOME"), ".map-test")
	ConfigFileName = "config"
)

// MapConfig parse config information from $HOME/.map-test/config
type MapConfig struct {
	APIKey          string `json:"apiKey"`
	ClientID        string `json:"clientID"`
	ClientSignature string `json:"clientSignature"`
}

// GetConfig read $HOME/.map-test/config and returns config informations in MapConfig struct
func GetConfig() (*MapConfig, error) {
	ret := &MapConfig{}
	configPath := filepath.Join(ConfigPath, ConfigFileName)
	log.Println("Reading config file from: ", configPath)
	fileByte, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(fileByte, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
