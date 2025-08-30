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
		context.JSON(http.StatusBadRequest,gin.H{"message":"400 Bad Request"})
		return
	}

	product,err := models.GetProduct(id)
	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"500 Internal Server Error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "product", "products": product})
}

func getProducts(context *gin.Context) {
	products, err := models.GetProducts()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "500 Internal Server Error"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "product", "products": products})
}

func updateProduct(context *gin.Context) {
	id,err := strconv.ParseInt(context.Param("id"),10,64)
	if err != nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"400 Bad Request"})
		return
	}

	var product models.Product
	product.Id = id

	err = context.ShouldBindJSON(&product)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "400 Bad Request"})
		return
	}

	err = product.UpdateProduct()
	  if err != nil {
        context.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
        return
    }
	context.JSON(http.StatusOK, gin.H{"message": "product updated", "products": product})

}

func deleteProduct(context *gin.Context) {
	id,err := strconv.ParseInt(context.Param("id"),10,64)
	if err != nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"400 Bad Request"})
		return
	}

	var product models.Product
	product.Id = id
	err  = product.DeleteProduct()
	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":err.Error()})
		return
	}

	   context.JSON(http.StatusOK, gin.H{
        "message": "Product deleted successfully",
        "deleted_id": id,
    })
}

func createProduct(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == ""{
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return
	}
	var product models.Product
	err := context.ShouldBindJSON(&product)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "400 Bad Request"})
		return
	}
	err = product.CreateProduct()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "500 Internal Server Error"})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"message": "Products Created!"})
}


