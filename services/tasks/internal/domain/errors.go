package domain

import (
	"errors"
)

var (
	ErrInternal     = errors.New("server internal error")
	ErrAccessDenied = errors.New("access denied")
)
