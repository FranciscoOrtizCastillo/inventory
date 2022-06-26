package service

import (
	"context"
	"errors"

	"github.com/FranciscoOrtizCastillo/inventory/encryption"
	"github.com/FranciscoOrtizCastillo/inventory/internal/models"
)

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

func (s *serv) RegisterUser(ctx context.Context, email, name, password string) error {

	user, _ := s.repo.GetUserByEmail(ctx, email)

	if user != nil {
		return ErrUserAlreadyExists
	}

	//TODO hash password
	bb, err := encryption.Encrypt([]byte(password))

	if err != nil {
		return err
	}

	encryptedPassword := encryption.ToBase64(bb)

	return s.repo.SaveUser(ctx, email, name, encryptedPassword)
}

func (s *serv) LoginUser(ctx context.Context, email, password string) (*models.User, error) {

	user, err := s.repo.GetUserByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	//TODO decrypt password
	bb, err := encryption.FromBase64(user.Password)

	if err != nil {
		return nil, err
	}

	decryptedPassword, err := encryption.Decrypt(bb)

	if err != nil {
		return nil, err
	}

	if string(decryptedPassword) != password {
		return nil, ErrInvalidCredentials
	}

	return &models.User{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}
