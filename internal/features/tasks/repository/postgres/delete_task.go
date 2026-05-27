package tasks_repository_postgres

import (
	"context"
	"fmt"

	core_errors "github.com/glebateee/taskapp/internal/core/errors"
)

func (r *TasksRepository) DeleteTask(
	ctx context.Context,
	taskID int,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `DELETE
			  FROM taskapp.tasks
			  WHERE id = $1;`
	tag, err := r.pool.Exec(ctx, query, taskID)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("task with id='%d: %w", taskID, core_errors.ErrNotFound)
	}
	return nil
}
