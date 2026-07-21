package repository

import "database/sql"
import "context"

type repository struct {
	db *sql.DB
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

func (r repository) SaveShortUrl(ctx context.Context, db *sql.DB, url *string) error {
	return nil
}
