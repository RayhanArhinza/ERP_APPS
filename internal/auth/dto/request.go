package dto

type ChangePasswordRequest struct {
	Password string `json:"password"`
}
type ForgotPasswordRequest struct {
	Email string `json:"email"`
}
type ResetPasswordRequest struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}
type LogoutRequest struct {
	Token string `json:"token"`
}
type RefreshTokenRequest struct {
	Token string `json:"token"`
}
type LoginRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}
type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}