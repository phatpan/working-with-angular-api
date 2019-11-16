package outcome

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type deleteOutcomeReq struct {
	UserID int `json:"userId"`
}

func (h *Handler) deleteOutcomeByUserID(c echo.Context) error {
	uid := c.Param("userId")
	id := c.Param("id")
	oc := &deleteOutcomeReq{}
	if err := c.Bind(oc); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"code":    "O-5003",
			"message": "RequestBody invalid",
		})
	}

	return h.deleteOutcomeByIDTable(c, uid, id)
}

func (h *Handler) deleteOutcomeByIDTable(c echo.Context, uid interface{}, id interface{}) error {
	ct = time.Now()
	ct.Format(time.RFC3339)

	stmtDelete := `Delete from outcome where user_id = ? and id = ?`

	res, err := h.DB.Exec(stmtDelete, uid, id)

	if err != nil {
		h.Logger(c).Errorf("deleteOutcomeByIDTable error: %v", err)
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
		h.Logger(c).Errorf("deleteOutcomeByIDTable error: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"code":    "O-5004",
			"message": "System error, please try again",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{})
}
