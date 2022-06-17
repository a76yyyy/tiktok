package pack

import (
	"context"
	"errors"

	"github.com/a76yyyy/tiktok/kitex_gen/user"
	"gorm.io/gorm"

	"github.com/a76yyyy/tiktok/dal/db"
)

func FollowingList(ctx context.Context, vs []*db.Relation, fromID int64) ([]*user.User, error) {
	users := make([]*db.User, 0)
	for _, v := range vs {
		user2, err := db.GetUserByID(ctx, int64(v.ToUserID))
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		users = append(users, user2)
	}

	return Users(ctx, users, fromID)
}

func FollowerList(ctx context.Context, vs []*db.Relation, fromID int64) ([]*user.User, error) {
	users := make([]*db.User, 0)
	for _, v := range vs {
		user2, err := db.GetUserByID(ctx, int64(v.UserID))
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		users = append(users, user2)
	}

	return Users(ctx, users, fromID)
}
