package main

import (
	"log"
	"net"

	_ "github.com/lib/pq"

	"github.com/AlexEkdahl/kango/config"
	"github.com/AlexEkdahl/kango/internal/repository"
	"github.com/AlexEkdahl/kango/internal/server/app"
	"github.com/AlexEkdahl/kango/internal/server/service"
	"github.com/AlexEkdahl/kango/pkg/contract"
	"google.golang.org/grpc"
)

func main() {
	c := config.New()
	err := repository.NewDB(*c)
	if err != nil {
		log.Fatalf("Failed: %v", err)
	}
	dao := repository.NewDAO()

	taskService := service.NewTaskService(dao)

	// Create a gRPC server and register the KanbanServer with it.
	grpcServer := grpc.NewServer()
	contract.RegisterKanbanServer(grpcServer, app.NewMicroservice(taskService))

	// Start the gRPC server.
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("Starting gRPC server on", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
