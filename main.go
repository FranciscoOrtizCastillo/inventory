package main

import (
	"context"

	"github.com/FranciscoOrtizCastillo/inventory/database"
	"github.com/FranciscoOrtizCastillo/inventory/internal/repository"
	"github.com/FranciscoOrtizCastillo/inventory/internal/service"
	"github.com/FranciscoOrtizCastillo/inventory/settings"
	"go.uber.org/fx"
)

func main() {

	//Inyeccion de dependencias con fx
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
			repository.New,
			service.New,
		),
		fx.Invoke(
			/*func(s *settings.Settings) {
				log.Println(s)
			},
			func(db *sqlx.DB) {
				_, err := db.Query("SELECT * FROM USERS")
				if err != nil {
					panic(err)
				}
			},*/
			func(ctx context.Context, serv service.Service) {
				//TODO
				err := serv.RegisterUser(ctx, "test@email.com", "TestUser", "password123")
				if err != nil {
					panic(err)
				}

				u, err := serv.LoginUser(ctx, "test@email.com", "password123")
				if err != nil {
					panic(err)
				}

				if u.Name != "TestUser" {
					panic("Invalid user")
				}
			},
		),
	)

	app.Run()
}
