package virtualbox

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"

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

func (r *Repository) DeleteInstance(ctx context.Context, id string) error {
	cmd := exec.CommandContext(ctx, "vboxmanage", "unregistervm", "--delete-all", id)

	var stderr bytes.Buffer

	cmd.Stderr = &stderr

	output, err := cmd.Output()
	if err != nil {
		if strings.Contains(stderr.String(), "Could not find a registered machine named") {
			return repository.ErrInstanceNotFound
		}

		return fmt.Errorf("failed to delete instance: %w: %s", err, string(output))
	}

	log.Println(string(output))

	return nil
}
