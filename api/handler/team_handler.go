package handler

import (
	"net/http"

	apiErrors "github.com/GabrielLoureiroGomes/basket-collection/api/handler/errors"
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
	teamService service.TeamService
}

func NewTeamHandler(service service.TeamService) TeamHandler {
	return TeamHandler{
		teamService: service,
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

	teamInserted, err := t.teamService.InsertTeam(ctx.Request.Context(), teamToInsert)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Error("error to insert team", zap.Field{Type: zapcore.StringType, String: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"Team data saved with success!": teamInserted})
}

func (t TeamHandler) GetAllTeams(ctx *gin.Context) {
	teamsToReturn, err := t.teamService.GetAllTeams(ctx.Request.Context())
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

func (t TeamHandler) GetOneTeam(ctx *gin.Context) {
	teamName := request.GetOneTeamSchema{}
	if err := ctx.ShouldBindQuery(&teamName); err != nil {
		buildErrorResponse(ctx, http.StatusBadRequest, apiErrors.ErrBindParams)
	}

	teamToReturn, err := t.teamService.GetOneTeam(ctx.Request.Context(), teamName.Name)
	if len([]*domain.Team{teamToReturn}) == 0 {
		buildErrorResponse(ctx, http.StatusNotFound, domain.ErrNotFound)
	}

	if err != nil {
		buildErrorResponse(ctx, http.StatusInternalServerError, err)
		log.Error("error to get one team", zap.Field{Type: zapcore.StringType, String: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, teamToReturn)
}

func buildErrorResponse(gin *gin.Context, statusCode int, err error) {
	errToReturn := err
	if err == nil {
		errToReturn = apiErrors.ErrUnknown
	}

	if statusCode == http.StatusInternalServerError {
		gin.String(statusCode, errToReturn.Error())
	}
}
