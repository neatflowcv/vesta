package client

import (
	"context"

	"github.com/neatflowcv/vesta/internal/pkg/domain"
)

type Client interface {
	ListInstances(ctx context.Context) ([]*domain.Instance, error)
	DeleteInstance(ctx context.Context, id string) error
}
