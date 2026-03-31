package service

import (
	"context"

	authdomain "github.com/ferilee/api-idetech/backend/internal/auth/domain"
)

type repository interface {
	ListByTenant(ctx context.Context, tenantSlug string) ([]authdomain.User, error)
}

type Service struct {
	repository repository
}

func NewService(repository repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) ListByTenant(ctx context.Context, tenantSlug string) ([]authdomain.User, error) {
	users, err := s.repository.ListByTenant(ctx, tenantSlug)
	if err != nil {
		return nil, err
	}

	for index := range users {
		users[index].PasswordHash = ""
	}

	return users, nil
}
