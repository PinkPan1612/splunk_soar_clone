package main

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"splunk_soar_clone/config"
	"splunk_soar_clone/internal/delivery/http/handler"
	"splunk_soar_clone/internal/delivery/http/router"
	"splunk_soar_clone/internal/repository/postgres"
	"splunk_soar_clone/internal/usecase/user"

	"github.com/joho/godotenv"
)

func main() {
	_, b, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(b), "../..")

	// Load .env from project root
	if err := godotenv.Load(filepath.Join(projectRoot, ".env")); err != nil {
		log.Fatalf("❌ Error loading .env file: %v", err)
	}

	db := config.ConnectionDatabase()
	jwtKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	userRepo := postgres.NewUserRepository(db)
	userUseCase := user.NewUserUseCase(userRepo, jwtKey)
	authHandler := handler.NewAuthHandler(userUseCase)

	r := router.SetupRouter(authHandler, jwtKey)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
