package main

import (
	routes "ECommerceGo/handler/routes"
	"fmt"
)

func main() {
	fmt.Println("Started")

	//? start server
	routes.HandleRoutes()
}
