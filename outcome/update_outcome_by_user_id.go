package outcome

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type updateOutcomeReq struct {
	UserID         int       `json:"userId"`
	OutcomeGroupID int       `json:"outcomeGroupId"`
	Amount         int       `json:"amount"`
	Name           string    `json:"name"`
	Date           time.Time `json:"date"`
}

func (h *Handler) updateOutcomeByUserID(c echo.Context) error {
	id := c.Param("id")

	oc := &updateOutcomeReq{}
	if err := c.Bind(oc); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"code":    "O-5003",
			"message": "RequestBody invalid",
		})
	}

	return h.updateOutcomeByIDTable(c, oc, id)
}

func (h *Handler) updateOutcomeByIDTable(c echo.Context, req *updateOutcomeReq, id interface{}) error {
	ct = time.Now()
	ct.Format(time.RFC3339)

	stmtUpdate := `UPDATE outcome set
		outcome_group_id = ?, name = ?, amount = ?, date = ?, updated_date = ?, updated_by = ?
		where user_id = ? and id = ?`

	res, err := h.DB.Exec(stmtUpdate,
		req.OutcomeGroupID,
		req.Name,
		req.Amount,
		req.Date,
		ct,
		req.UserID,
		req.UserID,
		id,
	)

	if err != nil {
		h.Logger(c).Errorf("updateOutcomeByIDTable error: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"code":    "O-5004",
			"message": "System error, please try again",
		})
	}

	count, err := res.RowsAffected()
	if count == 0 {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"code":    "O-5004",
			"message": "Invalid UserId or ID, please try again",
		})
	}

	if err != nil {
		h.Logger(c).Errorf("updateOutcomeByIDTable error: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"code":    "O-5004",
			"message": "System error, please try again",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{})
}
