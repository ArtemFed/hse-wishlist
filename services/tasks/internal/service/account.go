package service

import (
	"context"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/domain"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/service/adapters"
	"github.com/google/uuid"
)

const (
	spanDefaultAccount = "account/service."
)

var _ adapters.AccountService = &accountService{}

type accountService struct {
	r adapters.AccountRepository
}

func NewAccountService(accountRepository adapters.AccountRepository) adapters.AccountService {
	return &accountService{r: accountRepository}
}

func (s *accountService) List(ctx context.Context, filter domain.AccountFilter) ([]domain.AccountGet, error) {
	_, newCtx, span := domain.GetTracerSpan(ctx, adapters.ServiceAccount, spanDefaultAccount, ".List")
	defer span.End()

	accounts, err := s.r.List(newCtx, filter)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (s *accountService) Create(ctx context.Context, account domain.AccountCreate) (*uuid.UUID, error) {
	_, newCtx, span := domain.GetTracerSpan(ctx, adapters.ServiceAccount, spanDefaultAccount, ".Create")
	defer span.End()

	id, err := s.r.Create(newCtx, account)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (s *accountService) Update(ctx context.Context, account domain.AccountUpdate) error {
	_, newCtx, span := domain.GetTracerSpan(ctx, adapters.ServiceAccount, spanDefaultAccount, ".Update")
	defer span.End()

	err := s.r.Update(newCtx, account)
	if err != nil {
		return err
	}

	return nil
}
