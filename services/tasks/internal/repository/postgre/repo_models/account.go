package repo_models

import (
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/domain"
	"github.com/google/uuid"
	"time"
)

type Account struct {
	UUID      uuid.UUID `db:"uuid"`
	Login     string    `db:"login"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func CreateToAccountPostgres(model domain.AccountCreate) *Account {
	id, _ := uuid.NewUUID()
	return &Account{
		UUID:      id,
		Login:     model.Login,
		Password:  model.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func UpdateToAccountPostgres(model domain.AccountUpdate) *Account {
	return &Account{
		UUID:     model.UUID,
		Login:    model.Login,
		Password: model.Password,
	}
}

func ToAccountDomain(model Account) domain.AccountGet {
	return domain.AccountGet{
		UUID:      model.UUID,
		Login:     model.Login,
		Password:  model.Password,
		CreatedAt: model.CreatedAt,
		UpdateAt:  model.UpdatedAt,
	}
}
