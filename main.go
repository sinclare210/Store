package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sinclare210/Store.git/db"
	"github.com/sinclare210/Store.git/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
