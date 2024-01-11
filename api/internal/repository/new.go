package repository

import (
	"assignment/internal/model"
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Repository interface {
	AddUser(context.Context, model.User) error
	CheckUserByEmail(context.Context, string) (bool, error)
	AddFriendship(context.Context, []string) error
	CheckFriendship(context.Context, []string) (bool, error)
	GetFriendsList(context.Context, string) ([]string, error)
	CheckSubscription(context.Context, []string) (string, error)
	CheckIfBlocked(context.Context, []string) (bool, error)
	SubscribeToBlocked(context.Context, []string) error
	SubscribeToNonBlocked(context.Context, []string) error
	BlockNonSubscribedUser(context.Context, []string) error
	BlockUserInsideSubscription(context.Context, []string) error
	UpdateTopic(context.Context, model.UpdateInfo) error
	GetReceivedUpdateList(context.Context, string) ([]string, error)
}

type RepoImplement struct {
	pgConn boil.ContextExecutor
}

func New(pgConn boil.ContextExecutor) Repository {
	return RepoImplement{pgConn: pgConn}
}
