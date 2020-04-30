package postgres

import (
	"context"

	"github.com/eriktate/watdo"
	"github.com/eriktate/watdo/uid"
)

func (s *Store) CreateAccount(ctx context.Context, account watdo.Account) (uid.UID, error) {
	query := getQuery("createAccount")
	if account.ID.Empty() {
		account.ID = uid.New()
	}

	return account.ID, runNamedTx(ctx, s.db, query, account)
}

func (s *Store) UpdateAccount(ctx context.Context, account watdo.Account) error {
	query := getQuery("updateAccount")

	return runNamedTx(ctx, s.db, query, account)
}

func (s *Store) FetchAccount(ctx context.Context, id uid.UID) (watdo.Account, error) {
	query := getQuery("fetchAccount")

	var account watdo.Account
	if err := s.db.GetContext(ctx, &account, query, id); err != nil {
		return account, err
	}

	return account, nil
}

func (s *Store) ListAccounts(ctx context.Context, req watdo.ListAccountsReq) ([]watdo.Account, error) {
	query := getQuery("listAccounts")

	var accounts []watdo.Account
	if err := s.db.SelectContext(ctx, &accounts, query, req); err != nil {
		return nil, err
	}

	return accounts, nil
}
