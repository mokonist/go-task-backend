package domain

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type TaskStatus int

const (
	UNKNOWN TaskStatus = iota + 1
	TODO
	WIP
	DONE
)

type Task struct {
	Id          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   string     `json:"created_at"`
	UpdatedAt   string     `json:"updated_at"`
}

func NewTask(title, description string, status int) (*Task, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	taskStatus, err := validateTaskStatus(status)
	if err != nil {
		return nil, err
	}

	return &Task{
		Id:          id.String(),
		Title:       title,
		Description: description,
		Status:      *taskStatus,
		CreatedAt:   fmt.Sprintf("%s", time.Now().Format("2006-01-02T15:04:05+09:00")),
		UpdatedAt:   fmt.Sprintf("%s", time.Now().Format("2006-01-02T15:04:05+09:00")),
	}, nil
}

func (t *Task) Update(title, description string, status int) (*Task, error) {
	taskStatus, err := validateTaskStatus(status)
	if err != nil {
		return nil, err
	}

	t.Title = title
	t.Description = description
	t.Status = *taskStatus
	t.UpdatedAt = fmt.Sprintf("%s", time.Now().Format("2006-01-02T15:04:05+09:00"))
	return t, nil
}

func validateTaskStatus(status int) (*TaskStatus, error) {
	if status < int(TODO) || status > int(DONE) {
		return nil, errors.New("invalid status value")
	}
	taskStatus := TaskStatus(status)
	return &taskStatus, nil
}
