package in_memory

import (
	"github.com/frantacer/go-backend-template/src/application"
	"github.com/frantacer/go-backend-template/src/domain"
)

type UnitOfWork struct {
	tasks []domain.Task
}

var _ application.UnitOfWork = &UnitOfWork{}

func NewUnitOfWork() UnitOfWork {
	return UnitOfWork{
		tasks: []domain.Task{},
	}
}

func (w *UnitOfWork) InsertTask(t domain.Task) error {
	w.tasks = append(w.tasks, t)
	return nil
}

func (w *UnitOfWork) FindTasks() ([]domain.Task, error) {
	return w.tasks, nil
}
