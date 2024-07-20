package domain

import (
	"github.com/mokocm/todo-backend/pkg/infrastructure"
	"go.uber.org/fx"
)

type TaskRepository interface {
	CreateTask(task Task) (*Task, error)
	GetTaskByID(id string) (*Task, error)
	GetAllTasks() (*map[Task]Task, error)
}

type taskRepositoryParams struct {
	fx.In
	DB *infrastructure.DB
}

func newTaskRepository(params taskRepositoryParams) TaskRepository {
	return &taskRepositoryParams{DB: params.DB}
}

func (r *taskRepositoryParams) CreateTask(task Task) (*Task, error) {
	r.DB.Save(task.Id, task)
	data := r.DB.Get(task.Id).(Task)

	return &data, nil
}

func (r *taskRepositoryParams) GetTaskByID(id string) (*Task, error) {
	res := r.DB.Get(id).(Task)
	return &res, nil
}

func (r *taskRepositoryParams) GetAllTasks() (*map[Task]Task, error) {
	res := r.DB.GetAll()
	result := make(map[Task]Task, len(res))
	for _, v := range res {
		result[v.(Task)] = v.(Task)
	}
	return &result, nil
}
