package postgre

import (
	"context"
	"github.com/ArtemFed/hse-wishlist/pkg/xcommon"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/domain"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/repository/postgre/repo_models"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/service/adapters"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	global "go.opentelemetry.io/otel"
	"time"
)

const (
	spanDefaultAccount = "account/repository.postgre"
)

const (
	baseAccountGetQuery = `SELECT uuid, login, created_at, updated_at FROM accounts`
	createAccountQuery  = `INSERT INTO accounts (login, password) VALUES ($1, $2) RETURNING uuid`
	updateAccountQuery  = `UPDATE accounts SET login = $2, password = $3, updated_at = $4 WHERE uuid = $1`
)

var _ adapters.AccountRepository = (*accountRepository)(nil)

type accountRepository struct {
	db *sqlx.DB
}

func NewAccountRepository(db *sqlx.DB) adapters.AccountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) List(ctx context.Context, filter domain.AccountFilter) ([]domain.AccountGet, error) {
	tr := global.Tracer(adapters.ServiceAccount)
	_, span := tr.Start(ctx, spanDefaultAccount+".List")
	defer span.End()

	paramsMap := mapGetAccountRequestParams(filter)
	query, args := xcommon.QueryWhereAnd(baseAccountGetQuery, paramsMap)
	var accounts []repo_models.Account
	err := r.db.SelectContext(ctx, &accounts, query, args...)
	if err != nil {
		return nil, err
	}
	return xcommon.ConvertSliceP(accounts, repo_models.ToAccountDomain), nil
}

func (r *accountRepository) Create(ctx context.Context, account domain.AccountCreate) (*uuid.UUID, error) {
	tr := global.Tracer(adapters.ServiceAccount)
	_, span := tr.Start(ctx, spanDefaultAccount+".Create")
	defer span.End()

	accountPostgres := repo_models.CreateToAccountPostgres(account)
	row := r.db.QueryRowx(
		createAccountQuery,
		accountPostgres.Login,
		accountPostgres.Password,
	)

	var id uuid.UUID
	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func (r *accountRepository) Update(ctx context.Context, account domain.AccountUpdate) error {
	tr := global.Tracer(adapters.ServiceAccount)
	_, span := tr.Start(ctx, spanDefaultAccount+".Update")
	defer span.End()

	accountPostgres := repo_models.UpdateToAccountPostgres(account)
	_, err := r.db.ExecContext(
		ctx,
		updateAccountQuery,
		accountPostgres.UUID,
		accountPostgres.Login,
		accountPostgres.Password,
		time.Now(),
	)
	return err
}

func mapGetAccountRequestParams(params domain.AccountFilter) map[string]xcommon.FilterItem {
	paramsMap := make(map[string]xcommon.FilterItem)
	if params.UUID != nil {
		paramsMap["uuid"] = xcommon.FilterItem{Value: *params.UUID, Operator: "="}
	}
	if params.Login != nil {
		paramsMap["login"] = xcommon.FilterItem{Value: *params.Login, Operator: "="}
	}
	return paramsMap
}
