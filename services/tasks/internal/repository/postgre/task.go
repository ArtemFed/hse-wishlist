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
)

const (
	spanDefaultTask = "task/repository.postgre"
)

const (
	baseTaskGetQuery = `
		SELECT uuid, name, text, status, created_by, started_at, ended_at, created_at, updated_at
		FROM tasks
	`
	createTaskQuery = `
		INSERT INTO tasks (name, text, status, started_at, ended_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING uuid
	`
	updateTaskQuery = `
		UPDATE tasks
		SET 
		    name = $2,
		    text = $3,
		    status = $4,
		    started_at = $5,
		    ended_at = $6,
		WHERE uuid = $1;
	`
	patchStatusTaskQuery = `
		UPDATE tasks
		SET 
		    status = $2,
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

func (r *taskRepository) List(ctx context.Context, filter *domain.TaskFilter) ([]domain.TaskGet, error) {
	tr := global.Tracer(adapters.ServiceTask)
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
	tr := global.Tracer(adapters.ServiceTask)
	_, span := tr.Start(ctx, spanDefaultTask+".Create")
	defer span.End()

	taskPostgres := repo_models.CreateToTaskPostgres(task)
	row := r.db.QueryRow(
		createTaskQuery,
		taskPostgres.Name,
		taskPostgres.Text,
		taskPostgres.Status,
		taskPostgres.CreatedBy,
		taskPostgres.StartedAt,
		taskPostgres.EndedAt,
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
		taskPostgres.Name,
		taskPostgres.Text,
		taskPostgres.Status,
		taskPostgres.CreatedBy,
		taskPostgres.StartedAt,
		taskPostgres.EndedAt,
	)
	return err
}

func (r *taskRepository) PatchStatus(ctx context.Context, id uuid.UUID, status string) error {
	_, err := r.db.ExecContext(
		ctx,
		patchStatusTaskQuery,
		id,
		status,
	)
	return err
}

func mapGetTaskRequestParams(params *domain.TaskFilter) map[string]xcommon.FilterItem {
	if params == nil {
		return map[string]xcommon.FilterItem{}
	}
	paramsMap := make(map[string]xcommon.FilterItem)
	if params.UUID != nil {
		paramsMap["uuid"] = xcommon.FilterItem{Value: *params.UUID, Operator: "="}
	}
	if params.CreatedBy != nil {
		paramsMap["created_by"] = xcommon.FilterItem{Value: *params.CreatedBy, Operator: "="}
	}
	if params.Status != nil {
		paramsMap["status"] = xcommon.FilterItem{Value: *params.Status, Operator: "="}
	}
	if params.StartedAtLeftBound != nil {
		paramsMap["started_at"] = xcommon.FilterItem{Value: *params.StartedAtLeftBound, Operator: ">"}
	}
	if params.StartedAtRightBound != nil {
		paramsMap["started_at"] = xcommon.FilterItem{Value: *params.StartedAtRightBound, Operator: "<"}
	}
	if params.EndedAtLeftBound != nil {
		paramsMap["ended_at"] = xcommon.FilterItem{Value: *params.EndedAtLeftBound, Operator: ">"}
	}
	if params.EndedAtRightBound != nil {
		paramsMap["ended_at"] = xcommon.FilterItem{Value: *params.EndedAtRightBound, Operator: "<"}
	}
	return paramsMap
}
