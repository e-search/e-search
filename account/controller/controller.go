package controller

import (
	"net/http"

	"github.com/e-search/e-search/account"
	"github.com/e-search/e-search/account/usecase"
	"github.com/labstack/echo"
)

// AccountController controller struct
type AccountController struct {
	AccountUsecase usecase.AccountUsecase
}

// Create account create controller
func (c *AccountController) Create(ctx echo.Context) error {
	model := account.Account{}
	err := ctx.Bind(&model)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	res, err := c.AccountUsecase.Create(&model)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, res)
}

// NewAccountController mount accountcontroller
func NewAccountController(e *echo.Echo, us usecase.AccountUsecase) {
	handler := &AccountController{
		AccountUsecase: us,
	}

	e.POST("/accounts", handler.Create)
}
