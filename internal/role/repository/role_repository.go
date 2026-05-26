package repository

import (
	role "ERP_APPS/internal"
	"context"
)

type RoleRepository interface {
	GetAll(ctx context.Context) ([]*role.Role, error)
	GetByID(ctx context.Context, id uint) (*role.Role, error)
	Create(ctx context.Context, role *role.Role) (*role.Role, error)
	Update(ctx context.Context, role *role.Role) (*role.Role, error)
	Delete(ctx context.Context, id uint) error
}