package service

import (
	"context"
	"one-lab-final/internal/entity"
	"one-lab-final/pkg/util"
)

func (m *Manager) CreateUser(ctx context.Context, u *entity.User) error {
	hash, err := util.HashPassword(*u.Password.Plaintext)
	if err != nil {
		return err
	}

	u.Password.Hash = &hash

	return m.Repository.CreateUser(ctx, u)
}

func (m *Manager) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	return m.Repository.GetUserByUsername(ctx, username)
}

func (m *Manager) GetUserByCredentials(ctx context.Context, credentials string) (*entity.User, error) {
	return m.Repository.GetUserByCredentials(ctx, credentials)
}

func (m *Manager) GetUserByToken(ctx context.Context, token string) (*entity.User, error) {
	return m.Repository.GetUserByToken(ctx, token)
}

func (m *Manager) UpdateUser(ctx context.Context, u *entity.User) error {
	if u.Password.Plaintext != nil {
		hash, err := util.HashPassword(*u.Password.Plaintext)
		if err != nil {
			return err
		}

		u.Password.Hash = &hash
	}

	return m.Repository.UpdateUser(ctx, u)
}

func (m *Manager) DeleteUser(ctx context.Context, userID int64) error {
	return m.Repository.DeleteUser(ctx, userID)
}

func (m *Manager) GrantRoleToUser(ctx context.Context, userID int64, role entity.Role) error {
	return m.Repository.GrantRoleToUser(ctx, userID, role)
}
