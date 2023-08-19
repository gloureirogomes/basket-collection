package mock

import (
	"context"

	"github.com/GabrielLoureiroGomes/basket-collection/core/domain"
)

// TeamService interface is used to manage team functions
type TeamService interface {

	// InsertTeam function is used to create a team register
	InsertTeam(ctx context.Context, team *domain.Team) (*domain.Team, error)

	//GetAllTeams used to get all team data
	GetAllTeams(ctx context.Context) ([]*domain.Team, error)

	//GetOneTeam used to get one team data
	GetOneTeam(ctx context.Context, teamName string) (*domain.Team, error)

	//DeleteTeam used to delete one team data
	DeleteTeam(ctx context.Context, teamName string) error
}
