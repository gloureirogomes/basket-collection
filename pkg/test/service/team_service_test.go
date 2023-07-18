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

func TestInsertTeam(t *testing.T) {
	givenCtx := context.Background()
	givenTeam := &domain.Team{
		Name:       "Los Angeles Lakers",
		Conference: "West",
		State:      "California",
	}

	testCases := map[string]func(*testing.T, *mocktest.TeamDatabaseRepositoryMock){
		"should save team data with success": func(t *testing.T, teamRepository *mocktest.TeamDatabaseRepositoryMock) {
			service := service.NewTeamService(teamRepository)

			teamRepository.On("InsertTeam", givenCtx, givenTeam).Return(nil)

			teamInserted, err := service.InsertTeam(givenCtx, givenTeam)
			assert.NotEmpty(t, teamInserted)
			assert.NoError(t, err)
		},
		"should return error when try to save empty team data": func(t *testing.T, teamRepository *mocktest.TeamDatabaseRepositoryMock) {
			service := service.NewTeamService(teamRepository)
			err := errors.New("error to insert empty data")
			givenEmptyTeam := &domain.Team{}

			teamRepository.On("InsertTeam", givenCtx, givenEmptyTeam).Return(err)

			teamInserted, err := service.InsertTeam(givenCtx, givenEmptyTeam)
			assert.Empty(t, teamInserted)
			assert.Error(t, err)
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			teamRepository := new(mocktest.TeamDatabaseRepositoryMock)

			testCase(t, teamRepository)
		})
	}
}

func TestGetAllTeams(t *testing.T) {
	givenCtx := context.Background()
	teamsToReturn := []*domain.Team{
		{
			Name:       "Los Angeles Lakers",
			Conference: "West",
			State:      "California",
		},
		{
			Name:       "Golden State Warriors",
			Conference: "West",
			State:      "California",
		},
	}

	testCases := map[string]func(*testing.T, *mocktest.TeamDatabaseRepositoryMock){
		"should return error when an error occurs on database": func(t *testing.T, teamRepository *mocktest.TeamDatabaseRepositoryMock) {
			unexpectedError := errors.New("unexpected error")
			service := service.NewTeamService(teamRepository)

			teamRepository.On("GetAll", givenCtx).Return([]*domain.Team{}, unexpectedError)

			teamsReturned, err := service.GetAllTeams(givenCtx)
			assert.Empty(t, teamsReturned)
			assert.Error(t, err)
		},
		"should return not found error when not found data on database": func(t *testing.T, teamRepository *mocktest.TeamDatabaseRepositoryMock) {
			service := service.NewTeamService(teamRepository)

			teamRepository.On("GetAll", givenCtx).Return([]*domain.Team{}, domain.ErrNotFound)

			teamsReturned, err := service.GetAllTeams(givenCtx)
			assert.Empty(t, teamsReturned)
			assert.ErrorIs(t, err, domain.ErrNotFound)
		},
		"should return team data with success": func(t *testing.T, teamRepository *mocktest.TeamDatabaseRepositoryMock) {
			service := service.NewTeamService(teamRepository)

			teamRepository.On("GetAll", givenCtx).Return(teamsToReturn, nil)

			teamsReturned, err := service.GetAllTeams(givenCtx)
			assert.Equal(t, 2, len(teamsReturned))
			assert.NotEmpty(t, teamsReturned)
			assert.NoError(t, err)
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			teamRepository := new(mocktest.TeamDatabaseRepositoryMock)

			testCase(t, teamRepository)
		})
	}
}

func TestGetOneTeam(t *testing.T) {
	givenCtx := context.Background()
	teamToReturn := &domain.Team{
		Name:       "Golden State Warriors",
		Conference: "West",
		State:      "California",
	}

	testCases := map[string]func(*testing.T, *mocktest.TeamDatabaseRepositoryMock){
		"should return error when error occurs on database": func(t *testing.T, teamRepository *mocktest.TeamDatabaseRepositoryMock) {
			unexpectedError := errors.New("unexpected error")
			notSavedTeamName := "Portland Trail Blazers"
			service := service.NewTeamService(teamRepository)

			teamRepository.On("GetOne", givenCtx, notSavedTeamName).Return(&domain.Team{}, unexpectedError)

			teamReturned, err := service.GetOneTeam(givenCtx, notSavedTeamName)
			assert.Empty(t, teamReturned)
			assert.Error(t, err)
		},
		"should return not found error when data not found on database": func(t *testing.T, teamRepository *mocktest.TeamDatabaseRepositoryMock) {
			notSavedTeamName := "Portland Trail Blazers"
			service := service.NewTeamService(teamRepository)

			teamRepository.On("GetOne", givenCtx, notSavedTeamName).Return(&domain.Team{}, domain.ErrNotFound)

			teamReturned, err := service.GetOneTeam(givenCtx, notSavedTeamName)
			assert.Empty(t, teamReturned)
			assert.ErrorIs(t, err, domain.ErrNotFound)
		},
		"should return team data with success": func(t *testing.T, teamRepository *mocktest.TeamDatabaseRepositoryMock) {
			service := service.NewTeamService(teamRepository)

			teamRepository.On("GetOne", givenCtx, teamToReturn.GetName()).Return(teamToReturn, nil)

			teamReturned, err := service.GetOneTeam(givenCtx, teamToReturn.GetName())
			assert.NotEmpty(t, teamReturned)
			assert.NoError(t, err)
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			teamRepository := new(mocktest.TeamDatabaseRepositoryMock)

			testCase(t, teamRepository)
		})
	}
}
