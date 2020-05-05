package service

import (
	"context"

	"github.com/eriktate/wrkhub"
	"github.com/eriktate/wrkhub/uid"
)

type UserService struct {
	store wrkhub.UserStore
}

func NewUserService(store wrkhub.UserStore) UserService {
	return UserService{store}
}

func (s UserService) SaveUser(ctx context.Context, user wrkhub.User) (uid.UID, error) {
	if user.ID.Empty() {
		return s.store.CreateUser(ctx, user)
	}

	return user.ID, s.store.UpdateUser(ctx, user)
}

func (s UserService) FetchUser(ctx context.Context, id uid.UID) (wrkhub.User, error) {
	return s.store.FetchUser(ctx, id)
}

func (s UserService) ListUsers(ctx context.Context, req wrkhub.ListUsersReq) ([]wrkhub.User, error) {
	return s.store.ListUsers(ctx, wrkhub.ListUsersReq{})
}
