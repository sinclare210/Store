package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string)(string,error){
	bytes,err := bcrypt.GenerateFromPassword([]byte(password),14)
	return string(bytes),err
}

func CheckHashedPassword(password string,hashedPassword string)error{
	err := bcrypt.CompareHashAndPassword([]byte(password),[]byte(hashedPassword))
	return err
}
