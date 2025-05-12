package model

type LoginType string

const (
	GithubLogin LoginType = "github"
	EmailLogin  LoginType = "email"
)

type AuthReq interface {
	GetType() LoginType
}

type AuthRes struct {
	Token        string     `json:"token"`
	RefreshToken string     `json:"refresh_token"`
	Payload      JwtPayload `json:"payload"`
}

type EmailAuthReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type GithubAuthReq struct {
	Code string `form:"code" binding:"required"`
}

func (r *EmailAuthReq) GetType() LoginType {
	return EmailLogin
}

func (r *GithubAuthReq) GetType() LoginType {
	return GithubLogin
}
