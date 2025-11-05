package flow

import (
	"context"
	"errors"
	"fmt"

	"github.com/neatflowcv/vesta/internal/pkg/client"
	"github.com/neatflowcv/vesta/internal/pkg/domain"
	"github.com/neatflowcv/vesta/internal/pkg/repository"
)

type Service struct {
	client client.Client
	repo   repository.Repository
}

func NewService(client client.Client, repo repository.Repository) *Service {
	return &Service{
		client: client,
		repo:   repo,
	}
}

func (s *Service) ListInstances(ctx context.Context) ([]*domain.Instance, error) {
	instances, err := s.client.ListInstances(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list instances: %w", err)
	}

	return instances, nil
}

func (s *Service) DeleteInstance(ctx context.Context, id string) error {
	err := s.client.DeleteInstance(ctx, id)
	if err != nil {
		if errors.Is(err, client.ErrInstanceNotFound) {
			return ErrInstanceNotFound
		}

		return fmt.Errorf("failed to delete instance: %w", err)
	}

	return nil
}

func (s *Service) StartInstance(ctx context.Context, id string) (*domain.Instance, error) {
	err := s.client.StartInstance(ctx, id)
	if err != nil {
		if errors.Is(err, client.ErrInstanceNotFound) {
			return nil, ErrInstanceNotFound
		}

		if errors.Is(err, client.ErrInstanceAlreadyRunning) {
			return nil, ErrInstanceAlreadyRunning
		}

		return nil, fmt.Errorf("failed to start instance: %w", err)
	}

	instance, err := s.client.GetInstance(ctx, id)
	if err != nil {
		if errors.Is(err, client.ErrInstanceNotFound) {
			return nil, ErrInstanceNotFound
		}

		return nil, fmt.Errorf("failed to get instance: %w", err)
	}

	return instance, nil
}

func (s *Service) StopInstance(ctx context.Context, id string) (*domain.Instance, error) {
	err := s.client.StopInstance(ctx, id)
	if err != nil {
		if errors.Is(err, client.ErrInstanceNotFound) {
			return nil, ErrInstanceNotFound
		}

		if errors.Is(err, client.ErrInstanceNotRunning) {
			return nil, ErrInstanceNotRunning
		}

		return nil, fmt.Errorf("failed to stop instance: %w", err)
	}

	instance, err := s.client.GetInstance(ctx, id)
	if err != nil {
		if errors.Is(err, client.ErrInstanceNotFound) {
			return nil, ErrInstanceNotFound
		}

		return nil, fmt.Errorf("failed to get instance: %w", err)
	}

	return instance, nil
}

func (s *Service) PromoteInstance(ctx context.Context, id string) (*domain.Base, error) {
	err := s.client.PromoteInstance(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to promote instance: %w", err)
	}

	err = s.repo.CreateBase(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to create base: %w", err)
	}

	s.client.GetBase(ctx, id)

	base, err := s.repo.GetBase(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get base: %w", err)
	}

	return instance, nil
}
