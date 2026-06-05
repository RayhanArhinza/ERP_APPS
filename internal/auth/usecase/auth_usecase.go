package usecase

import (
	admin "ERP_APPS/internal/admin"
	"ERP_APPS/internal/auth/repository"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Login(ctx context.Context, email, password string) (*admin.Admin, error)
	Logout(ctx context.Context, id uint) error
	Register(ctx context.Context, adminData *admin.Admin) (*admin.Admin, error)
	ChangePassword(ctx context.Context, id uint, password string) error
	FindByEmail(ctx context.Context, email string) (*admin.Admin, error)
	SaveResetToken(ctx context.Context, id uint, token string) error
	UpdatePasswordWithToken(ctx context.Context, token, newPassword string) error
}

type authUsecase struct {
	authRepo repository.AuthRepository
}

func NewAuthUsecase(authRepo repository.AuthRepository) AuthUsecase {
	return &authUsecase{
		authRepo: authRepo,
	}
}

func (u *authUsecase) Login(
	ctx context.Context,
	email,
	password string,
) (*admin.Admin, error) {

	adminData, err := u.authRepo.Login(ctx, email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(adminData.Password),
		[]byte(password),
	)

	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	return adminData, nil
}

func (u *authUsecase) Logout(ctx context.Context, id uint) error {
	return nil
}

func (u *authUsecase) Register(
	ctx context.Context,
	adminData *admin.Admin,
) (*admin.Admin, error) {

	if existing, _ := u.authRepo.FindByEmail(ctx, adminData.Email); existing != nil {
		return nil, errors.New("email already in use")
	}

	if adminData.Email == "" {
		return nil, errors.New("email cannot be empty")
	}

	if adminData.Password == "" {
		return nil, errors.New("password cannot be empty")
	}

	if len(adminData.Password) < 6 {
		return nil, errors.New("password must be at least 6 characters")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(adminData.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return nil, err
	}

	adminData.Password = string(hashedPassword)

	return u.authRepo.Register(ctx, adminData)
}

func (u *authUsecase) ChangePassword(
	ctx context.Context,
	id uint,
	password string,
) error {

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	return u.authRepo.ChangePassword(
		ctx,
		id,
		string(hashedPassword),
	)
}

func (u *authUsecase) FindByEmail(
	ctx context.Context,
	email string,
) (*admin.Admin, error) {

	if email == "" {
		return nil, errors.New("email cannot be empty")
	}

	return u.authRepo.FindByEmail(ctx, email)
}

func (u *authUsecase) SaveResetToken(
	ctx context.Context,
	id uint,
	token string,
) error {

	if token == "" {
		return errors.New("token cannot be empty")
	}

	return u.authRepo.SaveResetToken(ctx, id, token)
}

func (u *authUsecase) UpdatePasswordWithToken(
	ctx context.Context,
	token string,
	newPassword string,
) error {

	if token == "" {
		return errors.New("token cannot be empty")
	}

	if newPassword == "" {
		return errors.New("new password cannot be empty")
	}

	if len(newPassword) < 6 {
		return errors.New("new password must be at least 6 characters")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(newPassword),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	return u.authRepo.UpdatePasswordWithToken(
		ctx,
		token,
		string(hashedPassword),
	)
}