package main

import (
	controller "go-todo-api/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Validator = controller.NewCustomValidator()

	e.POST("/regist", controller.RegistTodo)
	e.PUT("/update/:id", controller.UpdateTodo)
	e.DELETE("/delete/:id", controller.DeleteTodo)
	e.GET("/list", controller.TodoList)

	e.HTTPErrorHandler = controller.ApiErrorHandler

	e.Logger.Fatal(e.Start(":1323"))
}
