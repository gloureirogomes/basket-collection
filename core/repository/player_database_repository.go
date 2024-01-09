package repository

import (
	"context"

	"github.com/GabrielLoureiroGomes/basket-collection/core/domain"
)

// PlayerDatabaseRepository is a interface used to manage player database interactions
type PlayerDatabaseRepository interface {

	// InsertPlayer used to save on database player data
	InsertPlayer(ctx context.Context, player domain.Player) error

	// ListPlayers used to list all players on database
	ListPlayers(ctx context.Context) ([]domain.Player, error)
}