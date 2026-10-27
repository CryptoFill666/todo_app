package tasks_service

import (
	"context"

	tasks_repository "github.com/CryptoFill666/todo_app/internal/features/tasks/repository"
)

func (s *TaskService) CreateTask(ctx context.Context, title string, description *string) (tasks_repository.ModelTask, error) {
	task, err := s.tasksRepository.CreateTask(ctx, title, description)
	if err != nil {
		return tasks_repository.ModelTask{}, err
	}
	return task, nil
}
