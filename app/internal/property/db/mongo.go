package db

import (
	"github.com/senizdegen/sdu-housing/property-service/internal/property"
	"github.com/senizdegen/sdu-housing/property-service/pkg/logging"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ property.Storage = &db{}

type db struct {
	collection *mongo.Collection
	logger     logging.Logger
}

func NewStorage(storage *mongo.Database, collection string, logger logging.Logger) property.Storage {
	return &db{
		collection: storage.Collection(collection),
		logger:     logger,
	}
}

//methods
