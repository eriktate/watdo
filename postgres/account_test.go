package postgres_test

import (
	"context"
	"testing"

	"github.com/eriktate/watdo"
	"github.com/eriktate/watdo/postgres"
)

func Test_Accounts(t *testing.T) {
	// SETUP
	ctx := context.TODO()
	store, err := postgres.New(postgres.NewStoreOpts())
	if err != nil {
		t.Fatalf("failed to connect to postgres: %s", err)
	}

	account := watdo.Account{
		Name: "Test account1",
	}

	// RUN & ASSERT
	id, err := store.CreateAccount(ctx, account)
	if err != nil {
		t.Fatalf("failed to create test account: %s", err)
	}

	fetchedAccount, err := store.FetchAccount(ctx, id)
	if err != nil {
		t.Fatalf("failed to fetch test account: %s", err)
	}

	if fetchedAccount.Name != account.Name {
		t.Fatal("saved account name doesn't match source")
	}
}
