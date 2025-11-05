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
