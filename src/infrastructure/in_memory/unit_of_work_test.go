package in_memory

import (
	"github.com/frantacer/go-backend-template/src/domain"
	"github.com/google/uuid"
	"testing"
)
import "github.com/stretchr/testify/require"

func TestInMemoryUnitOfWork(t *testing.T) {
	t.Run(`Given an empty in memory unit of work
	when a task is inserted and the method FindTasks is called
	then it must return a list of tasks that only includes the inserted task`,
		func(t *testing.T) {
			uow := NewUnitOfWork()

			initialTasks, err := uow.FindTasks()
			require.NoError(t, err)
			require.Empty(t, initialTasks)

			id := uuid.MustParse("cd3dc666-f753-4bee-b550-52c61c60a640")
			newTask := domain.NewTask(id)
			err = uow.InsertTask(newTask)
			require.NoError(t, err)

			currentTasks, err := uow.FindTasks()
			require.NoError(t, err)
			require.Equal(t, []domain.Task{newTask}, currentTasks)
		})
}
