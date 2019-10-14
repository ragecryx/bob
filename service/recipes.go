package service

import (
	"flag"
	"io/ioutil"
	"log"

	types "github.com/ragecryx/bob/common"
	"gopkg.in/yaml.v2"
)

var (
	recipesFile   = flag.String("recipes", "./recipes.yaml", "Default recipes file")
	loadedRecipes = types.Recipes{}
)

// LoadRecipes reads the recipes file
// and stores the config in the related
// data structure
func LoadRecipes() *types.Recipes {
	data, err := ioutil.ReadFile(*recipesFile)

	if err != nil {
		log.Fatalf("Cannot read recipes file %s Error: %s", *recipesFile, err)
	}

	var recipes types.Recipes
	yamlErr := yaml.Unmarshal(data, &recipes)

	if yamlErr != nil {
		log.Fatalf("Cannot unmarshal recipes file! Error: %s", yamlErr)
	}

	loadedRecipes = recipes

	return &loadedRecipes
}

// GetRecipes provides the current
// recipes configuration object
func GetRecipes() *types.Recipes {
	return &loadedRecipes
}
