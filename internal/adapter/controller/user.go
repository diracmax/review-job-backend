package controller

import (
	"github.com/diracmax/review-job-backend/internal/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type User struct {
	Name        string `json:"name"`
	RawPassword string `json:"raw_password"`
}

type Message struct {
	Content string `json:"message"`
}

type userController struct {
	usecase usecase.UserUsecase
}

func NewUserController(usecase usecase.UserUsecase) *userController {
	return &userController{
		usecase: usecase,
	}
}

func (c *userController) Register(ctx echo.Context) error {
	u := new(User)
	if err := ctx.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err := c.usecase.Register(u.Name, u.RawPassword)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	m := &Message{
		Content: "success",
	}
	return ctx.JSON(http.StatusOK, m)
}
