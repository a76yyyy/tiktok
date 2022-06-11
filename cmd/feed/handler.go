package main

import (
	"context"
	"sort"
	"time"

	"github.com/a76yyyy/tiktok/cmd/user/command"
	"github.com/a76yyyy/tiktok/dal/db"
	"github.com/a76yyyy/tiktok/kitex_gen/feed"
	"github.com/a76yyyy/tiktok/kitex_gen/user"
	"github.com/a76yyyy/tiktok/pack"
	"github.com/a76yyyy/tiktok/pkg/errno"
)

// FeedSrvImpl implements the last service interface defined in the IDL.
type FeedSrvImpl struct{}

// GetUserFeed implements the FeedSrvImpl interface.
func (s *FeedSrvImpl) GetUserFeed(ctx context.Context, req *feed.DouyinFeedRequest) (resp *feed.DouyinFeedResponse, err error) {
	uid := 0
	if *req.Token != "" {
		claim, err := Jwt.ParseToken(*req.Token)
		if err != nil {
			resp = pack.BuildVideoResp(errno.ErrTokenInvalid)
			return resp, nil
		} else {
			uid = int(claim.Id)
		}
	}

	videos, err := db.MGetVideos(ctx, *req.LatestTime)
	if err != nil {
		resp = pack.BuildVideoResp(err)
		return resp, nil
	}
	if len(videos) == 0 {
		resp = pack.BuildVideoResp(errno.ErrVideoNotFound)
		return resp, nil
	}
	var vis []*feed.Video
	var nextTime int64
	if len(videos) > 0 {
		sort.Slice(videos, func(i, j int) bool {
			return videos[i].UpdatedAt.UnixMilli() > videos[j].UpdatedAt.UnixMilli()
		})
		nextTime = videos[len(videos)-1].UpdatedAt.UnixMilli()
	} else {
		nextTime = time.Now().UnixMilli()
	}
	for _, v := range videos {
		user, err := command.NewMGetUserService(ctx).MGetUser(&user.DouyinUserRequest{UserId: int64(v.AuthorID)})
		if err != nil {
			return pack.BuildVideoResp(err), nil
		}
		flag := false
		if uid != 0 {
			if result, err := db.GetFavoriteVideo(ctx, uid, int(v.ID)); err != nil {
				resp = pack.BuildVideoResp(err)
				return resp, nil
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

	resp = pack.BuildVideoResp(errno.Success)
	resp.VideoList = vis
	resp.NextTime = &nextTime
	return resp, nil
}

// GetVideoById implements the FeedSrvImpl interface.
func (s *FeedSrvImpl) GetVideoById(ctx context.Context, req *feed.VideoIdRequest) (resp *feed.Video, err error) {
	// TODO: Your code here...
	return
}
