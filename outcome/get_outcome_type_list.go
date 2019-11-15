package outcome

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

type typeResp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) getOutcomeTypeList(c echo.Context) error {
	stmt := "select id, name from outcome_type"
	t := []typeResp{}

	rows, err := h.DB.Query(stmt)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, nil)
		}
		h.Logger(c).Errorf("Get outcome type error: %v", err)
		c.JSON(http.StatusInternalServerError, err)
	}
	defer rows.Close()

	for rows.Next() {
		var res typeResp
		err := rows.Scan(&res.ID, &res.Name)
		if err != nil {
			h.Logger(c).Errorf("Get outcome type error: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"code":    "O-5003",
				"message": "System error, please try again",
			})
		}
		t = append(t, res)
	}
	err = rows.Err()
	if err != nil {
		h.Logger(c).Errorf("Cannot Get outcome type error: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "O-5004",
			"message": "System error, please try again",
		})
	}

	return c.JSON(http.StatusOK, t)
}
