package usermstcontroller

import (
	"cicd_api/commons/dbaccess"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsermstController(c *gin.Context) {
	db, error := dbaccess.DBAccess()

	if error != nil {
		c.JSON(http.StatusInternalServerError, "DB Access Error")
		fmt.Println(error)
		return
	}
	defer db.Close()
	c.JSON(http.StatusOK, "User List")

}
