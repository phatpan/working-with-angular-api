package user

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) getUserByEmail(c echo.Context) error {
	email := c.Param("email")
	stmt := "select id, name from users where email = ?"
	user := userResponse{}
	rows, err := h.DB.Query(stmt, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, nil)
		}
		h.Logger(c).Errorf("getUserByEmail error: %v", err)
		c.JSON(http.StatusInternalServerError, err)
	}
	defer rows.Close()

	for rows.Next() {
		var u userResponse
		err := rows.Scan(&u.ID, &u.Name)
		if err != nil {
			h.Logger(c).Errorf("getUserByEmail error: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"code":    "U-5001",
				"message": "System error, please try again",
			})
		}
		user = u
	}
	err = rows.Err()
	if err != nil {
		h.Logger(c).Errorf("Cannot getUserByEmail error: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "U-5002",
			"message": "System error, please try again",
		})
	}

	return c.JSON(http.StatusOK, user)
}
