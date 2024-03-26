package routes

import (
	"log"
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		log.Printf("Error signup user: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	user.ID = 1
	
	err = user.SaveAndUpdateUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User signed up!"})
}

func login(context *gin.Context) { 
	var user models.User

	err := context.ShouldBindJSON(&user)
	
	if err != nil {
		log.Printf("Error signup user: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.ValidateCredentials()
	
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful!"})
}