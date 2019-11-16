package outcome

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

type searchResp struct {
	ID               int    `json:"id"`
	OutcomeGroupID   int    `json:"outcomeGroupId"`
	OutcomeGroupName string `json:"OutcomeGroupName"`
	Amount           int    `json:"amount"`
	Name             string `json:"name"`
	Date             string `json:"date"`
}

func (h *Handler) searchOutcomeByUserIDWithText(c echo.Context) error {
	search := c.Param("search")
	uid := c.Param("id")

	stmt := `select i.id, i.outcome_group_id, g.name, i.name,  i.amount, i.date from outcome i
			JOIN outcome_group g on g.id = i.outcome_group_id
			where g.name LIKE CONCAT('%',?,'%') or i.amount LIKE CONCAT('%',?,'%') and i.user_id= ?`
	oc := []searchResp{}

	rows, err := h.DB.Query(stmt, search, search, uid)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, nil)
		}
		h.Logger(c).Errorf("searchOutcomeByUserIDWithText error: %v", err)
		c.JSON(http.StatusInternalServerError, err)
	}
	defer rows.Close()

	for rows.Next() {
		var res searchResp
		err := rows.Scan(&res.ID, &res.OutcomeGroupID, &res.OutcomeGroupName, &res.Name,
			&res.Amount, &res.Date)
		if err != nil {
			h.Logger(c).Errorf("searchOutcomeByUserIDWithText error: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"code":    "O-5005",
				"message": "System error, please try again",
			})
		}
		oc = append(oc, res)
	}
	err = rows.Err()
	if err != nil {
		h.Logger(c).Errorf("searchOutcomeByUserIDWithText error: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "O-5006",
			"message": "System error, please try again",
		})
	}

	return c.JSON(http.StatusOK, oc)
}
