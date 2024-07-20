package usecase

import (
	"context"
	"github.com/mokocm/todo-backend/pkg/domain"
	"github.com/mokocm/todo-backend/pkg/service"
	"go.uber.org/fx"
)

type CreateTaskParams struct {
	fx.In
	Service service.TaskService
}

type CreateTaskUseCase struct {
	service service.TaskService
}

func newCreateTaskUseCase(params CreateTaskParams) *CreateTaskUseCase {
	return &CreateTaskUseCase{
		service: params.Service,
	}
}

func (p *CreateTaskUseCase) Execute(ctx context.Context, title, description string, status int) (*domain.Task, error) {

	task, err := domain.NewTask(title, description, status)
	if err != nil {
		return nil, err
	}

	res, err := p.service.CreateTask(ctx, *task)
	if err != nil {
		return nil, err
	}

	return res, nil
}
