package dto

import "time"

type CreateAdminRequest struct {
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Password  string    `json:"password" binding:"required,min=8"`
	RoleID    uint      `json:"role_id" binding:"required"`
	TglLahir  time.Time `json:"tgl_lahir" binding:"required"`
	Alamat    string    `json:"alamat" binding:"required"`
}

type UpdateAdminRequest struct {
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Password  string    `json:"password" binding:"required,min=8"`
	RoleID    uint      `json:"role_id" binding:"required"`
	TglLahir  time.Time `json:"tgl_lahir" binding:"required"`
	Alamat    string    `json:"alamat" binding:"required"`
}

type PaginationRequest struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}