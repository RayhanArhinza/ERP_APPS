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
	var req dto.PaginationRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Limit <= 0 {
		req.Limit = 10
	}

	admins, total, err := h.adminUsecase.GetAll(c.Request.Context(), req.Page, req.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalPages := total / int64(req.Limit)
	if total%int64(req.Limit) != 0 {
		totalPages++
	}

	c.JSON(http.StatusOK, dto.AdminPaginatedResponse{
		Data: dto.ToAdminResponses(admins),
		Pagination: dto.PaginationResponse{
			Page:       req.Page,
			Limit:      req.Limit,
			TotalData:  total,
			TotalPages: totalPages,
		},
	})
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
func (h *AdminHandler) BulkCreate(c *gin.Context) {
    var requests []dto.CreateAdminRequest
    if err := c.ShouldBindJSON(&requests); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    var results []dto.AdminResponse
    for _, req := range requests {
        data := &admin.Admin{
            Name:     req.Name,
            Email:    req.Email,
            Password: req.Password,
            RoleID:   req.RoleID,
            TglLahir: req.TglLahir,
            Alamat:   req.Alamat,
        }
        res, err := h.adminUsecase.Create(c.Request.Context(), data)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": err.Error(),
                "failed_on": req.Email,
            })
            return
        }
        results = append(results, dto.ToAdminResponse(*res))
    }

    c.JSON(http.StatusCreated, results)
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