package handlers

import (
	"context"
	"strconv"

	"github.com/a76yyyy/tiktok/pkg/errno"

	"github.com/a76yyyy/tiktok/dal/pack"
	"github.com/a76yyyy/tiktok/kitex_gen/favorite"

	"github.com/a76yyyy/tiktok/cmd/api/rpc"

	"github.com/gin-gonic/gin"
)

func FavoriteAction(c *gin.Context) {
	var paramVar FavoriteActionParam
	token := c.Query("token")
	video_id := c.Query("video_id")
	action_type := c.Query("action_type")

	vid, err := strconv.Atoi(video_id)
	if err != nil {
		SendResponse(c, pack.BuildFavoriteActionResp(errno.ErrBind))
		return
	}
	act, err := strconv.Atoi(action_type)
	if err != nil {
		SendResponse(c, pack.BuildFavoriteActionResp(errno.ErrBind))
		return
	}

	paramVar.Token = token
	paramVar.VideoId = int64(vid)
	paramVar.ActionType = int32(act)

	resp, err := rpc.FavoriteAction(context.Background(), &favorite.DouyinFavoriteActionRequest{
		VideoId:    paramVar.VideoId,
		Token:      paramVar.Token,
		ActionType: paramVar.ActionType,
	})
	if err != nil {
		SendResponse(c, pack.BuildFavoriteActionResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}

func FavoriteList(c *gin.Context) {
	var paramVar UserParam
	userid, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		SendResponse(c, pack.BuildFavoriteListResp(errno.ErrBind))
		return
	}
	paramVar.UserId = int64(userid)
	paramVar.Token = c.Query("token")

	if len(paramVar.Token) == 0 || paramVar.UserId < 0 {
		SendResponse(c, pack.BuildFavoriteListResp(errno.ErrBind))
		return
	}

	resp, err := rpc.FavoriteList(context.Background(), &favorite.DouyinFavoriteListRequest{
		UserId: paramVar.UserId,
		Token:  paramVar.Token,
	})
	if err != nil {
		SendResponse(c, pack.BuildFavoriteListResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}
