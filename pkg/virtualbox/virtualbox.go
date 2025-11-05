package virtualbox

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) ListVMs(ctx context.Context) ([]*VM, error) {
	cmd := exec.CommandContext(ctx, "vboxmanage", "list", "vms", "--long")

	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to list instances: %w", err)
	}

	return parseVMs(output), nil
}

func (c *Client) UnregisterVM(ctx context.Context, id string) error {
	cmd := exec.CommandContext(ctx, "vboxmanage", "unregistervm", "--delete-all", id)

	var stderr bytes.Buffer

	cmd.Stderr = &stderr

	output, err := cmd.Output()
	if err != nil {
		if strings.Contains(stderr.String(), "Could not find a registered machine named") {
			return ErrVMNotFound
		}

		return fmt.Errorf("failed to delete instance: %w: %s", err, string(output))
	}

	return nil
}

func (c *Client) StartVM(ctx context.Context, id string) error {
	cmd := exec.CommandContext(ctx, "vboxmanage", "startvm", "--type", "headless", id)

	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to start VM: %w: %s", err, string(output))
	}

	return nil
}

func (c *Client) ShutdownVM(ctx context.Context, id string) error {
	cmd := exec.CommandContext(ctx, "vboxmanage", "controlvm", id, "shutdown")

	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to stop VM: %w: %s", err, string(output))
	}

	return nil
}

func (c *Client) ShowVMInfo(ctx context.Context, id string) (*VM, error) {
	cmd := exec.CommandContext(ctx, "vboxmanage", "showvminfo", "--machinereadable", id)

	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to show VM info: %w: %s", err, string(output))
	}

	return parseVMs(output)[0], nil
}
