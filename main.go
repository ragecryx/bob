package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/ragecryx/bob/service"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s %s\n", r.Method, r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func main() {
	config := service.LoadConfig()

	PORT := ":" + strconv.Itoa(config.Port)
	args := os.Args

	if len(args) == 1 {
		fmt.Println("* Using default port: ", PORT)
	} else {
		PORT = ":" + args[1]
	}

	http.HandleFunc(config.BasePath, rootHandler)

	serveErr := http.ListenAndServe(PORT, nil)

	if serveErr != nil {
		fmt.Println(serveErr)
		os.Exit(1)
	}

}
