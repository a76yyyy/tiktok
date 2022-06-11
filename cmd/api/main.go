// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/a76yyyy/tiktok/cmd/api/handlers"
	"github.com/a76yyyy/tiktok/cmd/api/rpc"
	"go.uber.org/zap"

	// jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/a76yyyy/tiktok/pkg/jwt"
	"github.com/a76yyyy/tiktok/pkg/ttviper"
	"github.com/cloudwego/kitex/pkg/klog"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

var (
	Config      = ttviper.ConfigInit("TIKTOK_API", "apiConfig")
	ServiceName = Config.Viper.GetString("Server.Name")
	ServiceAddr = fmt.Sprintf("%s:%d", Config.Viper.GetString("Server.Address"), Config.Viper.GetInt("Server.Port"))
	EtcdAddress = fmt.Sprintf("%s:%d", Config.Viper.GetString("Etcd.Address"), Config.Viper.GetInt("Etcd.Port"))
	Jwt         *jwt.JWT
)

func Init() {
	rpc.InitRPC(&Config)
	Jwt = jwt.NewJWT([]byte(Config.Viper.GetString("JWT.signingKey")))
}

func main() {
	logger := Config.InitLogger()
	defer logger.Sync()

	zap.ReplaceGlobals(logger)

	Init()

	r := gin.New()

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	r.Use(ginzap.Ginzap(zap.L(), time.RFC3339, false))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	r.Use(ginzap.RecoveryWithZap(zap.L(), true))

	douyin := r.Group("/douyin")
	user := douyin.Group("/user")
	user.POST("/login/", handlers.Login)
	user.POST("/register/", handlers.Register)
	user.GET("/", handlers.GetUserById)
	video := r.Group("/feed")
	video.GET("/", handlers.GetUserFeed)

	if err := http.ListenAndServe(ServiceAddr, r); err != nil {
		klog.Fatal(err)
	}
}
