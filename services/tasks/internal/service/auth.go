package service

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/domain"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/service/adapters"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"os"
	"time"
)

const (
	spanDefaultAuth = "auth/service."
	salt            = "hse2025"
	signingKey      = "2025hse"
)

type tokenClaims struct {
	jwt.StandardClaims
	AccountUUID uuid.UUID `json:"user_uuid"`
	Login       string    `json:"login"`
}

var _ adapters.AuthService = &authService{}

type authService struct {
	ac adapters.AccountService
}

func NewAuthService(accountService adapters.AccountService) adapters.AuthService {
	return &authService{accountService}
}

func (a *authService) Login(ctx context.Context, params domain.AccountAuth) (string, error) {
	_, _, span := domain.GetTracerSpan(ctx, adapters.ServiceAuth, spanDefaultAuth, ".List")
	defer span.End()

	list, err := a.ac.List(ctx, domain.AccountFilter{
		Login: &params.Login,
	})
	if err != nil {
		return "", err
	}

	if len(list) == 0 || !verifyUserPass(list[0], params.Password) {
		return "", errors.New("credentials are incorrect")
	}

	token, err := a.GenerateToken(list[0])
	if err != nil {
		return "", err
	}

	return token, nil
}

func (a *authService) ParseToken(accessToken string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return uuid.UUID{}, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return uuid.UUID{}, errors.New("token claims are not of type")
	}
	return claims.AccountUUID, nil
}

func (a *authService) GenerateToken(account domain.AccountGet) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		account.UUID,
		account.Login,
	})

	jwtSigningKey := signingKey
	if os.Getenv("AUTH_SIGNING_KEY") != "" {
		jwtSigningKey = os.Getenv("AUTH_SIGNING_KEY")
	}
	return token.SignedString([]byte(jwtSigningKey))
}

func verifyUserPass(account domain.AccountGet, password string) bool {
	return account.Password == generatePasswordHash(password)
}

func generatePasswordHash(password string) string {
	//TODO заменить на bcrypt
	hash := sha1.New()
	hash.Write([]byte(password))
	passwordSalt := salt
	if os.Getenv("AUTH_PASSWORD_SALT") != "" {
		passwordSalt = os.Getenv("AUTH_PASSWORD_SALT")
	}

	return fmt.Sprintf("%x", hash.Sum([]byte(passwordSalt)))
}
