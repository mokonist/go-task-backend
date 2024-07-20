package service

import (
	"context"
	domain2 "github.com/mokocm/todo-backend/pkg/domain"
)

type TaskService interface {
	GetTask(ctx context.Context, id string) (*domain2.Task, error)
	CreateTask(ctx context.Context, task domain2.Task) (*domain2.Task, error)
	GetAllTasks(ctx context.Context) (*map[domain2.Task]domain2.Task, error)
}
type taskService struct {
	Repo domain2.TaskRepository
}

func newTaskService(r domain2.TaskRepository) TaskService {
	return &taskService{Repo: r}
}

func (t taskService) GetTask(ctx context.Context, id string) (*domain2.Task, error) {
	res, _ := t.Repo.GetTaskByID(id)
	return res, nil
}

func (t taskService) CreateTask(ctx context.Context, task domain2.Task) (*domain2.Task, error) {
	res, _ := t.Repo.CreateTask(task)
	return res, nil
}

func (t taskService) GetAllTasks(ctx context.Context) (*map[domain2.Task]domain2.Task, error) {
	res, _ := t.Repo.GetAllTasks()
	return res, nil
}
