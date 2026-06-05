package dto

import "ERP_APPS/internal/admin"

type AdminResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	RoleID    uint   `json:"role_id"`
	RoleName  string `json:"role_name"`
	TglLahir  string `json:"tgl_lahir"`
	Alamat    string `json:"alamat"`
}

type PaginationResponse struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	TotalData  int64 `json:"total_data"`
	TotalPages int64 `json:"total_pages"`
}

type AdminPaginatedResponse struct {
	Data       []AdminResponse    `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}

func ToAdminResponse(admin admin.Admin) AdminResponse {
	return AdminResponse{
		ID:       admin.ID,
		Name:     admin.Name,
		Email:    admin.Email,
		RoleID:   admin.RoleID,
		RoleName: admin.Role.Name,
		TglLahir: admin.TglLahir.Format("2006-01-02"),
		Alamat:   admin.Alamat,
	}
}

func ToAdminResponses(admins []*admin.Admin) []AdminResponse {
	var responses []AdminResponse
	for _, adm := range admins {
		responses = append(responses, ToAdminResponse(*adm))
	}
	return responses
}