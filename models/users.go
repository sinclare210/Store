package models

import (
	"errors"


	"github.com/sinclare210/Store.git/db"
	"github.com/sinclare210/Store.git/utils"
)

type User struct {
	Id int64
	Email string
	Password string
}

func (user User)CreateUser()error{
	query := `
	INSERT INTO users(Email,Password)
	VALUES(?,?)
	`

	stmt,err := db.DB.Prepare(query)
	if err != nil{
		return errors.New("wrong Statement")
	}
	defer stmt.Close()

	hashedPassword,err := utils.HashPassword(user.Password)
	if err != nil{
		return errors.New("failed to hash password")
	}

	_, err = stmt.Exec(&user.Email,&hashedPassword)
	if err != nil{
		return errors.New("wrong Execution")
	}
	return nil

}

func GetUsers()([]User,error){
	query := `
	SELECT * FROM users
	`

	rows, err := db.DB.Query(query)
	if err != nil{
		return nil,errors.New("bad request")
	}

	var users []User

	for rows.Next(){
		var user User

		err := rows.Scan(&user.Id,&user.Email,&user.Password)
		if err != nil{
			return nil,err
		}
		users = append(users, user)
	}
	return users,nil
}

func (user User)ValidCredentials()error{
	query := `
	SELECT Password FROM users WHERE Email = ?
	`

	rows := db.DB.QueryRow(query,user.Email)
	var HashedPassword string
	err := rows.Scan(&HashedPassword)
	if err != nil{
			return err
	}

	err = utils.CheckHashedPassword(HashedPassword,user.Password)
	if err != nil{
			return err
	}
	return nil
}

