// Command api starts the HTTP server: database, Gin router, /auth routes.
package main

import (
	"log"

	"video-platform/internal/config"
	"video-platform/internal/database"
	handlerauth "video-platform/internal/handler/auth"
	"video-platform/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	if err != nil {
		log.Fatalf("database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("database ping: %v", err)
	}

	r := gin.Default()

	// Public auth endpoints (see internal/handler/auth).
	auth := r.Group("/auth")
	auth.POST("/register", handlerauth.PostRegister(db))
	auth.POST("/login", handlerauth.PostLogin(db, cfg.JWTSecret))

	// Example protected routes: send Authorization: Bearer <token> from POST /auth/login.
	api := r.Group("/api")
	api.Use(middleware.RequireJWT(cfg.JWTSecret))
	api.GET("/me", handlerauth.GetMe)

	if err := r.Run(cfg.HTTPAddr); err != nil {
		log.Fatal(err)
	}
}
