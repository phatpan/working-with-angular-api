package income

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type deleteIncomeReq struct {
	UserID int `json:"userId"`
}

func (h *Handler) deleteIncomeByUserID(c echo.Context) error {
	uid := c.Param("userId")
	id := c.Param("id")
	income := &updateIncomeReq{}
	if err := c.Bind(income); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"code":    "F-5003",
			"message": "RequestBody invalid",
		})
	}

	return h.deleteIncomeByIDTable(c, uid, id)
}

func (h *Handler) deleteIncomeByIDTable(c echo.Context, uid interface{}, id interface{}) error {
	ct = time.Now()
	ct.Format(time.RFC3339)

	stmtDelete := `Delete from income where user_id = ? and id = ?`

	res, err := h.DB.Exec(stmtDelete, uid, id)

	if err != nil {
		h.Logger(c).Errorf("deleteIncomeByIDTable error: %v", err)
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
		h.Logger(c).Errorf("deleteIncomeByIDTable error: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"code":    "F-5004",
			"message": "System error, please try again",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{})
}
