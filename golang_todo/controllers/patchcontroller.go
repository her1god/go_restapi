package controllers

import (
	"database/sql"
	"encoding/json"
	"golang_todo/entities"
	"net/http"

	"github.com/labstack/echo"
)

func PatchController(e *echo.Echo, db *sql.DB) {
	e.PATCH("/todos/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")

		var request entities.UpdateRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request)

		_, err := db.Exec(
			"UPDATE todos SET title = ?, description = ? WHERE id = ?",
			request.Title,
			request.Description,
			id,
		)

		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "OK")
	})
}
