package adapters

import (
	"context"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/domain"
	"github.com/google/uuid"
)

type DiscountRepository interface {
	Get(ctx context.Context, id uuid.UUID) (*domain.DiscountGet, error)
	List(ctx context.Context, filter *domain.DiscountFilter) ([]domain.DiscountGet, error)
	EndDiscount(ctx context.Context, id uuid.UUID) error
	Create(ctx context.Context, discount *domain.DiscountCreate) (*uuid.UUID, error)
	Update(ctx context.Context, discount *domain.DiscountUpdate) error
}
