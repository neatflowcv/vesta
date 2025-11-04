package virtualbox

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/neatflowcv/vesta/internal/pkg/domain"
	"github.com/neatflowcv/vesta/internal/pkg/repository"
)

var _ repository.Repository = (*Repository)(nil)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) ListInstances(ctx context.Context) ([]*domain.Instance, error) {
	cmd := exec.CommandContext(ctx, "vboxmanage", "list", "vms", "--long")

	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to list instances: %w", err)
	}

	parser := NewParser()
	instances := parser.Parse(output)

	return instances, nil
}
