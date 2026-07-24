package service

import "context"
import "github.com/edsown/urlzinha/business/repository"

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{repo: repo}
}

func (svc Service) SaveShortUrl(ctx context.Context, url *string) error {
	svc.repo.SaveShortUrl(ctx, url)
	return nil
}
