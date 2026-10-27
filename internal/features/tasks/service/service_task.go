package tasks_service

import (
	"context"

	tasks_repository "github.com/CryptoFill666/todo_app/internal/features/tasks/repository"
)

type TaskService struct {
	tasksRepository TasksRepository
}

type TasksRepository interface {
	CreateTask(ctx context.Context, title string, description *string) (tasks_repository.ModelTask, error)
}

func NewTaskService(tasksRepository TasksRepository) *TaskService {
	return &TaskService{
		tasksRepository: tasksRepository,
	}
}
