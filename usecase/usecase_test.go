package usecase_test

import (
	"errors"
	"os"
	"testing"

	"github.com/varopxndx/sample-api/model"
	"github.com/varopxndx/sample-api/usecase"
	"github.com/varopxndx/sample-api/usecase/mocks"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestUsecase_GetSample(t *testing.T) {
	tests := []struct {
		name    string
		resp    *model.SampleResponse
		wantErr bool
		errMock interface{}
	}{
		{
			name: "Success",
			resp: &model.SampleResponse{
				ID:   1,
				Name: "SomeName",
				Age:  30,
			},
		},
		{
			name:    "Fail, not found",
			resp:    nil,
			wantErr: true,
			errMock: errors.New("GetSample: getting sample data"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

			mockService := new(mocks.Service)
			mockService.On("GetSample").Return(tt.resp, tt.errMock)

			u := usecase.New(mockService, logger)

			result, err := u.GetSample()

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Equal(t, err, tt.errMock)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, result, tt.resp)
			}
		})
	}
}
