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
	InsertUrl(ctx context.Context, url string) error
	RetrieveUrl(ctx context.Context, shortUrl string) (string, error)
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
	err := h.svc.InsertUrl(ctx, *url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return nil
}

func (h *Handler) HandleRetrieve(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	shortUrl := r.PathValue("shortUrl")
	if shortUrl == "" {
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}
	originalUrl, err := h.svc.RetrieveUrl(ctx, shortUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	http.Redirect(w, r, originalUrl, http.StatusFound)
	return nil

}
