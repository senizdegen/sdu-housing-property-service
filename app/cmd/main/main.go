package main

import (
	"context"

	"github.com/julienschmidt/httprouter"
	"github.com/senizdegen/sdu-housing/property-service/internal/config"
	"github.com/senizdegen/sdu-housing/property-service/internal/property"
	"github.com/senizdegen/sdu-housing/property-service/internal/property/db"
	"github.com/senizdegen/sdu-housing/property-service/pkg/handlers/metric"
	"github.com/senizdegen/sdu-housing/property-service/pkg/logging"
	mongo "github.com/senizdegen/sdu-housing/property-service/pkg/mongodb"
)

func main() {
	logging.Init()
	logger := logging.GetLogger()
	logger.Println("logger initializing")

	logger.Println("config initializing")
	cfg := config.GetConfig()
	logger.Info(cfg)

	logger.Println("router initializing")
	router := httprouter.New()

	metricHandler := metric.Handler{Logger: logger}
	metricHandler.Register(router)

	mongoClient, err := mongo.NewClient(
		context.Background(),
		cfg.MongoDB.Host,
		cfg.MongoDB.Port,
		cfg.MongoDB.Username,
		cfg.MongoDB.Password,
		cfg.MongoDB.Database,
		cfg.MongoDB.AuthDB,
	)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Debug(mongoClient)
	propertyStorage := db.NewStorage(mongoClient, cfg.MongoDB.Collection, logger)

	propertyService, err := property.NewService(propertyStorage, logger)
	if err != nil {
		logger.Fatal(err)
	}

	propertyHandler := property.Handler{
		Logger:          logger,
		PropertyService: propertyService,
	}

	propertyHandler.Register(router)
	logger.Info("finished successfully")
}
