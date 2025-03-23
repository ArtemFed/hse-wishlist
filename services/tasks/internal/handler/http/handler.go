package http

import (
	"errors"
	"fmt"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/config"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/domain"
	task "github.com/ArtemFed/hse-wishlist/services/tasks/internal/handler/http/api"
	http2 "github.com/ArtemFed/hse-wishlist/services/tasks/internal/handler/http/utils"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/service/adapters"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	apiPrefix    = "api"
	version      = "1"
	authHeader   = "Authorization"
	CtxAccountId = "accountId"
)

type MiddlewareFunc func(c *gin.Context)

type Handler struct {
	cfg            *config.Config
	taskService    adapters.TaskService
	accountService adapters.AccountService
	authService    adapters.AuthService
}

func NewHandler(
	cfg *config.Config,
	taskService adapters.TaskService,
	accountService adapters.AccountService,
	authService adapters.AuthService,
) Handler {
	return Handler{
		cfg:            cfg,
		taskService:    taskService,
		accountService: accountService,
		authService:    authService,
	}
}

// HandleError is a sample error handler function
func HandleError(c *gin.Context, err error, statusCode int) {
	c.JSON(statusCode, gin.H{"error": err.Error()})
}

func ConvertToTask(middlewareMainArr []MiddlewareFunc) []task.MiddlewareFunc {
	result := make([]task.MiddlewareFunc, len(middlewareMainArr))
	for i, middlewareMain := range middlewareMainArr {
		result[i] = func(c *gin.Context) {
			middlewareMain(c)
		}
	}
	return result
}

func InitMainHandler(
	handler Handler,
	router gin.IRouter,
	middlewares []MiddlewareFunc,
	httpPrefix string,
) {
	baseUrl := fmt.Sprintf("%s/%s/%s", apiPrefix, getVersion(), httpPrefix)

	task.RegisterHandlersWithOptions(router,
		task.NewMainHandler(
			handler.taskService,
			handler.accountService,
			handler.authService,
		),
		task.GinServerOptions{
			BaseURL:      baseUrl,
			Middlewares:  ConvertToTask(middlewares),
			ErrorHandler: HandleError,
		})
}

//func InitAuthHandler(
//	handler Handler,
//	router gin.IRouter,
//	middlewares []MiddlewareFunc,
//	httpPrefix string,
//) {
//	baseUrl := fmt.Sprintf("%s/%s/%s", apiPrefix, getVersion(), httpPrefix)
//
//	task.RegisterHandlersWithOptions(router,
//		task.NewAuthHandler(
//			handler.taskService,
//			handler.accountService,
//			handler.authService,
//		),
//		task.GinServerOptions{
//			BaseURL:      baseUrl,
//			Middlewares:  ConvertToTask(middlewares),
//			ErrorHandler: HandleError,
//		})
//}

func getVersion() string {
	return fmt.Sprintf("v%s", strings.Split(version, ".")[0])
}

// IdentityMiddleware middleware для проверки авторизации
func (h *Handler) IdentityMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Print("FullPath", ctx.FullPath())
		if strings.HasPrefix(ctx.FullPath(), "/api/v1/hse/auth") {
			ctx.Next()
			return
		}
		header := ctx.GetHeader(authHeader)
		if header == "" {
			http2.AbortWithBadResponse(ctx, http.StatusUnauthorized, errors.New("empty auth header"))
			return
		}

		bearerToken := strings.TrimPrefix(header, "Bearer ")
		accountId, err := h.authService.ParseToken(bearerToken)
		if err != nil {
			http2.AbortWithBadResponse(ctx, http.StatusUnauthorized, errors.New("invalid token 1"))
			return
		}

		list, err := h.accountService.List(ctx, domain.AccountFilter{
			UUID: &accountId,
		})
		if len(list) == 0 {
			http2.AbortWithBadResponse(ctx, http.StatusUnauthorized, errors.New("invalid token 2"))
			return
		}

		ctx.Set(CtxAccountId, accountId)
		ctx.Next()
	}
}
