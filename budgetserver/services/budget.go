package services

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Budget struct {
	UserID        string  `json:"user_id"`
	Category      string  `json:"category"`
	BudgetsLimit  float64 `json:"budgets_limit"`
	CurrentAmount float64 `json:"current_amount"`
}

// POST /create-budget
func CreateBudget(c *fiber.Ctx) error {
	var budget Budget
	if err := c.BodyParser(&budget); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unable to parse request"})
	}

	if err := DB.Create(&budget).Error; err != nil {
		log.Println("Erreur lors de la création du budget:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create budget"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Budget successfully created !", "budgets": budget})
}

// GET /get-budgets/:userid
func GetBudgetsByUserID(c *fiber.Ctx) error {
	userID := c.Params("userid")

	type Budgets struct {
		BudgetsID     string  `json:"budgets_id"`
		UserID        string  `json:"user_id"`
		Category      string  `json:"category"`
		BudgetsLimit  float64 `json:"budgets_limit"`
		CurrentAmount float64 `json:"current_amount"`
	}

	var budgets []Budgets

	if err := DB.Where("user_id = ?", userID).Find(&budgets).Error; err != nil {
		log.Println("Erreur lors de la récupération des budgets:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch budgets"})
	}

	return c.JSON(budgets)
}

// PUT /update-budget/:userid/:budgetsid
func UpdateBudget(c *fiber.Ctx) error {
	userID := c.Params("userid")
	budgetID := c.Params("budgetsid")

	var budget Budget

	// checks budget existence
	if err := DB.Where("user_id = ? AND budgets_id = ?", userID, budgetID).First(&budget).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Budget not found"})
	}

	type UpdateRequest struct {
		Amount float64 `json:"amount"`
	}
	var updateData UpdateRequest

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// checks that data is positive
	if updateData.Amount <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Amount must be positive"})
	}

	budget.CurrentAmount += updateData.Amount

	if err := DB.Model(&Budget{}).
		Where("user_id = ? AND budgets_id = ?", userID, budgetID).
		Update("current_amount", budget.CurrentAmount).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update budget"})
	}

	return c.JSON(fiber.Map{
		"message":    "Amount successfully added to budget",
		"new_amount": budget.CurrentAmount,
	})
}

// DELETE /delete-budget/:userid/:budgetsid
func DeleteBudget(c *fiber.Ctx) error {
	userID := c.Params("userid")
	budgetID := c.Params("budgetsid")

	var budget Budget
	if err := DB.Where("user_id = ? AND budgets_id = ?", userID, budgetID).First(&budget).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Budget not found"})
	}

	if err := DB.Where("user_id = ? AND budgets_id = ?", userID, budgetID).Delete(&Budget{}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete budget"})
	}

	return c.JSON("Budget successfully deleted")
}
