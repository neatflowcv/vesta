package flow

import (
	"context"
	"errors"
	"fmt"

	"github.com/neatflowcv/vesta/internal/pkg/client"
	"github.com/neatflowcv/vesta/internal/pkg/domain"
)

type Service struct {
	client client.Client
}

func NewService(client client.Client) *Service {
	return &Service{
		client: client,
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
