package virtualbox

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/neatflowcv/vesta/internal/pkg/domain"
	"github.com/neatflowcv/vesta/internal/pkg/repository"
	"github.com/neatflowcv/vesta/pkg/virtualbox"
)

var _ repository.Repository = (*Repository)(nil)

type Repository struct {
	client *virtualbox.Client
}

func NewRepository() *Repository {
	return &Repository{
		client: virtualbox.NewClient(),
	}
}

func (r *Repository) ListInstances(ctx context.Context) ([]*domain.Instance, error) {
	vms, err := r.client.ListVMs(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list VMs: %w", err)
	}

	var ret []*domain.Instance
	for _, vm := range vms {
		ret = append(ret, domain.NewInstance(vm.ID, vm.Name, mapVBStateToInstanceStatus(vm.Status)))
	}

	return ret, nil
}

func (r *Repository) DeleteInstance(ctx context.Context, id string) error {
	err := r.client.UnregisterVM(ctx, id)
	if err != nil {
		if errors.Is(err, virtualbox.ErrVMNotFound) {
			return repository.ErrInstanceNotFound
		}

		return fmt.Errorf("failed to delete instance: %w", err)
	}

	return nil
}

func mapVBStateToInstanceStatus(item string) domain.InstanceStatus {
	idx := strings.Index(item, "(")
	if idx != -1 {
		item = item[:idx]
	}

	item = strings.TrimSpace(item)
	switch item {
	case "running":
		return domain.InstanceStatusRunning
	case "powered off", "saved":
		return domain.InstanceStatusStopped
	case "stopping", "powering off", "aborted":
		return domain.InstanceStatusStopping
	case "starting", "powering on", "booting":
		return domain.InstanceStatusBooting
	default:
		return domain.InstanceStatusUnknown
	}
}
