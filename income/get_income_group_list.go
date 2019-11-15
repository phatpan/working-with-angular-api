package income

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

type groupResp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) getIncomeGroupList(c echo.Context) error {
	stmt := "select id, name from income_group"
	g := []groupResp{}

	rows, err := h.DB.Query(stmt)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, nil)
		}
		h.Logger(c).Errorf("Get income group error: %v", err)
		c.JSON(http.StatusInternalServerError, err)
	}
	defer rows.Close()

	for rows.Next() {
		var res groupResp
		err := rows.Scan(&res.ID, &res.Name)
		if err != nil {
			h.Logger(c).Errorf("Get income group error: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"code":    "I-5001",
				"message": "System error, please try again",
			})
		}
		g = append(g, res)
	}
	err = rows.Err()
	if err != nil {
		h.Logger(c).Errorf("Cannot Get income group error: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "I-5002",
			"message": "System error, please try again",
		})
	}

	return c.JSON(http.StatusOK, g)
}
