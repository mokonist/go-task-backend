package usecase

import (
	"context"
	"github.com/mokocm/todo-backend/pkg/domain"
	"github.com/mokocm/todo-backend/pkg/service"
	"go.uber.org/fx"
)

type GetTaskParams struct {
	fx.In
	Service service.TaskService
}

type GetTaskUseCase struct {
	service service.TaskService
}

func newGetTaskUseCase(params GetTaskParams) *GetTaskUseCase {
	return &GetTaskUseCase{
		service: params.Service,
	}
}

func (p *GetTaskUseCase) Execute(ctx context.Context, id string) (*domain.Task, error) {
	res, err := p.service.GetTask(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
