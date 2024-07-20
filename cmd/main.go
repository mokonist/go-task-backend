package main

import (
	"github.com/mokocm/todo-backend/pkg/domain"
	"github.com/mokocm/todo-backend/pkg/handler"
	"github.com/mokocm/todo-backend/pkg/infrastructure"
	"github.com/mokocm/todo-backend/pkg/service"
	"github.com/mokocm/todo-backend/pkg/usecase"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

func main() {
	fx.New(
		usecase.Module,
		domain.Module,
		infrastructure.Module,
		service.Module,
		handler.Module,
		fx.Invoke(
			func(server *grpc.Server) {}),
	).Run()
}
