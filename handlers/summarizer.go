package handlers

import (
	"summarizer-backend/models"
	"summarizer-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SummarizeHandler(c *fiber.Ctx) error {
	req := new(models.SummarizeRequest)
	if err := c.BodyParser(req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"response": "Invalid-request",
		})
	}

	summary, err := services.Summarizer(req.Text)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"summary": summary,
	})

}
