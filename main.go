package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sinclare210/Store.git/db"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.GET("/product/:id", getProduct)
	server.GET("/product/", getProducts)
	server.PUT("/product/:id", updateProduct)
	server.DELETE("/product/:id", deleteProduct)
	server.POST("/product", createProduct)
	server.POST("/signup", signUp)
	server.POST("/login", login)
	server.Run(":8080")
}

func getProduct(context *gin.Context) {

}

func getProducts(context *gin.Context) {

}

func updateProduct(context *gin.Context) {

}

func deleteProduct(context *gin.Context) {

}

func createProduct(context *gin.Context) {

}

func login(context *gin.Context) {

}

func signUp(context *gin.Context) {

}
