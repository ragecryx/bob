package service

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	builder "github.com/ragecryx/bob/builder"
)

// StartServer loads all required configurations
// and initializes the http server
func StartServer() {
	flag.Parse()

	builder.ConfigureTasks(5)

	LoadConfig()
	LoadRecipes()
	SetupEndpoints()

	port := ":" + strconv.Itoa(currentConfig.Port)
	log.Printf("* Listening on %s\n", port)
	serveErr := http.ListenAndServe(port, nil)

	if serveErr != nil {
		log.Panicf(serveErr.Error())
	}
}
