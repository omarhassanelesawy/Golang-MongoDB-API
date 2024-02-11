package handlers

import (
	"fmt"
	"ideanesttask/pkg/database/mongodb/models"
	"ideanesttask/pkg/database/mongodb/repository"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	repository.ConnectDB()
	db := "admin"
	collection := "users"
	coll := repository.GlobalConnection.Database(db).Collection(collection)

	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Perform any necessary validation
	insertResult, err := coll.InsertOne(repository.GlobalContext, req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted document ID:", insertResult.InsertedID)

	// For demonstration purposes, just return a success message
	message := "User signed up successfully"
	resp := models.ResponseSchema{Message: message}

	// Return a JSON response
	c.JSON(http.StatusOK, resp)
}


func SignInHandler(c *gin.Context) {
	// Parse the request body into the SigninRequest struct
	var req models.SigninRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Perform signin logic (authentication)
	// For demonstration purposes, let's assume the signin is successful
	accessToken := "generated_access_token"
	refreshToken := "generated_refresh_token"

	// Create the response
	resp := models.SigninResponse{
		Message:       "Signin successful",
		AccessToken:   accessToken,
		RefreshToken:  refreshToken,
	}

	// Return the response
	c.JSON(http.StatusOK, resp)
}
