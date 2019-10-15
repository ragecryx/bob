package service

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	builder "github.com/ragecryx/bob/builder"
	common "github.com/ragecryx/bob/common"
)

// StartServer loads all required configurations
// and initializes the http server
func StartServer() {
	flag.Parse()

	builder.ConfigureTasks(5)

	common.LoadConfig()
	common.LoadRecipes()
	SetupEndpoints()

	port := ":" + strconv.Itoa(common.GetConfig().Port)
	log.Printf("* Listening on %s\n", port)
	serveErr := http.ListenAndServe(port, nil)

	if serveErr != nil {
		log.Panicf(serveErr.Error())
	}
}
