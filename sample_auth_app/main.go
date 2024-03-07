package main

import (
	"example_app/sample_auth_app/adapter/inbound/echo/user_service"
	"example_app/sample_auth_app/common/mycontext"
	"example_app/sample_auth_app/domain/session/session"
	"example_app/sample_auth_app/domain/user/user"
	"example_app/sample_auth_app/domain/user/user_finder"
	"example_app/sample_auth_app/domain_service/secure_hasher"
	"example_app/sample_auth_app/domain_service/user_validation_service"
	"example_app/sample_auth_app/usecase/user_authentication"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Sample App API
// @version 1.0
// @description This is a sample server for Echo framework.
// @host localhost:1323
// @BasePath /api/v1
func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello.\n")
	})

	userService := user_service.New(func(ctx mycontext.Context) user_authentication.Usecase {
		return user_authentication.New(
			ctx,
			&user.RepositoryMock{},
			&user_finder.FinderMock{},
			&session.RepositoryMock{},
			user_validation_service.New(
				&user_validation_service.PortMock{},
			),
			secure_hasher.New(),
		)
	})

	e.POST("/user/signup", userService.SignUp)

	e.Logger.Fatal(e.Start(":1323"))
}
