package income

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/phatpan/working-with-angular-api/logs"
	"github.com/sirupsen/logrus"
)

// Handler is wrapper for database connection using inside user method
type Handler struct {
	DB          *sql.DB
	FieldLogger logs.FieldLogger
}

// NewHandler is contructor for grouping echo route
func NewHandler(e *echo.Echo, db *sql.DB, logger logs.FieldLogger) {
	h := &Handler{
		DB:          db,
		FieldLogger: logger,
	}

	e.GET("/income/types", h.getIncomeTypeList)
}

// Logger return logger for given echo context
func (h *Handler) Logger(c echo.Context) logs.Logger {
	req := c.Request()
	return h.FieldLogger.WithFields(logrus.Fields{
		"id":       c.Response().Header().Get(echo.HeaderXRequestID),
		"method":   req.Method,
		"path_uri": req.RequestURI,
	})
}
