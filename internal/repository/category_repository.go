package repository

import (
	"github.com/devfajar/blog-app/internal/domain"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(c *domain.Category) error
	List() ([]domain.Category, error)
	FindByID(id uint) (*domain.Category, error)
	FindBySlug(slug string) (*domain.Category, error)
}

type categoryRepo struct{ db *gorm.DB }

func NewCategoryRepo(db *gorm.DB) CategoryRepository { return &categoryRepo{db: db} }

func (r *categoryRepo) Create(c *domain.Category) error {
	return r.db.Create(c).Error
}
func (r *categoryRepo) List() ([]domain.Category, error) {
	var out []domain.Category
	return out, r.db.Order("name asc").Find(&out).Error
}
func (r *categoryRepo) FindByID(id uint) (*domain.Category, error) {
	var c domain.Category
	if err := r.db.First(&c, id).Error; err != nil {
		return nil, err
	}
	return &c, nil
}
func (r *categoryRepo) FindBySlug(slug string) (*domain.Category, error) {
	var c domain.Category
	if err := r.db.Where("slug = ?", slug).First(&c).Error; err != nil {
		return nil, err
	}
	return &c, nil
}
