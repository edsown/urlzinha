package repository

import (
	"context"
	"database/sql"
	"fmt"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) repository {
	return repository{db: db}
}

func (r repository) SaveUrl(ctx context.Context, tx *sql.Tx, shortUrl string, originalUrl string) error {
	query := `INSERT INTO urls (original_url, short_url, created_at) VALUES (?, ?, datetime('now'))`
	_, err := tx.Exec(query, originalUrl, shortUrl)
	if err != nil {
		return fmt.Errorf("error inserting into the database: %w", err)
	}
	return nil
}

func (r repository) BeginTx(ctx context.Context) (*sql.Tx, error) {
	return r.db.BeginTx(ctx, nil)
}
