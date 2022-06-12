package handlers

import (
	"bytes"
	"context"
	"io"
	"strconv"

	"github.com/a76yyyy/tiktok/pkg/errno"

	"github.com/a76yyyy/tiktok/kitex_gen/publish"
	"github.com/a76yyyy/tiktok/pack"

	"github.com/a76yyyy/tiktok/cmd/api/rpc"

	"github.com/gin-gonic/gin"
)

func PublishAction(c *gin.Context) {
	var paramVar PublishActionParam
	token := c.PostForm("token")
	title := c.PostForm("title")

	file, _, err := c.Request.FormFile("data")
	if err != nil {
		SendResponse(c, pack.BuildPublishResp(errno.ErrDecodingFailed))
		return
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		SendResponse(c, pack.BuildPublishResp(err))
		return
	}

	paramVar.Token = token
	paramVar.Title = title

	resp, err := rpc.PublishAction(context.Background(), &publish.DouyinPublishActionRequest{
		Title: paramVar.Title,
		Token: paramVar.Token,
		Data:  buf.Bytes(),
	})
	if err != nil {
		SendResponse(c, pack.BuildPublishResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}

func PublishList(c *gin.Context) {
	var paramVar UserParam
	userid, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		SendResponse(c, pack.BuildPublishListResp(errno.ErrBind))
		return
	}
	paramVar.UserId = int64(userid)
	paramVar.Token = c.Query("token")

	if len(paramVar.Token) == 0 || paramVar.UserId < 0 {
		SendResponse(c, pack.BuildPublishListResp(errno.ErrBind))
		return
	}

	resp, err := rpc.PublishList(context.Background(), &publish.DouyinPublishListRequest{
		UserId: paramVar.UserId,
		Token:  paramVar.Token,
	})
	if err != nil {
		SendResponse(c, pack.BuildPublishListResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}
