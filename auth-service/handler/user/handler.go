package user

import (
	"app/internal/core/ports"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ohmspeed777/go-pkg/errorx"
	"github.com/ohmspeed777/go-pkg/logx"
)

type Dependencies struct {
	UserService ports.UserService
}

type Handler struct {
	UserService ports.UserService
	transformer *transformer
}

func NewHandler(d Dependencies) *Handler {
	return &Handler{
		UserService: d.UserService,
		transformer: new(transformer),
	}
}

func (h *Handler) SignUp(c echo.Context) error {
	req := SignUpReq{}
	if err := c.Bind(&req); err != nil {
		return errorx.NewInvalidRequest(err)
	}

	err := h.UserService.SignUp(c.Request().Context(), h.transformer.toSingUpEntity(req))
	if err != nil {
		logx.GetLog().Error(err)
		return err
	}

	return c.JSON(http.StatusCreated, nil)
}

func (h *Handler) SignIn(c echo.Context) error {
	req := SignInReq{}
	if err := c.Bind(&req); err != nil {
		return errorx.NewInvalidRequest(err)
	}

	res, err := h.UserService.SignIn(c.Request().Context(), h.transformer.toSingInEntity(req))
	if err != nil {
		logx.GetLog().Error(err)
		return err
	}

	return c.JSON(http.StatusOK, h.transformer.toSingInJSON(*res))
}
