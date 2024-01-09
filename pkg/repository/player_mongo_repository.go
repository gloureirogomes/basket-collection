package repository

import (
	"context"
	"errors"

	"github.com/GabrielLoureiroGomes/basket-collection/core/domain"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type PlayerMongoRepository struct {
	Mongo *mongoClient
}

func NewPlayerMongoRepository(ctx context.Context) PlayerMongoRepository {
	return PlayerMongoRepository{
		Mongo: newMongoClient(ctx),
	}
}

// InsertPlayer used to save on database player data
func (m PlayerMongoRepository) InsertPlayer(ctx context.Context, player domain.Player) error {

	playerDocumentToInsert := newPlayerMongoDocument(player)
	_, err := m.getCollection().InsertOne(ctx, playerDocumentToInsert)

	if err != nil {
		log.Error("error to insert player data on mongo", zap.Field{Type: zapcore.StringType, String: err.Error()})
		return err
	}

	return nil
}

func (m PlayerMongoRepository) ListPlayers(ctx context.Context) ([]domain.Player, error) {
	filter := bson.D{}

	cursor, err := m.getCollection().Find(ctx, filter)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		log.Error("error to get player data on mongo", zap.Field{Type: zapcore.StringType, String: err.Error()})
		return []domain.Player{}, err
	}

	if cursor == nil || errors.Is(err, mongo.ErrNoDocuments) {
		return []domain.Player{}, domain.ErrNotFound
	}

	playersMongoDocument := []PlayerMongoDocument{}
	if err = cursor.All(ctx, &playersMongoDocument); err != nil {
		log.Error("error to parse data", zap.Field{Type: zapcore.StringType, String: err.Error()})
		return []domain.Player{}, err
	}

	playersToReturn := newPlayerListByPlayerMongoDocument(playersMongoDocument)

	return playersToReturn, nil
}

func (m PlayerMongoRepository) getCollection() *mongo.Collection {
	databaseName := viper.GetString("MONGO_DATABASE_NAME")
	playerCollection := viper.GetString("MONGO_PLAYER_COLLECTION")

	return m.Mongo.Client.Database(databaseName).Collection(playerCollection)
}
