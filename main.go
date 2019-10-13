package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"

	svc "github.com/ragecryx/bob/service"
)

func main() {
	fmt.Printf("«It's better to do one job well than two jobs... not so well.»\n - Bob\n\n")

	// Load configuration and recipes files
	flag.Parse()
	config := svc.LoadConfig()
	svc.LoadRecipes()

	// Set endpoints
	svc.SetupEndpoints()

	// Startup the server
	port := ":" + strconv.Itoa(config.Port)
	fmt.Printf("* Listening on %s", port)
	serveErr := http.ListenAndServe(port, nil)

	if serveErr != nil {
		fmt.Println(serveErr)
		os.Exit(1)
	}
}
