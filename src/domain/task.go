package domain

import (
	"github.com/google/uuid"
)

type Task struct {
	id      uuid.UUID
	message string
}

func NewTask(id uuid.UUID, message string) Task {
	return Task{id: id, message: message}
}

func (t Task) ID() uuid.UUID {
	return t.id
}

func (t Task) Message() string {
	return t.message
}
