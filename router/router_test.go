package router_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/varopxndx/sample-api/router"
	"github.com/varopxndx/sample-api/router/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_New(t *testing.T) {
	testCases := []struct {
		name       string
		status     int
		path       string
		httpMethod string
		method     string
	}{
		{
			name:       "Ping, OK",
			status:     http.StatusOK,
			path:       "/v1/ping",
			httpMethod: http.MethodGet,
			method:     "Ping",
		},
		{
			name:       "GetSample, OK",
			status:     http.StatusOK,
			path:       "/v1/sample",
			httpMethod: http.MethodGet,
			method:     "GetSample",
		},
		{
			name:       "Not Found",
			status:     http.StatusNotFound,
			path:       "/v1/bad/endpoint",
			httpMethod: http.MethodGet,
			method:     "GetSample",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			recorder := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(recorder)

			mockController := new(mocks.Controller)
			mockController.On(tc.method, mock.Anything)

			var err error
			c.Request, err = http.NewRequest(tc.httpMethod, tc.path, nil)
			assert.Nil(t, err)

			r := router.New(mockController)

			r.ServeHTTP(recorder, c.Request)

			assert.Equal(t, tc.status, recorder.Code)
		})
	}
}
