package repository

import (
	"context"
	"github.com/GabrielLoureiroGomes/basket-collection/core/domain"
	repo "github.com/GabrielLoureiroGomes/basket-collection/pkg/repository"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"testing"
)

type TeamMongoRepositoryIntegrationTestSuite struct {
	suite.Suite
	testcontainers.Container
	repository repo.MongoRepository
}

func TestTeamMongoRepositoryIntegrationTestSuite(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	suite.Run(t, new(TeamMongoRepositoryIntegrationTestSuite))
}

func (suite *TeamMongoRepositoryIntegrationTestSuite) SetupSuite() {
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
	suite.repository = repo.NewMongoRepository(ctx)
}

func (suite *TeamMongoRepositoryIntegrationTestSuite) TearDownSuite() {
	err := suite.Terminate(context.Background())
	assert.NoError(suite.T(), err)

	viper.Reset()
}

func (suite *TeamMongoRepositoryIntegrationTestSuite) TestInsert() {
	ctx := context.Background()
	givenTeam := &domain.Team{
		Name:       "Los Angeles Lakers",
		Conference: "West",
		State:      "California",
	}

	suite.Suite.T().Run("should return error to try to insert on invalid database", func(t *testing.T) {
		defer func() {
			viper.Reset()
			suite.setupTestEnvironment()
		}()
		viper.Set("MONGO_USER", "")
		viper.Set("MONGO_PASSWORD", "")
		viper.Set("MONGO_DATABASE_NAME", "INVALID HOST")

		repository := repo.NewMongoRepository(ctx)
		err := repository.InsertTeam(ctx, givenTeam)

		assert.NotNil(suite.T(), err)
	})

	suite.Suite.T().Run("should insert team with success", func(t *testing.T) {
		err := suite.repository.InsertTeam(ctx, givenTeam)

		assert.NoError(suite.T(), err)
	})
}

func (suite *TeamMongoRepositoryIntegrationTestSuite) setupTestEnvironment() {
	viper.Set("MONGO_DATABASE_NAME", "basket-collection")
	viper.Set("MONGO_TEAM_COLLECTION", "team")
}
