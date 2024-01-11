package repository

import (
	"assignment/internal/model"
	"assignment/internal/repository/orm"
	"context"
)

func (i RepoImplement) UpdateTopic(ctx context.Context, info model.UpdateInfo) error {
	_, err := orm.UserAccounts(
		orm.UserAccountWhere.Email.EQ(info.Sender),
	).UpdateAll(ctx, i.pgConn, orm.M{
		"topic":      "updated",
		"topic_body": info.Text,
	})
	return err
}
