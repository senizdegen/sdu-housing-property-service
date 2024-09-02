package property

import (
	"context"
	"errors"
	"fmt"

	"github.com/senizdegen/sdu-housing/property-service/internal/apperror"
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
	GetMany(ctx context.Context) ([]Property, error)
}

func (s *service) GetMany(ctx context.Context) ([]Property, error) {
	property, err := s.storage.FindMany(ctx)

	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			return nil, err
		}
		return nil, fmt.Errorf("failed to find many property. error: %w", err)
	}

	return property, nil
}
