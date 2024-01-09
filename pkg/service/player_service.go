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

type PlayerService struct {
	databaseRepository repository.PlayerDatabaseRepository
}

func NewPlayerService(databaseRepository repository.PlayerDatabaseRepository) PlayerService {
	return PlayerService{
		databaseRepository: databaseRepository,
	}
}

// InsertPlayer function is used to create a player register
func (t PlayerService) InsertPlayer(ctx context.Context, player domain.Player) (domain.Player, error) {
	if err := t.databaseRepository.InsertPlayer(ctx, player); err != nil {
		log.Error("error to insert player", zap.Field{Type: zapcore.StringType, String: err.Error()})
		return domain.Player{}, err
	}

	return player, nil
}
