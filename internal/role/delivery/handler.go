package delivery

import (
	role "ERP_APPS/internal"
	"ERP_APPS/internal/role/dto"
	"ERP_APPS/internal/role/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
type RoleHandler struct {
	roleUsecase usecase.RoleUsecase
}

func (h *RoleHandler) GetAll(c *gin.Context){
	roles, err := h.roleUsecase.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, roles)
}
func (h *RoleHandler) GetByID(c *gin.Context){
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "invalid id",
		})
		return
		}
	roleData, err := h.roleUsecase.GetByID(c.Request.Context(), uint(id,))
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"error" : err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, dto.ToRoleResponse(*roleData))
}

func (h *RoleHandler) Create(c *gin.Context){
	var roleRequest dto.CreateRoleRequest
	if err := c.ShouldBindJSON(&roleRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
		data := &role.Role{
			Name : roleRequest.Name,
		
		}
		res, err := h.roleUsecase.Create(c.Request.Context(), data)
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
		})
			return
		}
		c.JSON(http.StatusOK, dto.ToRoleResponse(*res))
}
func (h *RoleHandler) Update(c *gin.Context){
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}
	var roleRequest dto.UpdateRoleRequest
	if err := c.ShouldBindJSON(&roleRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error(),
		})
	}
	data := &role.Role{
		ID : uint(id),
		Name : roleRequest.Name,
	}
	res, err := h.roleUsecase.Update(c.Request.Context(), data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, dto.ToRoleResponse(*res))
}	
func (h *RoleHandler) Delete(c *gin.Context){
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}
	err = h.roleUsecase.Delete(c.Request.Context(), uint(id))
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "role deleted",
	})
}