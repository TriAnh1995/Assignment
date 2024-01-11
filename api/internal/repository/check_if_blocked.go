package repository

import (
	"assignment/internal/repository/orm"
	"context"
)

func (i RepoImplement) CheckIfBlocked(ctx context.Context, userEmails []string) (bool, error) {

	exists, err := orm.Relationships(
		orm.RelationshipWhere.UserEmail1.EQ(userEmails[0]),
		orm.RelationshipWhere.UserEmail2.EQ(userEmails[1]),
		orm.RelationshipWhere.Subscription.EQ("blocked"),
	).Exists(ctx, i.pgConn)
	return exists, err
}
