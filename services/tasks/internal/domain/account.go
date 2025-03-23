package domain

import (
	"github.com/google/uuid"
	"time"
)

type AccountGet struct {
	UUID      uuid.UUID
	Login     string
	CreatedAt time.Time
	UpdateAt  time.Time
}

type AccountCreate struct {
	Login    string
	Password string
}

type AccountUpdate struct {
	UUID     uuid.UUID
	Login    string
	Password string
}

type AccountFilter struct {
	UUID  *uuid.UUID
	Login *string
}

type AccountAuth struct {
	Login    string
	Password string
}
