package property

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/senizdegen/sdu-housing/property-service/internal/apperror"
	"github.com/senizdegen/sdu-housing/property-service/pkg/logging"
)

const (
	propertysURL = "/api/property"
	propertyURL  = "/api/property/:uuid"
)

type Handler struct {
	Logger          logging.Logger
	PropertyService Service
}

func (h *Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodPost, propertysURL, apperror.Middleware(h.CreateProperty))
	router.HandlerFunc(http.MethodGet, propertyURL, apperror.Middleware(h.GetPropertyById))
	router.HandlerFunc(http.MethodGet, propertysURL, apperror.Middleware(h.GetAllProperty))
	router.HandlerFunc(http.MethodPut, propertyURL, apperror.Middleware(h.UpdateProperty))
	router.HandlerFunc(http.MethodDelete, propertyURL, apperror.Middleware(h.DeleteProperty))
}

func (h *Handler) CreateProperty(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *Handler) GetPropertyById(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *Handler) GetAllProperty(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *Handler) UpdateProperty(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *Handler) DeleteProperty(w http.ResponseWriter, r *http.Request) error {
	return nil
}
