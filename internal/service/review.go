package service

import (
	"context"
	"one-lab-final/internal/entity"
	"one-lab-final/pkg/util"
)

func (m *Manager) CreateReview(ctx context.Context, review *entity.Review) error {
	err := m.Repository.CreateReview(ctx, review)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) GetReviewsByBookID(ctx context.Context, bookID int64, filter util.Filter) ([]*entity.Review, *util.Metadata, error) {
	filterSafeList := []string{"rating", "date", "created_at", "updated_at"}

	if !filter.ValidateSort(filterSafeList) {
		return nil, nil, ErrInvalidSortValue
	}

	reviews, meta, err := m.Repository.GetReviewsByBookID(ctx, bookID, filter)
	if err != nil {
		return nil, nil, err
	}

	return reviews, meta, nil
}

func (m *Manager) UpdateReview(ctx context.Context, review *entity.Review) error {
	err := m.Repository.UpdateReview(ctx, review)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) DeleteReview(ctx context.Context, reviewID int64, userID int64) error {
	err := m.Repository.DeleteReview(ctx, reviewID, userID)
	if err != nil {
		return err
	}

	return nil
}
