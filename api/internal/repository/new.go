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
	GetFriendsList(context.Context, string) ([]string, error)
	GetBlockedList(context.Context, string) ([]string, error)
	UpdateTopic(context.Context, model.UpdateInfo) error
	GetFollowerList(context.Context, string) ([]string, error)
}

type RepoImplement struct {
	pgConn boil.ContextExecutor
}

func New(pgConn boil.ContextExecutor) Repository {
	return RepoImplement{pgConn: pgConn}
}

func (i RepoImplement) GetFriendsList(ctx context.Context, userEmail string) ([]string, error) {
	return []string{}, nil
}

func (i RepoImplement) GetBlockedList(ctx context.Context, userEmail string) ([]string, error) {
	return []string{}, nil
}

func (i RepoImplement) GetFollowerList(ctx context.Context, userEmail string) ([]string, error) {
	return []string{}, nil
}

func (i RepoImplement) UpdateTopic(ctx context.Context, updateInfo model.UpdateInfo) error {
	return nil
}
