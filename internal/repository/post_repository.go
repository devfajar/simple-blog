package repository

import (
	"github.com/devfajar/blog-app/internal/domain"
	"gorm.io/gorm"
)

type PostRepository interface {
	Create(p *domain.Post) error
	Update(p *domain.Post) error
	Delete(id uint) error
	GetByID(id uint) (*domain.Post, error)
	FindBySlug(slug string) (*domain.Post, error)
	ListPublic(search string, categoryID uint, page, pageSize int) ([]domain.Post, int64, error)
	ListAdmin(search string, status string, page, pageSize int) ([]domain.Post, int64, error)
}

type postRepo struct{ db *gorm.DB }

func NewPostRepo(db *gorm.DB) PostRepository { return &postRepo{db: db} }

func (r *postRepo) Create(p *domain.Post) error { return r.db.Create(p).Error }
func (r *postRepo) Update(p *domain.Post) error { return r.db.Save(p).Error }
func (r *postRepo) Delete(id uint) error        { return r.db.Delete(&domain.Post{}, id).Error }

func (r *postRepo) GetByID(id uint) (*domain.Post, error) {
	var p domain.Post
	if err := r.db.First(&p, id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}
func (r *postRepo) FindBySlug(slug string) (*domain.Post, error) {
	var p domain.Post
	if err := r.db.Where("slug = ?", slug).First(&p).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *postRepo) ListPublic(search string, categoryID uint, page, pageSize int) ([]domain.Post, int64, error) {
	var posts []domain.Post
	var total int64
	q := r.db.Model(&domain.Post{}).Where("status = ?", "published")
	if categoryID > 0 {
		q = q.Where("category_id = ?", categoryID)
	}
	if search != "" {
		like := "%" + search + "%"
		q = q.Where("title LIKE ? OR content LIKE ?", like, like)
	}
	q.Count(&total)
	err := q.Order("published_at desc, created_at desc").
		Limit(pageSize).Offset((page - 1) * pageSize).Find(&posts).Error
	return posts, total, err
}

func (r *postRepo) ListAdmin(search string, status string, page, pageSize int) ([]domain.Post, int64, error) {
	var posts []domain.Post
	var total int64
	q := r.db.Model(&domain.Post{})
	if status != "" {
		q = q.Where("status = ?", status)
	}
	if search != "" {
		like := "%" + search + "%"
		q = q.Where("title LIKE ? OR content LIKE ?", like, like)
	}
	q.Count(&total)
	err := q.Order("created_at desc").
		Limit(pageSize).Offset((page - 1) * pageSize).Find(&posts).Error
	return posts, total, err
}
