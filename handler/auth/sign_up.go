package auth

import (
	messages "ECommerceGo/messages"
	models "ECommerceGo/models"

	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"

	"strconv"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Sign Up")
	//? Get the body of the request
	var _user models.UserModel
	_ = json.NewDecoder(r.Body).Decode(&_user)

	//? signin input validation
	_, _emailErr := mail.ParseAddress(_user.Email)
	if _emailErr != nil {
		messages.ErrorMessage(w, r, "Invalid Email")
		return
	} else if _user.Password == "" || len(_user.Password) < 6 {
		messages.ErrorMessage(w, r, "Invalid Password")
		return
	} else if _user.Name == "" || len(_user.Name) < 6 {
		messages.ErrorMessage(w, r, "Name is too short")
		return
	}
	_age, _ageErr := strconv.Atoi(_user.Age)
	if _ageErr != nil {
		messages.ErrorMessage(w, r, "Invalid Age")
		return
	} else if _age < 1 {
		messages.ErrorMessage(w, r, "Invalid Age")
		return
	}

	messages.SendMessage(w, r, "Success")
}
