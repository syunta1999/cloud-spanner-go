package di

import (
	"go.uber.org/fx"

	"cloud-spanner-go/config"
	"cloud-spanner-go/controllers"
	"cloud-spanner-go/repositories"
	"cloud-spanner-go/usecases"
)

func BuildApp() *fx.App {
	return fx.New(
		fx.Provide(
			repositories.NewSpannerClient,
			usecases.NewUserInteractor,
			controllers.NewUserController,
			config.NewConfig,
		),
		fx.Invoke(controllers.NewController),
	)
}
