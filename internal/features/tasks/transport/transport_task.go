package tasks_transport_http

import (
	"context"

	core_logger "github.com/CryptoFill666/todo_app/internal/core/logger"
	tasks_repository "github.com/CryptoFill666/todo_app/internal/features/tasks/repository"
)

type TasksHTTPHandler struct {
	tasksService TasksService
	log          *core_logger.Logger
}

type TasksService interface {
	CreateTask(ctx context.Context, title string, description *string) (tasks_repository.ModelTask, error)
}

func NewTasksHTTPHandler(tasksService TasksService, log *core_logger.Logger) *TasksHTTPHandler {
	return &TasksHTTPHandler{
		tasksService: tasksService,
		log:          log,
	}
}
