package api

import (
	"github.com/ArtemFed/hse-wishlist/pkg/xhttp/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authHeader   = "Authorization"
	ctxAccountId = "accountId"
)

func (h *Handler) userIdentity(ctx *gin.Context) *response.HttpResponse {
	header := ctx.GetHeader(authHeader)
	if header == "" {
		return response.NewHttpResponseWithMessage(http.StatusUnauthorized, "empty auth header")
	}

	bearerToken := strings.TrimSuffix(header, "Bearer ")
	accountId, err := h.authService.ParseToken(bearerToken)
	if err != nil {
		return response.NewHttpResponseWithMessage(http.StatusUnauthorized, "invalid auth")
	}

	ctx.Set(ctxAccountId, accountId)
	return nil
}
