package handler

import (
	"context"
	todov1 "github.com/mokocm/todo-backend/gen/proto/task/v1"
	"github.com/mokocm/todo-backend/pkg/usecase"
	"go.uber.org/fx"
)

type taskHandler struct {
	todov1.UnimplementedTaskServiceServer
	params taskHandlerParams
}

type taskHandlerParams struct {
	fx.In
	CreateUseCase *usecase.CreateTaskUseCase
	GetUseCase    *usecase.GetTaskUseCase
	GetAllUseCase *usecase.GetAllTasksUseCase
}

func newTaskHandler(params taskHandlerParams) todov1.TaskServiceServer {
	return &taskHandler{params: params}
}

func (p *taskHandler) GetTask(ctx context.Context, req *todov1.GetTaskRequest) (*todov1.GetTaskResponse, error) {
	res, err := p.params.GetUseCase.Execute(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &todov1.GetTaskResponse{
		Task: &todov1.Task{
			Id:          res.Id,
			Status:      todov1.TaskStatus(res.Status),
			CreatedAt:   res.CreatedAt,
			Title:       res.Title,
			Description: res.Description,
			UpdatedAt:   res.UpdatedAt,
		},
	}, nil
}

func (p *taskHandler) GetTasks(ctx context.Context, req *todov1.GetTasksRequest) (*todov1.GetTasksResponse, error) {
	res, err := p.params.GetAllUseCase.Execute(ctx)
	if err != nil {
		return nil, err
	}

	tasks := make([]*todov1.Task, 0)
	for _, v := range *res {
		tasks = append(tasks, &todov1.Task{
			Id:          v.Id,
			Status:      todov1.TaskStatus(v.Status),
			CreatedAt:   v.CreatedAt,
			Title:       v.Title,
			Description: v.Description,
			UpdatedAt:   v.UpdatedAt,
		})
	}
	return &todov1.GetTasksResponse{
		Task: tasks,
	}, nil
}

func (p *taskHandler) CreateTask(ctx context.Context, req *todov1.CreateTaskRequest) (*todov1.CreateTaskResponse, error) {
	res, err := p.params.CreateUseCase.Execute(ctx, req.Title, req.Description, int(req.Status))
	if err != nil {
		return nil, err
	}
	return &todov1.CreateTaskResponse{
		Task: &todov1.Task{
			Id:          res.Id,
			Status:      todov1.TaskStatus(res.Status),
			CreatedAt:   res.CreatedAt,
			Title:       res.Title,
			Description: res.Description,
			UpdatedAt:   res.UpdatedAt,
		},
	}, nil
}
