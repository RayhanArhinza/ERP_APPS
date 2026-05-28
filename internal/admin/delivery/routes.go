package delivery

import (
	"ERP_APPS/internal/admin/usecase"

	"github.com/gin-gonic/gin"
)
func NewAdminHandler(r *gin.RouterGroup, adminUsecase usecase.AdminUsecase){
	handler := &AdminHandler{
		adminUsecase: adminUsecase,
	}
	adminGroup := r.Group("/admins")
	{
		adminGroup.GET("/", handler.GetAll)
		adminGroup.GET("/:id", handler.GetByID)
		adminGroup.POST("/", handler.Create)
		adminGroup.PUT("/:id", handler.Update)
		adminGroup.DELETE("/:id", handler.Delete)
	}
}