package repository

import (
	"ERP_APPS/internal/admin"
	"context"
)
type AdminRepository interface {
	GetAll(ctx context.Context) ([]*admin.Admin, error)
	GetByID(ctx context.Context, id uint) (*admin.Admin, error)
	Create(ctx context.Context, admin *admin.Admin) (*admin.Admin, error)
	Update(ctx context.Context, admin *admin.Admin) (*admin.Admin, error)
	Delete(ctx context.Context, id uint) error
	Login(ctx context.Context, email string) (*admin.Admin, error)
	Register(ctx context.Context, admin *admin.Admin) (*admin.Admin, error)
	ChangePassword(ctx context.Context, id uint, password string) error
	FindByEmail(ctx context.Context, email string) (*admin.Admin, error)
	SaveResetToken(ctx context.Context, id uint, token string) error
	UpdatePasswordWithToken(ctx context.Context, token, newPassword string) error
}
