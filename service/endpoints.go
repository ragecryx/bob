package service

import (
	"fmt"
	"net/http"

	"github.com/yosssi/ace"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s %s\n", r.Method, r.URL.Path)
	fmt.Fprintf(w, "Loaded Recipes\n")

	for k := range GetRecipes().All {
		fmt.Fprintf(w, "* %s", k)
	}

	fmt.Printf("Served: %s\n", r.Host)
}

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

func SetupEndpoints() {
	config := GetConfig()

	http.HandleFunc(config.BasePath, rootHandler)
	http.HandleFunc(config.BasePath+"/panel", uiMain)
}
