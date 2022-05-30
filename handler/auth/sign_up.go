package auth

import (
	data "ECommerceGo/data"
	messages "ECommerceGo/messages"
	models "ECommerceGo/models"
	"time"

	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"

	"strconv"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Sign Up")
	_time := time.Now()

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

	for _, _u := range data.UserList {
		if _u.Email == _user.Email {
			messages.ErrorMessage(w, r, "Email already exists")
			return
		}
	}
	if len(data.UserList) == 0 {
		_user.Id = strconv.Itoa(0)
		_user.CreatedAt = _time
		_user.UpdatedAt = _time
		data.UserList = append(data.UserList, _user)
	} else {
		_user.Id = strconv.Itoa(len(data.UserList))
		_user.CreatedAt = _time
		_user.UpdatedAt = _time
		data.UserList = append(data.UserList, _user)
	}

	fmt.Println(_user)
	_user.Password = ""

	//? send response
	messages.Ok(w)
	json.NewEncoder(w).Encode(_user)

}
