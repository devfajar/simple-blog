package handler

import (
	"strconv"

	"github.com/devfajar/blog-app/internal/service"
	"github.com/devfajar/blog-app/internal/util"
	"github.com/gofiber/fiber/v2"
)

type PostHandler struct{ svc service.PostService }

func NewPostHandler(svc service.PostService) *PostHandler { return &PostHandler{svc: svc} }

type createPostReq struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	Status     string `json:"status"` // draft/published
	CategoryID uint   `json:"category_id"`
}

// Create POST /api/v1/admin/posts
func (h *PostHandler) Create(c *fiber.Ctx) error {
	var req createPostReq
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid body")
	}
	claims := util.GetUserClaims(c)
	out, err := h.svc.Create(claims.UserID, req.Title, req.Content, req.Status, req.CategoryID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(out)
}

type updatePostReq struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	Status     string `json:"status"`
	CategoryID uint   `json:"category_id"`
}

// Update PUT /api/v1/admin/posts/:id
func (h *PostHandler) Update(c *fiber.Ctx) error {
	var req updatePostReq
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid body")
	}
	id64, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil || id64 == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}
	out, err := h.svc.Update(uint(id64), req.Title, req.Content, req.Status, req.CategoryID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.JSON(out)
}

// Delete DELETE /api/v1/admin/posts/:id
func (h *PostHandler) Delete(c *fiber.Ctx) error {
	id64, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil || id64 == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}
	if err := h.svc.Delete(uint(id64)); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// ListPublic GET /api/v1/public/posts
func (h *PostHandler) ListPublic(c *fiber.Ctx) error {
	search := c.Query("search", "")
	catID := util.ParseUintDefault(c.Query("category_id", "0"), 0)
	page := util.ParseIntDefault(c.Query("page", "1"), 1)
	size := util.ParseIntDefault(c.Query("page_size", "10"), 10)

	items, total, err := h.svc.ListPublic(search, uint(catID), page, size)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(fiber.Map{
		"items": items,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"page_size": size,
		},
	})
}

// ListAdmin GET /api/v1/admin/posts
func (h *PostHandler) ListAdmin(c *fiber.Ctx) error {
	search := c.Query("search", "")
	status := c.Query("status", "")
	page := util.ParseIntDefault(c.Query("page", "1"), 1)
	size := util.ParseIntDefault(c.Query("page_size", "10"), 10)

	items, total, err := h.svc.ListAdmin(search, status, page, size)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(fiber.Map{
		"items": items,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"page_size": size,
		},
	})
}

// GetBySlug GET /api/v1/public/posts/:slug
func (h *PostHandler) GetBySlug(c *fiber.Ctx) error {
	p, err := h.svc.GetBySlug(c.Params("slug"))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "post not found")
	}
	return c.JSON(p)
}
