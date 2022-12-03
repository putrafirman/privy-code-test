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
		result, err := repository.GetAll(db)
		if err != nil {
			c.JSON(http.StatusBadRequest, nil)
		}
		return c.JSON(http.StatusOK, result)
	}
}

func GetCakeDetail(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idInt, _ := strconv.Atoi(id)

		result, err := repository.GetOne(db, idInt)
		if err != nil {
			c.JSON(http.StatusBadRequest, nil)
		}
		return c.JSON(http.StatusOK, result)
	}
}

func CreateCake(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		var cake repository.Cake

		c.Bind(&cake)
		err := repository.Create(db, cake)
		if err != nil {
			c.JSON(http.StatusBadRequest, nil)
		}
		return c.JSON(http.StatusOK, "")
	}
}

func UpdateCake(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		id, _ := strconv.Atoi(c.Param("id"))

		var cake repository.Cake

		c.Bind(&cake)

		cake.ID = id

		err := repository.Update(db, cake)
		if err != nil {
			c.JSON(http.StatusBadRequest, nil)
		}

		return c.JSON(http.StatusOK, "")
	}
}

func DeleteCake(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		err := repository.Delete(db, id)
		if err != nil {
			c.JSON(http.StatusBadRequest, nil)
		}
		return c.JSON(http.StatusOK, "")
	}
}
