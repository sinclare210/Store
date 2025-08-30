package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func load()(string){
	err := godotenv.Load()
	if err != nil{
		return ""
	}
	return os.Getenv("SECRET")
}

func GenerateToken(email string, user_id int64)(string,error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Email":email,
		"User_id":user_id,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(load()))
}

func VerifyToken(token string)(error){
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_,ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}


		return load(),nil
	})

	if err != nil{
		return errors.New("could not parse token")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid{
		return errors.New("invalid token")
	}

	// claims,ok := parsedToken.Claims.(jwt.MapClaims)
	// if !ok {
	// 	return errors.New("invalid token claims")
	// }

	// email := claims["Email"].(string)
	// user_id := claims["User_id"].(int64)

	return nil

}