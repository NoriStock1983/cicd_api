package usermstcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsermstController(c *gin.Context) {
	c.JSON(http.StatusOK, "User List")

}
