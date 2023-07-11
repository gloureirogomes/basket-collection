package repository

import (
	"context"

	"github.com/GabrielLoureiroGomes/basket-collection/core/domain"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type MongoRepository struct {
	mongo *mongoClient
}

type TeamMongoDocument struct {
	Name       string `bson:"name"`
	Conference string `bson:"conference"`
	State      string `bson:"state"`
}

func newTeamMongoDocument(team *domain.Team) TeamMongoDocument {
	return TeamMongoDocument{
		Name:       team.GetName(),
		Conference: team.GetConference(),
		State:      team.GetState(),
	}
}

func NewMongoRepository(ctx context.Context) MongoRepository {
	return MongoRepository{
		mongo: newMongoClient(ctx),
	}
}

// InsertTeam used to save on database team data
func (m MongoRepository) InsertTeam(ctx context.Context, team *domain.Team) error {

	teamDocumentToInsert := newTeamMongoDocument(team)
	_, err := m.getCollection().InsertOne(ctx, teamDocumentToInsert)

	if err != nil {
		log.Error("error to insert team data on mongo", zap.Field{Type: zapcore.StringType, String: err.Error()})
		return err
	}

	return nil
}

func (m MongoRepository) getCollection() *mongo.Collection {
	databaseName := viper.GetString("MONGO_DATABASE_NAME")
	teamCollection := viper.GetString("MONGO_TEAM_COLLECTION")

	return m.mongo.client.Database(databaseName).Collection(teamCollection)
}
