package usecase

import (
	"context"
	"github.com/mokocm/todo-backend/pkg/domain"
	"github.com/mokocm/todo-backend/pkg/service"
	"go.uber.org/fx"
)

type GetAllTasksParams struct {
	fx.In
	Service service.TaskService
}

type GetAllTasksUseCase struct {
	service service.TaskService
}

func newGetAllTasksUseCase(params GetAllTasksParams) *GetAllTasksUseCase {
	return &GetAllTasksUseCase{
		service: params.Service,
	}
}

func (p *GetAllTasksUseCase) Execute(ctx context.Context) (*map[domain.Task]domain.Task, error) {
	res, err := p.service.GetAllTasks(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}
