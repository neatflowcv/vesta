package client

import (
	"context"

	"github.com/neatflowcv/vesta/internal/pkg/domain"
)

type Client interface {
	ListInstances(ctx context.Context) ([]*domain.Instance, error)
	GetInstance(ctx context.Context, id string) (*domain.Instance, error)
	DeleteInstance(ctx context.Context, id string) error
	StartInstance(ctx context.Context, id string) error
	StopInstance(ctx context.Context, id string) error
	GetBase(ctx context.Context, id string) (*domain.Base, error)
}
