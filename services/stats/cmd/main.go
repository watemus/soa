package main

import (
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"stats/internal"
	taskspb "stats/proto"
	"syscall"
)

// @title Todo App API
// @version 1.0
// @description API Server for TodoList Application

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	repo, err := internal.NewRepo()
	if err != nil {
		log.Fatalf("error while connecting to postgres: %v", err)
	}
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

	srv := internal.NewSrv(repo)

	taskspb.RegisterStatsServiceServer(grpcServer, srv)

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
