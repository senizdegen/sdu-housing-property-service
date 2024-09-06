package property

import (
	"encoding/json"
	"fmt"
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
	h.Logger.Info("CREATE PROPERTY")
	w.Header().Set("Content-Type", "application/json")

	h.Logger.Debug("decode create property dto")
	var crProperty CreatePropertyDTO
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&crProperty); err != nil {
		return apperror.BadRequestError("invalid JSON scheme. check swagger API")
	}
	propertyUUID, err := h.PropertyService.Create(r.Context(), crProperty)
	if err != nil {
		return err
	}

	w.Header().Set("Location", fmt.Sprintf("%s/%s", propertyURL, propertyUUID))
	w.WriteHeader(http.StatusCreated)

	return nil
}

func (h *Handler) GetPropertyById(w http.ResponseWriter, r *http.Request) error {
	h.Logger.Info("GET PROPERTY BY ID")
	w.Header().Set("Content-Type", "application/json")

	params := httprouter.ParamsFromContext(r.Context())
	uuid := params.ByName("uuid")

	property, err := h.PropertyService.GetOne(r.Context(), uuid)
	if err != nil {
		return err
	}

	h.Logger.Debug("marshal property")
	propertyBytes, err := json.Marshal(property)
	if err != nil {
		return fmt.Errorf("failed to marshal property. err: %w", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(propertyBytes)

	return nil
}

func (h *Handler) GetAllProperty(w http.ResponseWriter, r *http.Request) error {
	h.Logger.Info("GET ALL PROPERTY")
	w.Header().Set("Content-Type", "application/json")

	property, err := h.PropertyService.GetMany(r.Context())
	if err != nil {
		return err
	}

	h.Logger.Debug("marshal property")
	propertyBytes, err := json.Marshal(property)
	if err != nil {
		return fmt.Errorf("failed to marshal property. err: %w", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(propertyBytes)

	return nil
}

func (h *Handler) UpdateProperty(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *Handler) DeleteProperty(w http.ResponseWriter, r *http.Request) error {
	return nil
}
