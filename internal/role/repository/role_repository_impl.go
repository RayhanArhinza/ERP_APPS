package repository

import (
	role "ERP_APPS/internal"
	"context"

	"gorm.io/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{
		db: db,
	}
}

func (r *roleRepository) GetAll(ctx context.Context) ([]*role.Role, error){
	var roles []*role.Role
	err := r.db.WithContext(ctx).Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}
func (r *roleRepository) GetByID(ctx context.Context, id uint) (*role.Role, error){
	var role role.Role
	err := r.db.WithContext(ctx).First(&role, id).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}
func (r *roleRepository) Create(ctx context.Context, role *role.Role)(*role.Role, error){
	err := r.db.WithContext(ctx).Create(role).Error
	if err != nil {
		return nil, err
	}
	return role, nil
}
func (r *roleRepository) Update(ctx context.Context, role *role.Role)(*role.Role, error){
	err := r.db.WithContext(ctx).Model(&role).Where("id = ?", role.ID).Updates(role).Error
	if err != nil {
		return nil, err
	}
	return role, nil

}
func (r *roleRepository) Delete(ctx context.Context, id uint)(error){
	err := r.db.WithContext(ctx).Delete(&role.Role{}, id).Error
	if err != nil {
		return err
	}
	return nil
}