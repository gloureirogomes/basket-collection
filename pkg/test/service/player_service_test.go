package service

import (
	"context"
	"errors"
	"testing"

	"github.com/GabrielLoureiroGomes/basket-collection/core/domain"
	"github.com/stretchr/testify/assert"

	"github.com/GabrielLoureiroGomes/basket-collection/pkg/service"
	mocktest "github.com/GabrielLoureiroGomes/basket-collection/pkg/test/mock"
)

func TestInsertPlayer(t *testing.T) {
	givenCtx := context.Background()

	givenTeam := domain.Team{
		Name:       "Los Angeles Lakers",
		Conference: "West",
		State:      "California",
	}

	givenPlayer := domain.Player{
		Name:     "LeBron James",
		Age:      38,
		Position: "PF",
		Country:  "EUA",
		Team:     givenTeam,
		Height:   206,
		Weight:   113,
	}

	testCases := map[string]func(*testing.T, *mocktest.PlayerDatabaseRepositoryMock){
		"should save player data with success": func(t *testing.T, playerRepository *mocktest.PlayerDatabaseRepositoryMock) {
			service := service.NewPlayerService(playerRepository)

			playerRepository.On("InsertPlayer", givenCtx, givenPlayer).Return(nil)

			playerInserted, err := service.InsertPlayer(givenCtx, givenPlayer)
			assert.NotEmpty(t, playerInserted)
			assert.NoError(t, err)
		},
		"should return error when try to save empty team data": func(t *testing.T, playerRepository *mocktest.PlayerDatabaseRepositoryMock) {
			service := service.NewPlayerService(playerRepository)
			err := errors.New("error to insert empty data")
			givenEmptyPlayer := domain.Player{}

			playerRepository.On("InsertPlayer", givenCtx, givenEmptyPlayer).Return(err)

			playerInserted, err := service.InsertPlayer(givenCtx, givenEmptyPlayer)
			assert.Empty(t, playerInserted)
			assert.Error(t, err)
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			playerRepository := new(mocktest.PlayerDatabaseRepositoryMock)

			testCase(t, playerRepository)
		})
	}
}
