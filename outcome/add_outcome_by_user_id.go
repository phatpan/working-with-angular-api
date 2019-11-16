package outcome

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

var ct time.Time

type outcomeReq struct {
	UserID         int       `json:"userId"`
	OutcomeGroupID int       `json:"outcomeGroupId"`
	Name           string    `json:"name"`
	Amount         int       `json:"amount"`
	Date           time.Time `json:"date"`
}

func (h *Handler) addOutcomeByUserID(c echo.Context) error {
	oc := &outcomeReq{}
	if err := c.Bind(oc); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"code":    "O-5003",
			"message": "RequestBody invalid",
		})
	}

	return h.insertOutcomeByUserIDTable(c, oc)
}

func (h *Handler) insertOutcomeByUserIDTable(c echo.Context, req *outcomeReq) error {
	ct = time.Now()
	ct.Format(time.RFC3339)

	stmtIns := `INSERT INTO outcome (
		user_id, outcome_group_id, name, amount, date, created_date, created_by)
		VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := h.DB.Exec(stmtIns,
		req.UserID,
		req.OutcomeGroupID,
		req.Name,
		req.Amount,
		req.Date,
		ct,
		req.UserID)

	if err != nil {
		h.Logger(c).Errorf("insertOutcomeByUserIDTable error: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"code":    "O-5004",
			"message": "System error, please try again",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{})
}
