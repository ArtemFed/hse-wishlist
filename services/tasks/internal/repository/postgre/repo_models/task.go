package repo_models

import (
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/domain"
	"github.com/google/uuid"
	"time"
)

const InitStatus = "Created"

type Task struct {
	UUID      uuid.UUID `db:"uuid"`
	Name      string    `db:"name"`
	Text      string    `db:"text"`
	Status    string    `db:"status"`
	CreatedBy uuid.UUID `db:"created_by"`
	StartedAt time.Time `db:"started_at"`
	EndedAt   time.Time `db:"ended_at"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func CreateToTaskPostgres(model domain.TaskCreate) *Task {
	id, _ := uuid.NewUUID()
	return &Task{
		UUID:      id,
		Name:      model.Name,
		Text:      model.Text,
		Status:    InitStatus,
		CreatedBy: model.CreatedBy,
		StartedAt: model.StartedAt,
		EndedAt:   model.EndedAt,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func UpdateToTaskPostgres(model domain.TaskUpdate) *Task {
	return &Task{
		UUID:      model.UUID,
		Name:      model.Name,
		Text:      model.Text,
		Status:    model.Status,
		CreatedBy: model.CreatedBy,
		StartedAt: model.StartedAt,
		EndedAt:   model.EndedAt,
	}
}

func ToTaskDomain(model Task) domain.TaskGet {
	return domain.TaskGet{
		UUID:      model.UUID,
		Name:      model.Name,
		Text:      model.Text,
		Status:    model.Status,
		CreatedBy: model.CreatedBy,
		StartedAt: model.StartedAt,
		EndedAt:   model.EndedAt,
		CreatedAt: model.CreatedAt,
		UpdateAt:  model.UpdatedAt,
	}
}
