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
	spanDefaultTask = "task/repository.postgre"
)

const (
	baseTaskGetQuery = `
		SELECT uuid, created_by, percent, started_at, ended_at, status, created_at, updated_at
		FROM tasks
	`
	createTaskQuery = `
		INSERT INTO tasks (created_by, percent, started_at, ended_at, status)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING uuid
	`
	updateTaskQuery = `
		UPDATE tasks
		SET 
		    created_by = $2,
		    percent = $3,
		    started_at = $4,
		    ended_at = $5,
		    status = $6
		WHERE uuid = $1;
	`
	finishTaskQuery = `
		UPDATE tasks
		SET 
		    ended_at = $2
		WHERE uuid = $1;
	`
)

var _ adapters.TaskRepository = &taskRepository{}

type taskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) adapters.TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Get(ctx context.Context, id uuid.UUID) (*domain.TaskGet, error) {
	tr := global.Tracer(adapters.ServiceNameTask)
	_, span := tr.Start(ctx, spanDefaultTask+".GetByLogin")
	defer span.End()

	tasks, err := r.List(ctx, &domain.TaskFilter{UUID: &id})
	if err != nil {
		return nil, err
	}
	return xcommon.EnsureSingle(tasks)
}

func (r *taskRepository) List(ctx context.Context, filter *domain.TaskFilter) ([]domain.TaskGet, error) {
	tr := global.Tracer(adapters.ServiceNameTask)
	_, span := tr.Start(ctx, spanDefaultTask+".List")
	defer span.End()

	paramsMap := mapGetTaskRequestParams(filter)
	query, args := xcommon.QueryWhereAnd(baseTaskGetQuery, paramsMap)
	var tasks []repo_models.Task
	err := r.db.SelectContext(ctx, &tasks, query, args...)
	if err != nil {
		return nil, err
	}
	return xcommon.ConvertSliceP(tasks, repo_models.ToTaskDomain), nil
}

func (r *taskRepository) Create(ctx context.Context, task *domain.TaskCreate) (*uuid.UUID, error) {
	tr := global.Tracer(adapters.ServiceNameTask)
	_, span := tr.Start(ctx, spanDefaultTask+".Create")
	defer span.End()

	taskPostgres := repo_models.CreateToTaskPostgres(task)
	row := r.db.QueryRow(
		createTaskQuery,
		taskPostgres.CreatedBy,
		taskPostgres.Percent,
		taskPostgres.StartedAt,
		taskPostgres.EndedAt,
		taskPostgres.Status,
	)

	var id uuid.UUID
	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}
	return &id, err
}

func (r *taskRepository) Update(ctx context.Context, task *domain.TaskUpdate) error {
	taskPostgres := repo_models.UpdateToTaskPostgres(task)
	_, err := r.db.ExecContext(
		ctx,
		updateTaskQuery,
		taskPostgres.UUID,
		taskPostgres.CreatedBy,
		taskPostgres.Percent,
		taskPostgres.StartedAt,
		taskPostgres.EndedAt,
		taskPostgres.Status,
	)
	return err
}

func (r *taskRepository) EndTask(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.ExecContext(
		ctx,
		finishTaskQuery,
		id,
		time.Now(),
	)
	return err
}

func mapGetTaskRequestParams(params *domain.TaskFilter) map[string]interface{} {
	if params == nil {
		return map[string]any{}
	}
	paramsMap := make(map[string]interface{})
	if params.UUID != nil {
		paramsMap["uuid"] = *params.UUID
	}
	if params.CreatedBy != nil {
		paramsMap["created_by"] = *params.CreatedBy
	}
	if params.Percent != nil {
		paramsMap["percent"] = *params.Percent
	}
	if params.Status != nil {
		paramsMap["status"] = *params.Status
	}
	return paramsMap
}
