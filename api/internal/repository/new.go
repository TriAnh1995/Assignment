package repository

import (
	"assignment/internal/model"
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Repository interface {
	AddUser(context.Context, model.User) error
	CheckUserByEmail(context.Context, string) (bool, error)
	AddFriendship(context.Context, string, string) error
	CheckFriendship(context.Context, []string) (bool, error)
	BlockSubscribedUser(ctx context.Context, userEmails []string) error
	BlockNonSubscribedUser(ctx context.Context, userEmails []string) error
}

type RepoImplement struct {
	pgConn boil.ContextExecutor
}

func New(pgConn boil.ContextExecutor) Repository {
	return RepoImplement{pgConn: pgConn}
}
