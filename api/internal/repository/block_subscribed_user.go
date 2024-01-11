package repository

import (
	"assignment/internal/repository/orm"
	"context"
)

func (i RepoImplement) BlockUserInsideSubscription(ctx context.Context, userEmails []string) error {
	_, err := orm.Relationships(
		orm.RelationshipWhere.UserEmail1.EQ(userEmails[0]),
		orm.RelationshipWhere.UserEmail2.EQ(userEmails[1]),
	).UpdateAll(ctx, i.pgConn, orm.M{"subscription": "blocked"})
	return err
}
