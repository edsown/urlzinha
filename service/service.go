package service

import "context"
import "github.com/edsown/urlzinha/business/repository"

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{repo: repo}
}

func (svc Service) SaveShortUrl(ctx context.Context, id int64) error {
	return svc.repo.SaveShortUrl(ctx, id)
}

func (svc Service) SaveLongUrl(ctx context.Context, url *string) (id int64, error error) {
	return svc.repo.SaveLongUrl(ctx, url)
}
