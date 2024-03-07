package user_service

import "example_app/sample_auth_app/usecase/user_authentication"

func UnmarshalSignInInput(j SignInInput) user_authentication.SignInInput {
	return user_authentication.SignInInput{
		Name:     j.Name,
		Password: j.Password,
	}
}

func UnmarshalSignUpInput(j SignUpInput) user_authentication.SignUpInput {
	return user_authentication.SignUpInput{
		Name:     j.Name,
		Password: j.Password,
	}
}

func MarshalSessionInfo(s user_authentication.SessionInfo) SessionInfo {
	return SessionInfo{
		Sid:    s.Id,
		Stoken: s.Token,
	}
}
