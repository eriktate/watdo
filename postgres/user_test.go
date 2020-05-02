package postgres_test

import (
	"context"
	"os"
	"testing"

	"github.com/eriktate/wrkhub"
	"github.com/eriktate/wrkhub/postgres"
)

func init() {
	os.Chdir("../")
}

func Test_Users(t *testing.T) {
	// SETUP
	ctx := context.TODO()
	store, err := postgres.New()
	if err != nil {
		t.Fatalf("failed to connect to postgres: %s", err)
	}

	user1 := wrkhub.User{
		Name:  "Test User",
		Email: "test@test.com",
	}

	user2 := wrkhub.User{
		Name:  "Test User 2",
		Email: "test2@test.com",
	}

	// RUN & ASSERT
	existing, err := store.ListUsers(ctx, wrkhub.ListUsersReq{})
	if err != nil {
		t.Fatal("failed to get existing list of users")
	}

	id1, err := store.CreateUser(ctx, user1)
	if err != nil {
		t.Fatalf("failed to create test user1: %s", err)
	}

	if _, err := store.CreateUser(ctx, user2); err != nil {
		t.Fatalf("failed to create test user2: %s", err)
	}

	fetchedUser, err := store.FetchUser(ctx, id1)
	if err != nil {
		t.Fatalf("failed to fetch test user: %s", err)
	}

	listUsers, err := store.ListUsers(ctx, wrkhub.ListUsersReq{})
	if err != nil {
		t.Fatalf("failed to list users: %s", err)
	}

	if fetchedUser.Name != user1.Name {
		t.Fatal("saved user name doesn't match source")
	}

	if fetchedUser.Email != user1.Email {
		t.Fatal("saved user email doesn't match source")
	}

	if len(listUsers) != len(existing)+2 {
		t.Fatalf("expected user listing to increase by 2")
	}
}
