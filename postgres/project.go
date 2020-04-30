package postgres

import (
	"context"

	"github.com/eriktate/wrkhub"
	"github.com/eriktate/wrkhub/uid"
)

func (s *Store) CreateProject(ctx context.Context, project wrkhub.Project) (uid.UID, error) {
	query := getQuery("createProject")
	if project.ID.Empty() {
		project.ID = uid.New()
	}

	return project.ID, runNamedTx(ctx, s.db, query, project)
}

func (s *Store) UpdateProject(ctx context.Context, project wrkhub.Project) error {
	query := getQuery("updateProject")

	return runNamedTx(ctx, s.db, query, project)
}

func (s *Store) FetchProject(ctx context.Context, id uid.UID) (wrkhub.Project, error) {
	query := getQuery("fetchProject")

	var project wrkhub.Project
	if err := s.db.GetContext(ctx, &project, query, id); err != nil {
		return project, err
	}

	return project, nil
}

func (s *Store) ListProjects(ctx context.Context, req wrkhub.ListProjectsReq) ([]wrkhub.Project, error) {
	query := getQuery("listProjects")

	var projects []wrkhub.Project
	if err := s.db.SelectContext(ctx, &projects, query); err != nil {
		return nil, err
	}

	return projects, nil
}
