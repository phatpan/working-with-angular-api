package outcome

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

	e.GET("/outcome/group", h.getOutcomeGroupList)

	e.GET("/outcome/id/:id", h.getOutcomeListByUserID)
	e.POST("/outcome", h.addOutcomeByUserID)
	e.PUT("/outcome/id/:id", h.updateOutcomeByUserID)
	e.DELETE("/outcome/user-id/:userId/id/:id", h.deleteOutcomeByUserID)

	e.GET("/outcome/user-id/:id/search/:search", h.searchOutcomeByUserIDWithText)
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
