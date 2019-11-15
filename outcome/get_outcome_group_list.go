package outcome

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) getOutcomeGroupList(c echo.Context) error {
	stmt := "select id, title from food"
	foods := []foodResp{}

	rows, err := h.DB.Query(stmt)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, nil)
		}
		h.Logger(c).Errorf("Get FoodList error: %v", err)
		c.JSON(http.StatusInternalServerError, err)
	}
	defer rows.Close()

	for rows.Next() {
		var f foodResp
		err := rows.Scan(&f.ID, &f.Title)
		if err != nil {
			h.Logger(c).Errorf("Get FoodList error: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"code":    "F-5001",
				"message": "System error, please try again",
			})
		}
		foods = append(foods, f)
	}
	err = rows.Err()
	if err != nil {
		h.Logger(c).Errorf("Cannot Get FoodList error: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "F-5002",
			"message": "System error, please try again",
		})
	}

	return c.JSON(http.StatusOK, foods)
}
