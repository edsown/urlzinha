package handler

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
)

type Handler struct {
	repo Repository
	db   *sql.DB
	svc  Service
}

type Repository interface {
	retrieveOriginalUrlDB(ctx context.Context, db *sql.DB, id uint64) error
	retrieveOriginalUrlCache(ctx context.Context, db *sql.DB, id uint64) error
	saveShortUrl(ctx context.Context, db *sql.DB, url string) error
}

type Service interface {
	saveShortUrl(ctx context.Context, db *sql.DB, url string) error
}

func NewHandler(repo Repository, db *sql.DB, svc Service) *Handler {

	return &Handler{
		repo: repo,
		db:   db,
		svc:  svc,
	}
}

func handleCreate(r http.Request, w http.ResponseWriter) error {
	ctx := r.Context()
	// TODO: create an utils package to gather info from the url params
	fmt.Println(ctx)
	return nil
}
