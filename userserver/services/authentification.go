package services

import (
	"github.com/gofiber/fiber/v2" // ORM for go
)

// User model
type User struct {
	UserID   string  `json:"user_id"`
	Nom      string  `json:"nom"`
	Password string  `json:"password"`
	Solde    float64 `json:"solde"`
}

// AuthRequest model (for authentification request)
type AuthRequest struct {
	Nom      string `json:"nom"`
	Password string `json:"password"`
}

// AuthRequest model (for authentification response)
type AuthResponse struct {
	UserID  string `json:"user_id,omitempty"`
	Message string `json:"message"`
}

func AuthenticateUser(c *fiber.Ctx) error {
	var request AuthRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(AuthResponse{Message: "invalid request"})
	}

	var user User
	result := DB.Where("nom = ? AND password = ?", request.Nom, request.Password).First(&user)

	if result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(AuthResponse{Message: "Name or password incorrect"})
	}

	return c.JSON(AuthResponse{UserID: user.UserID, Message: "Connexion succeed !"})
}
