package repository

import (
	"context"

	"github.com/FranciscoOrtizCastillo/inventory/internal/entity"
)

const (
	queryInsertUser = `
		insert into USERS (email,name,password)
		values (?,?,?);
		`

	queryGetUserByEmail = `
		select id,email,name,password
		from USERS
		where email = ?;
		`
)

func (r *repo) SaveUser(ctx context.Context, email, name, password string) error {

	_, err := r.db.ExecContext(ctx, queryInsertUser, email, name, password)
	return err
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {

	user := &entity.User{}

	err := r.db.GetContext(ctx, user, queryGetUserByEmail, email)

	if err != nil {
		return nil, err
	}

	return user, err
}
