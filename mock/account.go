package mock

import (
	"context"

	"github.com/eriktate/wrkhub"
	"github.com/eriktate/wrkhub/uid"
)

type AccountService struct {
	SaveAccountFn     func(ctx context.Context, account wrkhub.Account) (uid.UID, error)
	SaveAccountCalled int

	ListAccountsFn     func(ctx context.Context) ([]wrkhub.Account, error)
	ListAccountsCalled int

	FetchAccountFn     func(ctx context.Context, id uid.UID) (wrkhub.Account, error)
	FetchAccountCalled int

	Error error
}

func (m *AccountService) SaveAccount(ctx context.Context, account wrkhub.Account) (uid.UID, error) {
	m.SaveAccountCalled++

	if m.SaveAccountFn != nil {
		return m.SaveAccountFn(ctx, account)
	}

	return uid.UID{}, m.Error
}

func (m *AccountService) ListAccounts(ctx context.Context) ([]wrkhub.Account, error) {
	m.ListAccountsCalled++

	if m.ListAccountsFn != nil {
		return m.ListAccountsFn(ctx)
	}

	return nil, m.Error
}

func (m *AccountService) FetchAccount(ctx context.Context, id uid.UID) (wrkhub.Account, error) {
	m.FetchAccountCalled++

	if m.FetchAccountFn != nil {
		return m.FetchAccountFn(ctx, id)
	}

	return wrkhub.Account{}, m.Error
}
