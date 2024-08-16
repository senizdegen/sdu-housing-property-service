package property

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/senizdegen/sdu-housing/property-service/internal/apperror"
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
	router.HandlerFunc(http.MethodPost, propertyURL, apperror.Middleware(h.CreateProperty))
}

func (h *Handler) CreateProperty(w http.ResponseWriter, r *http.Request) error {
	return nil
}
