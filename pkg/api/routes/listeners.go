package routes

import (
	"ideanesttask/pkg/api/handlers"

	"github.com/gin-gonic/gin"
)

func RoutersUp(port string) {
	router := gin.Default()
	router.POST("/signup", handlers.SignUpHandler)
	router.POST("/signin", handlers.SignInHandler)
	router.POST("/refresh-token", handlers.RefreshHandler)
	router.POST("/organization", handlers.OrganizationHandler)
	router.Run(port)
}
