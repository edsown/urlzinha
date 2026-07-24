package repository

import (
	"context"
	"database/sql"
	"fmt"
)

type repository struct {
	db *sql.DB
}

type Repository interface {
	RetrieveOriginalUrlDB(ctx context.Context, db *sql.DB, id uint64) error
	RetrieveOriginalUrlCache(ctx context.Context, db *sql.DB, id uint64) error
	SaveShortUrl(ctx context.Context, url *string) error
	SaveLongUrl(ctx context.Context, db *sql.DB, url *string) (id uint64, error error)
}

func NewRepository(db *sql.DB) repository {
	return repository{db}
}
func (r repository) RetrieveOriginalUrlDB(ctx context.Context, db *sql.DB, id uint64) error {
	return nil
}

func (r repository) RetrieveOriginalUrlCache(ctx context.Context, db *sql.DB, id uint64) error {
	return nil
}

func (r repository) SaveShortUrl(ctx context.Context, url *string) error {
	query := `INSERT INTO ulrs (original_url, short_url, created_at) values (:1, :2, SYSDATE)`
	_, err := r.db.Exec(query, url, url)
	if err != nil {
		fmt.Errorf("error inserting into table %w", err)
	}
	return nil
}

func (r repository) SaveLongUrl(ctx context.Context, db *sql.DB, url *string) (id uint64, error error) {
	return 0, nil
}
