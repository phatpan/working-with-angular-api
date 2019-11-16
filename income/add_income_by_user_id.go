package income

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

var ct time.Time

type incomeReq struct {
	UserID        int       `json:"userId"`
	IncomeGroupID int       `json:"incomeGroupId"`
	Amount        int       `json:"amount"`
	Date          time.Time `json:"date"`
}

func (h *Handler) addIncomeByUserID(c echo.Context) error {
	income := &incomeReq{}
	if err := c.Bind(income); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"code":    "I-5003",
			"message": "RequestBody invalid",
		})
	}

	return h.insertIncomeByUserIDTable(c, income)
}

func (h *Handler) insertIncomeByUserIDTable(c echo.Context, req *incomeReq) error {
	ct = time.Now()
	ct.Format(time.RFC3339)

	stmtIns := `INSERT INTO income (
		user_id, income_group_id, amount, date, created_date, created_by)
		VALUES (?, ?, ?, ?, ?, ?)`

	_, err := h.DB.Exec(stmtIns,
		req.UserID,
		req.IncomeGroupID,
		req.Amount,
		req.Date,
		ct,
		req.UserID)

	if err != nil {
		h.Logger(c).Errorf("insertIncomeByUserIDTable error: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"code":    "I-5004",
			"message": "System error, please try again",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{})
}
