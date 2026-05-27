package tasks_service

import (
	"context"
	"fmt"

	"github.com/glebateee/taskapp/internal/core/domain"
)

func (s *TasksService) PatchTask(
	ctx context.Context,
	taskID int,
	taskPatch domain.TaskPatch,
) (domain.Task, error) {
	task, err := s.tasksRepository.GetTask(ctx, taskID)
	if err != nil {
		return domain.Task{}, fmt.Errorf("get task from repository: %w", err)
	}

	if err := task.ApplyPatch(taskPatch); err != nil {
		return domain.Task{}, fmt.Errorf("apply task patch: %w", err)
	}

	task, err = s.tasksRepository.PatchTask(ctx, taskID, task)
	if err != nil {
		return domain.Task{}, fmt.Errorf("save patched task: %w", err)
	}
	return task, nil
}
