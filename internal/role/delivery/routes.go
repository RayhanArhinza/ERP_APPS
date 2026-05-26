package delivery

import (
	"github.com/gin-gonic/gin"

	"ERP_APPS/internal/role/usecase"
)
func NewRoleHandler(r *gin.RouterGroup, roleUsecase usecase.RoleUsecase){
	handler := &RoleHandler{
		roleUsecase: roleUsecase,
	}
	roleGroup := r.Group("/roles")
	{
		roleGroup.GET("/", handler.GetAll)
		roleGroup.GET("/:id", handler.GetByID)
		roleGroup.POST("/", handler.Create)
		roleGroup.PUT("/:id", handler.Update)
		roleGroup.DELETE("/:id", handler.Delete)
	}

}