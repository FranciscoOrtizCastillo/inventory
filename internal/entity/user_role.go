package entity

type UserRole struct {
	ID     int64 `db:"id"` // json:"id"
	UserID int64 `db:"user_id"`
	RoleID int64 `db:"role_id"`
}
