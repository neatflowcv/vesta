package redis

import (
	"context"
	"fmt"

	"github.com/neatflowcv/vesta/internal/pkg/repository"
	goredis "github.com/redis/go-redis/v9"
)

const basesKey = "vesta:bases"

type Repository struct {
	client *goredis.Client
	key    string
}

func NewRepository(addr, password string, db int) (*Repository, error) {
	client := goredis.NewClient(&goredis.Options{ //nolint:exhaustruct
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	err := client.Ping(context.Background()).Err()
	if err != nil { // ensure connectivity
		return nil, fmt.Errorf("failed to ping Redis: %w", err)
	}

	return &Repository{
		client: client,
		key:    basesKey,
	}, nil
}

func (r *Repository) ListBaseIDs(ctx context.Context) ([]string, error) {
	ids, err := r.client.SMembers(ctx, r.key).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to list base IDs: %w", err)
	}

	return ids, nil
}

func (r *Repository) DeleteBase(ctx context.Context, id string) error {
	err := r.client.SRem(ctx, r.key, id).Err()
	if err != nil {
		return fmt.Errorf("failed to delete base: %w", err)
	}

	return nil
}

func (r *Repository) CreateBase(ctx context.Context, id string) error {
	err := r.client.SAdd(ctx, r.key, id).Err()
	if err != nil {
		return fmt.Errorf("failed to create base: %w", err)
	}

	return nil
}

var _ repository.Repository = (*Repository)(nil)
