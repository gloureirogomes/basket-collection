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

type TeamMongoRepositoryIntegrationTestSuite struct {
	suite.Suite
	testcontainers.Container
	repository repo.TeamMongoRepository
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
	suite.repository = repo.NewTeamMongoRepository(ctx)
}

func (suite *TeamMongoRepositoryIntegrationTestSuite) TearDownTest() {
	err := suite.repository.Mongo.Client.Database(viper.GetString("MONGO_DATABASE_NAME")).Drop(context.Background())
	assert.NoError(suite.T(), err)

	viper.Reset()
}

func (suite *TeamMongoRepositoryIntegrationTestSuite) TearDownSuite() {
	err := suite.Terminate(context.Background())
	assert.NoError(suite.T(), err)

	viper.Reset()
}

func (suite *TeamMongoRepositoryIntegrationTestSuite) TestInsertTeam() {
	givenCtx := context.Background()
	givenTeam := domain.Team{
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

		repository := repo.NewTeamMongoRepository(givenCtx)
		err := repository.InsertTeam(givenCtx, givenTeam)

		assert.NotNil(suite.T(), err)
	})

	suite.Suite.T().Run("should insert team with success", func(t *testing.T) {
		err := suite.repository.InsertTeam(givenCtx, givenTeam)

		assert.NoError(suite.T(), err)
	})
}

func (suite *TeamMongoRepositoryIntegrationTestSuite) TestGetAll() {
	givenCtx := context.Background()
	givenTeams := []domain.Team{
		{
			Name:       "Golden State Warriors",
			Conference: "West",
			State:      "California",
		},
		{
			Name:       "Los Angeles Lakers",
			Conference: "West",
			State:      "California",
		},
	}

	suite.Suite.T().Run("should return error to try to get on invalid database", func(t *testing.T) {
		defer func() {
			viper.Reset()
			suite.setupTestEnvironment()
		}()
		viper.Set("MONGO_USER", "")
		viper.Set("MONGO_PASSWORD", "")
		viper.Set("MONGO_DATABASE_NAME", "INVALID HOST")

		repository := repo.NewTeamMongoRepository(givenCtx)
		teams, err := repository.GetAll(givenCtx)

		assert.NotNil(suite.T(), err)
		assert.Empty(suite.T(), teams)
	})

	suite.Suite.T().Run("should not return error when not found team on database", func(t *testing.T) {
		emptyReturn, err := suite.repository.GetAll(givenCtx)

		assert.Empty(suite.T(), emptyReturn)
		assert.NoError(suite.T(), err)
	})

	suite.Suite.T().Run("should return team with success", func(t *testing.T) {
		suite.insertTeamsToTest(t, givenTeams)

		teamsReturned, err := suite.repository.GetAll(givenCtx)

		assert.ElementsMatch(suite.T(), givenTeams, teamsReturned)
		assert.NoError(suite.T(), err)
	})
}

func (suite *TeamMongoRepositoryIntegrationTestSuite) TestGetOne() {
	givenCtx := context.Background()
	givenTeam := domain.Team{
		Name:       "Los Angeles Lakers",
		Conference: "West",
		State:      "California",
	}

	suite.Suite.T().Run("should return error to try to get on invalid database", func(t *testing.T) {
		defer func() {
			viper.Reset()
			suite.setupTestEnvironment()
		}()
		viper.Set("MONGO_USER", "")
		viper.Set("MONGO_PASSWORD", "")
		viper.Set("MONGO_DATABASE_NAME", "INVALID HOST")

		repository := repo.NewTeamMongoRepository(givenCtx)
		team, err := repository.GetOne(givenCtx, givenTeam.Name)

		assert.NotNil(suite.T(), err)
		assert.Empty(suite.T(), team)
	})

	suite.Suite.T().Run("should return domain not found error when doesn't have data on database", func(t *testing.T) {
		emptyReturn, err := suite.repository.GetOne(givenCtx, givenTeam.Name)

		assert.Empty(suite.T(), emptyReturn)
		assert.ErrorIs(suite.T(), err, domain.ErrNotFound)
	})

	suite.Suite.T().Run("should return team with success", func(t *testing.T) {
		suite.insertTeamsToTest(t, []domain.Team{givenTeam})

		teamReturned, err := suite.repository.GetOne(givenCtx, givenTeam.Name)

		assert.Equal(suite.T(), givenTeam, teamReturned)
		assert.NoError(suite.T(), err)
	})
}

func (suite *TeamMongoRepositoryIntegrationTestSuite) TestDelete() {
	givenCtx := context.Background()
	givenTeam := domain.Team{
		Name:       "Los Angeles Lakers",
		Conference: "West",
		State:      "California",
	}

	suite.Suite.T().Run("should return error to try to delete on invalid database", func(t *testing.T) {
		defer func() {
			viper.Reset()
			suite.setupTestEnvironment()
		}()
		viper.Set("MONGO_USER", "")
		viper.Set("MONGO_PASSWORD", "")
		viper.Set("MONGO_DATABASE_NAME", "INVALID HOST")

		repository := repo.NewTeamMongoRepository(givenCtx)
		err := repository.Delete(givenCtx, givenTeam.Name)

		assert.NotNil(suite.T(), err)
	})

	suite.Suite.T().Run("should return domain not found error when doesn't have data on database", func(t *testing.T) {
		err := suite.repository.Delete(givenCtx, givenTeam.Name)

		assert.ErrorIs(suite.T(), err, domain.ErrNotFound)
	})

	suite.Suite.T().Run("should delete team with success", func(t *testing.T) {
		suite.insertTeamsToTest(t, []domain.Team{givenTeam})

		err := suite.repository.Delete(givenCtx, givenTeam.Name)

		assert.NoError(suite.T(), err)

		teams, err := suite.repository.GetAll(givenCtx)
		assert.Empty(suite.T(), teams)
		assert.NoError(suite.T(), err)
	})
}

func (suite *TeamMongoRepositoryIntegrationTestSuite) setupTestEnvironment() {
	viper.Set("MONGO_DATABASE_NAME", "basket-collection")
	viper.Set("MONGO_TEAM_COLLECTION", "team")
}

func (suite *TeamMongoRepositoryIntegrationTestSuite) insertTeamsToTest(t *testing.T, teams []domain.Team) {
	for _, team := range teams {
		err := suite.repository.InsertTeam(context.Background(), team)
		assert.NoError(t, err)
	}
}
