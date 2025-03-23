package domain

import (
	"github.com/google/uuid"
	"time"
)

type TaskGet struct {
	UUID      uuid.UUID
	Name      string
	Text      string
	Status    string
	CreatedBy uuid.UUID
	StartedAt time.Time
	EndedAt   time.Time
	CreatedAt time.Time
	UpdateAt  time.Time
}

type TaskCreate struct {
	Name      string
	Text      string
	CreatedBy uuid.UUID
	StartedAt time.Time
	EndedAt   time.Time
}

type TaskUpdate struct {
	UUID      uuid.UUID
	Name      string
	Text      string
	Status    string
	CreatedBy uuid.UUID
	StartedAt time.Time
	EndedAt   time.Time
}

type TaskFilter struct {
	UUID                *uuid.UUID
	CreatedBy           *uuid.UUID
	Status              *string
	StartedAtLeftBound  *time.Time
	StartedAtRightBound *time.Time
	EndedAtLeftBound    *time.Time
	EndedAtRightBound   *time.Time
}
