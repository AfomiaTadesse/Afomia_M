package main

import (
	"context"
	"log"
	"time"

	"github.com/AfomiaTadesse/Afomia_M/backend/config"
	"github.com/AfomiaTadesse/Afomia_M/backend/controller"
	"github.com/AfomiaTadesse/Afomia_M/backend/repository"
	"github.com/AfomiaTadesse/Afomia_M/backend/router"
	"github.com/AfomiaTadesse/Afomia_M/backend/usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg := config.Load()

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	db := client.Database("movie_collection")

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	movieRepo := repository.NewMovieRepository(db)

	// Initialize use cases
	userUsecase := usecase.NewUserUsecase(userRepo, cfg.JWTSecret, time.Hour)
	movieUsecase := usecase.NewMovieUsecase(movieRepo)

	// Initialize controllers
	userCtrl := controller.NewUserController(userUsecase)
	movieCtrl := controller.NewMovieController(movieUsecase)

	// Setup router with both controllers
	r := router.SetupRouter(userCtrl, movieCtrl,cfg.JWTSecret)

	// Start server
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}