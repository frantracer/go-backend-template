package in_memory

import (
	"github.com/frantacer/go-backend-template/src/application"
	"github.com/frantacer/go-backend-template/src/domain"
)

var tasks []domain.Task

type UnitOfWork struct {
	localTasks []domain.Task
}

var _ application.UnitOfWork = &UnitOfWork{}

func NewUnitOfWork() UnitOfWork {
	return UnitOfWork{
		localTasks: []domain.Task{},
	}
}

func (w *UnitOfWork) InsertTask(t domain.Task) error {
	w.localTasks = append(w.localTasks, t)
	return nil
}

func (w *UnitOfWork) FindTasks() ([]domain.Task, error) {
	return tasks, nil
}

func (w *UnitOfWork) Commit() error {
	tasks = append(tasks, w.localTasks...)
	return nil
}

func (w *UnitOfWork) Rollback() error {
	w.localTasks = []domain.Task{}
	return nil
}
