package repository

import (
	"user-services/internal/domain"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db}
}

// Create implements domain.UserRepository.
func (u *userRepository) Create(user *domain.User) error {
	return u.db.Create(user).Error
}

// Update implements domain.UserRepository.
func (u *userRepository) Update(user *domain.User) error {
	return u.db.Save(user).Error
}

// Delete implements domain.UserRepository.
func (u *userRepository) Delete(id string) error {
	return u.db.Delete(&domain.User{}, id).Error
}

// GetByEmail implements domain.UserRepository.
func (u *userRepository) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByID implements domain.UserRepository.
func (u *userRepository) GetByID(id string) (*domain.User, error) {
	var user domain.User
	err := u.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
