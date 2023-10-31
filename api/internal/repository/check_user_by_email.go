package repository

import (
	"assignment/internal/repository/orm"
	"context"
)

// CheckUserByEmail check whether user are already in db or not
func (i RepoImplement) CheckUserByEmail(ctx context.Context, userEmails string) (bool, error) {

	exist, err := orm.UserAccounts(
		orm.UserAccountWhere.Email.EQ(userEmails)).Exists(ctx, i.pgConn)
	return exist, err

}
