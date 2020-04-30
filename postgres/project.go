package postgres

import (
	"context"

	"github.com/eriktate/watdo"
	"github.com/eriktate/watdo/uid"
)

func (s *Store) CreateProject(ctx context.Context, project watdo.Project) (uid.UID, error) {
	query := getQuery("createProject")
	if project.ID.Empty() {
		project.ID = uid.New()
	}

	return project.ID, runNamedTx(ctx, s.db, query, project)
}

func (s *Store) UpdateProject(ctx context.Context, project watdo.Project) error {
	query := getQuery("updateProject")

	return runNamedTx(ctx, s.db, query, project)
}

func (s *Store) FetchProject(ctx context.Context, id uid.UID) (watdo.Project, error) {
	query := getQuery("fetchProject")

	var project watdo.Project
	if err := s.db.GetContext(ctx, &project, query, id); err != nil {
		return project, err
	}

	return project, nil
}

func (s *Store) ListProjects(ctx context.Context, req watdo.ListProjectsReq) ([]watdo.Project, error) {
	query := getQuery("listProjects")

	var projects []watdo.Project
	if err := s.db.SelectContext(ctx, &projects, query); err != nil {
		return nil, err
	}

	return projects, nil
}
