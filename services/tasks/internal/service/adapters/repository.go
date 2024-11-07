package adapters

import (
	"context"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/domain"
	"github.com/google/uuid"
)

type TaskRepository interface {
	Get(ctx context.Context, id uuid.UUID) (*domain.TaskGet, error)
	List(ctx context.Context, filter *domain.TaskFilter) ([]domain.TaskGet, error)
	EndTask(ctx context.Context, id uuid.UUID) error
	Create(ctx context.Context, task *domain.TaskCreate) (*uuid.UUID, error)
	Update(ctx context.Context, task *domain.TaskUpdate) error
}
