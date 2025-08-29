package db

import (
	"database/sql"

	_"github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB(){
	var err error
	DB, err = sql.Open("sqlite3","api.db")
	if err != nil{
		panic("Could not connect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables(){
	createUsersTable := `CREATE TABLE IF NOT EXISTS users(
	Id INTEGER PRIMARY KEY AUTOINCREMENT,
	Email TEXT NOT NULL,
	Password TEXT NOT NULL
)
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil{
		panic("Could not create user table")
	}

	createProductsTable := `CREATE TABLE IF NOT EXISTS products(
	Id INTEGER PRIMARY KEY AUTOINCREMENT,
	Name TEXT NOT NULL,
	Description TEXT NOT NULL,
	Price REAL NOT NULL, 
	User_Id INTEGER,
	FOREIGN KEY(User_Id) REFERENCES users(Id)
)
`

	_, err = DB.Exec(createProductsTable)
	if err != nil{
		panic("Could not create product table")
	}
}
