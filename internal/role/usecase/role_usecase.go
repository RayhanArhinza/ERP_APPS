package usecase

import (
	role "ERP_APPS/internal"
	"ERP_APPS/internal/role/repository"
	"context"
	"errors"
)
type RoleUsecase interface {
	GetAll(ctx context.Context) ([]*role.Role, error)
	GetByID(ctx context.Context, id uint) (*role.Role, error)
	Create(ctx context.Context, role *role.Role) (*role.Role, error)
	Update(ctx context.Context, role *role.Role) (*role.Role, error)
	Delete(ctx context.Context, id uint) error
}
type roleUsecase struct {
	roleRepo repository.RoleRepository	
}
func NewRoleUsecase(roleRepo repository.RoleRepository) RoleUsecase {
	return &roleUsecase{
		roleRepo: roleRepo,
	}
}
func (u *roleUsecase) GetAll(ctx context.Context) ([]*role.Role, error){
	return u.roleRepo.GetAll(ctx)
}
func (u *roleUsecase) GetByID(ctx context.Context, id uint)(*role.Role, error){
	return u.roleRepo.GetByID(ctx, id)
}
func (u *roleUsecase) Create(ctx context.Context, role *role.Role)(*role.Role, error){
	return u.roleRepo.Create(ctx, role)
}
func (u *roleUsecase) Update(ctx context.Context, role *role.Role)(*role.Role, error){
	existing, err := u.roleRepo.GetByID(ctx, role.ID)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, errors.New("role not found")
	}
	return u.roleRepo.Update(ctx, role)
}
func (u *roleUsecase) Delete(ctx context.Context, id uint)(error){
	return u.roleRepo.Delete(ctx, id)

}