package pack

import (
	"context"

	"github.com/a76yyyy/tiktok/kitex_gen/feed"

	"github.com/a76yyyy/tiktok/dal/db"
)

// Video pack feed info
func Video(ctx context.Context, v *db.Video, fromID int64) (*feed.Video, error) {
	if v == nil {
		return nil, nil
	}
	user, _ := db.MGetUser(ctx, int64(v.AuthorID))

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
				if result, err := db.GetFavoriteRelation(ctx, *fromID, int64(v.ID)); err != nil {
					flag = false
				} else if result.VideoID > 0 {
					flag = true
				} else {
					flag = false
				}

			}
			video2.IsFavorite = flag
			videos = append(videos, video2)
		}
	}
	return videos, nil
}
