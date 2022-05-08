package controller

import (
	"fmt"
	database "go-todo-api/db"
	model "go-todo-api/model"
	"net/http"
	"strconv"

	validator "github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type (
	Request struct {
		Content  string `json:"content" validate:"required"`
		Priority uint   `json:"priority" validate:"min=0"`
		Done     bool   `json:"is_done"`
	}

	Result struct {
		Status string `json:"status"`
		Id     uint   `json:"id"`
	}

	CustomValidator struct {
		validator *validator.Validate
	}
)

var db *gorm.DB

func init() {
	db = database.Connect()
}

func RegistTodo(c echo.Context) error {
	req := new(Request)
	if err := c.Bind(req); err != nil {
		return fmt.Errorf("bind error : %w", err)
	}

	if err := c.Validate(req); err != nil {
		return fmt.Errorf("validate error : %w", err)
	}

	todo := &model.Todo{
		Content:  req.Content,
		Priority: uint(req.Priority),
		Done:     req.Done,
	}
	qr := db.Create(todo)
	if qr.Error != nil {
		return fmt.Errorf("database error : %w", qr.Error)
	}

	todos, err := allList()
	if err != nil {
		return fmt.Errorf("database error : %w", err)
	}
	return c.JSON(http.StatusOK, todos)
}

func UpdateTodo(c echo.Context) error {
	p := c.Param("id")

	todo := &model.Todo{}
	id, err := strconv.Atoi(p)
	if err != nil {
		return fmt.Errorf("param error : %w", err)
	}

	qr := db.First(todo, id)
	if qr.Error != nil {
		return fmt.Errorf("database error : %w", qr.Error)
	}

	req := new(Request)
	if err := c.Bind(req); err != nil {
		return fmt.Errorf("bind error : %w", err)
	}
	if err := c.Validate(req); err != nil {
		return fmt.Errorf("validate error : %w", err)
	}

	todo.Content = req.Content
	todo.Priority = req.Priority
	todo.Done = req.Done
	qr = db.Save(todo)
	if qr.Error != nil {
		return fmt.Errorf("database error : %w", qr.Error)
	}

	return c.JSON(http.StatusOK, &Result{Status: "OK", Id: todo.Id})
}

func TodoList(c echo.Context) error {
	todos, err := allList()
	if err != nil {
		return fmt.Errorf("database error : %w", err)
	}
	return c.JSON(http.StatusOK, todos)
}

func DeleteTodo(c echo.Context) error {
	p := c.Param("id")

	id, err := strconv.Atoi(p)
	if err != nil {
		return fmt.Errorf("param error : %w", err)
	}

	qr := db.Delete(&model.Todo{}, id)
	if qr.Error != nil {
		return fmt.Errorf("database error : %w", qr.Error)
	}

	return c.JSON(http.StatusOK, &Result{Status: "OK", Id: uint(id)})
}

func allList() ([]model.Todo, error) {
	var todos []model.Todo
	qr := db.Order("id asc").Find(&todos)
	if qr.Error != nil {
		return nil, qr.Error
	}
	return todos, nil
}

func NewCustomValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}
