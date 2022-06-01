package handler

import (
	auth "ECommerceGo/handler/auth"
	methods "ECommerceGo/methods"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRoutes() {
	router := mux.NewRouter()
	///? routes
	router.HandleFunc("/", methods.HomePage).Methods("GET")
	router.HandleFunc("/signUp", auth.SignUp).Methods("POST")
	router.HandleFunc("/signIn", auth.SignIn).Methods("POST")

	///?
	log.Fatal(http.ListenAndServe(":8080", router))
}
