package service

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s %s\n", r.Method, r.URL.Path)
	fmt.Fprintf(w, "Loaded Recipes\n")

	for k, _ := range GetRecipes().All {
		fmt.Fprintf(w, "* %s", k)
	}

	fmt.Printf("Served: %s\n", r.Host)
}

func SetupEndpoints() {
	config := GetConfig()
	http.HandleFunc(config.BasePath, rootHandler)
}
