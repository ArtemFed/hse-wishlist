package repo_models

import (
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/domain"
	"github.com/google/uuid"
	"time"
)

type Task struct {
	UUID          uuid.UUID `db:"uuid"`
	CreatedBy     uuid.UUID `db:"created_by"`
	Percent       float32   `db:"percent"`
	StartedAt     time.Time `db:"started_at"`
	EndedAt       time.Time `db:"ended_at"`
	Status        string    `db:"status"`
	CreatedAt     time.Time `db:"created_at"`
	LastUpdatedAt time.Time `db:"updated_at"`
}

func CreateToTaskPostgres(model *domain.TaskCreate) *Task {
	id, _ := uuid.NewUUID()
	return &Task{
		UUID:      id,
		CreatedBy: model.CreatedBy,
		Percent:   model.Percent,
		StartedAt: model.StartedAt,
		EndedAt:   model.EndedAt,
		Status:    model.Status,
		CreatedAt: time.Now(),
	}
}

func UpdateToTaskPostgres(model *domain.TaskUpdate) *Task {
	return &Task{
		UUID:      model.UUID,
		CreatedBy: model.CreatedBy,
		Percent:   model.Percent,
		StartedAt: model.StartedAt,
		EndedAt:   model.EndedAt,
		Status:    model.Status,
	}
}

func ToTaskDomain(model *Task) *domain.TaskGet {
	return &domain.TaskGet{
		UUID:         model.UUID,
		CreatedBy:    model.CreatedBy,
		Percent:      model.Percent,
		StartedAt:    model.StartedAt,
		EndedAt:      model.EndedAt,
		Status:       model.Status,
		CreatedAt:    model.CreatedAt,
		LastUpdateAt: model.LastUpdatedAt,
	}
}
