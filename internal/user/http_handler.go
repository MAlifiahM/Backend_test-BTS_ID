package user

import (
	jwt "Intersolusi_Teknologi_Asia/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	repo *UserRepository
}

func NewUserHandler(repo *UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	err := h.repo.Register(request.Username, request.Password, request.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "user registered"})
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	userID, err := h.repo.Login(request.Username, request.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	token, err := jwt.GenerateJWT(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to generate token"})
	}

	return c.JSON(fiber.Map{"token": token})
}
