package main

import (
	"github.com/senizdegen/sdu-housing/property-service/internal/config"
	"github.com/senizdegen/sdu-housing/property-service/pkg/logging"
)

func main() {
	logging.Init()
	logger := logging.GetLogger()
	logger.Println("logger initializing")

	logger.Println("config initializing")
	cfg := config.GetConfig()
	logger.Info(cfg)
}
