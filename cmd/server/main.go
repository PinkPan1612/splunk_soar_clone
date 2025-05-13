package main

import (
	"log"
	"os"
	"splunk_soar_clone/config"
	"splunk_soar_clone/internal/delivery/http/handler"
	"splunk_soar_clone/internal/delivery/http/router"
	"splunk_soar_clone/internal/repository/postgres"
	"splunk_soar_clone/internal/usecase/user"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.ConnectionDatabase()

	userRepo := postgres.NewUserRepository(db)
	jwtKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	userUseCase := user.NewUserUseCase(userRepo, jwtKey)
	authHandler := handler.NewAuthHandler(userUseCase)

	r := router.SetupRouter(authHandler)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
