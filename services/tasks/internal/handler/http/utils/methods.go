package utils

import (
	"fmt"
	"github.com/ArtemFed/hse-wishlist/pkg/xerror"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/handler/http/dto"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/log"
	"github.com/gin-gonic/gin"
)

func AbortWithBadResponse(c *gin.Context, statusCode int, err error) {
	log.Logger.Debug(fmt.Sprintf("%s: %d %s", c.Request.URL, statusCode, xerror.GetLastMessage(err)))
	c.AbortWithStatusJSON(statusCode, dto.Error{Message: xerror.GetLastMessage(err)})
}

func AbortWithErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Logger.Error(fmt.Sprintf("%s: %d %s", c.Request.URL, statusCode, message))
	c.AbortWithStatusJSON(statusCode, dto.Error{Message: message})
}

func MapErrorToCode(err error) int {
	return xerror.GetCode(err)
}
