package main

import (
	"fmt"
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
	srv := service.NewService(repository.NewInMemoryRepo())

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8228))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
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
