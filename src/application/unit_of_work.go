package application

import "github.com/frantacer/go-backend-template/src/domain"

type UnitOfWork interface {
	InsertTask(t domain.Task) error
	FindTasks() ([]domain.Task, error)
}
