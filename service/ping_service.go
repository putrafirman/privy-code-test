package service

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
)

func GetPingAndPong(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		//db ping
		db.Ping()

		return c.String(http.StatusOK, "pong")
	}
}
