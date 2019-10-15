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

	config := common.LoadConfig()
	common.LoadRecipes()
	SetupEndpoints()

	builder.ConfigureTasks(config.TaskQueueSize)

	port := ":" + strconv.Itoa(common.GetConfig().Port)
	log.Printf("* Listening on %s\n", port)
	serveErr := http.ListenAndServe(port, nil)

	if serveErr != nil {
		log.Panicf(serveErr.Error())
	}
}
