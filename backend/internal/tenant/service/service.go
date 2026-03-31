package service

import (
	"context"

	"github.com/ferilee/api-idetech/backend/internal/tenant/domain"
)

type repository interface {
	FindBySlug(ctx context.Context, slug string) (domain.Tenant, error)
}

type Service struct {
	repository repository
}

func NewService(repository repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Bootstrap(ctx context.Context, slug string) (domain.Tenant, error) {
	return s.repository.FindBySlug(ctx, slug)
}
