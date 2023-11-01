package repository

import (
	"assignment/internal/model"
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Repository interface {
	AddUser(ctx context.Context, user model.User) error
	CheckUserByEmail(ctx context.Context, userEmails string) (bool, error)
}

type RepoImplement struct {
	pgConn boil.ContextExecutor
}

func New(pgConn boil.ContextExecutor) Repository {
	return RepoImplement{pgConn: pgConn}
}
