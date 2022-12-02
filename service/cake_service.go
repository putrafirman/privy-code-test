package service

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"putrafirman.com/playground/privy-code-test/repository"
)

func GetAllCake(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, repository.GetAll(db))
	}
}

func GetCakeDetail(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idInt, _ := strconv.Atoi(id)
		return c.JSON(http.StatusOK, repository.GetOne(db, idInt))
	}
}

func CreateCake(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		var cake repository.Cake

		c.Bind(&cake)
		repository.Create(db, cake)
		return nil
	}
}

func UpdateCake(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		id, _ := strconv.Atoi(c.Param("id"))

		var cake repository.Cake

		c.Bind(&cake)

		cake.ID = id

		repository.Update(db, cake)

		return nil
	}
}

func DeleteCake(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		repository.Delete(db, id)
		return nil
	}
}
