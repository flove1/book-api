package service

import (
	"context"
	"one-lab-final/internal/entity"
	"one-lab-final/pkg/util"
	"time"
)

func (m *Manager) Login(ctx context.Context, credentials string, password string) (*entity.Token, error) {
	user, err := m.Repository.GetUserByCredentials(ctx, credentials)
	if err != nil {
		return nil, err
	}

	err = util.CheckPassword(password, *user.Password.Hash)
	if err != nil {
		return nil, err
	}

	token, err := util.GenerateToken()
	if err != nil {
		return nil, err
	}

	tokenEntity := &entity.Token{
		Plaintext: token.Plaintext,
		Hash:      token.Hash,
		UserID:    user.ID,
		Expiry:    time.Now().Add(m.Config.AUTH.TokenExpiration),
	}

	err = m.Repository.CreateToken(ctx, tokenEntity)
	if err != nil {
		return nil, err
	}

	return tokenEntity, nil
}

func (m *Manager) DeleteExpiredTokens(ctx context.Context) error {
	err := m.Repository.DeleteExpiredTokens(ctx)
	if err != nil {
		return err
	}

	return nil
}
