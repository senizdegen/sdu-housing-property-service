package property

import (
	"github.com/senizdegen/sdu-housing/property-service/pkg/logging"
)

var _ Service = &service{}

type service struct {
	storage Storage
	logger  logging.Logger
}

func NewService(propertyStorage Storage, logger logging.Logger) (Service, error) {
	return &service{
		storage: propertyStorage,
		logger:  logger,
	}, nil
}

type Service interface {
}
