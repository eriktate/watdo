package service

import (
	"context"

	"github.com/eriktate/wrkhub"
	"github.com/eriktate/wrkhub/uid"
)

type AccountService struct {
	store wrkhub.AccountStore
}

func NewAccountService(store wrkhub.AccountStore) AccountService {
	return AccountService{store}
}

func (s AccountService) SaveAccount(ctx context.Context, account wrkhub.Account) (uid.UID, error) {
	if account.ID.Empty() {
		return s.store.CreateAccount(ctx, account)
	}

	return account.ID, s.store.UpdateAccount(ctx, account)
}

func (s AccountService) FetchAccount(ctx context.Context, id uid.UID) (wrkhub.Account, error) {
	return s.store.FetchAccount(ctx, id)
}

func (s AccountService) ListAccounts(ctx context.Context) ([]wrkhub.Account, error) {
	return s.store.ListAccounts(ctx, wrkhub.ListAccountsReq{})
}
