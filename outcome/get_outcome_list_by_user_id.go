package outcome

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

type incomeResp struct {
	IncomeGroupID int `json:"incomeGroupId"`
	Amount        int `json:"amount"`
}

func (h *Handler) getOutcomeListByUserID(c echo.Context) error {
	uid := c.Param("id")

	stmt := "select income_group_id, amount from income where email = ?"
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
		err := rows.Scan(&res.IncomeGroupID, &res.Amount)
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
