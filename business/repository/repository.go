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

func (r repository) RetrieveUrl(ctx context.Context, shortUrl string) (string, error) {
	query := `SELECT original_url FROM urls WHERE short_url = ?`
	var originalUrl string
	err := r.db.QueryRowContext(ctx, query, shortUrl).Scan(&originalUrl)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("url not found for %s", shortUrl)
		}
		return "", fmt.Errorf("error querying for original url: %w", err)
	}

	return originalUrl, nil
}
