package mock

import (
	"context"

	"github.com/GabrielLoureiroGomes/basket-collection/core/domain"
)

// PlayerService interface is used to manage players functions
type PlayerService interface {

	// InsertPlayer function is used to create a player register
	InsertPlayer(ctx context.Context, team *domain.Player) (*domain.Player, error)
}
