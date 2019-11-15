package outcome

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

type groupResp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) getOutcomeGroupList(c echo.Context) error {
	stmt := "select id, name from outcome_group"
	g := []groupResp{}

	rows, err := h.DB.Query(stmt)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, nil)
		}
		h.Logger(c).Errorf("Get outcome group error: %v", err)
		c.JSON(http.StatusInternalServerError, err)
	}
	defer rows.Close()

	for rows.Next() {
		var res groupResp
		err := rows.Scan(&res.ID, &res.Name)
		if err != nil {
			h.Logger(c).Errorf("Get outcome group error: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"code":    "O-5001",
				"message": "System error, please try again",
			})
		}
		g = append(g, res)
	}
	err = rows.Err()
	if err != nil {
		h.Logger(c).Errorf("Cannot Get outcome group error: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "O-5002",
			"message": "System error, please try again",
		})
	}

	return c.JSON(http.StatusOK, g)
}
