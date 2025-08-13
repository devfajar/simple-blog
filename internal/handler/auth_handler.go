package handler

import (
	"github.com/devfajar/blog-app/internal/service"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct{ svc service.AuthService }

func NewAuthHandler(svc service.AuthService) *AuthHandler { return &AuthHandler{svc: svc} }

type loginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login POST /api/v1/auth/login
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req loginReq
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid body")
	}
	if req.Email == "" || req.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "email & password required")
	}
	token, err := h.svc.Login(req.Email, req.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}
	return c.JSON(fiber.Map{"token": token})
}
