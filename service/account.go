package service

import (
	"context"

	"github.com/eriktate/watdo"
	"github.com/eriktate/watdo/uid"
)

type AccountService struct {
	store watdo.AccountStore
}

func NewAccountService(store watdo.AccountStore) AccountService {
	return AccountService{store}
}

func (s AccountService) SaveAccount(ctx context.Context, account watdo.Account) (uid.UID, error) {
	if account.ID.Empty() {
		return s.store.CreateAccount(ctx, account)
	}

	return account.ID, s.store.UpdateAccount(ctx, account)
}

func (s AccountService) FetchAccount(ctx context.Context, id uid.UID) (watdo.Account, error) {
	return s.store.FetchAccount(ctx, id)
}

func (s AccountService) ListAccounts(ctx context.Context) ([]watdo.Account, error) {
	return s.store.ListAccounts(ctx, watdo.ListAccountsReq{})
}
