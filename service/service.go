package service

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/edsown/urlzinha/business"
)

type Repository interface {
	SaveUrl(ctx context.Context, tx *sql.Tx, shortUrl string, originalUrl string) error
	RetrieveUrl(ctx context.Context, shortUrl string) (string, error)
	BeginTx(ctx context.Context) (*sql.Tx, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (svc Service) InsertUrl(ctx context.Context, url string) (string, error) {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}

	shortcode := business.Encode(time.Now().UnixNano())
	tx, err := svc.repo.BeginTx(ctx)
	if err != nil {
		return "", fmt.Errorf("error at begintx:  %w", err)
	}
	err = svc.repo.SaveUrl(ctx, tx, shortcode, url)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	tx.Commit()

	return shortcode, nil
}

func (svc Service) RetrieveUrl(ctx context.Context, shortUrl string) (string, error) {
	originalUrl, err := svc.repo.RetrieveUrl(ctx, shortUrl)

	if err != nil {
		return "", fmt.Errorf("error retrieving originalUrl: %w", err)

	}
	return originalUrl, nil

}
