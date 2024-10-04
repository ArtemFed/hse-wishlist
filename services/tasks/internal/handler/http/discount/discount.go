package discount

import (
	"context"
	http2 "github.com/ArtemFed/hse-wishlist/services/tasks/internal/handler/http/utils"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/service/adapters"
	"github.com/gin-gonic/gin"
	openapitypes "github.com/oapi-codegen/runtime/types"
	global "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"io"
	"net/http"
)

const (
	spanDefaultDiscount = "discount/handler."
)

var _ ServerInterface = &Handler{}

type Handler struct {
	discountService adapters.DiscountService
}

func NewDiscountHandler(discountService adapters.DiscountService) *Handler {
	return &Handler{discountService: discountService}
}

func getDiscountTracerSpan(ctx *gin.Context, name string) (trace.Tracer, context.Context, trace.Span) {
	tr := global.Tracer(adapters.ServiceNameDiscount)
	newCtx, span := tr.Start(ctx, spanDefaultDiscount+name)

	return tr, newCtx, span
}

func (h *Handler) GetDiscountsId(ctx *gin.Context, uuid openapitypes.UUID) {
	_, newCtx, span := getDiscountTracerSpan(ctx, ".GetByLogin")
	defer span.End()

	model, err := h.discountService.Get(newCtx, uuid)
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	response := DomainToGet(*model)

	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) GetDiscounts(ctx *gin.Context) {
	_, newCtx, span := getDiscountTracerSpan(ctx, ".GetList")
	defer span.End()

	var body *DiscountFilter
	if err := ctx.BindJSON(&body); err != nil && err != io.EOF {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	domains, err := h.discountService.List(newCtx, FilterToDomain(body))
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	list := make([]DiscountGet, len(domains))
	for i, item := range domains {
		list[i] = DomainToGet(item)
	}

	ctx.JSON(http.StatusOK, list)
}

func (h *Handler) PostDiscounts(ctx *gin.Context) {
	_, newCtx, span := getDiscountTracerSpan(ctx, ".Create")
	defer span.End()

	var body DiscountCreate
	if err := ctx.BindJSON(&body); err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	model := CreateToDomain(body)
	id, err := h.discountService.Create(newCtx, &model)
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	ctx.JSON(http.StatusOK, ModelUUID{UUID: *id})
}

func (h *Handler) PutDiscounts(ctx *gin.Context) {
	_, newCtx, span := getDiscountTracerSpan(ctx, ".Update")
	defer span.End()

	var body DiscountUpdate
	if err := ctx.BindJSON(&body); err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	model := UpdateToDomain(body)
	err := h.discountService.Update(newCtx, &model)
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	ctx.JSON(http.StatusOK, http.NoBody)
}

func (h *Handler) PatchDiscountsIdEnd(ctx *gin.Context, id openapitypes.UUID) {
	_, newCtx, span := getDiscountTracerSpan(ctx, ".End")
	defer span.End()

	err := h.discountService.EndDiscount(newCtx, id)
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	ctx.JSON(http.StatusOK, http.NoBody)
}
