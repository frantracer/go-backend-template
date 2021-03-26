package handlers

import (
	"context"

	"github.com/google/uuid"

	"github.com/frantacer/go-backend-template/src/application/repositories"
	"github.com/frantacer/go-backend-template/src/domain"
)

type InsertTaskHandler interface {
	Handle(ctx context.Context, cmd InsertTaskCommand) (domain.Task, error)
}

type InsertTaskCommand struct {
	Message string
}

type InsertTaskCommandHandler struct {
	creator repositories.UnitOfWorkCreator
}

func NewInsertTaskCommandHandler(creator repositories.UnitOfWorkCreator) InsertTaskCommandHandler {
	return InsertTaskCommandHandler{creator: creator}
}

func (h InsertTaskCommandHandler) Handle(ctx context.Context, cmd InsertTaskCommand) (domain.Task, error) {
	uow, err := h.creator.Create(ctx)
	if err != nil {
		return domain.Task{}, err
	}
	defer func() { _ = uow.Rollback() }()

	task := domain.NewTask(uuid.New(), cmd.Message)
	if err = uow.InsertTask(task); err != nil {
		return domain.Task{}, err
	}

	if err = uow.Commit(); err != nil {
		return domain.Task{}, err
	}

	return task, nil
}
