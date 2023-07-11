package repository

import (
	"context"

	"github.com/GabrielLoureiroGomes/basket-collection/core/domain"
)

// TeamDatabaseRepository is a interface used to manage team database interactions
type TeamDatabaseRepository interface {

	// InsertTeam used to save on database team data
	InsertTeam(ctx context.Context, team *domain.Team) error
}
