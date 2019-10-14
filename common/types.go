package types

// The Config stores the current server configuration.
// There are default values for each property but it's meant
// to be customized through a YAML file.
type Config struct {
	Port     int    `yaml:"port"`      // The port the service will use
	BasePath string `yaml:"base_path"` // The base path for the hooks
}

type Repository struct {
	Name   string `yaml:"name"`
	Branch string `yaml:"branch"`
	URL    string `yaml:"url"`
	VCS    string `yaml:"vcs"`
}

type Recipe struct {
	Repository Repository `yaml:"repository"`
	Command    string     `yaml:"command"`
}

type Recipes struct {
	All map[string]Recipe `yaml:",inline"`
}
