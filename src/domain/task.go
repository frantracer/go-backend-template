package domain

import "github.com/google/uuid"

type Task struct {
	id uuid.UUID
}

func NewTask(id uuid.UUID) Task {
	return Task{id: id}
}

func (t *Task) GetID() uuid.UUID {
	return t.id
}
