package main

import (
	"github.com/FranciscoOrtizCastillo/inventory/settings"
	"go.uber.org/fx"
)

func main() {

	app := fx.New(
		fx.Provide(
			settings.New,
		),
		fx.Invoke(
		/*func(s *settings.Settings) {
			log.Println(s)
		},*/
		),
	)

	app.Run()
}
