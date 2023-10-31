package repository

import (
	"assignment/internal/model"
	"assignment/internal/repository/orm"
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// AddUser insert user from input to db
func (i RepoImplement) AddUser(ctx context.Context, user model.User) error {
	account := &orm.UserAccount{
		Name:  user.Name,
		Email: user.Email,
	}
	return (*orm.UserAccount).
		Insert(account, ctx, i.pgConn, boil.Infer())
}
