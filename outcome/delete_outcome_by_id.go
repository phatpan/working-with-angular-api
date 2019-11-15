package outcome

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) deleteOutcomeByUserID(c echo.Context) error {

	return c.JSON(http.StatusOK, nil)
}
