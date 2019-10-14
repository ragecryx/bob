package main

import (
	"fmt"

	builder "github.com/ragecryx/bob/builder"
	svc "github.com/ragecryx/bob/service"
)

func main() {
	fmt.Printf("«It's better to do one job well than two jobs... not so well.»\n - Bob\n\n")
	builder.ConfigureTasks(5)
	svc.StartServer()
}
