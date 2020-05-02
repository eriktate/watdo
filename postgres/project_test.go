package postgres_test

import (
	"context"
	"testing"

	"github.com/eriktate/wrkhub"
	"github.com/eriktate/wrkhub/postgres"
	"github.com/eriktate/wrkhub/uid"
)

func Test_Projects(t *testing.T) {
	// SETUP
	ctx := context.TODO()
	store, err := postgres.New()
	if err != nil {
		t.Fatalf("failed to connect to postgres: %s", err)
	}

	account := wrkhub.Account{
		ID:   uid.New(),
		Name: "Test Project Account",
	}

	project1 := wrkhub.Project{
		Name:        "Test Project",
		Description: "A project for testing",
		AccountID:   account.ID,
	}

	project2 := wrkhub.Project{
		Name:        "Test Project 2",
		Description: "Another project for testing",
		AccountID:   account.ID,
	}

	if _, err := store.CreateAccount(ctx, account); err != nil {
		t.Fatal("failed to create owning account for project tests")
	}

	// RUN & ASSERT
	id1, err := store.CreateProject(ctx, project1)
	if err != nil {
		t.Fatalf("failed to create test project1: %s", err)
	}

	if _, err := store.CreateProject(ctx, project2); err != nil {
		t.Fatalf("failed to create test project2: %s", err)
	}

	fetchedProject, err := store.FetchProject(ctx, id1)
	if err != nil {
		t.Fatalf("failed to fetch test project: %s", err)
	}

	listProjects, err := store.ListProjects(ctx, wrkhub.ListProjectsReq{AccountID: account.ID})
	if err != nil {
		t.Fatalf("failed to list projects: %s", err)
	}

	allProjects, err := store.ListProjects(ctx, wrkhub.ListProjectsReq{})
	if err != nil {
		t.Fatalf("failed to list all projects: %s", err)
	}

	if fetchedProject.Name != project1.Name {
		t.Fatal("saved project name doesn't match source")
	}

	if fetchedProject.Description != project1.Description {
		t.Fatal("saved project email doesn't match source")
	}

	if len(listProjects) != 2 {
		t.Fatal("new project listing should be 2 more than original existing")
	}

	if len(allProjects) < len(listProjects) {
		t.Fatal("expected additional projects to exist")
	}
}
