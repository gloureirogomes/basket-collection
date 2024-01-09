package errors

import (
	"fmt"
	"net/http"

	"github.com/GabrielLoureiroGomes/basket-collection/core/domain"
	"github.com/gin-gonic/gin"
)

var (
	// ErrUnknown is an unknown internal server error
	ErrUnknown = fmt.Errorf("internal server error %w", domain.ErrBase)

	// ErrBindParams is an error to bind parameters
	ErrBindParams = fmt.Errorf("error to bind parameters %w", domain.ErrBase)
)

func BuildErrorResponse(gin *gin.Context, statusCode int, err error) {
	errToReturn := err
	if err == nil {
		errToReturn = ErrUnknown
	}

	if statusCode == http.StatusInternalServerError {
		gin.String(statusCode, errToReturn.Error())
	}
}
