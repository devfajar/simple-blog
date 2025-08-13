package handler

import (
	"github.com/devfajar/blog-app/internal/service"
	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct{ svc service.CategoryService }

func NewCategoryHandler(svc service.CategoryService) *CategoryHandler {
	return &CategoryHandler{svc: svc}
}

type createCatReq struct {
	Name string `json:"name"`
}

// Create POST /api/v1/admin/categories
func (h *CategoryHandler) Create(c *fiber.Ctx) error {
	var req createCatReq
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid body")
	}
	out, err := h.svc.Create(req.Name)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(out)
}

// List GET /api/v1/public/categories
func (h *CategoryHandler) List(c *fiber.Ctx) error {
	items, err := h.svc.List()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(items)
}
