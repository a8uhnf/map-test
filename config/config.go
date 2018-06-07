package config

import (
	"log"
	"os"
	"path/filepath"
)

// ConfigPath is the path of config file
// ConfigFileName is the config file name
var (
	ConfigPath     = filepath.Join(os.Getenv("HOME"), ".map-test")
	ConfigFileName = "config"
)

func GetConfig() {
	log.Println("Reading config file from: ", filepath.Join(ConfigPath, ConfigFileName))
}
