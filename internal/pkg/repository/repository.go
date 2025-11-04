package repository

import (
	"context"

	"github.com/neatflowcv/vesta/internal/pkg/domain"
)

type Repository interface {
	ListInstances(ctx context.Context) ([]*domain.Instance, error)
	DeleteInstance(ctx context.Context, id string) error
}
