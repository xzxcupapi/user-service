package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	Name      string    `json:"name" gorm:"size:100;not null"`
	Email     string    `json:"email" gorm:"size:100;unique;not null"`
	Password  string    `json:"-" gorm:"size:100;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepository interface {
	Create(user *User) error
	GetByID(id string) (*User, error)
	GetByEmail(email string) (*User, error)
	Update(user *User) error
	Delete(id string) error
}

type UserUsecase interface {
	Create(user *User) error
	GetByID(id string) (*User, error)
	GetByEmail(email string) (*User, error)
	Update(user *User) error
	Delete(id string) error
}
