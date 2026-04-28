package main

import (
	"fmt"
	"ner/http"
	"github/com/labstack/echo/v4"
)

func main(){
	e = echo.New()

	e.GET("/", func(c echo.Context) error{
		return c.String(http.StatusOK, "Привет !")
	})

	e.start(":8000")
}