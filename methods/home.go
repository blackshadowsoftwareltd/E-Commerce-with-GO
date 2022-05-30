package methods

import (
	"fmt"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Page") //? link : 127.0.0.1:8080/
}
