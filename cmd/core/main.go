package main

import (
	"fmt"
	"mood/internal/api"
)

func main() {
	a := api.New()
	e := a.RegisterRoutes()
	portStr := fmt.Sprintf(":%d", a.Port)

	e.Logger.Fatal(e.Start(portStr))
}
