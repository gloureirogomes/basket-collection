package repository

import (
	"context"

	"github.com/GabrielLoureiroGomes/basket-collection/core/domain"
	"github.com/spf13/viper"
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

func (m PlayerMongoRepository) getCollection() *mongo.Collection {
	databaseName := viper.GetString("MONGO_DATABASE_NAME")
	playerCollection := viper.GetString("MONGO_PLAYER_COLLECTION")

	return m.Mongo.Client.Database(databaseName).Collection(playerCollection)
}
