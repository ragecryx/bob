package main

import (
	"fmt"

	svc "github.com/ragecryx/bob/service"
)

func main() {
	fmt.Printf("«It's better to do one job well than two jobs... not so well.»\n - Bob\n\n")
	svc.StartServer()
}
