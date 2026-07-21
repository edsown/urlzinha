package handler

import (
	"context"
	"database/sql"
	utils "github.com/edsown/urlzinha/utils"
	"net/http"
)

type Handler struct {
	repo Repository
	db   *sql.DB
	svc  Service
}

type Repository interface {
	RetrieveOriginalUrlDB(ctx context.Context, db *sql.DB, id uint64) error
	RetrieveOriginalUrlCache(ctx context.Context, db *sql.DB, id uint64) error
	SaveShortUrl(ctx context.Context, db *sql.DB, url *string) error
}

type Service interface {
	SaveShortUrl(ctx context.Context, db *sql.DB, url *string) error
}

func NewHandler(repo Repository, db *sql.DB, svc Service) *Handler {
	return &Handler{
		repo: repo,
		db:   db,
		svc:  svc,
	}
}

func (h *Handler) HandleCreate(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	url := utils.GetValueFromQueryStr(r.URL.Query(), "url")
	h.svc.SaveShortUrl(ctx, h.db, url)

	return nil
}
