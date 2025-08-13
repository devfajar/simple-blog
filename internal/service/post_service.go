package service

import (
	"errors"
	"time"

	"github.com/devfajar/blog-app/internal/domain"
	"github.com/devfajar/blog-app/internal/repository"
	"github.com/devfajar/blog-app/internal/util"
)

type PostService interface {
	Create(authorID uint, title, content, status string, categoryID uint) (*domain.Post, error)
	Update(id uint, title, content, status string, categoryID uint) (*domain.Post, error)
	Delete(id uint) error
	GetBySlug(slug string) (*domain.Post, error)
	ListPublic(search string, categoryID uint, page, pageSize int) ([]domain.Post, int64, error)
	ListAdmin(search, status string, page, pageSize int) ([]domain.Post, int64, error)
}

type postService struct {
	posts repository.PostRepository
	cats  repository.CategoryRepository
}

func NewPostService(posts repository.PostRepository, cats repository.CategoryRepository) PostService {
	return &postService{posts: posts, cats: cats}
}

func (s *postService) Create(authorID uint, title, content, status string, categoryID uint) (*domain.Post, error) {
	if title == "" || content == "" {
		return nil, errors.New("title and content are required")
	}
	if status == "" {
		status = "draft"
	}
	if status != "draft" && status != "published" {
		return nil, errors.New("invalid status")
	}
	if _, err := s.cats.FindByID(categoryID); err != nil {
		return nil, errors.New("invalid category")
	}

	slug := util.Slugify(title)
	p := &domain.Post{
		Title:      title,
		Slug:       slug,
		Content:    content,
		Status:     status,
		AuthorID:   authorID,
		CategoryID: categoryID,
	}
	if status == "published" {
		now := time.Now()
		p.PublishedAt = &now
	}
	if err := s.posts.Create(p); err != nil {
		return nil, err
	}
	return p, nil
}

func (s *postService) Update(id uint, title, content, status string, categoryID uint) (*domain.Post, error) {
	p, err := s.posts.GetByID(id)
	if err != nil {
		return nil, err
	}
	if title != "" && title != p.Title {
		p.Title = title
		p.Slug = util.Slugify(title)
	}
	if content != "" {
		p.Content = content
	}
	if status != "" {
		if status != "draft" && status != "published" {
			return nil, errors.New("invalid status")
		}
		p.Status = status
		if status == "published" && p.PublishedAt == nil {
			now := time.Now()
			p.PublishedAt = &now
		}
	}
	if categoryID > 0 {
		if _, err := s.cats.FindByID(categoryID); err != nil {
			return nil, errors.New("invalid category")
		}
		p.CategoryID = categoryID
	}
	if err := s.posts.Update(p); err != nil {
		return nil, err
	}
	return p, nil
}

func (s *postService) Delete(id uint) error {
	return s.posts.Delete(id)
}
func (s *postService) GetBySlug(slug string) (*domain.Post, error) {
	return s.posts.FindBySlug(slug)
}
func (s *postService) ListPublic(search string, categoryID uint, page, pageSize int) ([]domain.Post, int64, error) {
	return s.posts.ListPublic(search, categoryID, page, pageSize)
}
func (s *postService) ListAdmin(search, status string, page, pageSize int) ([]domain.Post, int64, error) {
	return s.posts.ListAdmin(search, status, page, pageSize)
}
