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

// Context contains all the values needed by
// a page to be rendered.
type Context map[string]interface{}

func render(w http.ResponseWriter, page string, ctx Context) {
	tpl, err := ace.Load("web/base", fmt.Sprintf("web/%s", page), nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tpl.Execute(w, ctx); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handlePanel(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	xFwProto := r.Header.Get("X-Forwarded-Proto")

	prefix := "http://"
	if xFwProto == "https" {
		prefix = "https://"
	}

	apiPath := common.GetConfig().BasePath
	recipes := common.GetRecipes().All
	entries := make([]map[string]string, len(recipes))

	i := 0
	for k, v := range recipes {
		entries[i] = make(map[string]string)
		entries[i]["title"] = k
		entries[i]["repo"] = v.Repository.Name
		entries[i]["branch"] = v.Repository.Branch
		i++
	}

	render(w, "listing", Context{
		"recipes":  entries,
		"url_base": fmt.Sprintf("%s%s%s", prefix, host, apiPath),
	})
}

func handleRecipe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	recipeName := vars["recipe_name"]

	if val, ok := common.GetRecipes().All[recipeName]; ok {
		// Parse Github payload
		body, bodyErr := ioutil.ReadAll(r.Body)
		if bodyErr != nil {
			WriteError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse body. Error: %s", bodyErr))
			return
		}

		if builder.IsManualTrigger(val, body) {
			WriteBuildQueued(w, "Added by manual trigger")
			builder.Enqueue(val)
		} else if builder.IsGithubMerge(val, body) {
			WriteBuildQueued(w, "Added by Github merge event")
			builder.Enqueue(val)
		} else if builder.IsBitBucketMerge(val, body) {
			WriteBuildQueued(w, "Added by BitBucket merge event")
			builder.Enqueue(val)
		}
	} else {
		WriteError(w, http.StatusNotFound, fmt.Sprintf("Recipe '%s' does not exist", recipeName))
	}
}

// SetupEndpoints creates a router, registers
// both API and UI Panel endpoints and attaches
// the router to the global http request handler
func SetupEndpoints() {
	router := mux.NewRouter()

	router.HandleFunc("/", handlePanel)
	router.HandleFunc(common.GetConfig().BasePath+"/{recipe_name}", handleRecipe)

	http.Handle("/", router)
}
