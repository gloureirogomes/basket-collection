package api

import (
	"context"
	"fmt"

	"github.com/GabrielLoureiroGomes/basket-collection/api/handler"
	"github.com/GabrielLoureiroGomes/basket-collection/logger"
	"github.com/GabrielLoureiroGomes/basket-collection/pkg/repository"
	"github.com/GabrielLoureiroGomes/basket-collection/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log = logger.GetLogger()

type Server struct {
	teamHandler handler.TeamHandler
}

func NewServer() Server {
	teamDatabaseRepository := repository.NewMongoRepository(context.Background())
	teamService := service.NewTeamService(teamDatabaseRepository)
	teamHandler := handler.NewTeamHandler(teamService)

	return Server{teamHandler: teamHandler}
}

func (s Server) StartServer() {
	router := s.setupRoutes()

	if err := router.Run(fmt.Sprintf(":%s", viper.GetString("API_PORT"))); err != nil {
		log.Fatal("error to start server", zap.Field{Type: zapcore.StringType, String: err.Error()})
	}
}

func (s Server) setupRoutes() *gin.Engine {
	router := gin.Default()
	routerGroup := router.Group("/basket-collection")

	routerGroup.POST("/team", s.teamHandler.CreateTeam)
	routerGroup.GET("/teams", s.teamHandler.GetAllTeams)
	routerGroup.GET("/team", s.teamHandler.GetOneTeam)
	routerGroup.DELETE("/team", s.teamHandler.DeleteTeam)

	return router
}
