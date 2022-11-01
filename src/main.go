package main

import (
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
)

var ctx = context.Background()

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello World!")
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	e.GET("/database", func(c echo.Context) error {
		rdb := redis.NewClient(&redis.Options{
			Addr:     "redis:6379",
			Password: "",
			DB:       0,
		})

		result, err := rdb.Ping(ctx).Result()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, struct{ Status string }{Status: err.Error()})

		}
		return c.JSON(http.StatusOK, struct{ Status string }{Status: result})
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
