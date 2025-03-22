package task

import (
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/domain"
)

func CreateToDomain(create TaskCreate) domain.TaskCreate {
	return domain.TaskCreate{
		CreatedBy: create.CreatedBy,
		Percent:   create.Percent,
		StartedAt: create.StartedAt,
		EndedAt:   create.EndedAt,
		Status:    create.Status,
	}
}

func DomainToGet(get domain.TaskGet) TaskGet {
	return TaskGet{
		UUID:         get.UUID,
		CreatedBy:    get.CreatedBy,
		Percent:      get.Percent,
		StartedAt:    get.StartedAt,
		EndedAt:      get.EndedAt,
		Status:       get.Status,
		CreatedAt:    get.CreatedAt,
		LastUpdateAt: get.UpdateAt,
	}
}

func UpdateToDomain(update TaskUpdate) domain.TaskUpdate {
	return domain.TaskUpdate{
		UUID:      update.UUID,
		CreatedBy: update.CreatedBy,
		Percent:   update.Percent,
		StartedAt: update.StartedAt,
		EndedAt:   update.EndedAt,
		Status:    update.Status,
	}
}

func FilterToDomain(filter *TaskFilter) *domain.TaskFilter {
	if filter == nil {
		return nil
	}
	return &domain.TaskFilter{
		UUID:      filter.UUID,
		CreatedBy: filter.CreatedBy,
		Percent:   filter.Percent,
		Status:    filter.Status,
	}
}
