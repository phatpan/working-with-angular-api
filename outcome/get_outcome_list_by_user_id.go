package outcome

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

type outcomeResp struct {
	ID               int    `json:"id"`
	OutcomeGroupID   int    `json:"outcomeGroupId"`
	OutcomeGroupName string `json:"outcomeGroupName"`
	Name             string `json:"name"`
	Amount           int    `json:"amount"`
	Date             string `json:"date"`
}

func (h *Handler) getOutcomeListByUserID(c echo.Context) error {
	uid := c.Param("id")

	oc := []outcomeResp{}

	stmt := `select i.id, i.outcome_group_id, g.name, i.name,
	i.amount, i.date from outcome i
			JOIN outcome_group g on g.id = i.outcome_group_id
			where user_id = ? `

	rows, err := h.DB.Query(stmt, uid)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, nil)
		}
		h.Logger(c).Errorf("getOutcomeListByUserID error: %v", err)
		c.JSON(http.StatusInternalServerError, err)
	}
	defer rows.Close()

	for rows.Next() {
		var res outcomeResp
		groupName := ""
		err := rows.Scan(&res.ID, &res.OutcomeGroupID, &groupName, &res.Name, &res.Amount, &res.Date)
		if err != nil {
			h.Logger(c).Errorf("getOutcomeListByUserID error: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"code":    "O-5005",
				"message": "System error, please try again",
			})
		}
		res.OutcomeGroupName = groupName
		oc = append(oc, res)
	}
	err = rows.Err()
	if err != nil {
		h.Logger(c).Errorf("Cannot getOutcomeListByUserID error: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "O-5006",
			"message": "System error, please try again",
		})
	}

	return c.JSON(http.StatusOK, oc)
}
