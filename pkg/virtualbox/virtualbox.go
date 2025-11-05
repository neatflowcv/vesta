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
