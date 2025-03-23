package service

import (
	"context"
	"errors"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/domain"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/service/adapters"
	"github.com/google/uuid"
)

const (
	spanDefaultTask = "task/service."
)

var _ adapters.TaskService = &taskService{}

type taskService struct {
	r adapters.TaskRepository
}

func NewTaskService(taskRepository adapters.TaskRepository) adapters.TaskService {
	return &taskService{r: taskRepository}
}

func (s *taskService) List(ctx context.Context, filter domain.TaskFilter) ([]domain.TaskGet, error) {
	_, newCtx, span := domain.GetTracerSpan(ctx, adapters.ServiceTask, spanDefaultTask, ".List")
	defer span.End()

	tasks, err := s.r.List(newCtx, filter)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *taskService) Patch(ctx context.Context, id uuid.UUID, status string) error {
	_, newCtx, span := domain.GetTracerSpan(ctx, adapters.ServiceTask, spanDefaultTask, ".EndTask")
	defer span.End()

	if id == uuid.Nil {
		return errors.New("task ID cannot be nil")
	}

	err := s.r.PatchStatus(newCtx, id, status)
	if err != nil {
		return err
	}

	return nil
}

func (s *taskService) Create(ctx context.Context, task domain.TaskCreate) (*uuid.UUID, error) {
	_, newCtx, span := domain.GetTracerSpan(ctx, adapters.ServiceTask, spanDefaultTask, ".Create")
	defer span.End()

	id, err := s.r.Create(newCtx, task)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (s *taskService) Update(ctx context.Context, task domain.TaskUpdate) error {
	_, newCtx, span := domain.GetTracerSpan(ctx, adapters.ServiceTask, spanDefaultTask, ".Update")
	defer span.End()

	err := s.r.Update(newCtx, task)
	if err != nil {
		return err
	}

	return nil
}
