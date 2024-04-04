package router

import (
	companymstcontroller "cicd_api/controllers/companymstController"
	usermstcontroller "cicd_api/controllers/usermstController"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	authentication := router.Group("/auth")
	{
		authentication.GET("/login", usermstcontroller.GetAllUsermstController) // Fix the function call
	}
	companymst := router.Group("/companymst")
	{
		companymst.GET("/getallcompanymst", companymstcontroller.GetAllCompanymstController) // Fix the function call
	}

	return router

}
