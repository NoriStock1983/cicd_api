package usermstcontroller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllUsermstController(t *testing.T) {

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)

	GetAllUsermstController(c)

	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.EqualValues(t, "\"User List\"", response.Body.String())

}
