package models

import "time"

type UserModel struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Password string    `json:"password"`
	Age      string    `json:"age"`
	Date     time.Time `json:"date"`
}
