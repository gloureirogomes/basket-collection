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

type MongoRepository struct {
	mongo *mongoClient
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

// GetAll used to get all database team data
func (m MongoRepository) GetAll(ctx context.Context) ([]*domain.Team, error) {
	filter := bson.D{}
	cursor, err := m.getCollection().Find(ctx, filter)
	if err != nil && err != mongo.ErrNoDocuments {
		log.Error("error to get team data on mongo", zap.Field{Type: zapcore.StringType, String: err.Error()})
		return []*domain.Team{}, err
	}

	if cursor == nil || err == mongo.ErrNoDocuments {
		return []*domain.Team{}, domain.ErrNotFound
	}

	teamsMongoDocument := []TeamMongoDocument{}
	if err = cursor.All(ctx, &teamsMongoDocument); err != nil {
		log.Error("error to parse data", zap.Field{Type: zapcore.StringType, String: err.Error()})
		return []*domain.Team{}, err
	}

	teamsToReturn := newTeamListByTeamMongoDocument(teamsMongoDocument)

	return teamsToReturn, nil
}

// GetOne used to get one database team data
func (m MongoRepository) GetOne(ctx context.Context, teamName string) (*domain.Team, error) {
	filter := bson.D{{Key: "name", Value: teamName}}

	teamMongoDocument := TeamMongoDocument{}
	if err := m.getCollection().FindOne(ctx, filter).Decode(&teamMongoDocument); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return &domain.Team{}, domain.ErrNotFound
		}
		log.Error("error to get one team on mongo", zap.Field{Type: zapcore.StringType, String: err.Error()})
		return &domain.Team{}, err
	}

	teamToReturn := &domain.Team{
		Name:       teamMongoDocument.Name,
		Conference: teamMongoDocument.Conference,
		State:      teamMongoDocument.State,
	}

	return teamToReturn, nil
}

func (m MongoRepository) getCollection() *mongo.Collection {
	databaseName := viper.GetString("MONGO_DATABASE_NAME")
	teamCollection := viper.GetString("MONGO_TEAM_COLLECTION")

	return m.mongo.client.Database(databaseName).Collection(teamCollection)
}
