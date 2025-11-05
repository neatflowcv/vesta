package virtualbox

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/neatflowcv/vesta/internal/pkg/client"
	"github.com/neatflowcv/vesta/internal/pkg/domain"
	"github.com/neatflowcv/vesta/pkg/virtualbox"
)

var _ client.Client = (*Client)(nil)

type Client struct {
	client *virtualbox.Client
}

func NewClient() *Client {
	return &Client{
		client: virtualbox.NewClient(),
	}
}

func (r *Client) ListInstances(ctx context.Context) ([]*domain.Instance, error) {
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

func (r *Client) GetInstance(ctx context.Context, id string) (*domain.Instance, error) {
	vm, err := r.client.ShowVMInfo(ctx, id)
	if err != nil {
		if errors.Is(err, virtualbox.ErrVMNotFound) {
			return nil, client.ErrInstanceNotFound
		}

		return nil, fmt.Errorf("failed to get instance: %w", err)
	}

	return domain.NewInstance(vm.ID, vm.Name, mapVBStateToInstanceStatus(vm.Status)), nil
}

func (r *Client) StartInstance(ctx context.Context, id string) error {
	err := r.client.StartVM(ctx, id)
	if err != nil {
		if errors.Is(err, virtualbox.ErrVMNotFound) {
			return client.ErrInstanceNotFound
		}

		if errors.Is(err, virtualbox.ErrVMAlreadyLocked) {
			return client.ErrInstanceAlreadyRunning
		}

		return fmt.Errorf("failed to start instance: %w", err)
	}

	return nil
}

func (r *Client) StopInstance(ctx context.Context, id string) error {
	err := r.client.ShutdownVM(ctx, id)
	if err != nil {
		if errors.Is(err, virtualbox.ErrVMNotFound) {
			return client.ErrInstanceNotFound
		}

		if errors.Is(err, virtualbox.ErrVMNotRunning) {
			return client.ErrInstanceNotRunning
		}

		return fmt.Errorf("failed to stop instance: %w", err)
	}

	return nil
}

func (r *Client) DeleteInstance(ctx context.Context, id string) error {
	err := r.client.UnregisterVM(ctx, id)
	if err != nil {
		if errors.Is(err, virtualbox.ErrVMNotFound) {
			return client.ErrInstanceNotFound
		}

		return fmt.Errorf("failed to delete instance: %w", err)
	}

	return nil
}

func (r *Client) GetBase(ctx context.Context, id string) (*domain.Base, error) {
	vm, err := r.client.ShowVMInfo(ctx, id)
	if err != nil {
		if errors.Is(err, virtualbox.ErrVMNotFound) {
			return nil, client.ErrInstanceNotFound
		}

		return nil, fmt.Errorf("failed to get instance: %w", err)
	}

	return domain.NewBase(vm.ID, vm.Name, vm.CPU, vm.RAM, vm.Storage), nil
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
