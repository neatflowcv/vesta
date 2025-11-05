package flow

import "errors"

var (
	ErrInstanceNotFound       = errors.New("instance not found")
	ErrInstanceNotRunning     = errors.New("instance is not running")
	ErrInstanceAlreadyRunning = errors.New("instance is already running")
)
