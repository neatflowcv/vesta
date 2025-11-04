package flow

import (
	"context"
	"errors"
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

func (s *Service) DeleteInstance(ctx context.Context, id string) error {
	err := s.repository.DeleteInstance(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrInstanceNotFound) {
			return ErrInstanceNotFound
		}

		return fmt.Errorf("failed to delete instance: %w", err)
	}

	return nil
}
