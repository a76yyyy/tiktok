package pack

import (
	"context"

	"github.com/a76yyyy/tiktok/kitex_gen/user"

	"github.com/a76yyyy/tiktok/dal/db"
)

func FollowingList(ctx context.Context, vs []*db.Relation) []*user.User {
	users := make([]*db.User, 0)
	for _, v := range vs {
		user2, err := db.MGetUser(ctx, int64(v.ToUserID))
		if err != nil {
			return nil
		}
		users = append(users, user2)
	}

	packUsers := Users(users)

	return packUsers
}

func FollowerList(ctx context.Context, vs []*db.Relation) []*user.User {
	users := make([]*db.User, 0)
	for _, v := range vs {
		user2, err := db.MGetUser(ctx, int64(v.UserID))
		if err != nil {
			return nil
		}
		users = append(users, user2)
	}

	packUsers := Users(users)

	return packUsers
}
