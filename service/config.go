package service

import (
	"log"

	"gopkg.in/yaml.v2"
)

// The Config stores the current server configuration.
// There are default values for each property but it's meant
// to be customized through a YAML file.
type Config struct {
	Port     int    `yaml:"port"`      // The port the service will use
	BasePath string `yaml:"base_path"` // The base path for the hooks
}

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

// LoadConfig reads the configuration file
// from the default location or falls back
// to the default config values
func LoadConfig() Config {
	var config Config

	confErr := yaml.Unmarshal([]byte(DefaultConfig), &config)

	if confErr != nil {
		log.Fatalf("Cannot unmarshal default config data! Error: %s", confErr)
	}

	return config
}
