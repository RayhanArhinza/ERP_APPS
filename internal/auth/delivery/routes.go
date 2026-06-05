package delivery

import (
	"ERP_APPS/internal/auth/usecase"

	"github.com/gin-gonic/gin"
)
func NewAuthHandler(
	api *gin.RouterGroup,
	authUsecase usecase.AuthUsecase,
) {
	handler := &AuthHandler{
		authUsecase: authUsecase,
	}

	auth := api.Group("/auth")
	{
		auth.POST("/login", handler.Login)
		auth.POST("/register", handler.Register)
		auth.POST("/forgot-password", handler.ForgotPassword)
		auth.POST("/reset-password", handler.ResetPassword)

		auth.PUT("/change-password/:id", handler.ChangePassword)
		auth.POST("/logout/:id", handler.Logout)
	}
}