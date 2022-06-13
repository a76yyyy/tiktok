package handlers

import (
	"context"
	"strconv"

	"github.com/a76yyyy/tiktok/pkg/errno"

	"github.com/a76yyyy/tiktok/dal/pack"
	"github.com/a76yyyy/tiktok/kitex_gen/comment"

	"github.com/a76yyyy/tiktok/cmd/api/rpc"

	"github.com/gin-gonic/gin"
)

func CommentAction(c *gin.Context) {
	var paramVar CommentActionParam
	token := c.Query("token")
	video_id := c.Query("video_id")
	action_type := c.Query("action_type")

	vid, err := strconv.Atoi(video_id)
	if err != nil {
		SendResponse(c, pack.BuildCommentActionResp(errno.ErrBind))
		return
	}
	act, err := strconv.Atoi(action_type)
	if err != nil {
		SendResponse(c, pack.BuildCommentActionResp(errno.ErrBind))
		return
	}

	paramVar.Token = token
	paramVar.VideoId = int64(vid)
	paramVar.ActionType = int32(act)

	rpcReq := comment.DouyinCommentActionRequest{
		VideoId:    paramVar.VideoId,
		Token:      paramVar.Token,
		ActionType: paramVar.ActionType,
	}

	if act == 1 {
		comment_text := c.Query("comment_text")
		rpcReq.CommentText = &comment_text
	} else {
		comment_id := c.Query("comment_id")
		cid, err := strconv.Atoi(comment_id)
		if err != nil {
			SendResponse(c, pack.BuildCommentActionResp(errno.ErrBind))
			return
		}
		cid64 := int64(cid)
		rpcReq.CommentId = &cid64
	}

	resp, err := rpc.CommentAction(context.Background(), &rpcReq)
	if err != nil {
		SendResponse(c, pack.BuildCommentActionResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}

func CommentList(c *gin.Context) {
	var paramVar CommentListParam
	videoid, err := strconv.Atoi(c.Query("video_id"))
	if err != nil {
		SendResponse(c, pack.BuildCommentListResp(errno.ErrBind))
		return
	}
	paramVar.VideoId = int64(videoid)
	paramVar.Token = c.Query("token")

	if len(paramVar.Token) == 0 || paramVar.VideoId < 0 {
		SendResponse(c, pack.BuildCommentListResp(errno.ErrBind))
		return
	}

	resp, err := rpc.CommentList(context.Background(), &comment.DouyinCommentListRequest{
		VideoId: paramVar.VideoId,
		Token:   paramVar.Token,
	})
	if err != nil {
		SendResponse(c, pack.BuildCommentListResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}
