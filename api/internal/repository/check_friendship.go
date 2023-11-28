package repository

import (
	"assignment/internal/repository/orm"
	"context"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (i RepoImplement) CheckFriendship(ctx context.Context, userEmails []string) (bool, error) {
	// Using Query to find the rows that exists both emails input regardless the columns
	exist, err := orm.Friendships(
		qm.Where("(user_email_1 = ? AND user_email_2 = ?) OR (user_email_1 = ? AND user_email_2 = ?)",
			userEmails[0], userEmails[1], userEmails[1], userEmails[0]),
	).Exists(ctx, i.pgConn)
	return exist, err
}
