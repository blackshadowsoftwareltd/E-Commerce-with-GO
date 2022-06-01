package routes

import (
	methods "ECommerceGo/methods"
	auth "ECommerceGo/methods/auth"
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
