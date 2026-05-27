package tasks_repository_postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/glebateee/taskapp/internal/core/domain"
	core_errors "github.com/glebateee/taskapp/internal/core/errors"
	core_postgres_pool "github.com/glebateee/taskapp/internal/core/repository/postgres/pool"
)

func (r *TasksRepository) GetTask(
	ctx context.Context,
	taskID int,
) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `SELECT
				id,
				version,
				title, 
				description, 
				completed, 
				created_at, 
				completed_at, 
				author_user_id
			  FROM taskapp.tasks
			  WHERE id = $1;`

	row := r.pool.QueryRow(ctx, query, taskID)
	var taskModel TaskModel
	if err := row.Scan(
		&taskModel.ID,
		&taskModel.Version,
		&taskModel.Title,
		&taskModel.Description,
		&taskModel.Completed,
		&taskModel.CreatedAt,
		&taskModel.CompletedAt,
		&taskModel.AuthorID,
	); err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Task{}, fmt.Errorf(
				"%v: task with id='%d': %w",
				err,
				taskID,
				core_errors.ErrNotFound,
			)
		}
		return domain.Task{}, fmt.Errorf("scan task: %w", err)
	}
	taskDomain := taskDomainFromModel(taskModel)
	return taskDomain, nil
}
