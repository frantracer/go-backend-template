package handlers

import (
	"context"

	"github.com/frantacer/go-backend-template/src/application/repositories"
	"github.com/frantacer/go-backend-template/src/domain"
)

type FindTasksHandler interface {
	Handle(ctx context.Context) ([]domain.Task, error)
}

type FindTasksCommandHandler struct {
	creator repositories.UnitOfWorkCreator
}

func NewFindTasksCommandHandler(creator repositories.UnitOfWorkCreator) FindTasksCommandHandler {
	return FindTasksCommandHandler{creator: creator}
}

func (h FindTasksCommandHandler) Handle(ctx context.Context) ([]domain.Task, error) {
	uow, err := h.creator.Create(ctx)
	if err != nil {
		return []domain.Task{}, err
	}
	defer func() { _ = uow.Rollback() }()

	tasks, err := uow.FindTasks()
	if err != nil {
		return []domain.Task{}, err
	}

	return tasks, nil
}
