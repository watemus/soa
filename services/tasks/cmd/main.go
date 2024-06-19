package main

import (
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"tasks-service/internal/repository"
	service "tasks-service/internal/service"
	taskspb "tasks-service/proto"
)

func main() {
	cfg := repository.PostgresConfig{
		Host:     "db-tasks",
		Port:     "5432",
		Username: "postgres",
		Password: "qwerty",
		DBName:   "postgres",
		SSLMode:  "disable",
	}
	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("error while connecting to postgres: %v", err)
	}

	repo := repository.NewPostgresRepo(db)
	srv := service.NewService(repo)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8228))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			recovery.UnaryServerInterceptor(),
		),
		grpc.ChainStreamInterceptor(
			recovery.StreamServerInterceptor(),
		),
	)
	taskspb.RegisterTaskServiceServer(grpcServer, srv)

	go func() {
		log.Println("Starting serving...")
		err = grpcServer.Serve(lis)
		if err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Printf("Stopping server...")
	grpcServer.GracefulStop()
	log.Printf("Server stopped")
}
