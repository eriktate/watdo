package postgres

import (
	"context"
	"database/sql"

	"github.com/eriktate/wrkhub"
	"github.com/eriktate/wrkhub/uid"
)

type fetchUserRow struct {
	wrkhub.User
	AccountID        uid.UID        `db:"account_id"`
	DefaultProjectID uid.UID        `db:"default_project_id"`
	AccountName      sql.NullString `db:"account_name"`
}

type fetchUserRows []fetchUserRow

func (r fetchUserRows) ToUser() wrkhub.User {
	if len(r) == 0 {
		return wrkhub.User{}
	}

	user := r[0].User
	for _, u := range r {
		assoc := wrkhub.Association{
			AccountID:        u.AccountID,
			DefaultProjectID: u.DefaultProjectID,
			AccountName:      u.AccountName.String,
		}
		user.Associations = append(user.Associations, assoc)
	}

	return user
}

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

	var rows fetchUserRows
	if err := s.db.SelectContext(ctx, &rows, query, id); err != nil {
		return wrkhub.User{}, err
	}

	return rows.ToUser(), nil
}

func (s *Store) ListUsers(ctx context.Context, req wrkhub.ListUsersReq) ([]wrkhub.User, error) {
	query := getQuery("listUsers")

	var users []wrkhub.User
	if err := s.db.SelectContext(ctx, &users, query); err != nil {
		return nil, err
	}

	return users, nil
}
