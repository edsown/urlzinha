package handler

import (
	"context"
	utils "github.com/edsown/urlzinha/utils"
	"net/http"
)

type Handler struct {
	svc Service
}
type Service interface {
	SaveShortUrl(ctx context.Context, url *string) error
}

func NewHandler(svc Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) HandleCreate(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	url := utils.GetValueFromQueryStr(r.URL.Query(), "url")
	if url == nil {
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}
	h.svc.SaveShortUrl(ctx, url)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	return nil
}
