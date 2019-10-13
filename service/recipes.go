package service

import (
	"flag"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

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

var (
	recipesFile   = flag.String("recipes", "./recipes.yaml", "Default recipes file")
	loadedRecipes = Recipes{}
)

// LoadRecipes reads the recipes file
// and stores the config in the related
// data structure
func LoadRecipes() *Recipes {
	data, err := ioutil.ReadFile(*recipesFile)

	if err != nil {
		log.Fatalf("Cannot read recipes file %s Error: %s", *recipesFile, err)
	}

	var recipes Recipes
	yamlErr := yaml.Unmarshal(data, &recipes)

	if yamlErr != nil {
		log.Fatalf("Cannot unmarshal recipes file! Error: %s", yamlErr)
	}

	loadedRecipes = recipes

	return &loadedRecipes
}

// GetRecipes provides the current
// recipes configuration object
func GetRecipes() *Recipes {
	return &loadedRecipes
}
