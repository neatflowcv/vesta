package flow

import (
	"context"
	"fmt"

	"github.com/neatflowcv/vesta/internal/pkg/domain"
	"github.com/neatflowcv/vesta/internal/pkg/repository"
)

type Service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) ListInstances(ctx context.Context) ([]*domain.Instance, error) {
	instances, err := s.repository.ListInstances(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list instances: %w", err)
	}

	return instances, nil
}
