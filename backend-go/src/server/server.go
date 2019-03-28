package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"server/handlers"
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

	e.Static("/a", "public")

	e.Logger.Fatal(e.Start(":8080"))
}
