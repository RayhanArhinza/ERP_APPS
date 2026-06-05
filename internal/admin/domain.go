package admin

import (
	"ERP_APPS/internal/role"
	"time"
)

type Admin struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	RoleID   uint      `gorm:"foreignKey:RoleID" json:"role_id"`
	Role     role.Role `gorm:"foreignKey:RoleID" json:"role"`
	TglLahir time.Time `json:"tgl_lahir"`
	Alamat   string    `json:"alamat"`
	ResetToken          string    `json:"reset_token"`
	ResetTokenExpiresAt time.Time `json:"reset_token_expires_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}