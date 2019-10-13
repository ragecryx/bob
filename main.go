package main

import (
	"fmt"
	"net/http"
	"os"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s %s\n", r.Method, r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func main() {
	PORT := ":9000"
	args := os.Args

	if len(args) == 1 {
		fmt.Println("* Using default port: ", PORT)
	} else {
		PORT = ":" + args[1]
	}

	http.HandleFunc("/", rootHandler)

	err := http.ListenAndServe(PORT, nil)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
