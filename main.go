package main

import (
	"Project/routes"
)


func main() {
	routes := routes.InitRoutes()
	routes.Run("127.0.0.1:8888")
}
