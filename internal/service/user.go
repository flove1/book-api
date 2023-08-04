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

	err = m.Repository.CreateUser(ctx, u)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) GetUserByCredentials(ctx context.Context, username string) (*entity.User, error) {
	user, err := m.Repository.GetUserByCredentials(ctx, username)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (m *Manager) GetUserByToken(ctx context.Context, token string) (*entity.User, error) {
	user, err := m.Repository.GetUserByToken(ctx, token)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (m *Manager) UpdateUser(ctx context.Context, u *entity.User) error {
	if u.Password.Plaintext != nil {
		hash, err := util.HashPassword(*u.Password.Plaintext)
		if err != nil {
			return err
		}

		u.Password.Hash = &hash
	}

	err := m.Repository.UpdateUser(ctx, u)
	if err != nil {
		return err
	}
	return nil
}

func (m *Manager) DeleteUser(ctx context.Context, userID int64) error {
	err := m.Repository.DeleteUser(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}
