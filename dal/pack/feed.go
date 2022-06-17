package pack

import (
	"context"
	"errors"

	"github.com/a76yyyy/tiktok/kitex_gen/feed"
	"gorm.io/gorm"

	"github.com/a76yyyy/tiktok/dal/db"
)

// Video pack feed info
func Video(ctx context.Context, v *db.Video, fromID int64) (*feed.Video, error) {
	if v == nil {
		return nil, nil
	}
	user, err := db.GetUserByID(ctx, int64(v.AuthorID))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	author, err := User(ctx, user, fromID)
	if err != nil {
		return nil, err
	}
	favorite_count := int64(v.FavoriteCount)
	comment_count := int64(v.CommentCount)

	return &feed.Video{
		Id:            int64(v.ID),
		Author:        author,
		PlayUrl:       v.PlayUrl,
		CoverUrl:      v.CoverUrl,
		FavoriteCount: favorite_count,
		CommentCount:  comment_count,
		Title:         v.Title,
	}, nil
}

// Videos pack list of video info
func Videos(ctx context.Context, vs []*db.Video, fromID *int64) ([]*feed.Video, error) {
	videos := make([]*feed.Video, 0)
	for _, v := range vs {
		video2, err := Video(ctx, v, *fromID)
		if err != nil {
			return nil, err
		}

		if video2 != nil {
			flag := false
			if *fromID != 0 {
				results, err := db.GetFavoriteRelation(ctx, *fromID, int64(v.ID))
				if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, err
				}

				if results != nil {
					flag = true
				}
			}
			video2.IsFavorite = flag
			videos = append(videos, video2)
		}
	}
	return videos, nil
}
