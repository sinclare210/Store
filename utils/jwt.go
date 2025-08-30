package utils

import (

	"time"

	"github.com/golang-jwt/jwt/v5"
	
)

const secret = "supersecret"



func GenerateToken(email string, user_id int64)(string,error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"Email": email,
		"User_Id":user_id,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secret))
	
}