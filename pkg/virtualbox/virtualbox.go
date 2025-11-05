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

	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to list vms: %w: %s", err, stderr.String())
	}

	return parseVMs(stdout.Bytes()), nil
}

func (c *Client) UnregisterVM(ctx context.Context, id string) error {
	cmd := exec.CommandContext(ctx, "vboxmanage", "unregistervm", "--delete-all", id)

	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		if strings.Contains(stderr.String(), "Could not find a registered machine named") {
			return ErrVMNotFound
		}

		return fmt.Errorf("failed to delete instance: %w: %s", err, stderr.String())
	}

	return nil
}

func (c *Client) StartVM(ctx context.Context, id string) error {
	cmd := exec.CommandContext(ctx, "vboxmanage", "startvm", "--type", "headless", id)

	var stderr bytes.Buffer

	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		if strings.Contains(stderr.String(), "Could not find a registered machine named") {
			return ErrVMNotFound
		}

		if strings.Contains(stderr.String(), "is already locked by a session") {
			return ErrVMAlreadyLocked
		}

		return fmt.Errorf("failed to start VM: %w: %s", err, stderr.String())
	}

	return nil
}

func (c *Client) ShutdownVM(ctx context.Context, id string) error {
	cmd := exec.CommandContext(ctx, "vboxmanage", "controlvm", id, "shutdown")

	var stderr bytes.Buffer

	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		if strings.Contains(stderr.String(), "Could not find a registered machine named") {
			return ErrVMNotFound
		}

		if strings.Contains(stderr.String(), "is not currently running") {
			return ErrVMNotRunning
		}

		return fmt.Errorf("failed to stop VM: %w: %s", err, stderr.String())
	}

	return nil
}

func (c *Client) ShowVMInfo(ctx context.Context, id string) (*VM, error) {
	cmd := exec.CommandContext(ctx, "vboxmanage", "showvminfo", "--machinereadable", id)

	var (
		stderr bytes.Buffer
		stdout bytes.Buffer
	)

	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	err := cmd.Run()
	if err != nil {
		if strings.Contains(stderr.String(), "Could not find a registered machine named") {
			return nil, ErrVMNotFound
		}

		return nil, fmt.Errorf("failed to show VM info: %w: %s", err, stderr.String())
	}

	return parseVM(stdout.Bytes()), nil
}
