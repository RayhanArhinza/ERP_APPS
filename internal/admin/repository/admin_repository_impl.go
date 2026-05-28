package repository

import (
	"ERP_APPS/internal/admin"
	"context"

	"gorm.io/gorm"
)

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepository{
		db: db,
	}
}

func (r *adminRepository) GetAll(ctx context.Context) ([]*admin.Admin, error) {
	var admins []*admin.Admin

	if err := r.db.WithContext(ctx).
		Preload("Role").
		Find(&admins).Error; err != nil {
		return nil, err
	}

	return admins, nil
}

func (r *adminRepository) GetByID(ctx context.Context, id uint) (*admin.Admin, error) {
	var admin admin.Admin

	if err := r.db.WithContext(ctx).
		Preload("Role").
		First(&admin, id).Error; err != nil {
		return nil, err
	}

	return &admin, nil
}

func (r *adminRepository) Create(ctx context.Context, admin *admin.Admin) (*admin.Admin, error) {

	err := r.db.WithContext(ctx).Create(admin).Error
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (r *adminRepository) Update(ctx context.Context, admin *admin.Admin) (*admin.Admin, error) {

	err := r.db.WithContext(ctx).
		Model(&admin).
		Where("id = ?", admin.ID).
		Updates(admin).Error

	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (r *adminRepository) Delete(ctx context.Context, id uint) error {

	err := r.db.WithContext(ctx).
		Delete(&admin.Admin{}, id).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *adminRepository) Login(ctx context.Context, email string) (*admin.Admin, error) {

	var admin admin.Admin

	err := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&admin).Error

	if err != nil {
		return nil, err
	}

	return &admin, nil
}

func (r *adminRepository) Register(ctx context.Context, admin *admin.Admin) (*admin.Admin, error) {

	err := r.db.WithContext(ctx).Create(admin).Error
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (r *adminRepository) ChangePassword(
	ctx context.Context,
	id uint,
	password string,
) error {

	err := r.db.WithContext(ctx).
		Model(&admin.Admin{}).
		Where("id = ?", id).
		Update("password", password).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *adminRepository) FindByEmail(
	ctx context.Context,
	email string,
) (*admin.Admin, error) {

	var admin admin.Admin

	err := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&admin).Error

	if err != nil {
		return nil, err
	}

	return &admin, nil
}

func (r *adminRepository) SaveResetToken(
	ctx context.Context,
	id uint,
	token string,
) error {

	err := r.db.WithContext(ctx).
		Model(&admin.Admin{}).
		Where("id = ?", id).
		Update("reset_token", token).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *adminRepository) UpdatePasswordWithToken(
	ctx context.Context,
	token,
	newPassword string,
) error {

	err := r.db.WithContext(ctx).
		Model(&admin.Admin{}).
		Where("reset_token = ?", token).
		Updates(map[string]interface{}{
			"reset_token": nil,
			"password":    newPassword,
		}).Error

	if err != nil {
		return err
	}

	return nil
}