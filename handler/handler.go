package handler

import (
	methods "ECommerceGo/methods"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRoutes() {
	router := mux.NewRouter()
	///? routes
	router.HandleFunc("/", methods.HomePage).Methods("GET")

	///?
	log.Fatal(http.ListenAndServe(":8080", router))
}
