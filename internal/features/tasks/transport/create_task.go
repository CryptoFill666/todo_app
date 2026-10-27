package tasks_transport_http

import (
	"encoding/json"
	"net/http"
)

type CreateTaskRequest struct {
	Title       string  `json:"title"`
	Description *string `json:"description"`
}

func (h *TasksHTTPHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		h.log.Warn("Failed to decode request body:", "err", err)
		return
	}
	task, err := h.tasksService.CreateTask(r.Context(), req.Title, req.Description)
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		h.log.Error("Failed to create task:", "err", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	bytes, err := json.Marshal(task)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		h.log.Error("Failed to encode response:", "err", err)
		return
	}
	w.Write(bytes)
}
