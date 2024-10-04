package http

import (
	"fmt"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/config"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/handler/http/discount"
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
	cfg             *config.Config
	discountService adapters.DiscountService
}

func NewHandler(cfg *config.Config,
	discountService adapters.DiscountService,
) Handler {
	return Handler{
		cfg:             cfg,
		discountService: discountService,
	}
}

// HandleError is a sample error handler function
func HandleError(c *gin.Context, err error, statusCode int) {
	c.JSON(statusCode, gin.H{"error": err.Error()})
}

func ConvertToDiscount(middlewareMainArr []MiddlewareFunc) []discount.MiddlewareFunc {
	result := make([]discount.MiddlewareFunc, len(middlewareMainArr))
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

	discount.RegisterHandlersWithOptions(router,
		discount.NewDiscountHandler(handler.discountService),
		discount.GinServerOptions{
			BaseURL:      baseUrl,
			Middlewares:  ConvertToDiscount(middlewares),
			ErrorHandler: HandleError,
		})
}

func getVersion() string {
	return fmt.Sprintf("v%s", strings.Split(version, ".")[0])
}
