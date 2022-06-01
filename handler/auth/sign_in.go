package auth

import (
	data "ECommerceGo/data"
	helper "ECommerceGo/helper"
	messages "ECommerceGo/messages"
	models "ECommerceGo/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Sign In")
	var _emailFlag, _passwordFlag bool = false, false
	//? Get the body of the request
	var _user models.UserModel
	_ = json.NewDecoder(r.Body).Decode(&_user)

	//? signin input validation
	if !helper.IsEmailValid(_user.Email) {
		messages.ErrorMessage(w, r, "Invalid Email")
		return
	} else if !helper.IsPasswordValid(_user.Password) {
		messages.ErrorMessage(w, r, "Invalid Password")
		return
	}
	for _, _u := range data.UserList {
		if _user.Email == _u.Email {
			_emailFlag = true
			if helper.CheckPasswordAndHash(_user.Password, _u.Password) {
				_passwordFlag = true
				_user = _u
			}
			break
		}
	}
	if _emailFlag == false {
		messages.ErrorMessage(w, r, "Email not matches")
		return
	} else if _passwordFlag == false {
		messages.ErrorMessage(w, r, "Password not matches")
		return
	}
	_token, _tokenErr := helper.GenerateJWTToken(_user.Id)
	if _tokenErr != nil {
		fmt.Println("err ", _tokenErr)
		messages.ErrorMessage(w, r, "Failed Generating Token")
		return
	}
	_user.Id = ""
	_user.Password = ""
	_userModelResponse := models.UserModelResponse{Token: _token, User: _user}
	//? send response
	messages.Ok(w)
	json.NewEncoder(w).Encode(_userModelResponse)
	fmt.Println("Sign In Success")
}
