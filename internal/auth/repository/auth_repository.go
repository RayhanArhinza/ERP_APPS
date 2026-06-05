package repository

import (
	admin "ERP_APPS/internal/admin"
	"context"
)

type AuthRepository interface {
	Login(ctx context.Context, email string) (*admin.Admin, error)
	Register(ctx context.Context, admin *admin.Admin) (*admin.Admin, error)
	ChangePassword(ctx context.Context, id uint, password string) error
	FindByEmail(ctx context.Context, email string) (*admin.Admin, error)
	SaveResetToken(ctx context.Context, id uint, token string) error
	UpdatePasswordWithToken(ctx context.Context, token, newPassword string) error
}