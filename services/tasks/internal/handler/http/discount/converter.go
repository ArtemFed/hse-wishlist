package discount

import (
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/domain"
)

func CreateToDomain(create DiscountCreate) domain.DiscountCreate {
	return domain.DiscountCreate{
		CreatedBy: create.CreatedBy,
		Percent:   create.Percent,
		StartedAt: create.StartedAt,
		EndedAt:   create.EndedAt,
		Status:    create.Status,
	}
}

func DomainToGet(get domain.DiscountGet) DiscountGet {
	return DiscountGet{
		UUID:         get.UUID,
		CreatedBy:    get.CreatedBy,
		Percent:      get.Percent,
		StartedAt:    get.StartedAt,
		EndedAt:      get.EndedAt,
		Status:       get.Status,
		CreatedAt:    get.CreatedAt,
		LastUpdateAt: get.LastUpdateAt,
	}
}

func UpdateToDomain(update DiscountUpdate) domain.DiscountUpdate {
	return domain.DiscountUpdate{
		UUID:      update.UUID,
		CreatedBy: update.CreatedBy,
		Percent:   update.Percent,
		StartedAt: update.StartedAt,
		EndedAt:   update.EndedAt,
		Status:    update.Status,
	}
}

func FilterToDomain(filter *DiscountFilter) *domain.DiscountFilter {
	if filter == nil {
		return nil
	}
	return &domain.DiscountFilter{
		UUID:      filter.UUID,
		CreatedBy: filter.CreatedBy,
		Percent:   filter.Percent,
		Status:    filter.Status,
	}
}
