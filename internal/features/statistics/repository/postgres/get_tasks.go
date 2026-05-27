package statistics_repository_postgres

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/glebateee/taskapp/internal/core/domain"
)

func (r *StatisticsRepository) GetTasks(
	ctx context.Context,
	userID *int,
	from *time.Time,
	to *time.Time,
) ([]domain.Task, error) {
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
			  %s
			  ORDER BY id;`

	args := []any{}
	conditions := []string{}
	if userID != nil {
		args = append(args, userID)
		conditions = append(conditions, fmt.Sprintf("author_user_id = $%d", len(args)))
	}
	if from != nil {
		args = append(args, from)
		conditions = append(conditions, fmt.Sprintf("created_at >= $%d", len(args)))
	}
	if to != nil {
		args = append(args, to)
		conditions = append(conditions, fmt.Sprintf("completed_at < $%d", len(args)))
	}
	var insert string
	if len(conditions) > 0 {
		insert = " WHERE " + strings.Join(conditions, " AND ")
	}
	query = fmt.Sprintf(query, insert)
	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("select tasks: %w", err)
	}
	defer rows.Close()

	var taskModels []TaskModel
	for rows.Next() {
		var taskModel TaskModel
		if err := rows.Scan(
			&taskModel.ID,
			&taskModel.Version,
			&taskModel.Title,
			&taskModel.Description,
			&taskModel.Completed,
			&taskModel.CreatedAt,
			&taskModel.CompletedAt,
			&taskModel.AuthorID,
		); err != nil {
			return nil, fmt.Errorf("scan task: %w", err)
		}
		taskModels = append(taskModels, taskModel)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	taskDomains := taskDomainsFromModels(taskModels)
	return taskDomains, nil
}
