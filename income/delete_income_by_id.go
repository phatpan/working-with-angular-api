package income

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) deleteIncomeByUserID(c echo.Context) error {

	return c.JSON(http.StatusOK, nil)
}