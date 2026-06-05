package auth

import "time"

type Auth struct {
	ID                  uint      `json:"id"`
	Email               string    `json:"email"`
	Password            string    `json:"password"`
	RoleID              uint      `json:"role_id"`
	RoleName            string    `json:"role_name"`
	Token               string    `json:"token"`
	ResetToken          string    `json:"reset_token"`
	ResetTokenExpiresAt time.Time `json:"reset_token_expires_at"`
}