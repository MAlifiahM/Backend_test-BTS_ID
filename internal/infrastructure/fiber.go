package infrastructure

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"time"
)

func SetupFiberApp() *fiber.App {
	app := fiber.New(fiber.Config{
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${method} ${path} - ${status} - ${latency}\n",
	}))

	app.Use(recover.New())

	app.Use(func(c *fiber.Ctx) error {
		err := c.Next()
		if err != nil {
			log.Printf("Error: %v\n", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
		}
		return nil
	})

	return app
}
