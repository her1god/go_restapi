package controllers

import (
	"database/sql"
	"encoding/json"
	"golang_todo/entities"
	"net/http"

	"github.com/labstack/echo"
)

func PatchCheckController(e *echo.Echo, db *sql.DB) {
	e.PATCH("/todos/:id/check", func(ctx echo.Context) error {
		id := ctx.Param("id")

		var request entities.CheckRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request)

		var doneInt int

		if request.Done == true {
			doneInt = 1
		}
		_, err := db.Exec(
			"UPDATE todos SET done = ? WHERE id = ?",
			doneInt,
			id,
		)

		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "OK")
	})
}
