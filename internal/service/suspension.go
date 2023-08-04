package service

import (
	"context"
	"one-lab-final/internal/entity"
)

func (m *Manager) NewSuspension(ctx context.Context, suspension *entity.Suspension) error {
	return m.Repository.NewSuspension(ctx, suspension)
}

func (m *Manager) UpdateSuspension(ctx context.Context, suspension *entity.Suspension) error {
	return m.Repository.UpdateSuspension(ctx, suspension)
}

func (m *Manager) CheckSuspension(ctx context.Context, userID int64) ([]*entity.Suspension, error) {
	return m.Repository.CheckSuspension(ctx, userID)
}
