package controllers

import (
	"context"
	"log"

	"github.com/labstack/echo"
	"go.uber.org/fx"
)

func NewController(lc fx.Lifecycle, controller *UserController) *echo.Echo {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			e := echo.New()
			// rest api calls to echo
			Routes(e, controller)

			go func() {
				if err := e.Start(":8888"); err != nil {
					log.Fatalf("Echo server failed: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil // Cleanup logic here.
		},
	})
	return nil
}
