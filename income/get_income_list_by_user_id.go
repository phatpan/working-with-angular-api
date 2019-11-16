package income

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

type incomeResp struct {
	ID              int    `json:"id"`
	IncomeGroupID   int    `json:"incomeGroupId"`
	IncomeGroupName string `json:"incomeGroupName"`
	Amount          int    `json:"amount"`
	Date            string `json:"date"`
}

func (h *Handler) getIncomeListByUserID(c echo.Context) error {
	uid := c.Param("id")
	income := []incomeResp{}

	stmt := `select i.id, i.income_group_id, g.name,
			i.amount, i.date from income i
			JOIN income_group g on g.id = i.income_group_id
			where user_id = ? `

	rows, err := h.DB.Query(stmt, uid)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, nil)
		}
		h.Logger(c).Errorf("getIncomeListByUserID error: %v", err)
		c.JSON(http.StatusInternalServerError, err)
	}
	defer rows.Close()

	for rows.Next() {
		var res incomeResp
		name := ""
		err := rows.Scan(&res.ID, &res.IncomeGroupID, &name, &res.Amount, &res.Date)
		if err != nil {
			h.Logger(c).Errorf("getIncomeListByUserID error: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"code":    "I-5005",
				"message": "System error, please try again",
			})
		}
		res.IncomeGroupName = name
		income = append(income, res)
	}
	err = rows.Err()
	if err != nil {
		h.Logger(c).Errorf("Cannot getIncomeListByUserID error: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "I-5006",
			"message": "System error, please try again",
		})
	}

	return c.JSON(http.StatusOK, income)
}
