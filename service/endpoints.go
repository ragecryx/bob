package service

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yosssi/ace"
)

func uiMain(w http.ResponseWriter, r *http.Request) {
	tpl, err := ace.Load("templates/base", "templates/admin", nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tpl.Execute(w, map[string]string{"Msg": "Hello Ace"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func runRecipe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	recipeName := vars["recipe_name"]

	if val, ok := loadedRecipes.All[recipeName]; ok {
		fmt.Fprintf(w, "Will build %s", val)
	} else {
		fmt.Fprintf(w, "Recipe not found!")
	}
}

// SetupEndpoints creates a router, registers
// both API and UI Panel endpoints and attaches
// the router to the global http request handler
func SetupEndpoints() {
	router := mux.NewRouter()

	router.HandleFunc("/", uiMain)
	router.HandleFunc(currentConfig.BasePath+"/{recipe_name}", runRecipe)

	http.Handle("/", router)
}
