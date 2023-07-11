package repository

import (
	"context"
	"fmt"

	"github.com/GabrielLoureiroGomes/basket-collection/logger"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var log = logger.GetLogger()

// mongoClient is a database struct with client informations
type mongoClient struct {
	client *mongo.Client
}

// newMongoClient is used to connect on mongoDB
func newMongoClient(ctx context.Context) *mongoClient {
	user := viper.GetString("MONGO_USER")
	password := viper.GetString("MONGO_PASSWORD")
	databaseName := viper.GetString("MONGO_DATABASE_NAME")
	stringConnection := fmt.Sprintf("mongodb+srv://%s:%s@%s.rzlfysv.mongodb.net/?retryWrites=true&w=majority", user, password, databaseName)

	opts := options.Client().ApplyURI(stringConnection)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Error("error to connect on mongo", zap.Field{Type: zapcore.StringType, String: err.Error()})
	}

	return &mongoClient{
		client: client,
	}
}
