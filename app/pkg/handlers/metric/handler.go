package metric

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/senizdegen/sdu-housing/property-service/pkg/logging"
)

const (
	URL = "/api/heartbeat"
)

type Handler struct {
	Logger logging.Logger
}

func (h *Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, URL, h.Heartbeat)
}

func (h *Handler) Heartbeat(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
