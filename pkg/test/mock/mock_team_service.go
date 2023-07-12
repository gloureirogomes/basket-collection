package mock

import (
	"context"

	"github.com/GabrielLoureiroGomes/basket-collection/core/domain"
	"github.com/stretchr/testify/mock"
)

type TeamServiceMock struct {
	mock.Mock
}

// InsertTeam function is used to create a team register
func (m *TeamServiceMock) InsertTeam(ctx context.Context, team *domain.Team) (*domain.Team, error) {
	arguments := m.Called(ctx, team)
	return arguments.Get(0).(*domain.Team), arguments.Error(1)
}

// GetAllTeams used to get all database team data
func (m *TeamServiceMock) GetAllTeams(ctx context.Context) ([]*domain.Team, error) {
	arguments := m.Called(ctx)
	return arguments.Get(0).([]*domain.Team), arguments.Error(1)
}
