package service

import (
	"flag"
	"fmt"
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
	fmt.Printf("«It's better to do one job well than two jobs... not so well.»\n - Bob\n\n")

	common.InitLogging()

	config := common.LoadConfig()
	common.LoadRecipes()
	SetupEndpoints()

	builder.ConfigureTasks(config.TaskQueueSize)

	port := ":" + strconv.Itoa(common.GetConfig().Port)
	common.ServiceLog.Warnf("Listening on %s\n", port)
	serveErr := http.ListenAndServe(port, nil)

	if serveErr != nil {
		log.Panicf(serveErr.Error())
	}
}
