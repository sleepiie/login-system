package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sleepiie/login-system/adapter/handler"
	"github.com/sleepiie/login-system/adapter/repository"
	"github.com/sleepiie/login-system/usecase"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("FATAL ERROR: JWT_SECRET is not set in .env")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	userRepo := repository.NewMockUserRepository()

	authUseCase := usecase.NewAuthService(userRepo, jwtSecret)

	authHandler := handler.NewAuthHandler(authUseCase)

	http.HandleFunc("/login", authHandler.Login)

	fmt.Printf("Server is running on http://localhost:%s\n", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
