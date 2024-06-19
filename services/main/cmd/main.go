package main

import (
	"context"
	"log"
	"main-service/internal/handler"
	"main-service/internal/repository"
	"main-service/internal/server"
	"main-service/internal/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	authConfig := repository.PostgresConfig{
		Host:     "db",
		Port:     "5432",
		Username: "postgres",
		Password: "qwerty",
		DBName:   "postgres",
		SSLMode:  "disable",
	}

	authPostgres, err := repository.NewPostgresDB(authConfig)
	if err != nil {
		log.Fatalf("error while connecting postgres: %s", err.Error())
	}
	authRepo := repository.NewAuthPostgres(authPostgres)

	tasksRepo, err := repository.NewTasks("tasks:8228")
	if err != nil {
		log.Fatalf("error while connecting tasks: %s", err.Error())
	}

	authService := service.NewAuthService(authRepo)

	kafka, err := repository.NewStatistics()
	if err != nil {
		log.Fatalf("error while creating kafka: %v", err)
	}
	eventService := service.NewEvents(kafka)

	handlers := handler.NewHandler(authService, tasksRepo, eventService)

	routes := handlers.InitRoutes()

	srv := &server.Server{}
	go func() {
		if err := srv.Run("8080", routes); err != nil {
			log.Printf("error occurred on server running %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Printf("Stopping server...")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("error occurred on server shutting down: %s", err.Error())
	}
}
