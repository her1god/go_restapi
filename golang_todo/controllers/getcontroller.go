package controllers

import (
	"database/sql"
	"golang_todo/entities"
	"net/http"

	"github.com/labstack/echo"
)

func GetController(e *echo.Echo, db *sql.DB) {
	e.GET("/todos", func(ctx echo.Context) error {
		rows, err := db.Query("SELECT * FROM todos")
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		var res []entities.TodoResponse
		for rows.Next() {
			var id int
			var title string
			var description string
			var done int
			err := rows.Scan(&id, &title, &description, &done)
			if err != nil {
				return ctx.String(http.StatusInternalServerError, err.Error())
			}

			var todo entities.TodoResponse
			todo.Id = id
			todo.Title = title
			todo.Description = description
			if done == 1 {
				todo.Done = true
			}

			res = append(res, todo)
		}

		return ctx.JSON(http.StatusOK, res)
	})
}
