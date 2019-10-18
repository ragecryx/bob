package common

import (
	"flag"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var (
	recipesFile   *string
	recipesFlag   = flag.String("recipes", "./recipes.yaml", "Default recipes file")
	loadedRecipes = Recipes{}
)

// LoadRecipes reads the recipes file
// and stores the config in the related
// data structure
func LoadRecipes() *Recipes {
	// If recipes file is defined in config
	if len(*recipesFlag) > 0 {
		recipesFile = recipesFlag
	} else {
		*recipesFile = currentConfig.RecipesFilePath
	}

	data, err := ioutil.ReadFile(*recipesFile)

	if err != nil {
		commonLog.Fatalf("Cannot read recipes file %s Error: %s", *recipesFile, err)
	}

	var recipes Recipes
	yamlErr := yaml.Unmarshal(data, &recipes)

	if yamlErr != nil {
		commonLog.Fatalf("Cannot unmarshal recipes file! Error: %s", yamlErr)
	}

	loadedRecipes = recipes

	return &loadedRecipes
}

// GetRecipes provides the current
// recipes configuration object
func GetRecipes() *Recipes {
	return &loadedRecipes
}
