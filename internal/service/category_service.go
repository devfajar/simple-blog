package service

import (
	"errors"

	"github.com/devfajar/blog-app/internal/domain"
	"github.com/devfajar/blog-app/internal/repository"
	"github.com/devfajar/blog-app/internal/util"
)

type CategoryService interface {
	Create(name string) (*domain.Category, error)
	List() ([]domain.Category, error)
}

type categoryService struct{ cats repository.CategoryRepository }

func NewCategoryService(cats repository.CategoryRepository) CategoryService {
	return &categoryService{cats: cats}
}

func (s *categoryService) Create(name string) (*domain.Category, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}
	c := &domain.Category{
		Name: name,
		Slug: util.Slugify(name),
	}
	if err := s.cats.Create(c); err != nil {
		return nil, err
	}
	return c, nil
}
func (s *categoryService) List() ([]domain.Category, error) {
	return s.cats.List()
}
