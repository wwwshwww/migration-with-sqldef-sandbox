package user_service

import (
	"example_app/sample_auth_app/common/mycontext"
	"example_app/sample_auth_app/usecase/user_authentication"
	"net/http"

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

// @Summary CreateNewUser
// @Description you don't need session's id and token when request
// @Tags user
// @Accept  json
// @Produce  json
// @Success 204 "" ""
// @Router /user/signup [post]
func (s service) SignUp(c echo.Context) error {
	uu := s.newUserAuthenticationUsecase(mycontext.From(c.Request().Context()))

	var signUpInput SignUpInput
	if err := c.Bind(&signUpInput); err != nil {
		return err
		// return c.String(http.StatusBadRequest, "invalid param")
	}

	if err := uu.SignUp(UnmarshalSignUpInput(signUpInput)); err != nil {
		return err
	}
	return c.String(http.StatusNoContent, "")
}
