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
	defer stmt.Close()
	_,err = stmt.Exec(&product.Name,&product.Description,&product.Price,&product.User_Id)


	if err != nil{
		return errors.New("wrong Execution")
	}
	return nil
}

func GetProducts()([]Product,error){
	query := `
	SELECT * FROM products
	`

	rows,err := db.DB.Query(query)
	if err != nil{
		return nil,errors.New("bad request")
	}

	var products []Product

	for rows.Next(){
		var product Product
		err := rows.Scan(&product.Id,&product.Name,&product.Description,&product.Price,&product.User_Id)
		if err != nil{
			return nil,err
		}
		products = append(products, product)
	}
	return products,nil

}

func GetProduct(id int64) (Product,error){
	query := `
	SELECT * FROM products WHERE Id = ?
	`
	rows := db.DB.QueryRow(query,id)

	var product Product

	err := rows.Scan(&product.Id,&product.Name,&product.Description,&product.Price,&product.User_Id)
	if err != nil{
		return product,err
	}

	return product,nil
	
}

func (product Product)DeleteProduct() error{
	query := `
	DELETE FROM products WHERE Id = ?
	`

	stmt,err := db.DB.Prepare(query)
	if err != nil{
		return errors.New("bad request")
	}
	defer stmt.Close()

	_,err = stmt.Exec(product.Id)
	if err != nil{
		return errors.New("bad request")
	}
	return nil

	
}


