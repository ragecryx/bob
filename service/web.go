package service

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func StartServer() {
	flag.Parse()

	LoadConfig()
	LoadRecipes()
	SetupEndpoints()

	port := ":" + strconv.Itoa(currentConfig.Port)
	fmt.Printf("* Listening on %s\n", port)
	serveErr := http.ListenAndServe(port, nil)

	if serveErr != nil {
		fmt.Println(serveErr)
		os.Exit(1)
	}
}
