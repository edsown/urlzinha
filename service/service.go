package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/edsown/urlzinha/business"
)

type Repository interface {
	SaveUrl(ctx context.Context, tx *sql.Tx, shortUrl string, originalUrl string) error
	BeginTx(ctx context.Context) (*sql.Tx, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (svc Service) InsertUrl(ctx context.Context, url string) error {
	shortcode := business.Encode(time.Now().UnixNano())
	tx, err := svc.repo.BeginTx(ctx)
	if err != nil {
		return fmt.Errorf("error at begintx:  %w", err)
	}
	err = svc.repo.SaveUrl(ctx, tx, shortcode, url)
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()

	return nil
}
