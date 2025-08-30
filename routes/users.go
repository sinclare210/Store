package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sinclare210/Store.git/models"
	"github.com/sinclare210/Store.git/utils"
)

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "400 Bad Request"})
		return
	}

	err = user.ValidCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Credentials"})
		return
	}

	token, err := utils.GenerateToken(user.Email,user.Id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "400 Bad Request"})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"message": "Login Successful","token":token})
}

func signUp(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "400 Bad Request"})
		return
	}

	err = user.CreateUser()
		if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "500 Internal Server Error"})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"message": "User Created!"})

}

func getUsers(context *gin.Context) {
	users, err := models.GetUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "500 Internal Server Error"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "users", "products": users})
}




