package adapters

import (
	"context"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/domain"
	"github.com/google/uuid"
)

const (
	ServiceTask    = "task-service"
	ServiceAccount = "account-service"
	ServiceAuth    = "auth-service"
)

type TaskService interface {
	List(ctx context.Context, filter *domain.TaskFilter) ([]domain.TaskGet, error)
	Patch(ctx context.Context, id uuid.UUID, status string) error
	Create(ctx context.Context, task *domain.TaskCreate) (*uuid.UUID, error)
	Update(ctx context.Context, task *domain.TaskUpdate) error
}

type AccountService interface {
	List(ctx context.Context, filter *domain.AccountFilter) ([]domain.AccountGet, error)
	Create(ctx context.Context, task *domain.AccountCreate) (*uuid.UUID, error)
	Update(ctx context.Context, task *domain.AccountUpdate) error
}

type AuthService interface {
	Login(ctx context.Context, params *domain.AccountAuth) (string, error)
}
