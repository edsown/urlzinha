package service

import "fmt"
import "context"
import "database/sql"
import "github.com/edsown/urlzinha/business"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (svc Service) SaveShortUrl(ctx context.Context, db *sql.DB, url *string) error {
	business.Decode(*url)
	fmt.Println("decoding")
	return nil
}
