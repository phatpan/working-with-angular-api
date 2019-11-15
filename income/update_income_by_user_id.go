package income

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type updateIncomeReq struct {
	UserID        int       `json:"userId"`
	IncomeGroupID int       `json:"incomeGroupId"`
	Amount        int       `json:"amount"`
	Date          time.Time `json:"date"`
}

func (h *Handler) updateIncomeByUserID(c echo.Context) error {
	id := c.Param("id")

	income := &updateIncomeReq{}
	if err := c.Bind(income); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"code":    "F-5003",
			"message": "RequestBody invalid",
		})
	}

	return h.updateIncomeByIDTable(c, income, id)
}

func (h *Handler) updateIncomeByIDTable(c echo.Context, req *updateIncomeReq, id interface{}) error {
	ct = time.Now()
	ct.Format(time.RFC3339)

	stmtIns := `UPDATE income set
		income_group_id = ?, amount = ?, date = ?, updated_date = ?, updated_by = ?
		where user_id = ? and id = ?`

	res, err := h.DB.Exec(stmtIns,
		req.IncomeGroupID,
		req.Amount,
		req.Date,
		ct,
		req.UserID,
		req.UserID,
		id,
	)

	if err != nil {
		h.Logger(c).Errorf("updateIncomeByIDTable error: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"code":    "F-5004",
			"message": "System error, please try again",
		})
	}

	count, err := res.RowsAffected()
	if count == 0 {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"code":    "F-5004",
			"message": "Invalid UserId or ID, please try again",
		})
	}

	if err != nil {
		h.Logger(c).Errorf("updateIncomeByIDTable error: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"code":    "F-5004",
			"message": "System error, please try again",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{})
}
