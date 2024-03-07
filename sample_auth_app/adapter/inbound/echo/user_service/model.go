package user_service

type SignUpInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type SignInInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type SessionInfo struct {
	Sid    string `json:"sid"`
	Stoken string `json:"stoken"`
}
