package domain

import (
	"fmt"
	"time"

	core_errors "github.com/glebateee/taskapp/internal/core/errors"
)

type Task struct {
	ID          int
	Version     int
	Title       string
	Description *string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
	AuthorID    int
}

func NewTask(
	id int,
	version int,
	title string,
	description *string,
	completed bool,
	createdAt time.Time,
	completedAt *time.Time,
	authorID int,
) Task {
	return Task{
		ID:          id,
		Version:     version,
		Title:       title,
		Description: description,
		Completed:   completed,
		CreatedAt:   createdAt,
		CompletedAt: completedAt,
		AuthorID:    authorID,
	}
}
func NewTaskUninitialized(

	title string,
	description *string,
	authorID int,
) Task {
	return NewTask(
		UninitilizedID,
		UninitilizedVersion,
		title,
		description,
		false,
		time.Now(),
		nil,
		authorID,
	)
}

func (t *Task) Validate() error {
	titleLen := len([]rune(t.Title))
	if titleLen < 1 || titleLen > 100 {
		return fmt.Errorf(
			"invalid 'title' length: %d: %w",
			titleLen,
			core_errors.ErrInvalidArgument,
		)
	}
	if t.Description != nil {
		descriptionLen := len([]rune(*t.Description))
		if descriptionLen < 1 || descriptionLen > 1000 {
			return fmt.Errorf(
				"invalid 'description' length: %d: %w",
				descriptionLen,
				core_errors.ErrInvalidArgument,
			)
		}
	}
	if t.Completed {
		if t.CompletedAt == nil {
			return fmt.Errorf("'completed_at' can't be 'nil' if task completed")
		}

		if t.CompletedAt.Before(t.CreatedAt) {
			return fmt.Errorf("'completed_at' can't be before 'created_at'")

		}
	} else {
		if t.CompletedAt != nil {
			return fmt.Errorf("'completed_at' not set, but task completed")
		}
	}
	return nil
}
