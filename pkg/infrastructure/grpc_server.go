package infrastructure

import (
	"context"
	todov1 "github.com/mokocm/todo-backend/gen/proto/task/v1"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"log"
	"net"
)

func newGrpcServer(lc fx.Lifecycle, taskHandler todov1.TaskServiceServer) *grpc.Server {
	server := grpc.NewServer()
	port := ":50051"
	listen, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			log.Println("Starting gRPC server...")
			go func() {
				todov1.RegisterTaskServiceServer(server, taskHandler)
				server.Serve(listen)
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			log.Println("Stopping gRPC server...")
			server.GracefulStop()
			return nil
		},
	})
	return server
}
