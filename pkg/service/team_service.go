package service

import (
	"context"

	"github.com/GabrielLoureiroGomes/basket-collection/core/domain"
	"github.com/GabrielLoureiroGomes/basket-collection/core/repository"
	"github.com/GabrielLoureiroGomes/basket-collection/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log = logger.GetLogger()

type TeamService struct {
	databaseRepository repository.TeamDatabaseRepository
}

func NewTeamService(databaseRepository repository.TeamDatabaseRepository) TeamService {
	return TeamService{
		databaseRepository: databaseRepository,
	}
}

// InsertTeam function is used to create a team register
func (t TeamService) InsertTeam(ctx context.Context, team *domain.Team) (*domain.Team, error) {
	if err := t.databaseRepository.InsertTeam(ctx, team); err != nil {
		log.Error("error to insert team", zap.Field{Type: zapcore.StringType, String: err.Error()})
		return &domain.Team{}, err
	}

	return team, nil
}

// GetAllTeams used to get all team data
func (t TeamService) GetAllTeams(ctx context.Context) ([]*domain.Team, error) {
	teamsToReturn, err := t.databaseRepository.GetAll(ctx)
	if err != nil {
		log.Error("error to get all teams", zap.Field{Type: zapcore.StringType, String: err.Error()})
		return []*domain.Team{}, err
	}

	return teamsToReturn, nil
}

// GetOneTeam used to get one team data
func (t TeamService) GetOneTeam(ctx context.Context, teamName string) (*domain.Team, error) {
	teamToReturn, err := t.databaseRepository.GetOne(ctx, teamName)
	if err != nil {
		log.Error("error to get one team", zap.Field{Type: zapcore.StringType, String: err.Error()})
		return &domain.Team{}, err
	}

	return teamToReturn, nil
}
