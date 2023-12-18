package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetHealth
//
//	@Tags 			basic
//	@Summary		health
//	@Description	get health
//	@Produce		json
//	@Success		200
//	@Router			/health [get]
func GetHealth(context echo.Context) error {
	return context.JSON(http.StatusOK, nil)
}
