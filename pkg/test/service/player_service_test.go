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
	}

	givenPlayer := domain.Player{
		Name:     "LeBron James",
		Age:      38,
		Position: "PF",
		Team:     givenTeam,
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

func TestListPlayers(t *testing.T) {
	givenCtx := context.Background()
	givenPlayers := []domain.Player{
		{
			Name:     "LeBron James",
			Age:      39,
			Position: "SF",
			Team: domain.Team{
				Name:       "Los Angeles Lakers",
				Conference: "West",
			},
		},
		{
			Name:     "Stephen Curry",
			Age:      33,
			Position: "PG",
			Team: domain.Team{
				Name:       "Golden State Warriors",
				Conference: "West",
			},
		},
	}

	testCases := map[string]func(*testing.T, *mocktest.PlayerDatabaseRepositoryMock){
		"should return error when an error occurs on database": func(t *testing.T, playerRepository *mocktest.PlayerDatabaseRepositoryMock) {
			unexpectedError := errors.New("unexpected error")
			service := service.NewPlayerService(playerRepository)

			playerRepository.On("ListPlayers", givenCtx).Return([]domain.Player{}, unexpectedError)

			playersReturned, err := service.ListPlayers(givenCtx)
			assert.Empty(t, playersReturned)
			assert.Error(t, err)
		},
		"should return not found error when not found data on database": func(t *testing.T, playerRepository *mocktest.PlayerDatabaseRepositoryMock) {
			service := service.NewPlayerService(playerRepository)

			playerRepository.On("ListPlayers", givenCtx).Return([]domain.Player{}, domain.ErrNotFound)

			playersReturned, err := service.ListPlayers(givenCtx)
			assert.Empty(t, playersReturned)
			assert.ErrorIs(t, err, domain.ErrNotFound)
		},
		"should return list of players with success": func(t *testing.T, playerRepository *mocktest.PlayerDatabaseRepositoryMock) {
			service := service.NewPlayerService(playerRepository)

			playerRepository.On("ListPlayers", givenCtx).Return(givenPlayers, nil)

			playersReturned, err := service.ListPlayers(givenCtx)
			assert.ElementsMatch(t, givenPlayers, playersReturned)
			assert.NoError(t, err)
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			playerRepository := new(mocktest.PlayerDatabaseRepositoryMock)

			testCase(t, playerRepository)
		})
	}
}
