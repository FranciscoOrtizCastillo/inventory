package repository

import (
	"context"

	"github.com/FranciscoOrtizCastillo/inventory/internal/entity"
)

const (
	queryInsertUser = `
		insert into USERS (email,name,password)
		values (?,?,?);`

	queryGetUserByEmail = `
		select id,email,name,password
		from USERS
		where email = ?;`

	queryInsertUserRoles = `
		insert into USER_ROLES (user_id,role_id) values (:user_id, :role_id);`

	queryRemoveUserRoles = `
		delete from USER_ROLES where user_id = :user_id and role_id = :role_id`

	queryGetUserRoles = `
		select user_id,role_id from USER_ROLES where user_id = ?;`
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

func (r *repo) SaveUserRole(ctx context.Context, userID, roleID int64) error {

	data := entity.UserRole{
		UserID: userID,
		RoleID: roleID,
	}

	_, err := r.db.NamedExecContext(ctx, queryInsertUserRoles, data)

	return err
}

func (r *repo) RemoveUserRole(ctx context.Context, userID, roleID int64) error {

	data := entity.UserRole{
		UserID: userID,
		RoleID: roleID,
	}

	_, err := r.db.NamedExecContext(ctx, queryRemoveUserRoles, data)

	return err
}

func (r *repo) GetUserRoles(ctx context.Context, userID int64) ([]entity.UserRole, error) {

	var roles []entity.UserRole

	err := r.db.SelectContext(ctx, &roles, queryGetUserRoles, userID)

	if err != nil {
		return nil, err
	}

	return roles, err
}
