package domain

import (
	"errors"
	"fmt"
)

var (
	// ErrBase is the error base
	ErrBase = errors.New("")

	// ErrNotFound is not found error
	ErrNotFound = fmt.Errorf("not found %w", ErrBase)
)
