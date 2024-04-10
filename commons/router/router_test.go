package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_Router(t *testing.T) {
	test_success_urls := []string{
		"/auth/login",
	}

	router := Router()

	w := httptest.NewRecorder()

	for getdatacount := 0; getdatacount < len(test_success_urls)-1; getdatacount++ {
		req, _ := http.NewRequest("GET", test_success_urls[getdatacount], nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	}
}
