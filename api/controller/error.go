package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ApiErrorHandler(err error, c echo.Context) {
	c.String(http.StatusInternalServerError, err.Error())
}
