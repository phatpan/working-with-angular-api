package income

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

type incomeResp struct {
	ID            int    `json:"id"`
	IncomeGroupID int    `json:"incomeGroupId"`
	Amount        int    `json:"amount"`
	Name          string `json:"name"`
	Date          string `json:"date"`
}

func (h *Handler) getIncomeListByUserID(c echo.Context) error {
	uid := c.Param("id")

	stmt := "select id, income_group_id, amount, name, date from income where user_id = ?"
	income := []incomeResp{}

	rows, err := h.DB.Query(stmt, uid)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, nil)
		}
		h.Logger(c).Errorf("getIncomeByEmail error: %v", err)
		c.JSON(http.StatusInternalServerError, err)
	}
	defer rows.Close()

	for rows.Next() {
		var res incomeResp
		err := rows.Scan(&res.ID, &res.IncomeGroupID, &res.Amount, &res.Name, &res.Date)
		if err != nil {
			h.Logger(c).Errorf("getIncomeByEmail error: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"code":    "I-5005",
				"message": "System error, please try again",
			})
		}
		income = append(income, res)
	}
	err = rows.Err()
	if err != nil {
		h.Logger(c).Errorf("Cannot getIncomeByEmail error: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "I-5006",
			"message": "System error, please try again",
		})
	}

	return c.JSON(http.StatusOK, income)
}
