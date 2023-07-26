package product

import (
	"app/handler/common"
	"app/internal/core/ports"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type Dependencies struct {
	ProductService ports.ProductService
}

type Handler struct {
	ProductService ports.ProductService
	transformer    *transformer
}

func NewHandler(d Dependencies) *Handler {
	return &Handler{
		ProductService: d.ProductService,
		transformer:    new(transformer),
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

	entity, err := h.ProductService.FindAll(ctx.Request().Context(), h.transformer.toQueryRequest(req))
	if err != nil {
		return err
	}

	resp := h.transformer.toPaginateResponse(entity.Entities, common.PageInfo(entity.PageInfo))
	return ctx.JSON(http.StatusOK, resp)
}



func (h *Handler) Create(ctx echo.Context) error {
	var (
		req = Product{}
	)

	errs := ctx.Bind(&req)
	if errs != nil {
		return errors.WithStack(errs)
	}

	entity, err := h.ProductService.Create(ctx.Request().Context(), h.transformer.toRequest(req))
	if err != nil {
		return err
	}

	resp := h.transformer.toResponse(entity)
	return ctx.JSON(http.StatusOK, resp)
}