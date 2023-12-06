package repository

import (
	"context"

	"github.com/GabrielLoureiroGomes/basket-collection/core/domain"
)

// TeamDatabaseRepository is a interface used to manage team database interactions
type TeamDatabaseRepository interface {

	// InsertTeam used to save on database team data
	InsertTeam(ctx context.Context, team domain.Team) error

	//GetAll used to get all database team data
	GetAll(ctx context.Context) ([]domain.Team, error)

	//GetOne used to get one database team data
	GetOne(ctx context.Context, teamName string) (domain.Team, error)

	//Delete used to delete one database team data
	Delete(ctx context.Context, teamName string) error
}
