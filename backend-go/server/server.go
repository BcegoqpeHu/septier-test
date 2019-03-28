package main

import (
	"github.com/BcegoqpeHu/septier-test/backend-go/server/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)

	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	h := handlers.NewHandler()

	go h.Hub.Run()

	e.POST("/addmany", h.AddManyHandler)
	e.GET("/ws", h.NotificationsHandler)

	e.Static("/", "public")

	e.Logger.Fatal(e.Start(":8080"))
}
