package handler

import (
	"net/http"

	"github.com/GabrielLoureiroGomes/basket-collection/api/schema/request"
	"github.com/GabrielLoureiroGomes/basket-collection/core/domain"
	"github.com/GabrielLoureiroGomes/basket-collection/logger"
	"github.com/GabrielLoureiroGomes/basket-collection/pkg/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log = logger.GetLogger()

type TeamHandler struct {
	TeamService service.TeamService
}

func NewTeamHandler(service service.TeamService) TeamHandler {
	return TeamHandler{
		TeamService: service,
	}
}

func (t TeamHandler) CreateTeam(ctx *gin.Context) {
	teamSchemaToInsert := request.InsertTeamSchema{}
	if err := ctx.ShouldBindJSON(&teamSchemaToInsert); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Error("error to bind json", zap.Field{Type: zapcore.StringType, String: err.Error()})
		return
	}

	teamToInsert := &domain.Team{
		Name:       teamSchemaToInsert.Name,
		Conference: teamSchemaToInsert.Conference,
		State:      teamSchemaToInsert.State,
	}

	teamInserted, err := t.TeamService.InsertTeam(ctx.Request.Context(), teamToInsert)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Error("error to insert team", zap.Field{Type: zapcore.StringType, String: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"Team data saved with success!": teamInserted})
}

func (t TeamHandler) GetAllTeams(ctx *gin.Context) {
	teamsToReturn, err := t.TeamService.GetAllTeams(ctx.Request.Context())
	if len(teamsToReturn) == 0 {
		ctx.JSON(http.StatusNotFound, nil)
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Error("error to get all teams", zap.Field{Type: zapcore.StringType, String: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"teams": teamsToReturn})
}
