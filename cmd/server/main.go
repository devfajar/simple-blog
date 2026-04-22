package main

import (
	"log"
	"os"

	"github.com/devfajar/blog-app/internal/config"
	"github.com/devfajar/blog-app/internal/database"
	"github.com/devfajar/blog-app/internal/handler"
	"github.com/devfajar/blog-app/internal/middleware"
	"github.com/devfajar/blog-app/internal/repository"
	"github.com/devfajar/blog-app/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	cfg := config.Load()
	db := database.Connect(*cfg)

	// Auto-migrate models
	if err := database.AutoMigrate(db); err != nil {
		log.Fatalf("migration error: %v", err)
	}

	// Seed admin if empty
	if err := database.SeedAdmin(db, *cfg); err != nil {
		log.Fatalf("seed admin error: %v", err)
	}

	// DI wiring
	userRepo := repository.NewUserRepo(db)
	postRepo := repository.NewPostRepo(db)
	catRepo := repository.NewCategoryRepo(db)

	authSvc := service.NewAuthService(userRepo, *cfg)
	postSvc := service.NewPostService(postRepo, catRepo)
	catSvc := service.NewCategoryService(catRepo)

	authH := handler.NewAuthHandler(authSvc)
	postH := handler.NewPostHandler(postSvc)
	catH := handler.NewCategoryHandler(catSvc)

	app := fiber.New(fiber.Config{
		AppName:      "Portfolio Blog API",
		ErrorHandler: handler.FiberErrorHandler,
	})
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173, http://localhost:4173",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: false,
		MaxAge:           300,
	}))

	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Auth
	v1.Post("/auth/login", authH.Login)

	// Public
	pub := v1.Group("/public")
	pub.Get("/posts", postH.ListPublic)
	pub.Get("/posts/:slug", postH.GetBySlug)
	pub.Get("/categories", catH.List)

	// Admin (JWT)
	admin := v1.Group("/admin", middleware.JWT(cfg.JWTSecret))
	admin.Post("/categories", catH.Create)
	admin.Post("/posts", postH.Create)
	admin.Put("/posts/:id", postH.Update)
	admin.Delete("/posts/:id", postH.Delete)
	admin.Get("/posts", postH.ListAdmin)

	port := cfg.AppPort
	if port == "" {
		port = ":8080"
	}
	if port[0] != ':' {
		port = ":" + port
	}
	log.Printf("listening on %s", port)
	if err := app.Listen(port); err != nil {
		log.Println("server stopped:", err)
		os.Exit(1)
	}
}
