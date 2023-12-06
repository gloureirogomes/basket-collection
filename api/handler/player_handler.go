package handler

import (
	"fmt"
	"net/http"

	apiErrors "github.com/GabrielLoureiroGomes/basket-collection/api/handler/errors"
	"github.com/GabrielLoureiroGomes/basket-collection/api/schema/request"
	"github.com/GabrielLoureiroGomes/basket-collection/core/domain"
	"github.com/GabrielLoureiroGomes/basket-collection/pkg/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type PlayerHandler struct {
	playerService service.PlayerService
}

func NewPlayerHandler(service service.PlayerService) PlayerHandler {
	return PlayerHandler{
		playerService: service,
	}
}

func (p PlayerHandler) CreatePlayer(ctx *gin.Context) {
	playerSchemaToInsert := request.InsertPlayerSchema{}
	if err := ctx.ShouldBindJSON(&playerSchemaToInsert); err != nil {
		buildErrorResponse(ctx, http.StatusBadRequest, apiErrors.ErrBindParams)
		log.Error("error to bind json", zap.Field{Type: zapcore.StringType, String: err.Error()})
		return
	}

	playerToInsert := domain.Player{
		Name:     playerSchemaToInsert.Name,
		Age:      playerSchemaToInsert.Age,
		Position: playerSchemaToInsert.Position,
		Country:  playerSchemaToInsert.Country,
		Team:     playerSchemaToInsert.Team,
		Height:   playerSchemaToInsert.Height,
		Weight:   playerSchemaToInsert.Weight,
	}

	playerInserted, err := p.playerService.InsertPlayer(ctx.Request.Context(), playerToInsert)
	if err != nil {
		buildErrorResponse(ctx, http.StatusInternalServerError, err)
		log.Error("error to insert team", zap.Field{Type: zapcore.StringType, String: err.Error()})
		return
	}

	ctx.String(http.StatusCreated, fmt.Sprintf("The %s was inserted with success", playerInserted.GetName()))
}
