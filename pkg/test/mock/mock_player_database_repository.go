package mock

import (
	"context"

	"github.com/GabrielLoureiroGomes/basket-collection/core/domain"
	"github.com/stretchr/testify/mock"
)

type PlayerDatabaseRepositoryMock struct {
	mock.Mock
}

// InsertPlayer used to save on database player data
func (m *PlayerDatabaseRepositoryMock) InsertPlayer(ctx context.Context, player domain.Player) error {
	arguments := m.Called(ctx, player)
	return arguments.Error(0)
}

// ListPlayers used to list all players
func (m *PlayerDatabaseRepositoryMock) ListPlayers(ctx context.Context) ([]domain.Player, error) {
	arguments := m.Called(ctx)
	return arguments.Get(0).([]domain.Player), arguments.Error(1)
}
