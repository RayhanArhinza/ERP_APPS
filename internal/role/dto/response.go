package dto

import (
	role "ERP_APPS/internal"
	"time"
)
type RoleResponse struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
func ToRoleResponse(role role.Role) RoleResponse{
	return RoleResponse{
		ID		: role.ID,
		Name	: role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}
}
func ToRoleResponses(roles []role.Role) []RoleResponse{
	var roleResponses []RoleResponse
	for _, role := range roles{
		roleResponses = append(roleResponses, ToRoleResponse(role))
	}
	return roleResponses
}
