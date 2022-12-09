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
 * @LastEditTime: 2022-06-18 23:32:21
 * @FilePath: /tiktok/cmd/api/handlers/handler.go
 * @Description: 定义 所有API handler 的 输入输出参数
 */

// 定义 Hertz HTTP API 的 handler
package handlers

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// SendResponse pack response
func SendResponse(c *app.RequestContext, response interface{}) {
	c.JSON(consts.StatusOK, response)
}

// 用户注册 handler 输入参数
type UserRegisterParam struct {
	UserName string `json:"username"` // 用户名
	PassWord string `json:"password"` // 用户密码
}

// 用户信息 输出参数
type UserParam struct {
	UserId int64  `json:"user_id,omitempty"` // 用户id
	Token  string `json:"token,omitempty"`   // 用户鉴权token
}

// 视频流 handler 输入参数
type FeedParam struct {
	LatestTime *int64  `json:"latest_time,omitempty"` // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	Token      *string `json:"token,omitempty"`       // 可选参数，登录用户设置
}

// 发布视频操作 handler 输入参数
type PublishActionParam struct {
	Token string `json:"token,omitempty"` // 用户鉴权token
	Data  []byte `json:"data,omitempty"`  // 视频数据
	Title string `json:"title,omitempty"` // 视频标题
}

// 点赞操作 handler 输入参数
type FavoriteActionParam struct {
	UserId     int64  `json:"user_id,omitempty"`     // 用户id
	Token      string `json:"token,omitempty"`       // 用户鉴权token
	VideoId    int64  `json:"video_id,omitempty"`    // 视频id
	ActionType int32  `json:"action_type,omitempty"` // 1-点赞，2-取消点赞
}

// 评论操作  handler 输入参数
type CommentActionParam struct {
	UserId      int64   `json:"user_id,omitempty"`      // 用户id
	Token       string  `json:"token,omitempty"`        // 用户鉴权token
	VideoId     int64   `json:"video_id,omitempty"`     // 视频id
	ActionType  int32   `json:"action_type,omitempty"`  // 1-发布评论，2-删除评论
	CommentText *string `json:"comment_text,omitempty"` // 用户填写的评论内容，在action_type=1的时候使用
	CommentId   *int64  `json:"comment_id,omitempty"`   // 要删除的评论id，在action_type=2的时候使用
}

// 获取评论列表 handler 输入参数
type CommentListParam struct {
	Token   string `json:"token,omitempty"`    // 用户鉴权token
	VideoId int64  `json:"video_id,omitempty"` // 视频id
}

// 关注操作 handler 输入参数
type RelationActionParam struct {
	UserId     int64  `json:"user_id,omitempty"`     // 用户id
	Token      string `json:"token,omitempty"`       // 用户鉴权token
	ToUserId   int64  `json:"to_user_id,omitempty"`  // 对方用户id
	ActionType int32  `json:"action_type,omitempty"` // 1-关注，2-取消关注
}
