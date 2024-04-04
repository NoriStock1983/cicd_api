package companymstcontroller

import (
	companymstdao "cicd_api/models/DAO/companymst_dao"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCompanymstController(c *gin.Context) {

	var companymst, error = companymstdao.GetAllCompanymst()

	if error != nil {
		c.JSON(http.StatusInternalServerError, error)
		return
	}

	if len(companymst) == 0 {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}
	c.JSON(http.StatusOK, companymst)

}
