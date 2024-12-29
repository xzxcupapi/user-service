package usecase

import (
	"user-services/internal/domain"

	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{userRepo}
}

// Create implements domain.UserUsecase.
func (u *userUsecase) Create(user *domain.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return u.userRepo.Create(user)
}

// Update implements domain.UserUsecase.
func (u *userUsecase) Update(user *domain.User) error {
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}
	return u.userRepo.Update(user)
}

// Delete implements domain.UserUsecase.
func (u *userUsecase) Delete(id string) error {
	return u.userRepo.Delete(id)
}

// GetByEmail implements domain.UserUsecase.
func (u *userUsecase) GetByEmail(email string) (*domain.User, error) {
	return u.userRepo.GetByEmail(email)
}

// GetByID implements domain.UserUsecase.
func (u *userUsecase) GetByID(id string) (*domain.User, error) {
	return u.userRepo.GetByID(id)
}
