package usecase

import (
	"github.com/varopxndx/sample-api/model"

	"github.com/rs/zerolog"
)

// Service has the service layer methods
type Service interface {
	GetSample() (*model.SampleResponse, error)
}

// Usecase structure
type Usecase struct {
	service Service
	logger  zerolog.Logger
}

// New creates a usecase
func New(service Service, logger zerolog.Logger) *Usecase {
	return &Usecase{
		service: service,
		logger:  logger,
	}
}

// GetSample gets sample data
func (u *Usecase) GetSample() (*model.SampleResponse, error) {
	// bussiness logic
	response, err := u.service.GetSample()
	if err != nil {
		u.logger.Error().Msg("GetSample: getting sample data")
		return nil, err
	}

	return response, nil
}
