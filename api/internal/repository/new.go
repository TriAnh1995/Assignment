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
	CheckIfBlocked(context.Context, []string) (bool, error)
	CheckIfFollowed(context.Context, []string) (bool, error)
	SubscribeToBlocked(context.Context, []string) error
	SubscribeToNonBlocked(ctx context.Context, emails []string) error
}

type RepoImplement struct {
	pgConn boil.ContextExecutor
}

func New(pgConn boil.ContextExecutor) Repository {
	return RepoImplement{pgConn: pgConn}
}
func (i RepoImplement) CheckIfFollowed(ctx context.Context, userEmails []string) (bool, error) {
	return false, nil
}

func (i RepoImplement) CheckIfBlocked(ctx context.Context, userEmails []string) (bool, error) {
	return false, nil
}

func (i RepoImplement) SubscribeToBlocked(ctx context.Context, userEmails []string) error {
	return nil
}

func (i RepoImplement) SubscribeToNonBlocked(ctx context.Context, emails []string) error {
	return nil
}
