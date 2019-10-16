package service

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yosssi/ace"

	builder "github.com/ragecryx/bob/builder"
	common "github.com/ragecryx/bob/common"
)

func uiMain(w http.ResponseWriter, r *http.Request) {
	tpl, err := ace.Load("web/base", "web/admin", nil)

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

	if val, ok := common.GetRecipes().All[recipeName]; ok {
		// Parse Github payload
		body, bodyErr := ioutil.ReadAll(r.Body)
		if bodyErr != nil {
			fmt.Fprintf(w, "Failed to parse body. Error: %s", bodyErr)
			return
		}

		if builder.IsGithubMerge(val, body) {
			fmt.Fprintf(w, "Will build %s", val)
			builder.Enqueue(val)
		}
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
	router.HandleFunc(common.GetConfig().BasePath+"/{recipe_name}", runRecipe)

	http.Handle("/", router)
}
