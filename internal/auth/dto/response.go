package dto

type ChangePasswordResponse struct {
	Message string `json:"message"`
}

type ForgotPasswordResponse struct {
	Message string `json:"message"`
}

type ResetPasswordResponse struct {
	Message string `json:"message"`
}

type LogoutResponse struct {
	Message string `json:"message"`
}

type RefreshTokenResponse struct {
	Token string `json:"token"`
}

type LoginResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	RoleID   uint   `json:"role_id"`
	RoleName string `json:"role_name"`
	Token    string `json:"token"`
}

type RegisterResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	RoleID   uint   `json:"role_id"`
	RoleName string `json:"role_name"`
	Token    string `json:"token"`
}

func ToLoginResponse(
	id uint,
	name string,
	email string,
	roleID uint,
	roleName string,
	token string,
) LoginResponse {
	return LoginResponse{
		ID:       id,
		Name:     name,
		Email:    email,
		RoleID:   roleID,
		RoleName: roleName,
		Token:    token,
	}
}

func ToRegisterResponse(
	id uint,
	name string,
	email string,
	roleID uint,
	roleName string,
	token string,
) RegisterResponse {
	return RegisterResponse{
		ID:       id,
		Name:     name,
		Email:    email,
		RoleID:   roleID,
		RoleName: roleName,
		Token:    token,
	}
}

func ToChangePasswordResponse(message string) ChangePasswordResponse {
	return ChangePasswordResponse{
		Message: message,
	}
}

func ToForgotPasswordResponse(message string) ForgotPasswordResponse {
	return ForgotPasswordResponse{
		Message: message,
	}
}

func ToResetPasswordResponse(message string) ResetPasswordResponse {
	return ResetPasswordResponse{
		Message: message,
	}
}

func ToLogoutResponse(message string) LogoutResponse {
	return LogoutResponse{
		Message: message,
	}
}

func ToRefreshTokenResponse(token string) RefreshTokenResponse {
	return RefreshTokenResponse{
		Token: token,
	}
}