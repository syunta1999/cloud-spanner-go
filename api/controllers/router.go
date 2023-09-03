package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

// routes sets up the routes for the API endpoints
func Routes(e *echo.Echo, c *UserController) {
	// ユーザーの一覧取得
	e.GET("/users", func(ctx echo.Context) error {
		err := c.GetUsers()
		if err != nil {
			return ctx.JSON(
				http.StatusInternalServerError,
				map[string]string{
					"error": err.Error(),
				},
			)
		}

		return ctx.JSON(
			http.StatusOK,
			map[string]string{
				"message": "Successfully fetched users",
			},
		)
	})

	// ユーザーの作成
	e.POST("/createusers", func(ctx echo.Context) error {
		err := c.CreateUsers()
		if err != nil {
			return ctx.JSON(
				http.StatusInternalServerError,
				map[string]string{
					"error": err.Error(),
				},
			)
		}

		return ctx.JSON(
			http.StatusCreated,
			map[string]string{
				"message": "Successfully created user",
			},
		)
	})
}
