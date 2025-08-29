package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine){
	server.GET("/product/:id", getProduct)
	server.GET("/product/", getProducts)
	server.PUT("/product/:id", updateProduct)
	server.DELETE("/product/:id", deleteProduct)
	server.POST("/product", createProduct)
	server.POST("/signup", signUp)
	server.POST("/login", login)
	server.GET("/user",getUsers)
}
