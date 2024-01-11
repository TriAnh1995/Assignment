package repository

import (
	"assignment/internal/repository/orm"
	"context"
)

func (i RepoImplement) CheckSubscription(ctx context.Context, userEmails []string) (string, error) {
	exist, err := orm.Relationships(
		orm.RelationshipWhere.UserEmail1.EQ(userEmails[0]),
		orm.RelationshipWhere.UserEmail2.EQ(userEmails[1]),
	).Exists(ctx, i.pgConn)
	if !exist {
		return "None", err
	} else {
		status, err := orm.Relationships(
			orm.RelationshipWhere.UserEmail1.EQ(userEmails[0]),
			orm.RelationshipWhere.UserEmail2.EQ(userEmails[1]),
		).One(ctx, i.pgConn)
		switch status.Subscription {
		case orm.StatusTypeFollowed:
			return "Followed", err
		case orm.StatusTypeBlocked:
			return "Blocked", err
		case orm.StatusTypeDefault:
			return "Default", err
		}
	}
	return "", err
}
