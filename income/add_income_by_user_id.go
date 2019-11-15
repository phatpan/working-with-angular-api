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
	food := &incomeReq{}
	if err := c.Bind(food); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"code":    "F-5003",
			"message": "RequestBody invalid",
		})
	}

	return h.insertManageFoodByUserIDTable(c, food)
}

func (h *Handler) insertManageFoodByUserIDTable(c echo.Context, req *incomeReq) error {
	ct = time.Now()
	ct.Format(time.RFC3339)

	stmtIns := `INSERT INTO income (
		user_id, income_group_id, amount, date, created_date)
		VALUES (?, ?, ?, ?, ?)`

	_, err := h.DB.Exec(stmtIns,
		req.UserID,
		req.IncomeGroupID,
		req.Amount,
		req.Date,
		ct)

	if err != nil {
		h.Logger(c).Errorf("insertManageFoodByUserIdTable error: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"code":    "F-5004",
			"message": "System error, please try again",
		})
	}

	return c.JSON(http.StatusNoContent, nil)
}
