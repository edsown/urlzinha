package service

import "fmt"
import "context"
import "github.com/edsown/urlzinha/business"
import "github.com/edsown/urlzinha/business/repository"

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{repo: repo}
}

func (svc Service) SaveShortUrl(ctx context.Context, url *string) error {
	business.Decode(*url)
	fmt.Println("decoding")
	return nil
}
