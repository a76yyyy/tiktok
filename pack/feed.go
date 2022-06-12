package pack

import (
	"context"

	"github.com/a76yyyy/tiktok/kitex_gen/feed"

	"github.com/a76yyyy/tiktok/dal/db"
)

// Video pack feed info
func Video(ctx context.Context, v *db.Video, uid int64) *feed.Video {
	if v == nil {
		return nil
	}
	user, _ := db.MGetUser(ctx, uid)

	author := User(user)
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
	}
}

// Videos pack list of video info
func Videos(ctx context.Context, vs []*db.Video, uid *int64) ([]*feed.Video, error) {
	videos := make([]*feed.Video, 0)
	for _, v := range vs {
		if video2 := Video(ctx, v, int64(v.AuthorID)); video2 != nil {
			flag := false
			if *uid != 0 {
				if result, err := db.GetFavoriteRelation(ctx, *uid, int64(v.ID)); err != nil {
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
