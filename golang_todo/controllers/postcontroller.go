package controllers

import (
	"database/sql"
	"encoding/json"
	"golang_todo/entities"
	"net/http"

	"github.com/labstack/echo"
)

func PostController(e *echo.Echo, db *sql.DB) {
	e.POST("/todos", func(ctx echo.Context) error {

		var request entities.CreateRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request)

		_, err := db.Exec(
			"INSERT INTO todos (title, description, done) VALUES (?, ?, 0)",
			request.Title,
			request.Description,
		)

		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "OK")
	})
}
