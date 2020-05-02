package service

import (
	"context"

	"github.com/eriktate/wrkhub"
	"github.com/eriktate/wrkhub/uid"
)

type ProjectService struct {
	store wrkhub.ProjectStore
}

func NewProjectService(store wrkhub.ProjectStore) ProjectService {
	return ProjectService{store}
}

func (s ProjectService) SaveProject(ctx context.Context, project wrkhub.Project) (uid.UID, error) {
	if err := project.Validate(); err != nil {
		return project.ID, err
	}

	if project.ID.Empty() {
		return s.store.CreateProject(ctx, project)
	}

	return project.ID, s.store.UpdateProject(ctx, project)
}

func (s ProjectService) FetchProject(ctx context.Context, id uid.UID) (wrkhub.Project, error) {
	return s.store.FetchProject(ctx, id)
}

func (s ProjectService) ListProjects(ctx context.Context) ([]wrkhub.Project, error) {
	return s.store.ListProjects(ctx, wrkhub.ListProjectsReq{})
}
