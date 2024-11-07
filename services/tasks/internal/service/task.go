package service

import (
	"context"
	"errors"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/domain"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/service/adapters"
	"github.com/google/uuid"
	global "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
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

func getTaskTracerSpan(ctx context.Context, name string) (trace.Tracer, context.Context, trace.Span) {
	tr := global.Tracer(adapters.ServiceNameTask)
	newCtx, span := tr.Start(ctx, spanDefaultTask+name)
	return tr, newCtx, span
}

func (s *taskService) Get(ctx context.Context, id uuid.UUID) (*domain.TaskGet, error) {
	_, newCtx, span := getTaskTracerSpan(ctx, ".GetByLogin")
	defer span.End()

	task, err := s.r.Get(newCtx, id)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *taskService) List(ctx context.Context, filter *domain.TaskFilter) ([]domain.TaskGet, error) {
	_, newCtx, span := getTaskTracerSpan(ctx, ".List")
	defer span.End()

	tasks, err := s.r.List(newCtx, filter)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *taskService) EndTask(ctx context.Context, id uuid.UUID) error {
	_, newCtx, span := getTaskTracerSpan(ctx, ".EndTask")
	defer span.End()

	if id == uuid.Nil {
		return errors.New("task ID cannot be nil")
	}

	err := s.r.EndTask(newCtx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *taskService) Create(ctx context.Context, task *domain.TaskCreate) (*uuid.UUID, error) {
	_, newCtx, span := getTaskTracerSpan(ctx, ".Create")
	defer span.End()

	id, err := s.r.Create(newCtx, task)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (s *taskService) Update(ctx context.Context, task *domain.TaskUpdate) error {
	_, newCtx, span := getTaskTracerSpan(ctx, ".Update")
	defer span.End()

	err := s.r.Update(newCtx, task)
	if err != nil {
		return err
	}

	return nil
}
