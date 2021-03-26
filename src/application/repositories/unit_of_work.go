package repositories

import (
	"context"

	"github.com/frantacer/go-backend-template/src/domain"
)

type UnitOfWork interface {
	InsertTask(t domain.Task) error
	FindTasks() ([]domain.Task, error)
	Commit() error
	Rollback() error
}

type UnitOfWorkCreator interface {
	Create(ctx context.Context) (UnitOfWork, error)
}
