package user

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

var ct time.Time

type UserReq struct {
	Email string `json:"email"`
}

type UserResponse struct {
	ID int64 `json:"id"`
}

func (h *Handler) saveUserByEmail(c echo.Context) error {
	u := &UserReq{}
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"code":    "U-5003",
			"message": "RequestBody invalid",
		})
	}

	return h.insertUserByDeviceIDTable(c, u)
}

func (h *Handler) insertUserByDeviceIDTable(c echo.Context, u *UserReq) error {
	ct = time.Now()
	ct.Format(time.RFC3339)
	resp := UserResponse{}

	stmt := "select id from users where email = ?"
	rows, _ := h.DB.Query(stmt, u.Email)
	for rows.Next() {
		_ = rows.Scan(&resp.ID)
		if resp.ID != 0 {
			return c.JSON(http.StatusOK, resp)
		}
	}
	defer rows.Close()

	stmtIns := `INSERT INTO users (
	email, created_date)
	VALUES (?, ?)`

	res, err := h.DB.Exec(stmtIns, u.Email, ct)
	uid, err := res.LastInsertId()

	if err != nil {
		h.Logger(c).Errorf("insertUserByEmailTable error: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"code":    "U-5004",
			"message": "System error, please try again",
		})
	}

	resp.ID = uid

	return c.JSON(http.StatusOK, resp)
}
