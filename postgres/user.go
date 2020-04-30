package postgres

import (
	"context"

	"github.com/eriktate/wrkhub"
	"github.com/eriktate/wrkhub/uid"
)

func (s *Store) CreateUser(ctx context.Context, user wrkhub.User) (uid.UID, error) {
	query := getQuery("createUser")
	if user.ID.Empty() {
		user.ID = uid.New()
	}

	return user.ID, runNamedTx(ctx, s.db, query, user)
}

func (s *Store) UpdateUser(ctx context.Context, user wrkhub.User) error {
	query := getQuery("updateUser")

	return runNamedTx(ctx, s.db, query, user)
}

func (s *Store) FetchUser(ctx context.Context, id uid.UID) (wrkhub.User, error) {
	query := getQuery("fetchUser")

	var user wrkhub.User
	if err := s.db.GetContext(ctx, &user, query, id); err != nil {
		return user, err
	}

	return user, nil
}

func (s *Store) ListUsers(ctx context.Context, req wrkhub.ListUsersReq) ([]wrkhub.User, error) {
	query := getQuery("listUsers")

	var users []wrkhub.User
	if err := s.db.SelectContext(ctx, &users, query); err != nil {
		return nil, err
	}

	return users, nil
}
