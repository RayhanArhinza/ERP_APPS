package delivery

import (
	admin "ERP_APPS/internal/admin"
	"ERP_APPS/internal/auth/dto"
	"ERP_APPS/internal/auth/usecase"
	jwtPkg "ERP_APPS/pkg/jwt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}



func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	adminData, err := h.authUsecase.Login(
		c.Request.Context(),
		req.Email,
		req.Password,
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	adminData.Password = ""
	token, err := jwtPkg.GenerateToken(adminData.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to generate token",
		})
		return
	}

	response := dto.ToLoginResponse(
		adminData.ID,
		adminData.Name,
		adminData.Email,
		adminData.RoleID,
		adminData.Role.Name,
		token,
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"data": response,
	})
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	adminData, err := h.authUsecase.Register(
		c.Request.Context(),
		&admin.Admin{
			Name:     req.Name,
			Email:    req.Email,
			Password: req.Password,
		},
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	adminData.Password = ""

	token, err := jwtPkg.GenerateToken(adminData.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to generate token",
		})
		return
	}

	response := dto.ToRegisterResponse(
		adminData.ID,
		adminData.Name,
		adminData.Email,
		adminData.RoleID,
		adminData.Role.Name,
		token,
	)

	c.JSON(http.StatusCreated, gin.H{
		"message": "registration successful",
		"data": response,
	})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	err = h.authUsecase.Logout(
		c.Request.Context(),
		uint(id),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "logout successful",
	})
}

func (h *AuthHandler) ChangePassword(c *gin.Context) {
	var req dto.ChangePasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	err = h.authUsecase.ChangePassword(
		c.Request.Context(),
		uint(id),
		req.Password,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "password changed successfully",
	})
}

func (h *AuthHandler) ForgotPassword(c *gin.Context) {
	var req dto.ForgotPasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	adminData, err := h.authUsecase.FindByEmail(
		c.Request.Context(),
		req.Email,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "email not found",
		})
		return
	}

	token := "reset-token"

	err = h.authUsecase.SaveResetToken(
		c.Request.Context(),
		adminData.ID,
		token,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "reset password token generated",
		"token":   token,
	})
}

func (h *AuthHandler) ResetPassword(c *gin.Context) {
	var req dto.ResetPasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.authUsecase.UpdatePasswordWithToken(
		c.Request.Context(),
		req.Token,
		req.Password,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "password reset successfully",
	})
}