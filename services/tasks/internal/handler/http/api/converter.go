package api

import (
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/domain"
)

func TaskCreateToDomain(create TaskCreate) domain.TaskCreate {
	return domain.TaskCreate{
		Name:      create.Name,
		Text:      create.Text,
		CreatedBy: create.CreatedBy,
		StartedAt: create.StartedAt,
		EndedAt:   create.EndedAt,
	}
}

func TaskDomainToGet(get domain.TaskGet) TaskGet {
	return TaskGet{
		Id:        get.UUID,
		Name:      get.Name,
		Text:      get.Text,
		Status:    get.Status,
		CreatedBy: get.CreatedBy,
		StartedAt: get.StartedAt,
		EndedAt:   get.EndedAt,
		CreatedAt: get.CreatedAt,
		UpdateAt:  get.UpdateAt,
	}
}

func TaskUpdateToDomain(update TaskUpdate) domain.TaskUpdate {
	return domain.TaskUpdate{
		Name:      update.Name,
		Text:      update.Text,
		Status:    update.Status,
		CreatedBy: update.CreatedBy,
		StartedAt: update.StartedAt,
		EndedAt:   update.EndedAt,
	}
}

func TaskFilterToDomain(params GetTasksParams) domain.TaskFilter {
	return domain.TaskFilter{
		UUID:                params.Id,
		CreatedBy:           params.CreatedBy,
		Status:              params.Status,
		StartedAtLeftBound:  params.StartedAtLB,
		StartedAtRightBound: params.StartedAtRB,
		EndedAtLeftBound:    params.EndedAtLB,
		EndedAtRightBound:   params.EndedAtRB,
	}
}

func AccountCreateToDomain(create AccountCreate) domain.AccountCreate {
	return domain.AccountCreate{
		Login:    create.Login,
		Password: create.Password,
	}
}

func AccountDomainToGet(get domain.AccountGet) AccountGet {
	return AccountGet{
		Id:        get.UUID,
		CreatedAt: get.CreatedAt,
	}
}

func AccountUpdateToDomain(update AccountUpdate) domain.AccountUpdate {
	return domain.AccountUpdate{
		Login:    update.Login,
		Password: update.Password,
	}
}

func AccountFilterToDomain(params GetAccountsParams) domain.AccountFilter {
	return domain.AccountFilter{
		UUID: params.Id,
	}
}

func AccountAuthToDomain(model AccountAuth) domain.AccountAuth {
	return domain.AccountAuth{
		Login:    model.Login,
		Password: model.Password,
	}
}
