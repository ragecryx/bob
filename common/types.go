package common

import (
	"fmt"
	"regexp"
)

// The Config stores the current server configuration.
// There are default values for each property but it's meant
// to be customized through a YAML file.
type Config struct {
	Port     int    `yaml:"port"`      // The port the service will use
	BasePath string `yaml:"base_path"` // The base path for the hooks

	RecipesFilePath string `yaml:"recipes_file_path"`
	WorkspacePath   string `yaml:"workspace_path"`
	BuilderHostname string `yaml:"builder_hostname"`
	TaskQueueSize   int    `yaml:"task_queue_size"`
	CleanupBuilds   bool   `yaml:"cleanup_builds"`
}

// Repository represents a code repo used in the
// building process
type Repository struct {
	Name   string `yaml:"name"`
	Branch string `yaml:"branch"`
	URL    string `yaml:"url"`
	SSH    *struct {
		KeyFile string `yaml:"keyfile"`
		// The environment variable that contains
		// the password
		PasswdEnv *string `yaml:"passenv"`
	} `yaml:"ssh"`
	VCS string `yaml:"vcs"`
}

// Recipe has a repo that is used to checkout the
// code and a command used to build it.
type Recipe struct {
	Repository Repository `yaml:"repository"`
	Command    string     `yaml:"command"`
}

// IsHostedIn checks if the repo resides (hosted)
// on a specific git service by checking the URL.
func (r Recipe) IsHostedIn(title string) bool {
	matched, _ := regexp.Match(fmt.Sprintf(".*%s\\..*", title), []byte(r.Repository.URL))
	return matched
}

// Recipes is the container struct for all the
// available recipes to the system
type Recipes struct {
	All map[string]Recipe `yaml:",inline"`
}
