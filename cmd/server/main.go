package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"

	"user-dob-api/config"
	db "user-dob-api/db/sqlc"
	"user-dob-api/internal/handler"
	"user-dob-api/internal/logger"
	"user-dob-api/internal/middleware"
	"user-dob-api/internal/repository"
	"user-dob-api/internal/routes"
	"user-dob-api/internal/service"
)

func main() {
	logger.Init()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	conn, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatal(err)
	}
	if err = conn.Ping(); err != nil {
		log.Fatal("db ping failed:", err)
	}

	queries := db.New(conn)
	repo := repository.NewUserRepository(queries)
	svc := service.NewUserService(repo)
	h := handler.NewUserHandler(svc)

	app := fiber.New(fiber.Config{ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}})

	app.Use(middleware.RequestLogger())
	routes.Register(app, h)

	log.Fatal(app.Listen(cfg.ServerPort))
}
