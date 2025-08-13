package repository

import (
	"github.com/devfajar/blog-app/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(email string) (*domain.User, error)
	GetByID(id uint) (*domain.User, error)
	Create(u *domain.User) error
}

type userRepo struct{ db *gorm.DB }

func NewUserRepo(db *gorm.DB) UserRepository { return &userRepo{db: db} }

func (r *userRepo) FindByEmail(email string) (*domain.User, error) {
	var u domain.User
	if err := r.db.Where("email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
func (r *userRepo) GetByID(id uint) (*domain.User, error) {
	var u domain.User
	if err := r.db.First(&u, id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
func (r *userRepo) Create(u *domain.User) error {
	return r.db.Create(u).Error
}
