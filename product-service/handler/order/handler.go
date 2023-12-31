package order

import (
	"app/handler/common"
	"app/internal/core/ports"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ohmspeed777/go-pkg/jwtx"
	"github.com/ohmspeed777/go-pkg/logx"
	"github.com/pkg/errors"
)

type Dependencies struct {
	OrderService ports.OrderService
}

type Handler struct {
	OrderService ports.OrderService
	transformer  *transformer
}

func NewHandler(d Dependencies) *Handler {
	return &Handler{
		OrderService: d.OrderService,
		transformer:  new(transformer),
	}
}

func (h *Handler) GetAll(ctx echo.Context) error {
	var (
		req = common.Query{}
	)

	errs := ctx.Bind(&req)
	if errs != nil {
		return errors.WithStack(errs)
	}

	e := h.transformer.toQueryRequest(req)
	user, ok := ctx.Get("user").(*jwtx.User)
	if ok {
		e.UserID = user.ID
	}

	entity, err := h.OrderService.FindAll(ctx.Request().Context(), e)
	if err != nil {
		logx.GetLog().Error(err)
		return err
	}

	resp := h.transformer.toPaginateResponse(entity.Entities, common.PageInfo(entity.PageInfo))
	return ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) Create(ctx echo.Context) error {
	var (
		req = Order{}
	)

	errs := ctx.Bind(&req)
	if errs != nil {
		return errors.WithStack(errs)
	}

	user, ok := ctx.Get("user").(*jwtx.User)
	if ok {
		req.UserID = user.ID
	}

	_, err := h.OrderService.Create(ctx.Request().Context(), h.transformer.toRequest(req))
	if err != nil {
		logx.GetLog().Error(err)
		return err
	}

	return ctx.JSON(http.StatusOK, nil)
}
