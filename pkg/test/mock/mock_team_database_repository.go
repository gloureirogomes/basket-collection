package mock

import (
	"context"

	"github.com/GabrielLoureiroGomes/basket-collection/core/domain"
	"github.com/stretchr/testify/mock"
)

type TeamDatabaseRepositoryMock struct {
	mock.Mock
}

// InsertTeam used to save on database team data
func (m *TeamDatabaseRepositoryMock) InsertTeam(ctx context.Context, team *domain.Team) error {
	arguments := m.Called(ctx, team)
	return arguments.Error(0)
}

// GetAll used to get all database team data
func (m *TeamDatabaseRepositoryMock) GetAll(ctx context.Context) ([]*domain.Team, error) {
	arguments := m.Called(ctx)
	return arguments.Get(0).([]*domain.Team), arguments.Error(1)
}

// GetOne used to get one database team data
func (m *TeamDatabaseRepositoryMock) GetOne(ctx context.Context, teamName string) (*domain.Team, error) {
	arguments := m.Called(ctx, teamName)
	return arguments.Get(0).(*domain.Team), arguments.Error(1)
}
