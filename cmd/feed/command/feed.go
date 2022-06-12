package command

import (
	"context"
	"sort"
	"time"

	"github.com/a76yyyy/tiktok/cmd/user/command"
	"github.com/a76yyyy/tiktok/pkg/errno"

	"github.com/a76yyyy/tiktok/kitex_gen/feed"
	"github.com/a76yyyy/tiktok/kitex_gen/user"

	"github.com/a76yyyy/tiktok/dal/db"
)

const (
	LIMIT = 30 // 单次返回最大视频数
)

type GetUserFeedService struct {
	ctx context.Context
}

// NewGetUserFeedService new GetUserFeedService
func NewGetUserFeedService(ctx context.Context) *GetUserFeedService {
	return &GetUserFeedService{ctx: ctx}
}

// GetUserFeed get feed info.
func (s *GetUserFeedService) GetUserFeed(req *feed.DouyinFeedRequest, uid int64) (vis []*feed.Video, nextTime int64, err error) {
	videos, err := db.MGetVideos(s.ctx, LIMIT, req.LatestTime)
	if err != nil {
		return vis, nextTime, err
	}
	if len(videos) == 0 {
		return vis, nextTime, errno.ErrVideoNotFound
	}

	if len(videos) > 0 {
		sort.Slice(videos, func(i, j int) bool {
			return videos[i].UpdatedAt.UnixMilli() > videos[j].UpdatedAt.UnixMilli()
		})
		nextTime = videos[len(videos)-1].UpdatedAt.UnixMilli()
	} else {
		nextTime = time.Now().UnixMilli()
	}
	for _, v := range videos {
		user, err := command.NewMGetUserService(s.ctx).MGetUser(&user.DouyinUserRequest{UserId: int64(v.AuthorID)})
		if err != nil {
			return vis, nextTime, err
		}
		flag := false
		if uid != 0 {
			if result, err := db.GetFavoriteVideo(s.ctx, uid, int64(v.ID)); err != nil {
				return vis, nextTime, err
			} else if result.VideoID > 0 {
				flag = true
			} else {
				flag = false
			}

		}
		vis = append(vis, &feed.Video{
			Id:            int64(v.ID),
			Author:        user,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: int64(v.FavoriteCount),
			CommentCount:  int64(v.CommentCount),
			IsFavorite:    flag, // TODO 判断这个视频是否自己喜欢
		})
	}

	return vis, nextTime, nil
}
