package main

import (
	"golang_todo/controllers"
	"golang_todo/database"

	"github.com/labstack/echo"
)

func main() {
	db := database.InitDb()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	controllers.GetController(e, db)

	controllers.PatchCheckController(e, db)

	controllers.PatchController(e, db)

	controllers.DeleteController(e, db)

	controllers.PostController(e, db)

	e.Start(":8080")
}
