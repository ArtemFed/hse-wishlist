package service

import (
	"context"
	"errors"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/domain"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/service/adapters"
	"github.com/google/uuid"
	global "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

const (
	spanDefaultDiscount = "discount/service."
)

var _ adapters.DiscountService = &discountService{}

type discountService struct {
	r adapters.DiscountRepository
}

func NewDiscountService(discountRepository adapters.DiscountRepository) adapters.DiscountService {
	return &discountService{r: discountRepository}
}

func getDiscountTracerSpan(ctx context.Context, name string) (trace.Tracer, context.Context, trace.Span) {
	tr := global.Tracer(adapters.ServiceNameDiscount)
	newCtx, span := tr.Start(ctx, spanDefaultDiscount+name)
	return tr, newCtx, span
}

func (s *discountService) Get(ctx context.Context, id uuid.UUID) (*domain.DiscountGet, error) {
	_, newCtx, span := getDiscountTracerSpan(ctx, ".GetByLogin")
	defer span.End()

	discount, err := s.r.Get(newCtx, id)
	if err != nil {
		return nil, err
	}

	return discount, nil
}

func (s *discountService) List(ctx context.Context, filter *domain.DiscountFilter) ([]domain.DiscountGet, error) {
	_, newCtx, span := getDiscountTracerSpan(ctx, ".List")
	defer span.End()

	discounts, err := s.r.List(newCtx, filter)
	if err != nil {
		return nil, err
	}

	return discounts, nil
}

func (s *discountService) EndDiscount(ctx context.Context, id uuid.UUID) error {
	_, newCtx, span := getDiscountTracerSpan(ctx, ".EndDiscount")
	defer span.End()

	if id == uuid.Nil {
		return errors.New("discount ID cannot be nil")
	}

	err := s.r.EndDiscount(newCtx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *discountService) Create(ctx context.Context, discount *domain.DiscountCreate) (*uuid.UUID, error) {
	_, newCtx, span := getDiscountTracerSpan(ctx, ".Create")
	defer span.End()

	id, err := s.r.Create(newCtx, discount)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (s *discountService) Update(ctx context.Context, discount *domain.DiscountUpdate) error {
	_, newCtx, span := getDiscountTracerSpan(ctx, ".Update")
	defer span.End()

	err := s.r.Update(newCtx, discount)
	if err != nil {
		return err
	}

	return nil
}
