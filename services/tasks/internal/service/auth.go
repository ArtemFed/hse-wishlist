package service

import (
	"context"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/domain"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/service/adapters"
)

const (
	spanDefaultAuth = "auth/service."
)

var _ adapters.AuthService = &authService{}

type authService struct{}

func NewAuthService() adapters.AuthService {
	return &authService{}
}

func (a *authService) Login(ctx context.Context, params domain.AccountAuth) (string, error) {
	_, _, span := domain.GetTracerSpan(ctx, adapters.ServiceAuth, spanDefaultAuth, ".List")
	defer span.End()

	return "", nil
}
