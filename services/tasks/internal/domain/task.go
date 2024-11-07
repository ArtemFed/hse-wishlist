package domain

import (
	"github.com/google/uuid"
	"time"
)

type TaskGet struct {
	UUID         uuid.UUID
	CreatedBy    uuid.UUID
	Percent      float32
	StartedAt    time.Time
	EndedAt      time.Time
	Status       string
	CreatedAt    time.Time
	LastUpdateAt time.Time
}

type TaskCreate struct {
	CreatedBy uuid.UUID
	Percent   float32
	StartedAt time.Time
	EndedAt   time.Time
	Status    string
}

type TaskUpdate struct {
	UUID      uuid.UUID
	CreatedBy uuid.UUID
	Percent   float32
	StartedAt time.Time
	EndedAt   time.Time
	Status    string
}

type TaskFilter struct {
	UUID      *uuid.UUID
	CreatedBy *uuid.UUID
	Percent   *float32
	Status    *string
}
