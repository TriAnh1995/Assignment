package repository

import "github.com/friendsofgo/errors"

var (
	FriendshipAlreadyExist       = errors.New("orm: unable to insert into friendships: pq: duplicate key value violates unique constraint \"unique_friends\"")
	UserAlreadyExist             = errors.New("orm: unable to insert into user_account: pq: duplicate key value violates unique constraint \"user_account_email_key\"")
	InternalErrorAddUser         = errors.New("orm: unable to insert into user_account: all expectations were already fulfilled, call to Query 'INSERT INTO \"user_account\" (\"name\",\"email\") VALUES ($1,$2) RETURNING \"user_id\"' with args [{Name: Ordinal:1 Value:InternalErrorUser} {Name: Ordinal:2 Value:internal-error-email@example.com}] was not expected")
	InternalErrorCheckEmail      = errors.New("orm: failed to check if user_account exists: all expectations were already fulfilled, call to Query 'SELECT COUNT(*) FROM \"user_account\" WHERE (\"user_account\".\"email\" = $1) LIMIT 1;' with args [{Name: Ordinal:1 Value:internal-error-email@example.com}] was not expected")
	InternalErrorAddFriendship   = errors.New("orm: unable to insert into friendships: all expectations were already fulfilled, call to Query 'INSERT INTO \"friendships\" (\"user_email_1\",\"user_email_2\") VALUES ($1,$2) RETURNING \"friendship_id\"' with args [{Name: Ordinal:1 Value:internal-error-email_1@example.com} {Name: Ordinal:2 Value:internal-error-email_2@example.com}] was not expected")
	InternalErrorCheckFriendship = errors.New("orm: failed to check if friendships exists: all expectations were already fulfilled, call to Query 'SELECT COUNT(*) FROM \"friendships\" WHERE ((user_email_1 = $1 AND user_email_2 = $2) OR (user_email_1 = $3 AND user_email_2 = $4)) LIMIT 1;' with args [{Name: Ordinal:1 Value:internal-error-email_1@example.com} {Name: Ordinal:2 Value:internal-error-email_2@example.com} {Name: Ordinal:3 Value:internal-error-email_2@example.com} {Name: Ordinal:4 Value:internal-error-email_1@example.com}] was not expected")
	InternalErrorGetFriendList   = errors.New("orm: failed to assign all query results to Friendship slice: bind failed to execute query: all expectations were already fulfilled, call to Query 'SELECT \"friendships\".* FROM \"friendships\" WHERE (user_email_1 = $1 OR user_email_2 = $2);' with args [{Name: Ordinal:1 Value:already_exist_email_1@example.com} {Name: Ordinal:2 Value:already_exist_email_1@example.com}] was not expected")
)
