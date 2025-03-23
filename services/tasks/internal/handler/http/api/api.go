package api

import (
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/domain"
	http2 "github.com/ArtemFed/hse-wishlist/services/tasks/internal/handler/http/utils"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/service/adapters"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	spanDefaultTask    = "task/handler."
	spanDefaultAccount = "account/handler."
	spanDefaultAuth    = "auth/handler."
)

var _ ServerInterface = &Handler{}

type Handler struct {
	taskService    adapters.TaskService
	accountService adapters.AccountService
	authService    adapters.AuthService
}

func NewTaskHandler(taskService adapters.TaskService) *Handler {
	return &Handler{taskService: taskService}
}

func (h *Handler) GetAccounts(ctx *gin.Context, params GetAccountsParams) {
	_, newCtx, span := domain.GetTracerSpan(ctx, adapters.ServiceAccount, spanDefaultAccount, ".Get")
	defer span.End()

	domains, err := h.accountService.List(newCtx, AccountFilterToDomain(params))
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	list := make([]AccountGet, len(domains))
	for i, item := range domains {
		list[i] = AccountDomainToGet(item)
	}

	ctx.JSON(http.StatusOK, list)
}

func (h *Handler) PostAccounts(ctx *gin.Context) {
	_, newCtx, span := domain.GetTracerSpan(ctx, adapters.ServiceAccount, spanDefaultAccount, ".Create")
	defer span.End()

	var body AccountCreate
	if err := ctx.BindJSON(&body); err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	model := AccountCreateToDomain(body)
	id, err := h.accountService.Create(newCtx, model)
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	ctx.JSON(http.StatusOK, ModelUUID{Id: *id})
}

func (h *Handler) PutAccounts(ctx *gin.Context) {
	_, newCtx, span := domain.GetTracerSpan(ctx, adapters.ServiceAccount, spanDefaultAccount, ".Update")
	defer span.End()

	var body AccountUpdate
	if err := ctx.BindJSON(&body); err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	model := AccountUpdateToDomain(body)
	err := h.accountService.Update(newCtx, model)
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	ctx.JSON(http.StatusOK, http.NoBody)
}

func (h *Handler) PostAuth(ctx *gin.Context) {
	_, newCtx, span := domain.GetTracerSpan(ctx, adapters.ServiceAuth, spanDefaultAuth, ".Login")
	defer span.End()

	var body AccountAuth
	if err := ctx.BindJSON(&body); err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	model := AccountAuthToDomain(body)
	token, err := h.authService.Login(newCtx, model)
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	ctx.JSON(http.StatusOK, JwtToken{Token: &token})
}

func (h *Handler) GetTasks(ctx *gin.Context, params GetTasksParams) {
	_, newCtx, span := domain.GetTracerSpan(ctx, adapters.ServiceTask, spanDefaultTask, ".Get")
	defer span.End()

	domains, err := h.taskService.List(newCtx, TaskFilterToDomain(params))
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	list := make([]TaskGet, len(domains))
	for i, item := range domains {
		list[i] = TaskDomainToGet(item)
	}

	ctx.JSON(http.StatusOK, list)
}

func (h *Handler) PatchTasks(ctx *gin.Context, params PatchTasksParams) {
	_, newCtx, span := domain.GetTracerSpan(ctx, adapters.ServiceTask, spanDefaultTask, ".Patch")
	defer span.End()

	err := h.taskService.Patch(newCtx, params.Id, params.Status)
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	ctx.JSON(http.StatusOK, http.NoBody)
}

func (h *Handler) PostTasks(ctx *gin.Context) {
	_, newCtx, span := domain.GetTracerSpan(ctx, adapters.ServiceTask, spanDefaultTask, ".Create")
	defer span.End()

	var body TaskCreate
	if err := ctx.BindJSON(&body); err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	model := TaskCreateToDomain(body)
	id, err := h.taskService.Create(newCtx, model)
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	ctx.JSON(http.StatusOK, ModelUUID{Id: *id})
}

func (h *Handler) PutTasks(ctx *gin.Context) {
	_, newCtx, span := domain.GetTracerSpan(ctx, adapters.ServiceTask, spanDefaultTask, ".Update")
	defer span.End()

	var body TaskUpdate
	if err := ctx.BindJSON(&body); err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	model := TaskUpdateToDomain(body)
	err := h.taskService.Update(newCtx, model)
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	ctx.JSON(http.StatusOK, http.NoBody)
}
