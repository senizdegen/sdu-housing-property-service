package db

import (
	"context"
	"time"

	"github.com/senizdegen/sdu-housing/property-service/internal/property"
	"github.com/senizdegen/sdu-housing/property-service/pkg/logging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (s *db) FindMany(ctx context.Context) ([]property.Property, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var properties []property.Property

	findOptions := options.Find()
	//Set the limit of the number of record to find
	findOptions.SetLimit(5)
	//Define an array in which you can store the decoded documents

	cursor, err := s.collection.Find(ctx, bson.D{}, findOptions) //???
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		s.logger.Debug("decoding")
		var prop property.Property
		if err := cursor.Decode(&prop); err != nil {
			return nil, err
		}
		properties = append(properties, prop)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	s.logger.Debug("achieved")
	return properties, nil
}
