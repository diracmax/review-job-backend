package main

import (
	"github.com/diracmax/review-job-backend/internal/adapter/controller"
	"github.com/diracmax/review-job-backend/internal/adapter/repository"
	"github.com/diracmax/review-job-backend/internal/usecase"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/labstack/echo/v4/middleware"
)

var e = createMux()

func main() {
	e.GET("/health", healthCheck)

	repo := repository.NewUserRepository()
	uc := usecase.NewUserUsecase(repo)
	uct := controller.NewUserController(uc)

	e.POST("/user", uct.Register)
	e.Logger.Fatal(e.Start(":8080"))
}

func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	return e
}

func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
