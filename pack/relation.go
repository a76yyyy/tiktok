package pack

import (
	"context"

	"github.com/a76yyyy/tiktok/kitex_gen/user"

	"github.com/a76yyyy/tiktok/dal/db"
)

func FollowingList(ctx context.Context, vs []*db.Relation) []*user.User {
	users := make([]*db.User, 0)
	for _, v := range vs {
		users = append(users, &v.ToUser)
	}

	packUsers := Users(users)

	return packUsers
}

func FollowerList(ctx context.Context, vs []*db.Relation) []*user.User {
	users := make([]*db.User, 0)
	for _, v := range vs {
		users = append(users, &v.User)
	}

	packUsers := Users(users)

	return packUsers
}
