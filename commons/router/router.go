package router

import (
	usermstcontroller "cicd_api/controllers/usermstController"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	authentication := router.Group("/auth")
	{
		authentication.GET("/login", usermstcontroller.GetAllUsermstController) // Fix the function call
	}

	return router

}
