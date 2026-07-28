package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/edsown/urlzinha/business"
)

type repository struct {
	db *sql.DB
}

type Repository interface {
	RetrieveOriginalUrlDB(ctx context.Context, id int64) error
	RetrieveOriginalUrlCache(ctx context.Context, id int64) error
	SaveShortUrl(ctx context.Context, id int64) error
	SaveLongUrl(ctx context.Context, url *string) (id int64, error error)
}

func NewRepository(db *sql.DB) repository {
	return repository{db: db}
}
func (r repository) RetrieveOriginalUrlDB(ctx context.Context, id int64) error {
	return nil
}

func (r repository) RetrieveOriginalUrlCache(ctx context.Context, id int64) error {
	return nil
}

func (r repository) SaveShortUrl(ctx context.Context, id int64) error {
	shortUrl := business.Encode(id)
	fmt.Println(shortUrl)
	query := `UPDATE urls set short_url = ? WHERE id = ?`
	_, err := r.db.Exec(query, shortUrl, id)
	if err != nil {
		return fmt.Errorf("error inserting into the database: %w", err)
	}
	return nil
}

func (r repository) SaveLongUrl(ctx context.Context, url *string) (id int64, error error) {
	query := `INSERT INTO urls (original_url, created_at) VALUES (?, datetime('now'))`
	res, err := r.db.Exec(query, *url)
	if err != nil {
		return 0, fmt.Errorf("error inserting original url into database: %w", err)
	}
	id, err = res.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("error retrieving last insert id: %w", err)
	}
	return id, nil
}
