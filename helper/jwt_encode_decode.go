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

///? Decode
func DecriptJWTTokenToId(tokenString string) (string, string, error) {
	fmt.Println("Decript Token")

	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(data.SecretKey), nil //? []byte(SecretKey)
	})

	if err != nil {
		fmt.Println("err ", err)
		return "", "", err
	}

	claims := token.Claims.(*jwt.StandardClaims)
	return claims.Id, claims.Issuer, nil
}
