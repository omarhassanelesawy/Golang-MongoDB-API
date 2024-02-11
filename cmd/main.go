package main

import (
	"ideanesttask/pkg/api/routes"
	"ideanesttask/pkg/database/mongodb/models"
)

func main() {
	// router := gin.Default()
	// router.GET("/items", getRequest)
	// router.POST("/items", postRequest)
	// router.Run("localhost:4321")
	newUser := models.User{
		Email:    "john@example.com",
		Password: "password123",
	}
	models.PrintUserInfo(newUser)
	routes.RoutersUp("localhost:8080")

}
