package usermstcontroller

import "github.com/gin-gonic/gin"

func GetAllUsermstController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})

}
