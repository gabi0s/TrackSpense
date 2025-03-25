package services

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// route handler to get a user solde
func GetSoldeUser(c *fiber.Ctx) error {
	userID := c.Params("userid") // gathers user_id in the URL

	var user User
	if err := DB.Where("user_id = ?", userID).First(&user).Error; err != nil {
		log.Println("User error:", err)
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
			"solde":   nil,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Solde gathered with success",
		"solde":   user.Solde,
	})
}

// route handler to update user solde (adding or removing money from it)
func UpdateSoldeUser(c *fiber.Ctx) error {
	type UpdateSoldeRequest struct {
		UserID string  `json:"user_id"`
		Amount float64 `json:"amount"`
	}

	var req UpdateSoldeRequest

	if err := c.BodyParser(&req); err != nil { // checks validity of data
		log.Println(err)
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid data",
		})
	}

	var user User
	if err := DB.Where("user_id = ?", req.UserID).First(&user).Error; err != nil { // checks user existence in database
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	newSolde := user.Solde + req.Amount
	if newSolde < 0 { // if solde is negative -> error
		return c.Status(400).JSON(fiber.Map{
			"message":          "Insufficient funds: solde cannot be negative",
			"current_solde":    user.Solde,
			"attempted_change": req.Amount,
		})
	}

	user.Solde = newSolde

	// registers the update in the database
	if err := DB.Model(&User{}).Where("user_id = ?", user.UserID).Updates(map[string]interface{}{"solde": user.Solde}).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error while updating solde",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Solde successfully updated",
		"solde":   user.Solde,
	})
}
