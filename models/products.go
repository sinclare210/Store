package models

import (
	"errors"

	"github.com/sinclare210/Store.git/db"
)

type Product struct{
	Id int64
	Name string `binding:"required"`
	Description string `binding:"required"`
	Price float64 `binding:"required"`
	User_Id int64
}

func (product Product)CreateProduct()error{
	query := `
	INSERT INTO products(Name,Description,Price,User_Id)
	VALUES(?,?,?,?)
	`
	stmt,err := db.DB.Prepare(query)
	if err != nil{
		return errors.New("wrong Statement")
	}
	_,err = stmt.Exec(&product.Name,&product.Description,&product.Price,&product.User_Id)


	if err != nil{
		return errors.New("wrong Execution")
	}

	
	return nil
}

