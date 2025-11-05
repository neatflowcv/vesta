package repository

import "context"

type Repository interface {
	ListBaseIDs(ctx context.Context) ([]string, error)
	DeleteBase(ctx context.Context, id string) error
	CreateBase(ctx context.Context, id string) error
}
