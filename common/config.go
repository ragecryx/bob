package common

import (
	"flag"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// DefaultConfig contains the default server configuration
// in YAML format.
var DefaultConfig = `
# REST API
port: 9000
base_path: "/api"

# BUILDING
recipes_file_path: "config/recipes.yaml"
workspace_path: "workspace/"
builder_hostname: localhost
task_queue_size: 10
cleanup_builds: true
`

var (
	configFile    = flag.String("config", "./config.yaml", "Default configuration file")
	currentConfig = Config{}
)

// LoadConfig reads the configuration file
// and stores the data in related struct
func LoadConfig() *Config {
	var config Config
	var confErr error

	log.Printf("* Loading %s", *configFile)
	data, fileErr := ioutil.ReadFile(*configFile)

	if fileErr != nil {
		log.Printf("! Error: %s\n", fileErr)

		log.Printf("* Falling back to server defaults...\n")
		confErr = yaml.Unmarshal([]byte(DefaultConfig), &config)
	} else {
		confErr = yaml.Unmarshal(data, &config)
	}

	if confErr != nil {
		log.Panicf("Cannot unmarshal config data! Error: %s\n", confErr)
	}

	currentConfig = config

	return &currentConfig
}

// GetConfig provides te current
// server configuration object
func GetConfig() *Config {
	return &currentConfig
}