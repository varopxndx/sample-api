package controller_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/varopxndx/sample-api/controller"
	"github.com/varopxndx/sample-api/controller/mocks"
	"github.com/varopxndx/sample-api/model"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestController_GetSample(t *testing.T) {
	tests := []struct {
		name       string
		path       string
		resp       *model.SampleResponse
		statusCode int
		errMock    interface{}
	}{
		{
			name: "Success 200",
			path: "/v1/sample",
			resp: &model.SampleResponse{
				ID:   1,
				Name: "SomeName",
				Age:  30,
			},
			statusCode: http.StatusOK,
		},
		{
			name:       "Fail, not found",
			path:       "/v1/sample",
			resp:       nil,
			statusCode: http.StatusNotFound,
			errMock:    errors.New("bad...."),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			var err error
			logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

			mockUsecase := new(mocks.Usecase)
			mockUsecase.On("GetSample").Return(tt.resp, tt.errMock)

			rr := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(rr)

			ctx.Request, err = http.NewRequest(http.MethodGet, tt.path, nil)
			assert.Nil(t, err)

			c := controller.New(mockUsecase, logger)

			c.GetSample(ctx)
			assert.Equal(t, rr.Code, tt.statusCode)
		})
	}
}

func TestController_Ping(t *testing.T) {
	tests := []struct {
		name       string
		path       string
		resp       string
		statusCode int
	}{
		{
			name:       "Success 200",
			path:       "/v1/ping",
			resp:       "service healthy",
			statusCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			var err error
			logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

			rr := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(rr)

			ctx.Request, err = http.NewRequest(http.MethodGet, tt.path, nil)
			assert.Nil(t, err)

			c := controller.New(nil, logger)

			c.Ping(ctx)
			assert.Equal(t, rr.Code, tt.statusCode)
		})
	}
}
