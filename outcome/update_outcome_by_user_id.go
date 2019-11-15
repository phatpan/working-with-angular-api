package outcome

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (h *Handler) updateOutcomeByUserID(c echo.Context) error {
	uid := c.Param("id")

	food := &foodSelectedIDReq{}
	if err := c.Bind(food); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"code":    "F-5003",
			"message": "RequestBody invalid",
		})
	}

	return h.updateIncomeByIDTable(c, food, uid)
}

func (h *Handler) updateIncomeByIDTable(c echo.Context, food *foodSelectedIDReq, uid interface{}) error {
	ct = time.Now()
	ct.Format(time.RFC3339)

	stmtDeletePayment := `DELETE FROM m_food WHERE user_id = ?;`
	_, err := h.DB.Exec(stmtDeletePayment, uid)
	if err != nil {
		h.Logger(c).Errorf("deleteFoodByUserID error: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"code":    "F-5004",
			"message": err,
		})
	}

	for _, fid := range food.ID {

		stmtIns := `INSERT INTO m_food (
		food_id, user_id, created_date)
		VALUES (?, ?, ?)`

		_, err := h.DB.Exec(stmtIns,
			fid,
			uid,
			ct)

		if err != nil {
			h.Logger(c).Errorf("insertManageFoodByUserIdTable error: %v", err)
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"code":    "F-5004",
				"message": "System error, please try again",
			})
		}
	}

	return c.JSON(http.StatusNoContent, nil)
}
