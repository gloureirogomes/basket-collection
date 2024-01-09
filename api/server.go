package api

import (
	"context"
	"fmt"

	"github.com/GabrielLoureiroGomes/basket-collection/api/handler"
	"github.com/GabrielLoureiroGomes/basket-collection/logger"
	"github.com/GabrielLoureiroGomes/basket-collection/pkg/repository"
	"github.com/GabrielLoureiroGomes/basket-collection/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log = logger.GetLogger()

type Server struct {
	playerHandler handler.PlayerHandler
}

func NewServer() Server {
	playerDatabaseRepository := repository.NewPlayerMongoRepository(context.Background())
	playerService := service.NewPlayerService(playerDatabaseRepository)
	playerHandler := handler.NewPlayerHandler(playerService)

	return Server{playerHandler: playerHandler}
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

	routerGroup.POST("/player", s.playerHandler.CreatePlayer)

	return router
}
