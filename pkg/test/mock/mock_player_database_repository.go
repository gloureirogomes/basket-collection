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
