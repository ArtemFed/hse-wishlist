package task

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
	spanDefaultTask = "task/handler."
)

var _ ServerInterface = &Handler{}

type Handler struct {
	taskService adapters.TaskService
}

func NewTaskHandler(taskService adapters.TaskService) *Handler {
	return &Handler{taskService: taskService}
}

func getTaskTracerSpan(ctx *gin.Context, name string) (trace.Tracer, context.Context, trace.Span) {
	tr := global.Tracer(adapters.ServiceNameTask)
	newCtx, span := tr.Start(ctx, spanDefaultTask+name)

	return tr, newCtx, span
}

func (h *Handler) GetTasksId(ctx *gin.Context, uuid openapitypes.UUID) {
	_, newCtx, span := getTaskTracerSpan(ctx, ".GetByLogin")
	defer span.End()

	model, err := h.taskService.Get(newCtx, uuid)
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	response := DomainToGet(*model)

	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) GetTasks(ctx *gin.Context) {
	_, newCtx, span := getTaskTracerSpan(ctx, ".GetList")
	defer span.End()

	var body *TaskFilter
	if err := ctx.BindJSON(&body); err != nil && err != io.EOF {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	domains, err := h.taskService.List(newCtx, FilterToDomain(body))
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	list := make([]TaskGet, len(domains))
	for i, item := range domains {
		list[i] = DomainToGet(item)
	}

	ctx.JSON(http.StatusOK, list)
}

func (h *Handler) PostTasks(ctx *gin.Context) {
	_, newCtx, span := getTaskTracerSpan(ctx, ".Create")
	defer span.End()

	var body TaskCreate
	if err := ctx.BindJSON(&body); err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	model := CreateToDomain(body)
	id, err := h.taskService.Create(newCtx, &model)
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	ctx.JSON(http.StatusOK, ModelUUID{UUID: *id})
}

func (h *Handler) PutTasks(ctx *gin.Context) {
	_, newCtx, span := getTaskTracerSpan(ctx, ".Update")
	defer span.End()

	var body TaskUpdate
	if err := ctx.BindJSON(&body); err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	model := UpdateToDomain(body)
	err := h.taskService.Update(newCtx, &model)
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	ctx.JSON(http.StatusOK, http.NoBody)
}

func (h *Handler) PatchTasksIdEnd(ctx *gin.Context, id openapitypes.UUID) {
	_, newCtx, span := getTaskTracerSpan(ctx, ".End")
	defer span.End()

	err := h.taskService.EndTask(newCtx, id)
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	ctx.JSON(http.StatusOK, http.NoBody)
}
