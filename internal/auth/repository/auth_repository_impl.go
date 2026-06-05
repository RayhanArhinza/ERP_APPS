package repository

import (
	admin "ERP_APPS/internal/admin"
	"context"

	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) Login(ctx context.Context, email string) (*admin.Admin, error) {
	var adminData admin.Admin

	err := r.db.WithContext(ctx).
		Preload("Role").
		Where("email = ?", email).
		First(&adminData).Error

	if err != nil {
		return nil, err
	}

	return &adminData, nil
}

func (r *authRepository) Register(
	ctx context.Context,
	adminData *admin.Admin,
) (*admin.Admin, error) {

	err := r.db.WithContext(ctx).
		Create(adminData).Error

	if err != nil {
		return nil, err
	}

	err = r.db.WithContext(ctx).
		Preload("Role").
		First(adminData, adminData.ID).Error

	if err != nil {
		return nil, err
	}

	return adminData, nil
}

func (r *authRepository) ChangePassword(
	ctx context.Context,
	id uint,
	password string,
) error {

	return r.db.WithContext(ctx).
		Model(&admin.Admin{}).
		Where("id = ?", id).
		Update("password", password).Error
}

func (r *authRepository) FindByEmail(
	ctx context.Context,
	email string,
) (*admin.Admin, error) {

	var adminData admin.Admin

	err := r.db.WithContext(ctx).
		Preload("Role").
		Where("email = ?", email).
		First(&adminData).Error

	if err != nil {
		return nil, err
	}

	return &adminData, nil
}

func (r *authRepository) SaveResetToken(
	ctx context.Context,
	id uint,
	token string,
) error {

	return r.db.WithContext(ctx).
		Model(&admin.Admin{}).
		Where("id = ?", id).
		Update("reset_token", token).Error
}

func (r *authRepository) UpdatePasswordWithToken(
	ctx context.Context,
	token string,
	newPassword string,
) error {

	return r.db.WithContext(ctx).
		Model(&admin.Admin{}).
		Where("reset_token = ?", token).
		Updates(map[string]interface{}{
			"password":    newPassword,
			"reset_token": "",
		}).Error
}