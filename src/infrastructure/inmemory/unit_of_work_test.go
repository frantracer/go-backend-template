package inmemory

import (
	"testing"

	"github.com/frantacer/go-backend-template/src/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestInMemoryUnitOfWork(t *testing.T) {
	t.Run(`Given an empty in memory unit of work
	when a task is inserted and the method FindTasks is called
	then it must return a list of tasks that only includes the inserted task`,
		func(t *testing.T) {
			uow1 := NewUnitOfWork()

			initialTasks, err := uow1.FindTasks()
			require.NoError(t, err)
			require.Empty(t, initialTasks)

			id := uuid.MustParse("cd3dc666-f753-4bee-b550-52c61c60a640")
			message := "work on the project"
			newTask := domain.NewTask(id, message)
			err = uow1.InsertTask(newTask)
			require.NoError(t, err)

			err = uow1.Commit()
			require.NoError(t, err)

			uow2 := NewUnitOfWork()

			currentTasks, err := uow2.FindTasks()
			require.NoError(t, err)
			require.Equal(t, []domain.Task{newTask}, currentTasks)
		})
}
