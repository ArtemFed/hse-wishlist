package adapters

import (
	"context"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/domain"
	"github.com/google/uuid"
)

type TaskRepository interface {
	List(ctx context.Context, filter *domain.TaskFilter) ([]domain.TaskGet, error)
	PatchStatus(ctx context.Context, id uuid.UUID, status string) error
	Create(ctx context.Context, task *domain.TaskCreate) (*uuid.UUID, error)
	Update(ctx context.Context, task *domain.TaskUpdate) error
}

type AccountRepository interface {
	List(ctx context.Context, filter *domain.TaskFilter) ([]domain.TaskGet, error)
	Create(ctx context.Context, task *domain.TaskCreate) (*uuid.UUID, error)
	Update(ctx context.Context, task *domain.TaskUpdate) error
}
