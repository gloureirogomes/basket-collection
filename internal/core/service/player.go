package service

import (
	"context"

	"github.com/GabrielLoureiroGomes/basket-collection/internal/core/domain"
)

// Player interface is used to manage player functions
type Player interface {

	// InsertPlayer function is used to create a player register
	InsertPlayer(ctx context.Context, player domain.Player) error
}
