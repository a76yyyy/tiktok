package handlers

import (
	"context"
	"strconv"

	"github.com/a76yyyy/tiktok/pkg/errno"

	"github.com/a76yyyy/tiktok/kitex_gen/feed"
	"github.com/a76yyyy/tiktok/pack"

	"github.com/a76yyyy/tiktok/cmd/api/rpc"

	"github.com/gin-gonic/gin"
)

func GetUserFeed(c *gin.Context) {
	var feedVar FeedParam
	var laststTime int64
	var token string
	lastst_time := c.Query("latest_time")
	if len(lastst_time) != 0 {
		if latesttime, err := strconv.Atoi(lastst_time); err != nil {
			SendResponse(c, pack.BuildVideoResp(errno.ErrDecodingFailed))
			return
		} else {
			laststTime = int64(latesttime)
		}
	}

	feedVar.LatestTime = &laststTime

	token = c.Query("token")
	feedVar.Token = &token

	resp, err := rpc.GetUserFeed(context.Background(), &feed.DouyinFeedRequest{
		LatestTime: feedVar.LatestTime,
		Token:      feedVar.Token,
	})
	if err != nil {
		SendResponse(c, pack.BuildVideoResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}
