package http

import (
	"fmt"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/config"
	task "github.com/ArtemFed/hse-wishlist/services/tasks/internal/handler/http/task"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/service/adapters"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	apiPrefix = "api"
	version   = "1"
)

type MiddlewareFunc func(c *gin.Context)

type Handler struct {
	cfg         *config.Config
	taskService adapters.TaskService
}

func NewHandler(cfg *config.Config,
	taskService adapters.TaskService,
) Handler {
	return Handler{
		cfg:         cfg,
		taskService: taskService,
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

func InitHandler(
	handler Handler,
	router gin.IRouter,
	middlewares []MiddlewareFunc,
	httpPrefix string,
) {
	baseUrl := fmt.Sprintf("%s/%s/%s", apiPrefix, getVersion(), httpPrefix)

	task.RegisterHandlersWithOptions(router,
		task.NewTaskHandler(handler.taskService),
		task.GinServerOptions{
			BaseURL:      baseUrl,
			Middlewares:  ConvertToTask(middlewares),
			ErrorHandler: HandleError,
		})
}

func getVersion() string {
	return fmt.Sprintf("v%s", strings.Split(version, ".")[0])
}
