package helper

import (
	data "ECommerceGo/data"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

///? Encode
func GenerateJWTToken(id string) (string, error) {
	fmt.Println("Generate Token")

	claims := jwt.NewWithClaims(
		jwt.SigningMethodHS256, //? method
		jwt.StandardClaims{
			Id:        id,                                    //? ID
			Issuer:    id,                                    //? Issuer
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //? Expire time
		},
	)

	token, err := claims.SignedString([]byte(data.SecretKey)) //? []byte(SecretKey)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return token, nil
}
