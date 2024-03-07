package user_service

import (
	"example_app/sample_auth_app/common/mycontext"
	"example_app/sample_auth_app/usecase/user_authentication"

	"github.com/labstack/echo/v4"
)

type service struct {
	newUserAuthenticationUsecase func(ctx mycontext.Context) user_authentication.Usecase
}

func New(newUserAuthenticationUsecase func(ctx mycontext.Context) user_authentication.Usecase) service {
	return service{
		newUserAuthenticationUsecase: newUserAuthenticationUsecase,
	}
}

func (s service) SignUp(c echo.Context) error {
	// uu := s.newUserAuthenticationUsecase(mycontext.From(c.Request().Context()))

	// qp := c.QueryParams()
	// var name, password string
	return nil
}
