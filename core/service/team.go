package service

import (
	"context"

	"github.com/GabrielLoureiroGomes/basket-collection/core/domain"
)

// Team interface is used to manage team functions
type Team interface {

	// InsertTeam function is used to create a team register
	InsertTeam(ctx context.Context, team domain.Team) (*domain.Team, error)
}
