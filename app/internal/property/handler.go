package property

import (
	"github.com/julienschmidt/httprouter"
	"github.com/senizdegen/sdu-housing/property-service/pkg/logging"
)

const (
	propertyURL  = "/api/property"
	propertysURL = "/api/property/:uuid"
)

type Handler struct {
	Logger          logging.Logger
	PropertyService Service
}

func (h *Handler) Register(router *httprouter.Router) {
	//register
}

//handler methods
