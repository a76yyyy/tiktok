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
 * @Date: 2022-06-11 10:14:40
 * @LastEditors: a76yyyy q981331502@163.com
 * @LastEditTime: 2022-06-18 23:48:52
 * @FilePath: /tiktok/cmd/api/main.go
 * @Description: 使用 Gin 提供 API 服务将 HTTP 请求发送给 RPC 微服务端
 */

// 使用 Gin 提供 API 服务将 HTTP 请求发送给 RPC 微服务端
package main

import (
	"fmt"
	"net/http"

	"github.com/a76yyyy/tiktok/cmd/api/handlers"
	"github.com/a76yyyy/tiktok/cmd/api/rpc"
	"github.com/a76yyyy/tiktok/pkg/dlog"
	"github.com/a76yyyy/tiktok/pkg/jwt"
	"github.com/a76yyyy/tiktok/pkg/ttviper"

	// jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
)

var (
	Config      = ttviper.ConfigInit("TIKTOK_API", "apiConfig")
	ServiceName = Config.Viper.GetString("Server.Name")
	ServiceAddr = fmt.Sprintf("%s:%d", Config.Viper.GetString("Server.Address"), Config.Viper.GetInt("Server.Port"))
	EtcdAddress = fmt.Sprintf("%s:%d", Config.Viper.GetString("Etcd.Address"), Config.Viper.GetInt("Etcd.Port"))
	Jwt         *jwt.JWT
)

// 初始化 API 配置
func Init() {
	rpc.InitRPC(&Config)
	Jwt = jwt.NewJWT([]byte(Config.Viper.GetString("JWT.signingKey")))
}

// 初始化 GIN API 及 Router
func main() {
	logger := dlog.InitLog(3)
	defer logger.Sync()

	klog.SetLogger(logger)

	Init()

	r := gin.New()

	douyin := r.Group("/douyin")

	user := douyin.Group("/user")
	user.POST("/login/", handlers.Login)
	user.POST("/register/", handlers.Register)
	user.GET("/", handlers.GetUserById)

	video := douyin.Group("/feed")
	video.GET("/", handlers.GetUserFeed)

	publish := douyin.Group("/publish")
	publish.POST("/action/", handlers.PublishAction)
	publish.GET("/list/", handlers.PublishList)

	favorite := douyin.Group("/favorite")
	favorite.POST("/action/", handlers.FavoriteAction)
	favorite.GET("/list/", handlers.FavoriteList)

	comment := douyin.Group("/comment")
	comment.POST("/action/", handlers.CommentAction)
	comment.GET("/list/", handlers.CommentList)

	relation := douyin.Group("/relation")
	relation.POST("/action/", handlers.RelationAction)
	relation.GET("/follow/list/", handlers.RelationFollowList)
	relation.GET("/follower/list/", handlers.RelationFollowerList)

	if err := http.ListenAndServe(ServiceAddr, r); err != nil {
		klog.Fatal(err)
	}
}
