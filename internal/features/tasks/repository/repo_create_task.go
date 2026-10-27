package tasks_repository

import "context"

func (r *RepositoryTask) CreateTask(ctx context.Context, title string, description *string) (ModelTask, error) {
	sqlQuery := `
	INSERT INTO todoapp.task (user_id, title, description) 
	VALUES ($1, $2, $3) 
	RETURNING id, title, description, created_at, done_at`

	var task ModelTask
	err := r.db.QueryRow(ctx, sqlQuery, 1, title, description).Scan(&task.ID, &task.Title, &task.Description, &task.CreatedAt, &task.DoneAt)
	if err != nil {
		return ModelTask{}, err
	}
	return task, nil
}
