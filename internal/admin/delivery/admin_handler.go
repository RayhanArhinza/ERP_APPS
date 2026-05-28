package delivery

import (
	"ERP_APPS/internal/admin"
	"ERP_APPS/internal/admin/dto"
	"ERP_APPS/internal/admin/usecase"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminUsecase usecase.AdminUsecase
}

func (h *AdminHandler) GetAll(c *gin.Context) {
	admins, err := h.adminUsecase.GetAll(
		c.Request.Context(),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(
		http.StatusOK,
		dto.ToAdminResponses(admins),
	)
}

func (h *AdminHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}
	adminData, err := h.adminUsecase.GetByID(
		c.Request.Context(),
		uint(id),
	)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(
		http.StatusOK,
		dto.ToAdminResponse(*adminData),
	)
}

func (h *AdminHandler) Create(c *gin.Context) {
	var adminRequest dto.CreateAdminRequest
	if err := c.ShouldBindJSON(
		&adminRequest,
	); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	data := &admin.Admin{
		Name:      adminRequest.Name,
		Email:     adminRequest.Email,
		Password:  adminRequest.Password,
		RoleID:    adminRequest.RoleID,
		TglLahir:  adminRequest.TglLahir,
		Alamat:    adminRequest.Alamat,
	}
	res, err := h.adminUsecase.Create(
		c.Request.Context(),
		data,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(
		http.StatusCreated,
		dto.ToAdminResponse(*res),
	)
}

func (h *AdminHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}
	var adminRequest dto.UpdateAdminRequest
	if err := c.ShouldBindJSON(
		&adminRequest,
	); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	data := &admin.Admin{
		ID:         uint(id),
		Name:       adminRequest.Name,
		Email:      adminRequest.Email,
		Password:   adminRequest.Password,
		RoleID:     adminRequest.RoleID,
		TglLahir:   adminRequest.TglLahir,
		Alamat:     adminRequest.Alamat,
	}
	res, err := h.adminUsecase.Update(
		c.Request.Context(),
		data,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(
		http.StatusOK,
		dto.ToAdminResponse(*res),
	)
}

func (h *AdminHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}
	err = h.adminUsecase.Delete(
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
		"message": "admin deleted",
	})
}