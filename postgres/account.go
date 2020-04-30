package postgres

import (
	"context"

	"github.com/eriktate/wrkhub"
	"github.com/eriktate/wrkhub/uid"
)

func (s *Store) CreateAccount(ctx context.Context, account wrkhub.Account) (uid.UID, error) {
	query := getQuery("createAccount")
	if account.ID.Empty() {
		account.ID = uid.New()
	}

	return account.ID, runNamedTx(ctx, s.db, query, account)
}

func (s *Store) UpdateAccount(ctx context.Context, account wrkhub.Account) error {
	query := getQuery("updateAccount")

	return runNamedTx(ctx, s.db, query, account)
}

func (s *Store) FetchAccount(ctx context.Context, id uid.UID) (wrkhub.Account, error) {
	query := getQuery("fetchAccount")

	var account wrkhub.Account
	if err := s.db.GetContext(ctx, &account, query, id); err != nil {
		return account, err
	}

	return account, nil
}

func (s *Store) ListAccounts(ctx context.Context, req wrkhub.ListAccountsReq) ([]wrkhub.Account, error) {
	query := getQuery("listAccounts")

	var accounts []wrkhub.Account
	if err := s.db.SelectContext(ctx, &accounts, query, req); err != nil {
		return nil, err
	}

	return accounts, nil
}
