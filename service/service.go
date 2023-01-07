package service

import "github.com/varopxndx/sample-api/model"

// Service contains service client
type Service struct{}

// New returns a service struct
func New() *Service {
	return &Service{}
}

// GetSample gets sample data
func (s *Service) GetSample() (*model.SampleResponse, error) {
	// mocked data
	response := &model.SampleResponse{
		ID:   1,
		Name: "SomeName",
		Age:  30,
	}

	return response, nil
}
