package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sinclare210/Store.git/models"
)


func getProduct(context *gin.Context) {
	id,err := strconv.ParseInt(context.Param("id"),10,64)
	if err != nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Invalid request"})
		return
	}

	product,err := models.GetProduct(id)
	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Database error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "product", "products": product})
}

func getProducts(context *gin.Context) {
	products, err := models.GetProducts()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "product", "products": products})
}

func updateProduct(context *gin.Context) {

}

func deleteProduct(context *gin.Context) {
	id,err := strconv.ParseInt(context.Param("id"),10,64)
	if err != nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Invalid request"})
		return
	}

	var product models.Product
	product.Id = id
	err  = product.DeleteProduct()
	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Databese error"})
		return
	}

	   context.JSON(http.StatusOK, gin.H{
        "message": "Product deleted successfully",
        "deleted_id": id,
    })
}

func createProduct(context *gin.Context) {
	var product models.Product
	err := context.ShouldBindJSON(&product)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid inputs"})
		return
	}
	err = product.CreateProduct()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Database error"})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"message": "Products Created!"})
}

func login(context *gin.Context) {

}

func signUp(context *gin.Context) {

}
