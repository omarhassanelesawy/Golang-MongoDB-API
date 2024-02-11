package handlers

import (
	"fmt"
	"ideanesttask/pkg/api/middleware"
	"ideanesttask/pkg/controllers"
	"ideanesttask/pkg/database/mongodb/models"
	"ideanesttask/pkg/database/mongodb/repository"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	repository.ConnectDB()
	db := "admin"
	collection := "users"
	coll := repository.GlobalConnection.Database(db).Collection(collection)

	// Parse the request body into the SigninRequest struct
	var req models.SigninRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var result models.User
	// Perform signin logic (authentication)
	if err := coll.FindOne(repository.GlobalContext, bson.M{"email": req.Email}).Decode(&result); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	// For demonstration purposes, let's assume the signin is successful
	token, err := controllers.GenerateAccessToken(result)
	if err != nil {
		log.Fatal(err)
	}

	// Create the response
	resp := models.SigninResponse{
		Message:      "Signin successful",
		AccessToken:  token,
		RefreshToken: token,
	}

	var access_token models.Token

	access_token.Value = token
	access_token.ExpiryTime = time.Now().Add(time.Hour * 24) // Token expires in 24 hours

	middleware.AddToken(access_token)

	// Return the response
	c.JSON(http.StatusOK, resp)
}

func RefreshHandler(c *gin.Context) {
	// Return the response
	c.JSON(http.StatusOK, nil)
}

func OrganizationHandler(c *gin.Context) {
	repository.ConnectDB()
	db := "admin"
	collection := "organizations"
	coll := repository.GlobalConnection.Database(db).Collection(collection)

	tokens := repository.GlobalConnection.Database(db).Collection("tokens")

	token := c.GetHeader("Authorization")
	// Validate bearer token
	filter := bson.M{"value": token}
	// Define projection to include only the value field
	projection := bson.M{"value": 1, "_id": 0}
	var result models.Token
	if err := tokens.FindOne(repository.GlobalContext, filter, options.FindOne().SetProjection(projection)).Decode(&result); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
		return
	}

	// Parse the request body into an Organization struct
	var org models.Organization
	if err := c.ShouldBindJSON(&org); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Perform any necessary validation
	insertResult, err := coll.InsertOne(repository.GlobalContext, org)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted document ID:", insertResult.InsertedID)

	// For demonstration purposes, just return a success message
	message := "Organization added successfully"
	resp := models.ResponseSchema{Message: message}

	// Return a JSON response
	c.JSON(http.StatusOK, resp)
}
