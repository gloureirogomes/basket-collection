package repository

import (
	"context"

	"github.com/GabrielLoureiroGomes/basket-collection/internal/core/domain"
)

// PlayerDatabaseRepository is a interface used to manage player database interactions
type PlayerDatabaseRepository interface {

	// InsertPlayer used to save on database player data
	InsertPlayer(ctx context.Context, player *domain.Player) error
}