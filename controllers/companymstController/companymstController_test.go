package companymstcontroller

import (
	"bytes"
	"cicd_api/models/DTO/companymst_dto"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllCompanymst(t *testing.T) {
	expected_data := []companymst_dto.Companymst{
		{Companycd: "0010", Company_name: "test", Company_nm_short: "test", Active_from_date: time.Date(2022, 07, 1, 0, 0, 0, 0, time.UTC), Active_to_date: time.Date(2022, 07, 1, 0, 0, 0, 0, time.UTC), Created_by: "", Created_date: time.Date(2022, 07, 1, 0, 0, 0, 0, time.UTC), Updated_by: "", Updated_date: time.Date(2022, 07, 1, 0, 0, 0, 0, time.UTC), Update_counter: 0},
	}
	byteProduct, _ := json.Marshal(expected_data)
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(
		http.MethodPost,
		"/api/mst/user/users",
		bytes.NewBuffer(byteProduct),
	)
	GetAllCompanymstController(c)

	var actual_data []companymst_dto.Companymst
	err := json.Unmarshal(response.Body.Bytes(), &actual_data)
	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.Nil(t, err)

	for i, expected := range expected_data {

		assert.EqualValues(t, expected.Companycd, actual_data[i].Companycd)
		assert.EqualValues(t, expected.Company_name, actual_data[i].Company_name)
		assert.EqualValues(t, expected.Company_nm_short, actual_data[i].Company_nm_short)
		//		assert.EqualValues(t, expected.Active_from_date, actual_data[i].Active_from_date)
		//		assert.EqualValues(t, expected.Active_to_date, actual_data[i].Active_to_date)
		assert.EqualValues(t, expected.Created_by, actual_data[i].Created_by)
		//		assert.EqualValues(t, expected.Created_date, actual_data[i].Created_date)
		assert.EqualValues(t, expected.Updated_by, actual_data[i].Updated_by)
		//		assert.EqualValues(t, expected.Updated_date, actual_data[i].Updated_date)
		assert.EqualValues(t, expected.Update_counter, actual_data[i].Update_counter)
	}
}
