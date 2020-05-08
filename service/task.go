package service

import (
	"context"

	"github.com/eriktate/wrkhub"
	"github.com/eriktate/wrkhub/uid"
)

type TaskService struct {
	store wrkhub.TaskStore
}

func NewTaskService(store wrkhub.TaskStore) TaskService {
	return TaskService{store}
}

func (s TaskService) SaveTask(ctx context.Context, task wrkhub.Task) (uid.UID, error) {
	if task.ID.Empty() {
		return s.store.CreateTask(ctx, task)
	}

	return task.ID, s.store.UpdateTask(ctx, task)
}

func (s TaskService) FetchTask(ctx context.Context, id uid.UID) (wrkhub.Task, error) {
	return s.store.FetchTask(ctx, id)
}

func (s TaskService) ListTasks(ctx context.Context, req wrkhub.ListTasksReq) ([]wrkhub.Task, error) {
	return s.store.ListTasks(ctx, wrkhub.ListTasksReq{})
}
