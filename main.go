package main

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"

	"boilerplate/internal/controller"
	"boilerplate/internal/repository"
	"github.com/klass-lk/ginboot"
)

func main() {
	client, err := mongo.Connect(nil, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(nil)

	// Initialize repositories
	userRepository := repository.NewUserRepository(client.Database("example"))

	// Initialize server
	server := ginboot.New()

	// Set base path for all routes
	server.SetBasePath("/api/v1")

	// Configure CORS with custom settings
	server.CustomCORS(
		[]string{"http://localhost:3000", "https://yourdomain.com"},   // Allow specific origins
		[]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},           // Allow specific methods
		[]string{"Origin", "Content-Type", "Authorization", "Accept"}, // Allow specific headers
		24*time.Hour, // Max age of preflight requests
	)

	// Initialize and register controllers
	userController := controller.NewUserController(*userRepository)
	pingController := controller.NewPingController()

	server.RegisterController("/users", userController)
	server.RegisterController("/ping", pingController)

	if err := server.Start(8080); err != nil {
		log.Fatal(err)
	}
}
