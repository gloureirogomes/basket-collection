package repository

import (
	"context"
	"testing"

	"github.com/GabrielLoureiroGomes/basket-collection/core/domain"
	repo "github.com/GabrielLoureiroGomes/basket-collection/pkg/repository"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type PlayerMongoRepositoryIntegrationTestSuite struct {
	suite.Suite
	testcontainers.Container
	repository repo.PlayerMongoRepository
}

func TestPlayerMongoRepositoryIntegrationTestSuite(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	suite.Run(t, new(PlayerMongoRepositoryIntegrationTestSuite))
}

func (suite *PlayerMongoRepositoryIntegrationTestSuite) SetupSuite() {
	suite.setupTestEnvironment()

	ctx := context.Background()

	containerReq := testcontainers.ContainerRequest{
		Image:        "mongo",
		ExposedPorts: []string{"27017/tcp"},
		WaitingFor:   wait.ForExposedPort(),
	}

	genericRequest := testcontainers.GenericContainerRequest{
		ContainerRequest: containerReq,
		Started:          true,
	}

	container, err := testcontainers.GenericContainer(ctx, genericRequest)
	assert.NoError(suite.T(), err)

	host, err := container.Host(ctx)
	assert.NoError(suite.T(), err)
	viper.Set("MONGO_HOST", host)

	port, err := container.MappedPort(ctx, "27017")
	assert.NoError(suite.T(), err)
	viper.Set("MONGO_PORT", string(port))

	suite.Container = container
	suite.repository = repo.NewPlayerMongoRepository(ctx)
}

func (suite *PlayerMongoRepositoryIntegrationTestSuite) TearDownTest() {
	err := suite.repository.Mongo.Client.Database(viper.GetString("MONGO_DATABASE_NAME")).Drop(context.Background())
	assert.NoError(suite.T(), err)

	viper.Reset()
}

func (suite *PlayerMongoRepositoryIntegrationTestSuite) TearDownSuite() {
	err := suite.Terminate(context.Background())
	assert.NoError(suite.T(), err)

	viper.Reset()
}

func (suite *PlayerMongoRepositoryIntegrationTestSuite) TestInsertPlayer() {
	givenCtx := context.Background()

	givenTeam := domain.Team{
		Name:       "Los Angeles Lakers",
		Conference: "West",
	}

	givenPlayer := domain.Player{
		Name:     "LeBron James",
		Age:      38,
		Position: "PF",
		Number:   23,
		Team:     givenTeam,
	}

	suite.Suite.T().Run("should return error to try to insert on invalid database", func(t *testing.T) {
		defer func() {
			viper.Reset()
			suite.setupTestEnvironment()
		}()
		viper.Set("MONGO_USER", "")
		viper.Set("MONGO_PASSWORD", "")
		viper.Set("MONGO_DATABASE_NAME", "INVALID HOST")

		repository := repo.NewPlayerMongoRepository(givenCtx)
		err := repository.InsertPlayer(givenCtx, givenPlayer)

		assert.NotNil(suite.T(), err)
	})

	suite.Suite.T().Run("should insert player with success", func(t *testing.T) {
		err := suite.repository.InsertPlayer(givenCtx, givenPlayer)

		assert.NoError(suite.T(), err)
	})
}

func (suite *PlayerMongoRepositoryIntegrationTestSuite) TestListPlayers() {
	givenCtx := context.Background()
	givenPlayers := []domain.Player{
		{
			Name:     "LeBron James",
			Age:      39,
			Position: "SF",
			Number:   23,
			Team: domain.Team{
				Name:       "Los Angeles Lakers",
				Conference: "West",
			},
		},
		{
			Name:     "Stephen Curry",
			Age:      33,
			Position: "PG",
			Number:   30,
			Team: domain.Team{
				Name:       "Golden State Warriors",
				Conference: "West",
			},
		},
	}

	suite.Suite.T().Run("should return error to try to list players on invalid database", func(t *testing.T) {
		defer func() {
			viper.Reset()
			suite.setupTestEnvironment()
		}()
		viper.Set("MONGO_USER", "")
		viper.Set("MONGO_PASSWORD", "")
		viper.Set("MONGO_DATABASE_NAME", "INVALID HOST")

		repository := repo.NewPlayerMongoRepository(givenCtx)
		_, err := repository.ListPlayers(givenCtx)

		assert.NotNil(suite.T(), err)
	})

	suite.Suite.T().Run("should not return error when not found players on database", func(t *testing.T) {
		emptyReturn, err := suite.repository.ListPlayers(givenCtx)

		assert.Empty(suite.T(), emptyReturn)
		assert.NoError(suite.T(), err)
	})

	suite.Suite.T().Run("should list players with success", func(t *testing.T) {
		suite.insertPlayersToTest(t, givenPlayers)

		playersReturned, err := suite.repository.ListPlayers(givenCtx)

		assert.ElementsMatch(suite.T(), givenPlayers, playersReturned)
		assert.NoError(suite.T(), err)
	})
}

func (suite *PlayerMongoRepositoryIntegrationTestSuite) setupTestEnvironment() {
	viper.Set("MONGO_DATABASE_NAME", "basket-collection")
	viper.Set("MONGO_PLAYER_COLLECTION", "player")
}

func (suite *PlayerMongoRepositoryIntegrationTestSuite) insertPlayersToTest(t *testing.T, players []domain.Player) {
	for _, player := range players {
		err := suite.repository.InsertPlayer(context.Background(), player)
		assert.NoError(t, err)
	}
}
