package main

import (
	handler "ECommerceGo/handler"
	"fmt"
)

func main() {
	fmt.Println("Started")

	//? start server
	handler.HandleRoutes()
}
