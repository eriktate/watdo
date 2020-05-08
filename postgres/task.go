package postgres

import (
	"context"

	"github.com/eriktate/wrkhub"
	"github.com/eriktate/wrkhub/uid"
)

func (s *Store) CreateTask(ctx context.Context, task wrkhub.Task) (uid.UID, error) {
	query := getQuery("createTask")
	if task.ID.Empty() {
		task.ID = uid.New()
	}

	return task.ID, runNamedTx(ctx, s.db, query, task)
}

func (s *Store) UpdateTask(ctx context.Context, task wrkhub.Task) error {
	query := getQuery("updateTask")

	return runNamedTx(ctx, s.db, query, task)
}

func (s *Store) FetchTask(ctx context.Context, id uid.UID) (wrkhub.Task, error) {
	query := getQuery("fetchTask")

	var task wrkhub.Task
	if err := s.db.GetContext(ctx, &task, query, id); err != nil {
		return task, err
	}

	return task, nil
}

func (s *Store) ListTasks(ctx context.Context, req wrkhub.ListTasksReq) ([]wrkhub.Task, error) {
	query := getQuery("listTasks")

	var tasks []wrkhub.Task
	if err := s.db.SelectContext(ctx, &tasks, query, req.ProjectID); err != nil {
		return nil, err
	}

	return tasks, nil
}
