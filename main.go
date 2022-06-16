package main

import (
	"context"

	"github.com/FranciscoOrtizCastillo/inventory/database"
	"github.com/FranciscoOrtizCastillo/inventory/settings"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

func main() {

	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
		),
		fx.Invoke(
			/*func(s *settings.Settings) {
				log.Println(s)
			},*/
			func(db *sqlx.DB) {
				_, err := db.Query("SELECT * FROM USERS")
				if err != nil {
					panic(err)
				}
			},
		),
	)

	app.Run()
}
