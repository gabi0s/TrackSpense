package services

import "github.com/gofiber/fiber/v2"

type Expense struct {
	UserID          string  `json:"user_id"`
	BudgetID        string  `json:"budget_id"`
	BudgetCategorie string  `json:"budget_categorie"`
	Price           float64 `json:"price"`
}

func AddExpense(c *fiber.Ctx) error {
	var expense Expense
	if err := c.BodyParser(&expense); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}

	if err := DB.Create(&expense).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to add expense"})
	}

	return c.JSON(fiber.Map{"message": "Expense added successfully", "expense": expense})
}

func GetExpense(c *fiber.Ctx) error {
	userID := c.Params("userid")

	// checks userid existence
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "User ID is required"})
	}

	type Expenses struct {
		ExpenseID       string  `json:"expense_id"`
		UserID          string  `json:"user_id"`
		BudgetID        string  `json:"budget_id"`
		BudgetCategorie string  `json:"budget_categorie"`
		Price           float64 `json:"price"`
	}

	var expenses []Expenses

	if err := DB.Where("user_id = ?", userID).Find(&expenses).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error retrieving expenses"})
	}

	return c.JSON(expenses)
}
