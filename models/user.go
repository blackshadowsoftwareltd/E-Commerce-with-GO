package models

import "time"

type UserModel struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Age       string    `json:"age"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
type UserModelResponse struct {
	Token string    `json:"token"`
	User  UserModel `json:"user"`
}
