package usecase

import (
	"ERP_APPS/internal/admin"
	"ERP_APPS/internal/admin/repository"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AdminUsecase interface {
	GetAll(ctx context.Context) ([]*admin.Admin, error)
	GetByID(ctx context.Context, id uint) (*admin.Admin, error)
	Create(ctx context.Context, admin *admin.Admin) (*admin.Admin, error)
	Update(ctx context.Context, admin *admin.Admin) (*admin.Admin, error)
	Delete(ctx context.Context, id uint) error
	Login(ctx context.Context, email, password string) (*admin.Admin, error)
	Register(ctx context.Context, admin *admin.Admin) (*admin.Admin, error)
	ChangePassword(ctx context.Context, id uint, password string) error
	FindByEmail(ctx context.Context, email string) (*admin.Admin, error)
	SaveResetToken(ctx context.Context, id uint, token string) error
	UpdatePasswordWithToken(ctx context.Context, token, newPassword string) error
}
type adminUsecase struct {
	adminRepo repository.AdminRepository
}
func NewAdminUsecase(adminRepo repository.AdminRepository) AdminUsecase {
	return &adminUsecase{
		adminRepo: adminRepo,
	}
}
func (u *adminUsecase) GetAll(ctx context.Context) ([]*admin.Admin, error){
	return u.adminRepo.GetAll(ctx)
}
func (u *adminUsecase) GetByID(ctx context.Context, id uint)(*admin.Admin, error){
	if id == 0 {
		return nil, errors.New("id cannot be empty")
	}
	return u.adminRepo.GetByID(ctx, id)
}
func (u *adminUsecase) Create(ctx context.Context, admin *admin.Admin)(*admin.Admin, error){
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	admin.Password = string(hashedPassword)
	if existing, _ := u.adminRepo.FindByEmail(ctx, admin.Email); existing != nil {
		return nil, errors.New("email already in use")
	}
	return u.adminRepo.Create(ctx, admin)
}
func (u *adminUsecase) Update(ctx context.Context, admin *admin.Admin)(*admin.Admin, error){
	existing, err := u.adminRepo.GetByID(ctx, admin.ID)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, errors.New("admin not found")
	}
	if admin.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		admin.Password = string(hashedPassword)
	}
	if existing.Email != admin.Email {
		if other, _ := u.adminRepo.FindByEmail(ctx, admin.Email); other != nil && other.ID != admin.ID {
			return nil, errors.New("email already in use")
		}
	}else {
		admin.Email = existing.Email
	}
	return u.adminRepo.Update(ctx, admin)
}
func (u *adminUsecase) Delete(ctx context.Context, id uint)(error){
	return u.adminRepo.Delete(ctx, id)
}
func (u *adminUsecase) Login(ctx context.Context, email, password string)(*admin.Admin, error){
	admin, err := u.adminRepo.Login(ctx, email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(admin.Password),
		[]byte(password),
	)

	if err != nil {
		return nil, errors.New("invalid password or email")
	}

	return admin, nil
}
func (u *adminUsecase) Register(ctx context.Context, admin *admin.Admin)(*admin.Admin, error){
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	admin.Password = string(hashedPassword)
	if existing, _ := u.adminRepo.FindByEmail(ctx, admin.Email); existing != nil {
		return nil, errors.New("email already in use")
	}
	return u.adminRepo.Register(ctx, admin)
}
func (u *adminUsecase) ChangePassword(ctx context.Context, id uint, password string)(error){
	if password == "" {
		return errors.New("password cannot be empty")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	password = string(hashedPassword)
	return u.adminRepo.ChangePassword(ctx, id, password)
}
func (u *adminUsecase) FindByEmail(ctx context.Context, email string)(*admin.Admin, error){
	if email == "" {
		return nil, errors.New("email cannot be empty")
	}
	return u.adminRepo.FindByEmail(ctx, email)
}
func (u *adminUsecase) SaveResetToken(ctx context.Context, id uint, token string)(error){
	if token == "" {
		return errors.New("token cannot be empty")
	}
	return u.adminRepo.SaveResetToken(ctx, id, token)
}
func (u *adminUsecase) UpdatePasswordWithToken(ctx context.Context, token, newPassword string)(error){
	if token == "" {
		return errors.New("token cannot be empty")
	}
	if newPassword == "" {
		return errors.New("new password cannot be empty")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newPassword = string(hashedPassword)
	return u.adminRepo.UpdatePasswordWithToken(ctx, token, newPassword)
}