package income

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

type searchResp struct {
	ID                int    `json:"id"`
	IncomeGroupID     int    `json:"incomeGroupId"`
	IncomeNameGroupID string `json:"incomeNameGroupId"`
	Amount            int    `json:"amount"`
	Date              string `json:"date"`
}

func (h *Handler) searchIncomeByUserIDWithText(c echo.Context) error {
	search := c.Param("search")
	uid := c.Param("id")

	stmt := `select i.id, i.income_group_id, g.name, i.amount, i.date from income i
			JOIN income_group g on g.id = i.income_group_id
			where g.name LIKE CONCAT('%',?,'%') or i.amount LIKE CONCAT('%',?,'%') and i.user_id= ?`
	income := []searchResp{}

	rows, err := h.DB.Query(stmt, search, search, uid)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, nil)
		}
		h.Logger(c).Errorf("searchIncomeByUserIDWithText error: %v", err)
		c.JSON(http.StatusInternalServerError, err)
	}
	defer rows.Close()

	for rows.Next() {
		var res searchResp
		err := rows.Scan(&res.ID, &res.IncomeGroupID, &res.IncomeNameGroupID, &res.Amount, &res.Date)
		if err != nil {
			h.Logger(c).Errorf("searchIncomeByUserIDWithText error: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"code":    "I-5005",
				"message": "System error, please try again",
			})
		}
		income = append(income, res)
	}
	err = rows.Err()
	if err != nil {
		h.Logger(c).Errorf("searchIncomeByUserIDWithText error: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "I-5006",
			"message": "System error, please try again",
		})
	}

	return c.JSON(http.StatusOK, income)
}
