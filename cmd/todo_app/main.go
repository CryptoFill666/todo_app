package main

import (
	"net/http"

	core_logger "github.com/CryptoFill666/todo_app/internal/core/logger"
	tasks_repository "github.com/CryptoFill666/todo_app/internal/features/tasks/repository"
	tasks_service "github.com/CryptoFill666/todo_app/internal/features/tasks/service"
	tasks_transport_http "github.com/CryptoFill666/todo_app/internal/features/tasks/transport"
)

func main() {
	logger, err := core_logger.NewLogger()
	if err != nil {
		panic(err)
	}

	repo := tasks_repository.NewRepositoryTask()
	service := tasks_service.NewTaskService(repo)

	handler := tasks_transport_http.NewTasksHTTPHandler(service, logger)
	http.HandleFunc("/tasks", handler.CreateTask)
	http.ListenAndServe(":9091", nil)

}
