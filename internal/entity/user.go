package entity

type User struct {
	ID       int64  `db:"id"` // json:"id"
	Email    string `db:"email"`
	Name     string `db:"name"`
	Password string `db:"password"` // json:"-" Ponemos - en json para que no se muestre
}
