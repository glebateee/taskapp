package tasks_repository_postgres

import (
	"time"

	"github.com/glebateee/taskapp/internal/core/domain"
)

type TaskModel struct {
	ID          int
	Version     int
	Title       string
	Description *string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
	AuthorID    int
}

func taskDomainFromModel(taskModel TaskModel) domain.Task {
	return domain.NewTask(
		taskModel.ID,
		taskModel.Version,
		taskModel.Title,
		taskModel.Description,
		taskModel.Completed,
		taskModel.CreatedAt,
		taskModel.CompletedAt,
		taskModel.AuthorID,
	)
}

func taskDomainsFromModels(taskModel []TaskModel) []domain.Task {
	domains := make([]domain.Task, len(taskModel))
	for i, model := range taskModel {
		domains[i] = taskDomainFromModel(model)
	}
	return domains
}
