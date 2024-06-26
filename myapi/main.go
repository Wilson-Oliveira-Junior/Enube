package main

import (
	"myapi/routes"
)

func main() {
	r := routes.SetupRoutes()
	r.Run(":8080")
}
