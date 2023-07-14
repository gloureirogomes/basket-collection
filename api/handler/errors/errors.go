package errors

import (
	"fmt"

	"github.com/GabrielLoureiroGomes/basket-collection/core/domain"
)

var (
	// ErrUnknown is an unknown internal server error
	ErrUnknown = fmt.Errorf("internal server error %w", domain.ErrBase)

	// ErrBindParams is an error to bind parameters
	ErrBindParams = fmt.Errorf("error to bind parameters %w", domain.ErrBase)
)
