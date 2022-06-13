package command

import (
	"context"
	"sort"
	"time"

	"github.com/a76yyyy/tiktok/dal/pack"

	"github.com/a76yyyy/tiktok/kitex_gen/feed"

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
		return vis, nextTime, nil
	}

	if vis, err = pack.Videos(s.ctx, videos, &uid); err != nil {
		return vis, nextTime, nil
	}

	if len(videos) > 0 {
		sort.Slice(videos, func(i, j int) bool {
			return videos[i].UpdatedAt.UnixMilli() > videos[j].UpdatedAt.UnixMilli()
		})
		nextTime = videos[len(videos)-1].UpdatedAt.UnixMilli()
	} else {
		nextTime = time.Now().UnixMilli()
	}

	return vis, nextTime, nil
}
