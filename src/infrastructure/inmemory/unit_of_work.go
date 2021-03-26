package inmemory

import (
	"context"

	"github.com/frantacer/go-backend-template/src/application/repositories"
	"github.com/frantacer/go-backend-template/src/domain"
)

// Unit of Work Creator
type UnitOfWorkCreator struct{}

func NewUnitOfWorkCreator() UnitOfWorkCreator {
	return UnitOfWorkCreator{}
}

func (c UnitOfWorkCreator) Create(context.Context) (repositories.UnitOfWork, error) {
	uow := NewUnitOfWork()
	return &uow, nil
}

var _ repositories.UnitOfWorkCreator = &UnitOfWorkCreator{}

var tasks []domain.Task

// Unit of Work
type UnitOfWork struct {
	localTasks []domain.Task
}

func NewUnitOfWork() UnitOfWork {
	return UnitOfWork{
		localTasks: []domain.Task{},
	}
}

func (w *UnitOfWork) InsertTask(t domain.Task) error {
	w.localTasks = append(w.localTasks, t)
	return nil
}

func (w UnitOfWork) FindTasks() ([]domain.Task, error) {
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

var _ repositories.UnitOfWork = &UnitOfWork{}
