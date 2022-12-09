// Copyright 2022 a76yyyy && CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
 * @Author: a76yyyy q981331502@163.com
 * @Date: 2022-06-11 10:10:05
 * @LastEditors: a76yyyy q981331502@163.com
 * @LastEditTime: 2022-06-18 23:33:49
 * @FilePath: /tiktok/cmd/api/handlers/user.go
 * @Description: 定义 User API 的 handler
 */

package handlers

import (
	"context"
	"strconv"

	"github.com/a76yyyy/tiktok/cmd/api/rpc"
	"github.com/a76yyyy/tiktok/dal/pack"
	"github.com/a76yyyy/tiktok/kitex_gen/user"
	"github.com/a76yyyy/tiktok/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
)

// 传递 注册用户操作 的上下文至 User 服务的 RPC 客户端, 并获取相应的响应.
func Register(ctx context.Context, c *app.RequestContext) {
	var registerVar UserRegisterParam
	registerVar.UserName = c.Query("username")
	registerVar.PassWord = c.Query("password")

	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		SendResponse(c, pack.BuilduserRegisterResp(errno.ErrBind))
		return
	}

	resp, err := rpc.Register(ctx, &user.DouyinUserRegisterRequest{
		Username: registerVar.UserName,
		Password: registerVar.PassWord,
	})
	if err != nil {
		SendResponse(c, pack.BuilduserRegisterResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}

// 传递 注册用户登录操作 的上下文至 User 服务的 RPC 客户端, 并获取相应的响应.
func Login(ctx context.Context, c *app.RequestContext) {
	var registerVar UserRegisterParam
	registerVar.UserName = c.Query("username")
	registerVar.PassWord = c.Query("password")

	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		SendResponse(c, pack.BuilduserRegisterResp(errno.ErrBind))
		return
	}

	resp, err := rpc.Login(ctx, &user.DouyinUserRegisterRequest{
		Username: registerVar.UserName,
		Password: registerVar.PassWord,
	})
	if err != nil {
		SendResponse(c, pack.BuilduserRegisterResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}

// 传递 获取注册用户`UserID`操作 的上下文至 User 服务的 RPC 客户端, 并获取相应的响应.
func GetUserById(ctx context.Context, c *app.RequestContext) {
	var userVar UserParam
	userid, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		SendResponse(c, pack.BuilduserUserResp(errno.ErrBind))
		return
	}
	userVar.UserId = int64(userid)
	userVar.Token = c.Query("token")

	if len(userVar.Token) == 0 || userVar.UserId < 0 {
		SendResponse(c, pack.BuilduserUserResp(errno.ErrBind))
		return
	}

	resp, err := rpc.GetUserById(ctx, &user.DouyinUserRequest{
		UserId: userVar.UserId,
		Token:  userVar.Token,
	})
	if err != nil {
		SendResponse(c, pack.BuilduserUserResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}
