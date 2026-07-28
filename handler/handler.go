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
	SaveShortUrl(ctx context.Context, id int64) error
	SaveLongUrl(ctx context.Context, url *string) (int64, error)
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
	id, err := h.svc.SaveLongUrl(ctx, url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = h.svc.SaveShortUrl(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return nil
}
